package entities

/*
* OrganizationInfo store some organizaions information
 */
type OrganizationInfo struct {
	OrganizationId   int `xorm:"'id' autoincr pk"`
	OrganizationName string
	Introduction     string
}

func (u OrganizationInfo) TableName() string {
	return "organization"
}
