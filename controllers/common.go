package controllers

import (
	"strings"

	"github.com/astaxie/beego"
)

const (
	msgOk         = 0
	msgError      = -1
	msgNoLogin    = -2
	timeLayoutStr = "2006-01-02 15:04:05"
)

//BaseController ,base controller
type BaseController struct {
	beego.Controller
	controllerName string
	actionName     string
	userID         int64
	userName       string
	pageSize       int
}

//Prepare ,前期准备
func (b *BaseController) Prepare() {
	b.pageSize = 20
	controllerName, actionName := b.GetControllerAndAction()
	b.controllerName = strings.ToLower(controllerName[0 : len(controllerName)-10])
	b.actionName = strings.ToLower(actionName)

	sessionUserID := b.GetSession("userID")

	sid, ok := sessionUserID.(int64)
	if ok {
		b.userID = sid
	} else {
		b.userID = 0
	}

	sessionUserName := b.GetSession("userName")
	sname, ok := sessionUserName.(string)
	if ok {
		b.userName = sname
	} else {
		b.userName = ""
	}

	b.Data["siteName"] = beego.AppConfig.String("site.name")
	b.Data["userName"] = b.userName
}

//是否POST提交
func (b *BaseController) isPost() bool {
	return b.Ctx.Request.Method == "POST"
}

//获取用户IP地址
func (b *BaseController) getClientIP() string {
	s := b.Ctx.Request.RemoteAddr
	l := strings.LastIndex(s, ":")
	return s[0:l]
}

//重定向
func (b *BaseController) redirect(url string) {
	b.Redirect(url, 302)
	b.StopRun()
}

//加载模板
func (b *BaseController) display(tpl ...string) {
	var tplname string
	if len(tpl) > 0 {
		tplname = strings.Join([]string{tpl[0], "html"}, ".")
	} else {
		tplname = b.controllerName + "/" + b.actionName + ".html"
	}
	b.Layout = "layout.html"
	b.TplName = tplname
}

//ajax返回
func (b *BaseController) ajaxMsg(msg interface{}, msgno int) {
	out := make(map[string]interface{})
	out["status"] = msgno
	out["message"] = msg
	b.Data["json"] = out
	b.ServeJSON()
	b.StopRun()
}

//ajax返回 列表
func (b *BaseController) ajaxList(msg interface{}, msgno int, count int64, data interface{}) {
	out := make(map[string]interface{})
	out["code"] = msgno
	out["msg"] = msg
	out["count"] = count
	out["data"] = data
	b.Data["json"] = out
	b.ServeJSON()
	b.StopRun()
}
