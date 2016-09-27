package controllers

import (
	"github.com/revel/revel"
	"rightpalm/rightpalm.serviceagent/facebook"
)

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
	oauthURL := facebook.GenerateUrl()

	return c.Render(isNotDisplayMenu, oauthURL)
}
