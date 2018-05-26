package entities

/*
* DepartmentInfo store approval department information
 */
type DepartmentInfo struct {
	DepartmentId    int    `xorm:"'id' autoincr pk"`
	DepartmentName  string `xorm:"unique"`
	Introduction    string
	DepartmentOrder int //不能命名为order,它是mysql的关键字
	Note            string
	//Note, note whether it is Initial approval department or final approval department
	//value: initial, middle, final
}

func (u DepartmentInfo) TableName() string {
	return "department"
}
