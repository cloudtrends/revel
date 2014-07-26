package controllers

import "github.com/cloudtrends/revel"

type Application struct {
	*revel.Controller
}

func (c Application) Index() revel.Result {
	return c.Render()
}
