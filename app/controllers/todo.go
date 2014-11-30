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

    succeeds := false

    if c.Request.Method == "POST" {

        // Validation check here
        accessToken := "this_is_access_token_sample"
        if c.Request.PostForm.Get("access_token") == accessToken {
            succeeds = true
            b := models.FindTodoListByProjectId(c.Database, c.Request.PostForm.Get("project_id"))
            Max := len(b) - 1
            return c.Render(succeeds, b, Max)
        }
    }

    return c.Render(succeeds)
}

