package service

import (
	"github.com/kangbb/class-reservation/core/models/entities"
)

type ApproverInfoService struct{}

var ApproverService = ApproverInfoService{}

//Create a new Approver
func (*ApproverInfoService) NewApprover(id int, pwd string, name string, departmentId int, perm bool) entities.ApproverInfo {
	approver := entities.ApproverInfo{
		ApproverId:   id,
		ApproverPwd:  pwd,
		ApproverName: name,
		DepartmentId: departmentId,
		Permission:   perm,
	}
	return approver
}

//Insert Approver information to the db
func (*ApproverInfoService) SaveAInfo(approver *entities.ApproverInfo) (bool, error) {
	_, err := entities.MasterEngine.InsertOne(approver)
	if err != nil {
		return false, err
	}
	return true, nil
}

//Get a Approver information by ID
func (*ApproverInfoService) FindInfoById(id int) *entities.ApproverInfo {
	approver := new(entities.ApproverInfo)
	_, err := entities.SlaveEngine.Id(id).Get(approver)
	if err != nil {
		panic(err)
	}
	return approver
}

//Get Approver information by Department Id
func (*ApproverInfoService) FindInfoByDepartmentId(departmentId int) []entities.ApproverInfo {
	approver := make([]entities.ApproverInfo, 0)
	err := entities.SlaveEngine.Where("department_id = ?", departmentId).Find(approver)
	if err != nil {
		panic(err)
	}
	return approver
}

//Find all Approver information
func (*ApproverInfoService) FindAllInfo() []entities.ApproverInfo {
	approver := make([]entities.ApproverInfo, 0)
	err := entities.SlaveEngine.Find(&approver)
	if err != nil {
		panic(err)
	}
	return approver
}

//Update Approver password information
func (*ApproverInfoService) UpdatePasswordInfo(id int, pwd string) (bool, error) {
	approver := new(entities.ApproverInfo)
	approver.ApproverPwd = pwd

	_, err := entities.MasterEngine.Id(id).Update(approver)
	if err != nil {
		panic(err)
	}
	return true, nil
}

//Update Approver department information
func (*ApproverInfoService) UpdateDepartmentInfo(id int, departmentId int) (bool, error) {
	approver := new(entities.ApproverInfo)
	approver.DepartmentId = departmentId

	_, err := entities.MasterEngine.Id(id).Update(approver)
	if err != nil {
		panic(err)
	}
	return true, nil
}

//Delete Approver information
func (*ApproverInfoService) DeleteInfo(id int) (bool, error) {
	approver := new(entities.ApproverInfo)
	_, err := entities.MasterEngine.Id(id).Delete(approver)
	if err != nil {
		panic(err)
	}
	return false, nil
}
