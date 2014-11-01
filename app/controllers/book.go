package controllers

import (
	"github.com/Tsuguya/revmgo/app"
	"yarukoto/app/models"
	"github.com/revel/revel"
)

type Book struct {
	*revel.Controller
	mgo.MongoController
}

func (c Book) Index() revel.Result {
    b := models.FindAll(c.Database)
	return c.Render(b)
}

func (c Book) View(id string) revel.Result {
    b := models.FindByTitle(c.Database, id)
	return c.Render(b)
}
