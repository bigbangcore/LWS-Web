package controllers

import (
	"LWS_Web/libs"
	"LWS_Web/models"
	"strings"
	"time"
)

//RegisterController ,be used to registe a new user
type RegisterController struct {
	BaseController
}

//Register a new user
func (obj *RegisterController) Register() {
	obj.TplName = "register.html"
}

//AjaxReg is used to add a new user
func (obj *RegisterController) AjaxReg() {
	user := new(models.User)
	user.LoginName = strings.TrimSpace(obj.GetString("username"))
	user.Phone = strings.TrimSpace(obj.GetString("phone"))
	user.State = 1
	user.Email = strings.TrimSpace(obj.GetString("email"))

	user.CreateTime = time.Now().Unix()
	user.UpdateTime = time.Now().Unix()

	pwd := strings.TrimSpace(obj.GetString("password"))
	user.Pwd = libs.MD5(pwd)

	if _, err := models.UserAdd(user); err != nil {
		obj.ajaxMsg(err.Error(), msgError)
	}
	obj.ajaxMsg("", msgOk)
}
