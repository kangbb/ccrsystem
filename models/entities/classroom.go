package entities

/*
* ClassroomInfo store approver information
 */
type ClassroomInfo struct {
	ClassroomId       int    `xorm:"pk autoincr 'id'"`
	ClassroomCampus   string `xorm:"unique(classroom)"`
	ClassroomBuilding string `xorm:"unique(classroom)"`
	ClassroomNum      string `xorm:"unique(classroom)"`
	Capacity          int
}

func (u ClassroomInfo) TableName() string {
	return "classroom"
}
