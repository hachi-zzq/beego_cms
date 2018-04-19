package backend

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

type AdminController struct {
	beego.Controller
}


func (c *AdminController)PostLogin(){

}


func (c *AdminController)Logout(){

}

func (c *AdminController)Index(){

	logs.SetLogger("console")
	l := logs.GetLogger()
	l.Println(c.Ctx.Input.GetData("name"))


}

func (c *AdminController)Create(){

}

func (c *AdminController)Detail(){

}

func (c *AdminController)Update(){

}

func (c *AdminController)Delete(){

}
