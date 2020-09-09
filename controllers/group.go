package controllers

//GroupController is used to manage groups
type GroupController struct {
	BaseController
}

//List ,groups list
func (obj *GroupController) List() {
	obj.Data["pageTitle"] = "分组设置"
	obj.display()
}

//Add a group
func (obj *GroupController) Add() {
	obj.Data["pageTitle"] = "新增分组"
	obj.display()
}

//Edit a group
func (obj *GroupController) Edit() {
	obj.Data["pageTitle"] = "编辑分组"
	obj.display()
}

//Table ,return the groups list
func (obj *GroupController) Table() {
	//列表
}

//AjaxSave ,save group info by ajax
func (obj *GroupController) AjaxSave() {

}

//AjaxDel ,del group info by ajax
func (obj *GroupController) AjaxDel() {

}
