package controllers

import (
	"github.com/Tsuguya/revmgo/app"
	//"yarukoto/app/models"
	"github.com/revel/revel"
)

type Project struct {
	*revel.Controller
	mgo.MongoController
}

func (c Project) List() revel.Result {
	return c.Render()
}

