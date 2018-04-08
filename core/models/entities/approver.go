package entities

//ApproverInfo store approver information
type ApproverInfo struct {
	ApproverId   int `xorm:"pk 'id'"`
	ApproverPwd  string
	ApproverName string
	Permission   bool
}

func (u ApproverInfo) TableName() string {
	return "approver"
}
