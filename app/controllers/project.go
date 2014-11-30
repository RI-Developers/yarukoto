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

    if c.Request.Method == "POST" {
        //b := models.FindProjectListByTeamId(c.Database, "54577d4ce4b0c733f78cb7a7")
        b := models.FindProjectListByTeamId(c.Database, c.Request.PostForm.Get("team_id"))
        Max := len(b) - 1
        return c.Render(b, Max)
    }
 
	return c.Render()
    //return c.RenderJson(b)
}

