package controllers

import "github.com/revel/revel"

type Chart struct {
	BaseController
}

func (c Chart) Index() revel.Result {
	return c.Render()
}
