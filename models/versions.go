package models

import (
	"github.com/astaxie/beego/orm"
)

//Versions ,for device versions
type Versions struct {
	ID                int64   `orm:"column(ID)"`
	VersionTitle      string  `orm:"column(VersionTitle)"`
	VersionAddTime    int64   `orm:"column(VersionAddTime)"`
	VersionDes        string  `orm:"column(VersionDes)"`
	VersionURL        string  `orm:"column(VersionURL)"`
	VersionUpdateTime int64   `orm:"column(VersionUpdateTime)"`
	VersionSize       float64 `orm:"column(VersionSize)"`
}

//TableName ,LWS_Versions
func (d *Versions) TableName() string {
	return TableName("versions")
}

//VerGetList ,Get Versions list
func VerGetList(page, pageSize int, filters ...interface{}) ([]Versions, int64) {
	offset := (page - 1) * pageSize
	list := make([]Versions, 0)
	query := orm.NewOrm().QueryTable(TableName("versions"))
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

//VerAdd ,add Versions
func VerAdd(dev *Versions) (int64, error) {
	id, err := orm.NewOrm().Insert(dev)
	if err != nil {
		return 0, err
	}
	return id, nil
}

//VerGetByID ,Get Versions by id
func VerGetByID(id int64) (*Versions, error) {

	o := orm.NewOrm()
	dev := Versions{ID: id}
	err := o.Read(&dev)

	if err != nil {
		return nil, err
	}
	return &dev, err
}

//Update ,update data
func (d Versions) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(d, fields...); err != nil {
		return err
	}
	return nil
}

//VerDelete is used to delete a version
func VerDelete(id int64) (int64, error) {
	query := orm.NewOrm().QueryTable(TableName("versions"))
	return query.Filter("ID", id).Delete()
}
