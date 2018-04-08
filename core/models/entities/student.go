package entities

// StudentInfo store student information
type StudentInfo struct {
	StudentId   int `xorm:"pk 'id'"`
	StudentPwd  string
	StudentName string
	Permission  bool
}

//go xorm reflect regular
//define the table name as 'items'
func (u StudentInfo) TableName() string {
	return "student"
}
