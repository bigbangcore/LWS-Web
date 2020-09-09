package models

//Group ,used for group manage
type Group struct {
	ID         int    `orm:"column(ID)"`
	PID        int    `orm:"column(PID)"`
	GroupName  string `orm:"column(GroupName)"`
	GroupType  string `orm:"column(GroupType)"`
	GroupDes   string `orm:"column(GroupDes)"`
	CreateTime int64  `orm:"column(CreateTime)"`
}

//GroupUser ,Relationship about Group to User
type GroupUser struct {
	GroupID int `orm:"column(GroupID)"`
	UserID  int `orm:"column(UserID)"`
}

//GroupRole ,Relationship about Group to Role
type GroupRole struct {
	GroupID int `orm:"column(GroupID)"`
	RoleID  int `orm:"column(RoleID)"`
}

//TableName ,LWS_Group
func (u *Group) TableName() string {
	return TableName("group")
}

//TableName ,LWS_Group_User
func (g *GroupUser) TableName() string {
	return TableName("groupUser")
}

//TableName ,LWS_GroupToRole
func (g *GroupRole) TableName() string {
	return TableName("groupRole")
}
