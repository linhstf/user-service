package controllers

import "github.com/revel/revel"

// Home struct is a struct of Home controller
type Home struct {
	BaseController
}

// Index is a first at all route
func (c Home) Index() revel.Result {
	return c.Render()
}

func (c Home) Login() revel.Result {
	isNotDisplayMenu := true
	return c.Render(isNotDisplayMenu)
}
