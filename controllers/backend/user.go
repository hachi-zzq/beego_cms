package backend

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"cms/models"
	"strconv"
)

type UserController struct {
	beego.Controller
}


func (c *UserController)Login(){

}

func (c *UserController)Logout(){

}

func (c *UserController)Index(){
	o := orm.NewOrm()

	var users []models.User

	_ ,err := o.QueryTable("users").All(&users)

	if err != nil{
		c.Ctx.WriteString(err.Error())
		return
	}

	c.Data["users"] = users

	c.TplName = "admin/users.html"
}

func (c *UserController)Detail(){

	userId := c.Ctx.Input.Param(":id")

	intUserId ,_ := strconv.ParseInt(userId,10,64)
	user ,err :=models.GetUserById(intUserId)

	if err != nil{
		c.Abort("404")
	}

	c.Data["user"] = user

	c.TplName = "admin/user_detail.html"
}

func (c *UserController)Create(){
	c.TplName = "admin/user_create.html"

}

func (c *UserController)PostCreate(){
	name := c.GetString("name")
	pass := c.GetString("password")
	o := orm.NewOrm()
	_,err := o.Insert(&models.User{
		Name:name,
		Password:pass,
	})

	if err == nil{
		c.Ctx.Redirect(302,"/admin/users")
	}else {
		c.Ctx.WriteString(err.Error())
	}
}

func (c *UserController)Update(){

}

