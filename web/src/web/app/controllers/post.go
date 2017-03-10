package controllers

import (
    "github.com/revel/revel"
    "web/app/models"
)

type Post struct {
    GormController
}

func (c Post) View(id int) revel.Result {
	return c.Render()
}

func (c Post) List(page int) revel.Result {
    pageLen := revel.Config.IntDefault("page.len", 20)
    posts := []models.Post{}
    if err := c.Txn.Where("date_deleted is NULL").Order("id").Limit(pageLen).Find(&posts).Error; err != nil {
        revel.WARN.Printf("Fetch posts data failed: %+v", err)
    }
    revel.INFO.Println(posts)
    return c.Render(posts)
}

func (c Post) New(name string, content string) revel.Result {
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
