package models

import (
	"github.com/astaxie/beego/orm"
)

//Role ,RoleID -> User, RoleID <-> Group
type Role struct {
	ID       int    `orm:"column(ID)"`
	RoleName string `orm:"column(RoleName)"`
	IsShow   int    `orm:"column(IsShow)"`
}

//TableName ,LWS_Role
func (r *Role) TableName() string {
	return TableName("role")
}

//GetList ,Get Role list
func (r *Role) GetList(page, pageSize int, filters ...interface{}) ([]Role, int64) {
	offset := (page - 1) * pageSize
	list := make([]Role, 0)
	query := orm.NewOrm().QueryTable(r.TableName())
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

//Add ,add Role
func (r *Role) Add(dev Role) (int64, error) {
	id, err := orm.NewOrm().Insert(dev)
	if err != nil {
		return 0, err
	}
	return id, nil
}

//GetByID ,Get Role by id
func (r *Role) GetByID(id int) (Role, error) {

	o := orm.NewOrm()
	dev := Role{ID: id}
	err := o.Read(&dev)

	if err != nil {
		return Role{}, err
	}
	return dev, err
}

//Update ,update data
func (r Role) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(r, fields...); err != nil {
		return err
	}
	return nil
}
