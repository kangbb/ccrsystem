package entities

// StudentInfo store student information
type StudentInfo struct {
	StudentId   int `xorm:"pk 'id'"`
	StudentPwd  string
	StudentName string
	Permission  bool
}

//AdminInfo store admin information
type AdminInfo struct {
	AdminId    int `xorm:"pk 'id'"`
	AdminPwd   string
	AdminName  string
	Permission bool
}

//ApproverInfo store approver information
type ApproverInfo struct {
	ApproverId   int `xorm:"pk 'id'"`
	ApproverPwd  string
	ApproverName string
	Permission   bool
}

//ClassroomInfo store approver information
type ClassroomInfo struct {
	ClassroomId       int `xorm:"pk autoincr 'id'"`
	ClassroomCampus   string
	ClassroomBuilding string
	ClassroomNum      string
	Capicity          int
}

//ReservationInfo store reservation information
type ReservationInfo struct {
	ResId       int `xorm:"autoincr pk 'id'"`
	ResState    string
	StartTime   string
	EndTime     string
	ResReason   string
	StudentId   int
	ApproverId  int
	ClassroomId int `xorm:"'classroom_id'"`
}

//go xorm reflect regular
//define the table name as 'items'
func (u StudentInfo) TableName() string {
	return "student"
}
func (u AdminInfo) TableName() string {
	return "admin"
}
func (u ApproverInfo) TableName() string {
	return "approver"
}
func (u ClassroomInfo) TableName() string {
	return "classroom"
}
func (u ReservationInfo) TableName() string {
	return "reservation"
}
