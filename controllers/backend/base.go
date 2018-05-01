package backend

import (
	"github.com/astaxie/beego"
	"cms/models"
)

type BaseController struct {
	beego.Controller
	Admin *models.Admin
}

func (c *BaseController) Prepare() {
	adminId := c.GetSession("admin_id")
	if adminId != nil {
		admin, _ := models.GetAdminById(int64(adminId.(int)))
		if admin == nil {
			c.Ctx.WriteString("管理员不存在")
		}
		c.Data["current"] = admin
		c.Admin = admin
		c.Layout = "admin/layout.html"
	}
}
