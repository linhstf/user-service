package controllers

import (
	"github.com/revel/revel"
	"rightpalm/rightpalm.domain/dbcontext"
	"rightpalm/rightpalm.serviceagent/facebook"
)

type Product struct {
	BaseController
}

func (c Product) Index() revel.Result {
	return c.Render()
}

func (c Product) Synchonize() revel.Result {
	data := make(map[string]interface{}, 0)

	result := facebook.Synchonize(facebook.TestAccessToken)
	data["result"] = result

	return c.RenderJson(data)
}

func (c Product) Debug() revel.Result {
	data := make(map[string]interface{}, 0)
	db := new(dbcontext.DbContext)
	result := db.Connect()
	data["result"] = result

	return c.RenderJson(data)
}
