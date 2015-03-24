package controllers

import (
    "github.com/revel/revel"
    "web/app/models"
    "time"
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
    user := models.User{Name: "Soddy", Username: "Hello", DateCreated: time.Now(), Password: "123"}
    c.Txn.NewRecord(user)
    c.Txn.Create(&user)

    var users = []*models.User{}
    c.Txn.Limit(100).Find(&users)
    return c.RenderJson(users[0])
}