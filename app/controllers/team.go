package controllers

import (
	"github.com/Tsuguya/revmgo/app"
	//"yarukoto/app/models"
	"github.com/revel/revel"
)

type Team struct {
	*revel.Controller
	mgo.MongoController
}

func (c Team) List() revel.Result {
	return c.Render()
}

