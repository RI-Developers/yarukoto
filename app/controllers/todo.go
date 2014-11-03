package controllers

import (
	"github.com/Tsuguya/revmgo/app"
	//"yarukoto/app/models"
	"github.com/revel/revel"
)

type Todo struct {
	*revel.Controller
	mgo.MongoController
}

func (c Todo) List() revel.Result {
	return c.Render()
}

