package controllers

import (
	"github.com/Tsuguya/revmgo/app"
	"yarukoto/app/models/team"
	"github.com/revel/revel"
)

type Todo struct {
	*revel.Controller
	mgo.MongoController
}

func (c Todo) List() revel.Result {
    b := models.FindTodoListByTeamAndProjectId(c.Database, "54577d4ce4b0c733f78cb7a7", "t001")
	return c.RenderJson(b)
}

