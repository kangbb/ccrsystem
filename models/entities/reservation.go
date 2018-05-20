package entities

import (
	"time"
)

/*
* ReservationInfo store reservation information
 */
type ReservationInfo struct {
	ResId            int `xorm:"autoincr pk 'id'"`
	ResReason        string
	StartTime        time.Time
	EndTime          time.Time
	ClassroomId      int `xorm:"'classroom_id'"`
	StudentId        int
	OrganizationName string
	ApproverId       int
	ResState         int // 0, 未审批；1，审批中；2，审批通过；3，审批未通过
	ApprovalNote     string
}

func (u ReservationInfo) TableName() string {
	return "reservation"
}
