package controllers

import "github.com/astaxie/beego"

//MainController ,main controller
type MainController struct {
	BaseController
}

//Get ,get method
func (obj *MainController) Get() {
	obj.Data["pageTitle"] = "LWS management"

	if obj.userID == 0 {
		obj.redirect(beego.URLFor("LoginController.Login"))
	}

	obj.display()
	obj.TplName = "index.html" //如果用户不设置该参数，那么默认会去到模板目录的 Controller/<方法名>.tpl 查找
}

//Index ,web home
func (obj *MainController) Index() {
	obj.Data["pageTitle"] = "LWS management"
	if obj.userID == 0 {
		obj.redirect(beego.URLFor("LoginController.Login"))
	}
	obj.display()
	obj.TplName = "index.html"
}
