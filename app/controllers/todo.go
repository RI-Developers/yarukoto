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

    if c.Request.Method == "POST" {
        b := models.FindTodoListByProjectId(c.Database, c.Request.PostForm.Get("project_id"))
        Max := len(b) - 1
        return c.Render(b, Max)
    }

    return c.Render()
}

