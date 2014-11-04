package controllers

import (
	"github.com/Tsuguya/revmgo/app"
	"yarukoto/app/models/team"
	"github.com/revel/revel"
    //"net/http"
)

type Team struct {
	*revel.Controller
	mgo.MongoController
}

func (c Team) List() revel.Result {
    //c.Response.Status = http.StatusTeapot
    c.Response.ContentType = "application/json"
    b := models.FindTeamList(c.Database)
    Max := len(b) - 1
	return c.Render(b, Max)
}

