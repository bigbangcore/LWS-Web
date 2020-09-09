package models

import (
	"github.com/astaxie/beego/orm"
)

//User it is used for user admin.
type User struct {
	ID         int64  `orm:"column(ID)"`         //id
	LoginName  string `orm:"column(LoginName)"`  //user name
	RealName   string `orm:"column(RealName)"`   //real Name
	Pwd        string `orm:"column(Pwd)"`        //pwd
	RoleIds    string `orm:"column(RoleIds)"`    //role id
	Phone      string `orm:"column(Phone)"`      //phone
	Email      string `orm:"column(Email)"`      //email
	Salt       string `orm:"column(Salt)"`       //salt
	LastLogin  int64  `orm:"column(LastLogin)"`  //last login
	LastIP     string `orm:"column(LastIP)"`     //last ip
	State      int    `orm:"column(State)"`      //status
	CreateID   int64  `orm:"column(CreateID)"`   //create id
	UpdateID   int64  `orm:"column(UpdateID)"`   //update id
	CreateTime int64  `orm:"column(CreateTime)"` //create time
	UpdateTime int64  `orm:"column(UpdateTime)"` //update time
	UserDes    string `orm:"column(UserDes)"`    //user describe
}

//TableName ,LWS_User
func (u *User) TableName() string {
	return TableName("user")
}

//UserAdd ,func of User
func UserAdd(user *User) (int64, error) {
	return orm.NewOrm().Insert(user)
}

//UserGetByName ,func of User
func UserGetByName(loginName string) (*User, error) {
	a := new(User)
	err := orm.NewOrm().QueryTable(TableName("user")).Filter("LoginName", loginName).One(a)
	if err != nil {
		return nil, err
	}
	return a, nil
}

//UserGetList ,func of User
func UserGetList(page, pageSize int, filters ...interface{}) ([]*User, int64) {
	offset := (page - 1) * pageSize
	list := make([]*User, 0)
	query := orm.NewOrm().QueryTable(TableName("user"))
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

//UserGetByID ,func of User
func UserGetByID(id int) (*User, error) {
	r := new(User)
	err := orm.NewOrm().QueryTable(TableName("user")).Filter("ID", id).One(r)
	if err != nil {
		return nil, err
	}
	return r, nil
}

//Update ,func of User
func (u User) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(u, fields...); err != nil {
		return err
	}
	return nil
}
