package controllers

import "github.com/revel/revel"

type Connect struct {
	*revel.Controller
}

func (c Connect) Index() revel.Result {
	return c.Render()
}

func (c Connect) Weibo() revel.Result {
    return c.Render()
}