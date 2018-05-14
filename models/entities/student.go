package entities

/*
* StudentInfo store student information
 */
type StudentInfo struct {
	StudentId      int `xorm:"pk 'id'"`
	StudentPwd     string
	StudentName    string
	OrganizationId int
}

/*
* go xorm reflect regular
* define the table name as 'student'
 */
func (u StudentInfo) TableName() string {
	return "student"
}
