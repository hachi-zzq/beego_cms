package filters

import (
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego"
	"strconv"
)

const AdminUserId = "admin_id"

const UserId = "user_id"

func UserAuth(c *context.Context) {

	userId := beego.AppConfig.String("debug_user_id")

	if intUserId,err := strconv.Atoi(userId); err==nil && intUserId != 0{
		LoginUser(c,intUserId)
		return
	}

	requestUrl := c.Input.URL()
	passUrls := []string{"/login"}

	for _, value := range passUrls {
		if value == requestUrl {
			return
		}
	}

	if !userIsLogin(c) {
		c.Redirect(302, "/login")
	}
}

func userIsLogin(c *context.Context) bool {

	return c.Input.CruSession.Get(UserId) != nil
}

func LoginUser(c *context.Context, userId int) bool {

	c.Input.CruSession.Set(UserId, userId)
	return true
}

func AdminAuth(c *context.Context) {

	//debug
	adminId := beego.AppConfig.String("debug_admin_id")
	intAdminId, err := strconv.Atoi(adminId)

	if err == nil && intAdminId != 0 {
		LoginAdmin(c, intAdminId)
		return
	}

	requestUrl := c.Input.URL()

	passUrl := []string{"/admin/login"}

	for _, value := range passUrl {

		if value == requestUrl {
			return
		}
	}

	if !adminIsLogin(c) {
		c.Redirect(302, "/admin/login")
	}
}

func adminIsLogin(c *context.Context) bool {

	return c.Input.CruSession.Get(AdminUserId) != nil
}

func LoginAdmin(c *context.Context, adminId int) bool {

	c.Input.CruSession.Set(AdminUserId, adminId)
	return true
}
