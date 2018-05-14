package services

import (
	"github.com/kangbb/ccrsystem/models/entities"
)

type DepartmentInfoService struct{}

var DepartmentService = DepartmentInfoService{}

/*
* Create a new deparment
 */
func (*DepartmentInfoService) NewDeparment(name string, introduction string, order int, note string) *entities.DepartmentInfo {
	department := &entities.DepartmentInfo{
		DepartmentName: name,
		Introduction:   introduction,
		Order:          order,
		Note:           note,
	}
	return department
}

/*
* Insert department information to the db
 */
func (*DepartmentInfoService) SaveAInfo(deparment *entities.DepartmentInfo) error {
	_, err := entities.MasterEngine.InsertOne(deparment)

	return err
}

/*
* Get a department information by ID
 */
func (*DepartmentInfoService) FindInfoById(id int) (*entities.DepartmentInfo, error) {
	department := new(entities.DepartmentInfo)
	_, err := entities.SlaveEngine.Id(id).Get(department)

	return department, err
}

/*
* Get a department information by note
 */
func (*DepartmentInfoService) FindInfoByNote(note string) (*entities.DepartmentInfo, error) {
	department := new(entities.DepartmentInfo)
	_, err := entities.SlaveEngine.Where("note = ?", note).Get(department)

	return department, err
}

/*
* Get a department information by order
 */
func (*DepartmentInfoService) FindInfoByOrder(order int) (*entities.DepartmentInfo, error) {
	department := new(entities.DepartmentInfo)
	_, err := entities.SlaveEngine.Where("order = ?", order).Get(department)

	return department, err
}

/*
* Find all department information
 */
func (*DepartmentInfoService) FindAllInfo() ([]entities.DepartmentInfo, error) {
	department := make([]entities.DepartmentInfo, 0)
	err := entities.SlaveEngine.Asc("order").Find(&department)

	return department, err
}

/*
* Update department information, for name, introduction, order, and note
 */
func (*DepartmentInfoService) UpdateInfo(id int, name string, introduction string, order int, note string) error {
	department := new(entities.DepartmentInfo)
	department.DepartmentId = id
	department.DepartmentName = name
	department.Introduction = introduction
	department.Order = order
	department.Note = note

	_, err := entities.MasterEngine.Id(id).Update(department)

	return err
}

/*
* Delete department information
 */
func (*DepartmentInfoService) DeleteInfo(id int) error {
	department := new(entities.DepartmentInfo)
	_, err := entities.MasterEngine.Id(id).Delete(department)

	return err
}
