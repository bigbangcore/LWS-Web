package controllers

import "github.com/astaxie/beego"

//SysmanageController ,used for system manage
type SysmanageController struct {
	BaseController
}

//AlertSkin ,skin change
func (obj *SysmanageController) AlertSkin() {
	if obj.userID == 0 {
		obj.redirect(beego.URLFor("LoginController.Login"))
	}
	obj.TplName = "sysmanage/alertskin.html"
}
