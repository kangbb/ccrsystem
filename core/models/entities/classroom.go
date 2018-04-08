package entities

//ClassroomInfo store approver information
type ClassroomInfo struct {
	ClassroomId       int `xorm:"pk autoincr 'id'"`
	ClassroomCampus   string
	ClassroomBuilding string
	ClassroomNum      string
	Capicity          int
}

func (u ClassroomInfo) TableName() string {
	return "classroom"
}
