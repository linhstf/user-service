package controllers

import "github.com/revel/revel"

type Product struct {
	BaseController
}

func (c Product) Index() revel.Result {
	return c.Render()
}
