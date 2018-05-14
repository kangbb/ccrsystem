package services

import (
	"github.com/kangbb/ccrsystem/models/entities"
)

type ApproverInfoService struct{}

var ApproverService = ApproverInfoService{}

/*
* Create a new Approver
 */
func (*ApproverInfoService) NewApprover(id int, pwd string, name string, departmentId int) *entities.ApproverInfo {
	approver := &entities.ApproverInfo{
		ApproverId:   id,
		ApproverPwd:  pwd,
		ApproverName: name,
		DepartmentId: departmentId,
	}
	return approver
}

/*
* Insert Approver information to the db
 */
func (*ApproverInfoService) SaveAInfo(approver *entities.ApproverInfo) error {
	_, err := entities.MasterEngine.InsertOne(approver)

	return err
}

/*
* Get a Approver information by ID
 */
func (*ApproverInfoService) FindInfoById(id int) (*entities.ApproverInfo, error) {
	approver := new(entities.ApproverInfo)
	_, err := entities.SlaveEngine.Id(id).Get(approver)

	return approver, err
}

/*
* Get Approver information by Department Id
 */
func (*ApproverInfoService) FindInfoByDepartmentId(departmentId int) ([]entities.ApproverInfo, error) {
	approver := make([]entities.ApproverInfo, 0)
	err := entities.SlaveEngine.Where("department_id = ?", departmentId).Find(approver)

	return approver, err
}

/*
* Find all Approver information
 */
func (*ApproverInfoService) FindAllInfo() ([]entities.ApproverInfo, error) {
	approver := make([]entities.ApproverInfo, 0)
	err := entities.SlaveEngine.Find(&approver)

	return approver, err
}

/*
* Update Approver password information
 */
func (*ApproverInfoService) UpdatePasswordInfo(id int, pwd string) error {
	approver := new(entities.ApproverInfo)
	approver.ApproverPwd = pwd

	_, err := entities.MasterEngine.Id(id).Update(approver)

	return err
}

/*
* Update Approver department information
 */
func (*ApproverInfoService) UpdateDepartmentInfo(id int, departmentId int) error {
	approver := new(entities.ApproverInfo)
	approver.DepartmentId = departmentId

	_, err := entities.MasterEngine.Id(id).Update(approver)

	return err
}

/*
* Delete Approver information
 */
func (*ApproverInfoService) DeleteInfo(id int) error {
	approver := new(entities.ApproverInfo)
	_, err := entities.MasterEngine.Id(id).Delete(approver)

	return err
}
