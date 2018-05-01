package controllers

import (
	"github.com/astaxie/beego"
)

type UserController struct {
	beego.Controller
}


func (c *UserController)Login(){

}

func (c *UserController)Logout(){

}

func (c *UserController)Register(){

}

func (c *UserController)UserInfo(){

}

func (c *UserController)Home()  {

	c.Ctx.WriteString("home")
}


