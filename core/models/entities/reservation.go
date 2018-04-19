package entities

import (
	"time"
)

//ReservationInfo store reservation information
type ReservationInfo struct {
	ResId        int `xorm:"autoincr pk 'id'"`
	ResState     string
	StartTime    time.Time
	EndTime      time.Time
	ResReason    string
	ApprovalNote string
	DepartmentId int
	StudentId    int
	ApproverId   int
	ClassroomId  int `xorm:"'classroom_id'"`
}

func (u ReservationInfo) TableName() string {
	return "reservation"
}
