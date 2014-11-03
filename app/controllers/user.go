package controllers

import (
	"github.com/Tsuguya/revmgo/app"
	//"yarukoto/app/models"
	"github.com/revel/revel"
)

type User struct {
	*revel.Controller
	mgo.MongoController
}

func (c User) Login() revel.Result {
	return c.Render()
}

