package controllers

import (
	"github.com/Tsuguya/revmgo/app"
	"yarukoto/app/models/project"
	"github.com/revel/revel"
)

type Todo struct {
	*revel.Controller
	mgo.MongoController
}

func (c Todo) List() revel.Result {
    c.Response.ContentType = "application/json; charset=utf8"
    b := models.FindTodoListByProjectId(c.Database, "5467157be4b0c9468004aef6")
    Max := len(b) - 1
	return c.Render(b, Max)
}

