package models

import (
	"github.com/astaxie/beego/orm"
)

//Device ,used for device manage.为类型增加方法有一个限制，就是方法的定义必须和类型的定义在同一个包中
type Device struct {
	DeviceID         string `orm:"column(DeviceID)"`
	DeviceName       string `orm:"column(DeviceName)"`
	DeviceDes        string `orm:"column(DeviceDes)"`
	DeviceStatus     int    `orm:"column(DeviceStatus);type(int)"`
	DeviceUpdateTime int64  `orm:"column(DeviceUpdateTime);type(bigint)"`
	ID               int64  `orm:"column(ID)"`
	DeviceRegTime    int64  `orm:"column(DeviceRegTime);type(bigint)"`
	KeyID            int64  `orm:"column(KeyID);type(bigint)"`
	VersionID        int64  `orm:"column(VersionID);type(bigint)"`
	UserID           int    `orm:"column(UserID);type(int)"`
}

//TableName ,LWS_Device
func (d *Device) TableName() string {
	return TableName("device")
}

//DeviceGetList ,Get device list
func DeviceGetList(page, pageSize int, filters ...interface{}) ([]Device, int64) {
	offset := (page - 1) * pageSize
	list := make([]Device, 0)
	query := orm.NewOrm().QueryTable(TableName("device"))
	if len(filters) > 0 {
		l := len(filters)
		for k := 0; k < l; k += 2 {
			query = query.Filter(filters[k].(string), filters[k+1])
		}
	}
	total, _ := query.Count()
	query.OrderBy("-id").Limit(pageSize, offset).All(&list)
	return list, total
}

//DeviceExist ,check a data Exist or not
func DeviceExist(ID int64) bool {
	query := orm.NewOrm().QueryTable(TableName("device"))
	query = query.Filter("KeyID", ID)
	return query.Exist()
}

//DeviceAdd ,add device
func DeviceAdd(dev *Device) (int64, error) {
	id, err := orm.NewOrm().Insert(dev)
	if err != nil {
		return 0, err
	}
	return id, nil
}

//DeviceGetByID ,Get device by id
func DeviceGetByID(id int64) (*Device, error) {
	d := new(Device)
	err := orm.NewOrm().QueryTable(TableName("device")).Filter("id", id).One(d)
	if err != nil {
		return nil, err
	}
	return d, nil
}

//DeviceGetAllSelectedKeyID is used to get keyID lst what had been used in device
func DeviceGetAllSelectedKeyID() ([]int64, error) {
	var list orm.ParamsList
	_, err := orm.NewOrm().QueryTable(TableName("device")).Distinct().ValuesFlat(&list, "KeyID")
	if err != nil {
		return nil, err
	}
	sil := make([]int64, 0)
	for m := range list {
		re, ok := list[m].(int64)
		if ok && re != 0 {
			sil = append(sil, re)
		}
	}
	return sil, nil
}

//Update ,update data
func (d Device) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(&d, fields...); err != nil {
		return err
	}
	return nil
}

//DeviceDelete is used to delete a device
func DeviceDelete(id int64) (int64, error) {
	query := orm.NewOrm().QueryTable(TableName("device"))
	return query.Filter("ID", id).Delete()
}
