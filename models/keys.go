package models

import (
	"github.com/astaxie/beego/orm"
)

//DeviceKey ,DeviceID ->DeviceKeyID
type DeviceKey struct {
	ID         int64  `orm:"column(ID)"`
	Pubkey     string `orm:"column(Pubkey)"`
	Prikey     string `orm:"column(Prikey)"`
	CreateTime int64  `orm:"column(CreateTime)"`
	State      int    `orm:"column(State)"`
	Des        string `orm:"column(Des)"`
	UserID     int    `orm:"column(UserID);type(int)"`
}

//TableName ,LWS_DeviceKey
func (d *DeviceKey) TableName() string {
	return TableName("devicekey")
}

//KeyGetList ,Get key list
func KeyGetList(page, pageSize int, filters ...interface{}) ([]DeviceKey, int64) {
	offset := (page - 1) * pageSize
	list := make([]DeviceKey, 0)
	query := orm.NewOrm().QueryTable(TableName("devicekey"))
	if len(filters) > 0 {
		l := len(filters)
		for k := 0; k < l; k += 2 {
			if filters[k].(string) == "ID__in" {
				query = query.Exclude(filters[k].(string), filters[k+1])
			} else {
				query = query.Filter(filters[k].(string), filters[k+1])
			}
		}
	}
	total, _ := query.Count()
	query.OrderBy("-id").Limit(pageSize, offset).All(&list)
	return list, total
}

//KeyAdd ,add key
func KeyAdd(key *DeviceKey) (int64, error) {
	id, err := orm.NewOrm().Insert(key)
	if err != nil {
		return 0, err
	}
	return id, nil
}

//KeyInsertMulti ,insert multi
func KeyInsertMulti(arr []DeviceKey) (int64, error) {
	successNums, err := orm.NewOrm().InsertMulti(1, arr)
	if err != nil {
		return 0, err
	}
	return successNums, nil
}

//KeyGetByID ,Get key by id
func KeyGetByID(id int64) (*DeviceKey, error) {
	d := new(DeviceKey)
	err := orm.NewOrm().QueryTable(TableName("devicekey")).Filter("id", id).One(d)
	if err != nil {
		return nil, err
	}
	return d, nil
}

//Update ,update data
func (d DeviceKey) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(d, fields...); err != nil {
		return err
	}
	return nil
}

//KeyDelete is used to delete a key
func KeyDelete(id int64) (int64, error) {
	query := orm.NewOrm().QueryTable(TableName("devicekey"))
	return query.Filter("ID", id).Delete()
}
