package service

import (
	"github.com/kangbb/class-reservation/core/models/entities"
)

type DepartmentInfoService struct{}

var DepartmentService = DepartmentInfoService{}

//Create a new deparment
func (*DepartmentInfoService) NewDeparment(name string, order int, note string) entities.DepartmentInfo {
	department := entities.DepartmentInfo{
		DepartmentName: name,
		Order:          order,
		Note:           note,
	}
	return department
}

//Insert department information to the db
func (*DepartmentInfoService) SaveAInfo(deparment *entities.DepartmentInfo) (bool, error) {
	_, err := entities.MasterEngine.InsertOne(deparment)
	if err != nil {
		return false, err
	}
	return true, nil
}

//Get a department information by ID
func (*DepartmentInfoService) FindInfoById(id int) *entities.DepartmentInfo {
	department := new(entities.DepartmentInfo)
	_, err := entities.SlaveEngine.Id(id).Get(department)
	if err != nil {
		panic(err)
	}
	return department
}

//Get a department information by note
func (*DepartmentInfoService) FindInfoByNote(note string) *entities.DepartmentInfo {
	department := new(entities.DepartmentInfo)
	_, err := entities.SlaveEngine.Where("note = ?", note).Get(department)
	if err != nil {
		panic(err)
	}
	return department
}

//Get a department information by order
func (*DepartmentInfoService) FindInfoByOrder(order int) *entities.DepartmentInfo {
	department := new(entities.DepartmentInfo)
	_, err := entities.SlaveEngine.Where("order = ?", order).Get(department)
	if err != nil {
		panic(err)
	}
	return department
}

//Find all department information
func (*DepartmentInfoService) FindAllInfo() []entities.DepartmentInfo {
	department := make([]entities.DepartmentInfo, 0)
	err := entities.SlaveEngine.Asc("order").Find(&department)
	if err != nil {
		panic(err)
	}
	return department
}

//Update department information
func (*DepartmentInfoService) UpdateInfo(id int, name string, order int, note string) (bool, error) {
	department := new(entities.DepartmentInfo)
	department.DepartmentId = id
	department.DepartmentName = name
	department.Order = order
	department.Note = note

	_, err := entities.MasterEngine.Id(id).Update(department)
	if err != nil {
		panic(err)
	}
	return true, nil
}

//Delete department information
func (*DepartmentInfoService) DeleteInfo(id int) (bool, error) {
	department := new(entities.DepartmentInfo)
	_, err := entities.MasterEngine.Id(id).Delete(department)
	if err != nil {
		panic(err)
	}
	return false, nil
}
