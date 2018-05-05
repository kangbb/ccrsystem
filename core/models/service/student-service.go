package service

import (
	"github.com/kangbb/ccrsystem/core/models/entities"
)

type StudentInfoService struct{}

var StudentService = StudentInfoService{}

//Create a new student
func (*StudentInfoService) NewStudent(id int, pwd string, name string, perm bool) entities.StudentInfo {
	student := entities.StudentInfo{
		StudentId:   id,
		StudentPwd:  pwd,
		StudentName: name,
		Permission:  perm,
	}
	return student
}

//Insert student information to the db
func (*StudentInfoService) SaveAInfo(std *entities.StudentInfo) (bool, error) {
	_, err := entities.MasterEngine.InsertOne(std)
	if err != nil {
		return false, err
	}
	return true, nil
}

//Get a student information by ID
func (*StudentInfoService) FindInfoById(id int) *entities.StudentInfo {
	std := new(entities.StudentInfo)
	_, err := entities.SlaveEngine.Id(id).Get(std)
	if err != nil {
		panic(err)
	}
	return std
}

//Find all student information
func (*StudentInfoService) FindAllInfo() []entities.StudentInfo {
	std := make([]entities.StudentInfo, 0)
	err := entities.SlaveEngine.Find(&std)
	if err != nil {
		panic(err)
	}
	return std
}

//Update student information
//Just for password
func (*StudentInfoService) UpdateInfo(id int, pwd string) (bool, error) {
	std := new(entities.StudentInfo)
	std.StudentPwd = pwd

	_, err := entities.MasterEngine.Id(id).Update(std)
	if err != nil {
		panic(err)
	}
	return true, nil
}

//Delete student information
func (*StudentInfoService) DeleteInfo(id int) (bool, error) {
	std := new(entities.StudentInfo)
	_, err := entities.MasterEngine.Id(id).Delete(std)
	if err != nil {
		panic(err)
	}
	return false, nil
}
