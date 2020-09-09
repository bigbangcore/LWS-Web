package controllers

import (
	"LWS_Web/models"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/astaxie/beego"
)

//VersionsController is used to manage the device versions
type VersionsController struct {
	BaseController
}

//List ,return the versions list
func (obj *VersionsController) List() {
	if obj.userID == 0 {
		obj.redirect(beego.URLFor("LoginController.Login"))
	}
	obj.TplName = "versions/list.html"
}

//Add device info
func (obj *VersionsController) Add() {
	obj.Data["pageTitle"] = "add device"
	//lst, _ := models.DeviceGetList(1, 1000)
	//obj.Data["devLst"] = lst
	obj.TplName = "versions/add.html"
}

//Push software to devices
func (obj *VersionsController) Push() {
	obj.Data["pageTitle"] = "Push Software to Devices"
	obj.Data["versionId"], _ = obj.GetInt64("id")
	obj.TplName = "versions/push.html"
}

//PushTo ,used to push software to devices
func (obj *VersionsController) PushTo() {
	idsStr := strings.TrimSpace(obj.GetString("idsStr"))
	versionID, _ := obj.GetInt64("id")

	arr := strings.Split(idsStr, ",")
	failGetDevCount := 0
	failUpdateCount := 0
	for i := 0; i < len(arr)-1; i++ {
		id, _ := strconv.ParseInt(arr[i], 10, 64)
		dev, devErr := models.DeviceGetByID(id)
		if devErr != nil {
			failGetDevCount++
			continue
		}
		dev.VersionID = versionID
		dev.DeviceUpdateTime = time.Now().Unix()
		if err := dev.Update(); err != nil {
			failUpdateCount++
			continue
		}
	}
	if failGetDevCount > 0 || failUpdateCount > 0 {
		obj.ajaxMsg("GetDevice Failure:"+strconv.Itoa(failGetDevCount)+",Update Failure:"+strconv.Itoa(failUpdateCount), msgError)
	}
	obj.ajaxMsg("", msgOk)
}

//Edit device info
func (obj *VersionsController) Edit() {
	obj.Data["pageTitle"] = "edit device info"
	id, _ := obj.GetInt64("id", 0)
	mo, err := models.DeviceGetByID(id)
	if err != nil {
		obj.Data["devInfo"] = ""
		return
	}
	obj.Data["deviceInfo"] = mo
	obj.TplName = "versions/edit.html"
}

//Table ,return device list
func (obj *VersionsController) Table() {

	//列表
	page, err := obj.GetInt("page")
	if err != nil {
		page = 1
	}
	limit, err := obj.GetInt("limit")
	if err != nil {
		limit = 30
	}

	startTime := strings.TrimSpace(obj.GetString("startTime"))
	endTime := strings.TrimSpace(obj.GetString("endTime"))

	obj.pageSize = limit
	//查询条件
	filters := make([]interface{}, 0)
	//filters = append(filters, "status", 1)

	if startTime != "" {
		st, _ := time.Parse(timeLayoutStr, startTime)
		filters = append(filters, "VersionUpdateTime__gte", st.Unix())
	}

	if endTime != "" {
		et, _ := time.Parse(timeLayoutStr, endTime)
		filters = append(filters, "VersionUpdateTime__lte", et.Unix())
	}

	result, count := models.VerGetList(page, obj.pageSize, filters...)

	list := make([]map[string]interface{}, len(result))
	for k, v := range result {
		row := make(map[string]interface{})
		row["ID"] = v.ID
		row["VersionTitle"] = v.VersionTitle
		row["VersionSize"] = v.VersionSize
		row["VersionURL"] = v.VersionURL
		row["VersionDes"] = v.VersionDes
		row["VersionAddTime"] = beego.Date(time.Unix(v.VersionAddTime, 0), "Y-m-d H:i:s")
		row["VersionUpdateTime"] = beego.Date(time.Unix(v.VersionUpdateTime, 0), "Y-m-d H:i:s")
		list[k] = row
	}
	obj.ajaxList("Success", msgOk, count, list)
}

//UploadFile is used to upload files
func (obj *VersionsController) UploadFile() {
	file, head, err := obj.GetFile("file")
	if err != nil {
		obj.ajaxMsg("Get file failed.", msgError)
	}
	defer file.Close()

	ts := time.Now().Unix()
	filename := beego.Date(time.Unix(ts, 0), "YmdHis")

	sli := strings.Split(head.Filename, ".")
	strPath := "static/upload/" + filename + "." + strings.Join(sli[1:], ".")
	err = obj.SaveToFile("file", strPath)
	if err != nil {
		obj.ajaxMsg("Upload failed.", msgError)
	}
	obj.ajaxMsg(strPath+"|"+strconv.FormatInt(head.Size, 10), msgOk)
}

//DownloadFile is used to download file
func (obj *VersionsController) DownloadFile(str string) {
	fileName := filepath.Base(str)
	obj.Ctx.Output.Download(str, fileName)
}

//AjaxSave ,save info by ajax
func (obj *VersionsController) AjaxSave() {
	version := new(models.Versions)
	version.VersionTitle = strings.TrimSpace(obj.GetString("versionTitle"))
	version.VersionURL = strings.TrimSpace(obj.GetString("versionURL"))
	version.VersionSize, _ = obj.GetFloat("versionSize")
	version.VersionUpdateTime = time.Now().Unix()
	version.VersionDes = strings.TrimSpace(obj.GetString("versionDes"))
	versionID, _ := obj.GetInt64("id")

	if versionID == 0 {
		//add
		version.ID = 0
		version.VersionAddTime = time.Now().Unix()

		if _, err := models.VerAdd(version); err != nil {
			obj.ajaxMsg(err.Error(), msgError)
		}
		obj.ajaxMsg("", msgOk)
	}

	//modify
	version.ID = versionID
	if err := version.Update(); err != nil {
		obj.ajaxMsg(err.Error(), msgError)
	}
	obj.ajaxMsg("", msgOk)
}

//AjaxDel ,del info by ajax
func (obj *VersionsController) AjaxDel() {
	versionID, _ := obj.GetInt64("idsStr")
	_, err := models.VerDelete(versionID)
	if err != nil {
		obj.ajaxMsg(err.Error(), msgError)
	}
	obj.ajaxMsg("", msgOk)
}
