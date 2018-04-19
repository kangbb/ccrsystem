package entities

//DepartmentInfo store approval department information
type DepartmentInfo struct {
	DepartmentId   int `xorm:"'id' autoincr pk"`
	DepartmentName string
	Order          int
	Note           string //note whether it is Initial approval department or final approval department
	//value: initial, middle, final
}

func (u DepartmentInfo) TableName() string {
	return "department"
}
