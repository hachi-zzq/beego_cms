package routers

import (
	"github.com/astaxie/beego"
	"cms/controllers/backend"
	"cms/controllers"
	"cms/filters"
)

func init() {

	backendNs := beego.NewNamespace("admin",
		//beego.NSBefore(backend.AdminController.Auth),
		beego.NSBefore(filters.AdminAuth),
		beego.NSRouter("/", &backend.AdminController{}, "get:Home"),
		//登录
		beego.NSRouter("login", &backend.AdminController{}, "get:GetLogin;post:PostLogin"),
		//登出
		beego.NSRouter("logout", &backend.AdminController{}, "get:Logout"),
		beego.NSRouter("me", &backend.AdminController{}, "get:Me;post:UpdateMe"),
		//管理员
		beego.NSNamespace("admins",
			beego.NSRouter("/", &backend.AdminController{}, "get:Index;post:Create"),
			beego.NSRouter("/:id", &backend.AdminController{}, "get:Detail;patch:Update;delete:Delete"),
		),
		//用户
		beego.NSNamespace("users",
			beego.NSRouter("/", &backend.UserController{}, "get:Index;post:PostCreate"),
			beego.NSRouter("/create", &backend.UserController{}, "get:Create"),
			beego.NSRouter("/:id", &backend.UserController{}, "get:Detail;patch:Update;delete:Delete"),
			beego.NSRouter("/:id/delete", &backend.UserController{}, "get:Delete"),
		),
		//帖子
		beego.NSNamespace("articles",
			beego.NSRouter("/", &backend.ArticleController{}, "get:Index"),
			beego.NSRouter("/:id", &backend.ArticleController{}, "get:Detail;delete:Delete"),
			beego.NSRouter("/:id/approve", &backend.ArticleController{}, "patch:Approve"),
		),

		//评论
		beego.NSNamespace("comments",
			beego.NSRouter("/", &backend.CommentController{}, "get:Index"),
			beego.NSRouter("/:id", &backend.CommentController{}, "delete:Delete"),
			beego.NSRouter("/:id/approve", &backend.CommentController{}, "patch:Approve"),
		),
	)

	beego.Router("/", &controllers.UserController{}, "get:Home")
	beego.Router("login", &controllers.UserController{}, "post:Login")
	beego.Router("logout", &controllers.UserController{}, "post:Logout")
	beego.Router("register", &controllers.UserController{}, "post:Register")
	beego.Router("me", &controllers.UserController{}, "get:UserInfo")

	frontNs := beego.NewNamespace("articles",
		beego.NSRouter("/", &controllers.ArticleController{}, "get:Index"),
		beego.NSRouter("/:id", &controllers.ArticleController{}, "get:Detail;delete:Delete;patch:Update"),
		beego.NSRouter("/:id/comment", &controllers.ArticleController{}, "post:AddComment"),
		beego.NSRouter("/:id/comment/:comment_id", &controllers.ArticleController{}, "delete:DeleteComment"),
	)

	beego.AddNamespace(backendNs)
	beego.AddNamespace(frontNs)
}
