package controllers

import (
	"LWS_Web/libs"
	"LWS_Web/models"
	"strconv"
	"strings"

	"github.com/astaxie/beego"
)

//LoginController , used for login
type LoginController struct {
	BaseController
}

//Login TODO:XSRF过滤
func (obj *LoginController) Login() {
	if obj.userID > 0 {
		obj.redirect(beego.URLFor("MainController.Index"))
	}
	obj.TplName = "login.html"
}

//LoginIn , user login
func (obj *LoginController) LoginIn() {
	user := new(models.User)
	user.LoginName = strings.TrimSpace(obj.GetString("username"))
	pwd := strings.TrimSpace(obj.GetString("password"))
	user.Pwd = libs.MD5(pwd)

	//check
	tmp, err := models.UserGetByName(user.LoginName)
	if err != nil || user.LoginName != tmp.LoginName {
		obj.ajaxMsg(err.Error(), msgError)
	}

	if tmp.Pwd != user.Pwd {
		obj.ajaxMsg("User name or password error.", msgError)
	}

	obj.SetSession("userID", tmp.ID)
	obj.SetSession("userName", tmp.LoginName)

	obj.Ctx.SetCookie("auth", strconv.FormatInt(tmp.ID, 10)+"|"+tmp.LoginName, 7*86400)
	obj.ajaxMsg("", msgOk)
}

//LoginOut is used for login out
func (obj *LoginController) LoginOut() {
	obj.Ctx.SetCookie("auth", "")
	obj.userID = 0
	obj.userName = ""
	obj.redirect(beego.URLFor("LoginController.Login"))
}

//NoAuth ,show that you have no auth
func (obj *LoginController) NoAuth() {
	obj.Ctx.WriteString("没有权限")
}
