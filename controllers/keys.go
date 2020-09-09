package controllers

import (
	"LWS_Web/libs"
	"LWS_Web/models"
	"strconv"
	"strings"
	"time"

	"github.com/astaxie/beego"
)

//KeysController is used to manage keys
type KeysController struct {
	BaseController
}

//List ,keyices list
func (obj *KeysController) List() {
	obj.Data["pageTitle"] = "keys list"
	if obj.userID == 0 {
		obj.redirect(beego.URLFor("LoginController.Login"))
	}
	obj.TplName = "keys/list.html"
}

//Add keyice info
func (obj *KeysController) Add() {
	obj.Data["pageTitle"] = "add key"
	obj.TplName = "keys/add.html"
}

//Create keyice info
func (obj *KeysController) Create() {
	obj.Data["pageTitle"] = "create key"
	obj.TplName = "keys/create.html"
}

//Edit keyice info
func (obj *KeysController) Edit() {
	obj.Data["pageTitle"] = "Edit keys info"
	id, _ := obj.GetInt64("id", 0)
	mo, err := models.KeyGetByID(id)
	if err != nil {
		obj.Data["KeysInfo"] = ""
		return
	}
	obj.Data["KeysInfo"] = mo
	obj.TplName = "keys/edit.html"
}

//Table ,return keyice list
func (obj *KeysController) Table() {
	IsUsedForDeviceLst, _ := obj.GetBool("isUsedForDeviceLst")
	filters := make([]interface{}, 0)
	//filters = append(filters, "status", 1)
	if IsUsedForDeviceLst {
		keyLst, err := models.DeviceGetAllSelectedKeyID()

		if err == nil {
			filters = append(filters, "ID__in", keyLst)
		}

	}
	page, err := obj.GetInt("page")
	if err != nil {
		page = 1
	}
	limit, err := obj.GetInt("limit")
	if err != nil {
		limit = 30
	}

	pubKey := strings.TrimSpace(obj.GetString("pubKey"))
	priKey := strings.TrimSpace(obj.GetString("priKey"))

	obj.pageSize = limit

	if pubKey != "" {
		filters = append(filters, "pubKey__icontains", pubKey)
	}

	if priKey != "" {
		filters = append(filters, "priKey__icontains", priKey)
	}

	result, count := models.KeyGetList(page, obj.pageSize, filters...)

	list := make([]map[string]interface{}, len(result))
	for k, v := range result {
		row := make(map[string]interface{})
		row["ID"] = v.ID
		row["Pubkey"] = v.Pubkey
		row["Prikey"] = v.Prikey
		row["State"] = v.State
		row["Des"] = v.Des
		row["CreateTime"] = beego.Date(time.Unix(v.CreateTime, 0), "Y-m-d H:i:s")
		list[k] = row
	}
	obj.ajaxList("Success", msgOk, count, list)
}

//AjaxSave ,save info by ajax
func (obj *KeysController) AjaxSave() {
	if obj.userID == 0 {
		obj.redirect(beego.URLFor("LoginController.Login"))
	}
	key := new(models.DeviceKey)
	key.Pubkey = strings.TrimSpace(obj.GetString("pubkey"))
	key.Prikey = strings.TrimSpace(obj.GetString("prikey"))
	key.Des = strings.TrimSpace(obj.GetString("des"))
	key.State, _ = obj.GetInt("state")

	auth := strings.Split(obj.Ctx.GetCookie("auth"), "|")
	key.UserID, _ = strconv.Atoi(auth[0])

	keyID, _ := obj.GetInt64("id")
	if keyID == 0 {
		//add
		key.ID = 0
		key.CreateTime = time.Now().Unix()

		if _, err := models.KeyAdd(key); err != nil {
			obj.ajaxMsg(err.Error(), msgError)
		}
		obj.ajaxMsg("", msgOk)
	}

	//modify
	key.ID = keyID
	if err := key.Update(); err != nil {
		obj.ajaxMsg(err.Error(), msgError)
	}
	obj.ajaxMsg("", msgOk)
}

//AjaxCreateKey , batch creation keys
func (obj *KeysController) AjaxCreateKey() {
	if obj.userID == 0 {
		obj.redirect(beego.URLFor("LoginController.Login"))
	}
	auth := strings.Split(obj.Ctx.GetCookie("auth"), "|")
	userID, _ := strconv.Atoi(auth[0])
	keyCount, _ := obj.GetInt("keyCount")
	arr := make([]libs.KeyPair, 0, 10)
	for i := 0; i < keyCount; i++ {
		keyPair := libs.MakeKeyPair()
		arr = append(arr, keyPair)
	}
	successNums, err := InsertMulti(arr, userID)
	if err != nil {
		obj.ajaxMsg(err.Error(), msgError)
	}
	obj.ajaxMsg(successNums, msgOk)
}

//InsertMulti ,batch insert keys
func InsertMulti(arr []libs.KeyPair, userID int) (int64, error) {
	keyArr := make([]models.DeviceKey, 0)
	createTime := time.Now().Unix()
	for i := 0; i < len(arr); i++ {
		mo := models.DeviceKey{
			ID:         0,
			Pubkey:     arr[i].PublicKey,
			Prikey:     arr[i].PrivateKey,
			CreateTime: createTime,
			State:      1,
			Des:        "create",
			UserID:     userID,
		}
		keyArr = append(keyArr, mo)
	}
	successNums, err := models.KeyInsertMulti(keyArr)
	return successNums, err
}

//AjaxDel ,del info by ajax
func (obj *KeysController) AjaxDel() {
	if obj.userID == 0 {
		obj.redirect(beego.URLFor("LoginController.Login"))
	}
	ID, _ := obj.GetInt64("idsStr")

	//check device table, if it contains associated data, can't delete this key
	re := models.DeviceExist(ID)
	if re {
		obj.ajaxMsg(false, msgError)
	}

	_, err := models.KeyDelete(ID)
	if err != nil {
		obj.ajaxMsg(err.Error(), msgError)
	}
	obj.ajaxMsg("", msgOk)
}
