package controllers

import (
	"github.com/Tsuguya/revmgo/app"
	"yarukoto/app/models/team"
	"github.com/revel/revel"
)

type Project struct {
	*revel.Controller
	mgo.MongoController
}

func (c Project) List() revel.Result {
    c.Response.ContentType = "application/json; charset=utf8"
    b := models.FindProjectListById(c.Database, "54577d4ce4b0c733f78cb7a7")
    Max := len(b) - 1
	return c.Render(b, Max)
}

