package controllers

import "github.com/revel/revel"

type Post struct {
	*revel.Controller
}

func (c Post) View(id int) revel.Result {
	return c.Render()
}

func (c Post) PreNew() revel.Result {
    return c.Render()
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
