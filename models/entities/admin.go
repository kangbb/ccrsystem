package entities

/*
* AdminInfo store admin information
*/
type AdminInfo struct {
	AdminId   int `xorm:"pk 'id'"`
	AdminPwd  string
	AdminName string
}

func (u AdminInfo) TableName() string {
	return "admin"
}
