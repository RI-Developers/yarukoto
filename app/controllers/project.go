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

    succeeds := false

    if c.Request.Method == "POST" {
        // Validation check here
        accessToken := "this_is_access_token_sample"
        if c.Request.PostForm.Get("access_token") == accessToken {
            succeeds = true
            b := models.FindProjectListByTeamId(c.Database, c.Request.PostForm.Get("team_id"))
            Max := len(b) - 1
            return c.Render(succeeds, b, Max)
        }
    }

	return c.Render(succeeds)
}

