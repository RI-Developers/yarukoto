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
    c.Response.ContentType = "application/json; charset=utf8"

    succeeds := false

    id  := c.Request.PostForm.Get("id")
    pwd := c.Request.PostForm.Get("pwd")


    // authentication process here
    if id == "user" && pwd == "password" {
        succeeds = true
        accessToken := "this_is_access_token_sample"
        c.Render(succeeds, accessToken)
    }


	return c.Render(succeeds)
}

