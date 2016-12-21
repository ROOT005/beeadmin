package routers

import (
	"beeadmin/controllers"
	"beeadmin/models"
	"github.com/astaxie/beego"
)

func init() {
	//注册路由
	models.RegisterDB()
	/*设置路由器*/
	beego.Router("/", &controllers.LoginController{})
	beego.Router("/admin", &controllers.AdminController{})
	beego.Router("/admin/dowload", &controllers.AdminController{}, "get:Dowload")
	beego.Router("/admin/message", &controllers.AdminController{}, "get:Message")
	beego.AutoRouter(&controllers.LoginController{})
}
