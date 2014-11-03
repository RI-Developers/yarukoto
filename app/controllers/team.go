package controllers

import (
	"github.com/Tsuguya/revmgo/app"
	"yarukoto/app/models/team"
	"github.com/revel/revel"
)

type Team struct {
	*revel.Controller
	mgo.MongoController
}

func (c Team) List() revel.Result {
    // set content type to json
    //c.Response.Status = http.StatusTeapot
    //c.Response.ContentType = "application/json"

    b := models.FindTeamList(c.Database)
	return c.RenderJson(b)
}

