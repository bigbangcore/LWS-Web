package models

//Auth ,such as add,del,modify,search,view
type Auth struct {
	ID       int    `orm:"column(ID)"`
	AuthName string `orm:"column(AuthName)"`
	AuthURL  string `orm:"column(AuthURL)"`
	Sort     int    `orm:"column(Sort)"`
	Icon     string `orm:"column(Icon)"`
	IsShow   int    `orm:"column(IsShow)"`
	State    int    `orm:"column(State)"`
}

//TableName ,LWS_Auth
func (a *Auth) TableName() string {
	return TableName("auth")
}
