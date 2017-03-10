package controllers

import (
    "github.com/revel/revel"
    "web/app/models"
    "time"
    "github.com/Pallinder/go-randomdata"
)

type App struct {
    GormController
}

func (c App) Index() revel.Result {
	return c.Render()
}

func (c App) SavePost(name string, content string) revel.Result {
    revel.TRACE.Println("%s\n%s\n", name, content)
    c.Validation.Required(name).Message("必须")
    c.Validation.Required(content).Message("必须")

    if c.Validation.HasErrors() {
        c.Validation.Keep()
        c.FlashParams()
        return c.Redirect(App.Index)
    }
    return c.Redirect(App.Index)
}

func (c App) UserList() revel.Result {
    user := models.User{Name: "Soddy", Nickname: "Hello", DateCreated: time.Now(), Password: "123"}
    c.Txn.NewRecord(user)
    c.Txn.Create(&user)

    for i := 0; i < 10; i ++ {
        post := models.Post{
            Title: randomdata.Letters(10),
            Body:  randomdata.Paragraph(),
            User:  user,
        }
        c.Txn.Create(&post)
    }

    var users = []*models.User{}
    c.Txn.Limit(100).Find(&users)
    return c.RenderJson(users)
}