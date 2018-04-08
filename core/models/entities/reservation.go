package entities

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

func (u ReservationInfo) TableName() string {
	return "reservation"
}
