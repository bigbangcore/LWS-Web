package controllers

import (
	"LWS_Web/models"
	"strconv"
	"strings"
	"time"

	"github.com/astaxie/beego"
)

//DeviceController is used to manage devices
type DeviceController struct {
	BaseController
}

//List ,devices list
func (obj *DeviceController) List() {
	obj.Data["pageTitle"] = "device list"
	if obj.userID == 0 {
		obj.redirect(beego.URLFor("LoginController.Login"))
	}
	obj.TplName = "device/list.html"
}

//Add device info
func (obj *DeviceController) Add() {
	obj.Data["pageTitle"] = "add device"
	obj.TplName = "device/add.html"
}

//Edit device info
func (obj *DeviceController) Edit() {
	obj.Data["pageTitle"] = "edit device info"
	id, _ := obj.GetInt64("id", 0)
	mo, err := models.DeviceGetByID(id)
	if err != nil {
		obj.Data["devInfo"] = ""
		return
	}
	obj.Data["deviceInfo"] = mo
	obj.TplName = "device/edit.html"
}

//ChooseKey is used to choose key for device
func (obj *DeviceController) ChooseKey() {
	obj.Data["pageTitle"] = "Choose Key"
	obj.Data["deviceId"], _ = obj.GetInt64("id")
	obj.TplName = "device/choosekey.html"
}

//SubmitKey is used to choose key for the device
func (obj *DeviceController) SubmitKey() {
	keyID, _ := obj.GetInt64("keyID")
	devID, _ := obj.GetInt64("id")
	dev, devErr := models.DeviceGetByID(devID)
	if devErr != nil {
		obj.ajaxMsg(devErr.Error(), msgError)
	}

	dev.KeyID = keyID
	dev.DeviceUpdateTime = time.Now().Unix()
	if err := dev.Update(); err != nil {
		obj.ajaxMsg(err.Error(), msgError)
	}
	obj.ajaxMsg("", msgOk)
}

//ChooseSoftware is used to choose software for device
func (obj *DeviceController) ChooseSoftware() {
	obj.Data["pageTitle"] = "Choose Software"
	obj.TplName = "device/choosesoftware.html"
}

//Table ,return device list
func (obj *DeviceController) Table() {
	//列表
	page, err := obj.GetInt("page")
	if err != nil {
		page = 1
	}
	limit, err := obj.GetInt("limit")
	if err != nil {
		limit = 30
	}

	deviceName := strings.TrimSpace(obj.GetString("deviceName"))
	startTime := strings.TrimSpace(obj.GetString("startTime"))
	endTime := strings.TrimSpace(obj.GetString("endTime"))

	obj.pageSize = limit
	//查询条件
	filters := make([]interface{}, 0)
	//filters = append(filters, "status", 1)

	if startTime != "" {
		st, _ := time.Parse(timeLayoutStr, startTime)
		filters = append(filters, "DeviceUpdateTime__gte", st.Unix())
	}

	if endTime != "" {
		et, _ := time.Parse(timeLayoutStr, endTime)
		filters = append(filters, "DeviceUpdateTime__lte", et.Unix())
	}

	if deviceName != "" {
		filters = append(filters, "DeviceName__icontains", deviceName)
	}

	result, count := models.DeviceGetList(page, obj.pageSize, filters...)

	list := make([]map[string]interface{}, len(result))
	for k, v := range result {
		row := make(map[string]interface{})
		row["ID"] = v.ID
		row["DeviceName"] = v.DeviceName
		row["DeviceID"] = v.DeviceID
		row["DeviceStatus"] = v.DeviceStatus
		row["DeviceDes"] = v.DeviceDes
		row["DeviceRegTime"] = beego.Date(time.Unix(v.DeviceRegTime, 0), "Y-m-d H:i:s")
		row["DeviceUpdateTime"] = beego.Date(time.Unix(v.DeviceUpdateTime, 0), "Y-m-d H:i:s")

		row["KeyID"] = v.KeyID
		moKey, _ := models.KeyGetByID(v.KeyID)
		if moKey != nil {
			row["PublicKey"] = moKey.Pubkey
		}

		row["VersionID"] = v.VersionID
		moVersion, _ := models.VerGetByID(v.VersionID)
		if moVersion != nil {
			row["VersionTitle"] = moVersion.VersionTitle
		}

		list[k] = row
	}
	obj.ajaxList("Success", msgOk, count, list)
}

//AjaxSave ,save info by ajax
func (obj *DeviceController) AjaxSave() {
	if obj.userID == 0 {
		obj.ajaxMsg("", msgNoLogin)
	}
	dev := new(models.Device)
	dev.DeviceID = strings.TrimSpace(obj.GetString("deviceID"))
	dev.DeviceName = strings.TrimSpace(obj.GetString("deviceName"))
	dev.DeviceDes = strings.TrimSpace(obj.GetString("deviceDes"))
	dev.DeviceStatus, _ = obj.GetInt("deviceStatus")
	dev.DeviceUpdateTime = time.Now().Unix()
	dev.KeyID, _ = obj.GetInt64("keyID")
	dev.VersionID, _ = obj.GetInt64("versionID")

	auth := strings.Split(obj.Ctx.GetCookie("auth"), "|")
	dev.UserID, _ = strconv.Atoi(auth[0])

	devID, _ := obj.GetInt64("id")
	if devID == 0 {
		//add
		dev.ID = 0
		dev.DeviceRegTime = time.Now().Unix()

		if _, err := models.DeviceAdd(dev); err != nil {
			obj.ajaxMsg(err.Error(), msgError)
		}
		obj.ajaxMsg("", msgOk)
	}

	//modify
	dev.ID = devID
	if err := dev.Update(); err != nil {
		obj.ajaxMsg(err.Error(), msgError)
	}
	obj.ajaxMsg("", msgOk)
}

//AjaxDel ,del info by ajax
func (obj *DeviceController) AjaxDel() {
	if obj.userID == 0 {
		obj.ajaxMsg("", msgNoLogin)
	}
	ID, _ := obj.GetInt64("idsStr")

	//check versions table, if it contains associated data, can't delete this device
	_, err := models.DeviceDelete(ID)
	if err != nil {
		obj.ajaxMsg(err.Error(), msgError)
	}
	obj.ajaxMsg("", msgOk)
}
