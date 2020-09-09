package routers

import (
	"LWS_Web/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/index", &controllers.MainController{}, "*:Index")
	beego.Router("/login", &controllers.LoginController{}, "*:Login")
	beego.Router("/login_out", &controllers.LoginController{}, "*:LoginOut")
	beego.Router("/register", &controllers.RegisterController{}, "*:Register")

	beego.AutoRouter(&controllers.SysmanageController{})
	beego.AutoRouter(&controllers.DeviceController{})
	beego.AutoRouter(&controllers.KeysController{})
	beego.AutoRouter(&controllers.VersionsController{})
	beego.AutoRouter(&controllers.RegisterController{})
	beego.AutoRouter(&controllers.LoginController{})
}
