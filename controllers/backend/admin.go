package backend

import (
	"github.com/astaxie/beego/orm"
	"cms/filters"
	"cms/models"
)

type AdminController struct {
	BaseController
}

func (c *AdminController)GetLogin(){

	c.TplName = "admin/login.html"
}


func (c *AdminController)PostLogin(){
	name := c.GetString("name")
	password := c.GetString("password")

	o := orm.NewOrm()

	admin := models.Admin{}

	err := o.QueryTable(new(models.Admin)).Filter("name",name).Filter("password",password).One(&admin)

	if err == orm.ErrNoRows{
		//没有
		c.Ctx.WriteString("登录失败，请检查用户名")
		return
	}else {
		filters.LoginAdmin(c.Ctx,int(admin.Id))
		c.Redirect("/admin/users",302)
	}

}


func (c *AdminController)Logout(){
	filters.LogoutAdmin(c.Ctx)
	c.Redirect("/admin/users",302)
}

func (c *AdminController)Index(){
	c.Ctx.WriteString("admin home")
}

func (c *AdminController)Home(){
	c.Ctx.WriteString("admin home")
}

func (c *AdminController)Create(){

}

func (c *AdminController)Detail(){

}

func (c *AdminController)Update(){

}

func (c *AdminController)Delete(){

}

func(c *AdminController)Me(){
	c.TplName = "admin/me.html"
}

func(c *AdminController)UpdateMe(){

	admin := c.Admin
	admin.Name = c.GetString("name")

	o := orm.NewOrm()
	if _ , err :=o.Update(admin); err == nil{
		c.Redirect("/admin/me",302)
	}else {
		c.Ctx.WriteString("更新失败"+err.Error())
	}





}

