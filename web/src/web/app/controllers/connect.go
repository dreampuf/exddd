package controllers

import (
    "io"
    "fmt"
    "time"
    "strings"
    "net/url"
    "net/http"
    "io/ioutil"
    "encoding/json"

    "github.com/Pallinder/go-randomdata"
    "github.com/revel/revel"
    "golang.org/x/oauth2"

    "web/app/routes"
    "web/app/models"
)


type Connect struct {
    GormController
}

type tokenJSON struct {
    AccessToken  string         `json:"access_token"`
    Uid          string         `json:"uid"`
    ExpiresIn    expirationTime `json:"expires_in"` // at least PayPal returns string, while most return number
    Scope        string         `json:"scope"`
}

func (e *tokenJSON) expiry() (t time.Time) {
    if v := e.ExpiresIn; v != 0 {
        return time.Now().Add(time.Duration(v) * time.Second)
    }
    return
}

type expirationTime int32

func (e *expirationTime) UnmarshalJSON(b []byte) error {
    var n json.Number
    err := json.Unmarshal(b, &n)
    if err != nil {
        return err
    }
    i, err := n.Int64()
    if err != nil {
        return err
    }
    *e = expirationTime(i)
    return nil
}

type userJSON struct {
    ProfileURL string           `json:"profile_url"`
    Domain     string           `json:"domain"`
    ScreenName string           `json:"screen_name"`
    Thumbnail  string           `json:"avatar_large"`
    Gender     string           `json:"gender"`
}

func rootURL(c Connect, ctrol string) string {
    var schema string = "https://"
    if !(c.Request.URL.IsAbs() || revel.Config.BoolDefault("http.ssl", false)) {
        schema = "http://"
    }
    return fmt.Sprintf("%s%s%s", schema, c.Request.Host, ctrol)
}

func weiboConf(redirectUrl string) *oauth2.Config {
    return &oauth2.Config{
        ClientID:     revel.Config.StringDefault("oauth2.weibo.id", ""),
        ClientSecret: revel.Config.StringDefault("oauth2.weibo.secret", ""),
        RedirectURL:  redirectUrl,
        Scopes:       []string{"email"},
        Endpoint: oauth2.Endpoint{
            AuthURL:  "https://api.weibo.com/oauth2/authorize",
            TokenURL: "https://api.weibo.com/oauth2/access_token",
        },
    }
}



func (c Connect) Index() revel.Result {
	return c.Render()
}

func (c Connect) Weibo() revel.Result {
    var weibo = weiboConf(rootURL(c, routes.Connect.WeiboToken()))
    url := weibo.AuthCodeURL("auth", oauth2.AccessTypeOnline)
    return c.Redirect(url)
}

func condVal(v string) []string {
    if v == "" {
        return nil
    }
    return []string{v}
}

func StupidWeiboToken(c *oauth2.Config, code string) (*oauth2.Token, tokenJSON, error) {
    v := url.Values{
		"grant_type":   {"authorization_code"},
		"code":         {code},
		"redirect_uri": condVal(c.RedirectURL),
		"scope":        condVal(strings.Join(c.Scopes, " ")),
	}
    req, err := http.NewRequest("POST", c.Endpoint.TokenURL, strings.NewReader(v.Encode()))
    if err != nil {
        return nil, tokenJSON{}, err
    }
    req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
    req.SetBasicAuth(c.ClientID, c.ClientSecret)
    r, err := http.DefaultClient.Do(req)
    if err != nil {
        return nil, tokenJSON{}, err
    }
    defer r.Body.Close()
    body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1<<20))
    var tj tokenJSON
    revel.INFO.Println(string(body))
    if err = json.Unmarshal(body, &tj); err != nil {
        return nil, tokenJSON{}, err
    }
    revel.INFO.Println(tj)
    token := &oauth2.Token{
        AccessToken:  tj.AccessToken,
        TokenType:    "authorization_code",
        Expiry:       tj.expiry(),
    }
    return token, tj, nil
}

func (c Connect) WeiboToken() revel.Result {
    code := c.Params.Get("code")
    var weibo = weiboConf(rootURL(c, routes.Connect.WeiboToken()))

    //tok, err := weibo.Exchange(oauth2.NoContext, code)
    tok, raw, err := StupidWeiboToken(weibo, code)
    if err != nil {
        revel.WARN.Println(err)
        return c.Redirect(Connect.Index)
    }
    if !tok.Valid() {
        revel.WARN.Println("Token invalid:", tok)
    }

    client := weibo.Client(oauth2.NoContext, tok)
    profileInfoURL := fmt.Sprintf("https://api.weibo.com/2/users/show.json?uid=%s&access_token=%s", raw.Uid, tok.AccessToken)
    resp, err := client.Get(profileInfoURL)
    if err != nil {
        revel.WARN.Println(err)
        return c.Redirect(Connect.Failed)
    }

    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        revel.WARN.Println(err)
        return c.Redirect(Connect.Failed)
    }
    revel.INFO.Println(string(body))
    userinfo := userJSON{}

    if err := json.Unmarshal(body, &userinfo); err != nil {
        revel.WARN.Println(err)
        return c.Redirect(Connect.Failed)
    }

    gender := randomdata.Female
    if userinfo.Gender == "m" {
        gender = randomdata.Male
    }
    user := new(models.User)
    userTmp := new(models.User)
    c.Txn.Where("weibo_id = (?)", raw.Uid).First(user)
    if user == nil {
        revel.INFO.Printf("WeiboID %s didn't exists; Create new one.\n", raw.Uid)
        var postfix string = ""
        name := userinfo.ProfileURL
        if len(userinfo.Domain) > 0 {
            name = userinfo.Domain
        }
        for {
            if c.Txn.Where("name = ?", name + postfix).Find(userTmp).RecordNotFound() {
                user.Name = name + postfix
                break
            }
            postfix = randomdata.FirstName(gender)
        }
        postfix = ""
        name = userinfo.ScreenName
        for {
            if c.Txn.Where("nickname = ?", name + postfix).Find(userTmp).RecordNotFound() {
                user.Nickname = name + postfix
                break
            }
            postfix = randomdata.FirstName(gender)
        }
        user.WeiboID = raw.Uid
        user.WeiboToken = tok.AccessToken
        user.WeiboExpires = tok.Expiry

        if count := c.Txn.Save(user).RowsAffected; count != 1 {
            revel.WARN.Printf("User %s (Weibo) create failed.\n")
            return c.Redirect(Connect.Failed)
        }
        revel.WARN.Printf("User [%d]%s created", user.Id, user.Name)
    }

    c.Session["uid"] = string(user.Id)
    return c.Redirect(Connect.Index)
}

func (c Connect) Failed() revel.Result {
    return c.Render()
}
