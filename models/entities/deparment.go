package entities

/*
* DepartmentInfo store approval department information
 */
type DepartmentInfo struct {
	DepartmentId   int `xorm:"'id' autoincr pk"`
	DepartmentName string
	Introduction   string
	Order          int
	Note           string
	//Note, note whether it is Initial approval department or final approval department
	//value: initial, middle, final
}

func (u DepartmentInfo) TableName() string {
	return "department"
}
