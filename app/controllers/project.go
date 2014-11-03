package controllers

import (
	"github.com/Tsuguya/revmgo/app"
	"yarukoto/app/models/team"
	"github.com/revel/revel"
    "fmt"
)

type Project struct {
	*revel.Controller
	mgo.MongoController
}

func (c Project) List() revel.Result {
    b := models.FindProjectListById(c.Database, "54577d4ce4b0c733f78cb7a7")
    fmt.Printf("%+v", b)
	return c.RenderJson(b)
}

