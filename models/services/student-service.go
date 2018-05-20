package services

import (
	"github.com/kangbb/ccrsystem/models/entities"
)

type StudentInfoService struct{}

var StudentService = StudentInfoService{}

/*
* Create a new student information.
 */
func (*StudentInfoService) NewStudent(id int, pwd string, name string) *entities.StudentInfo {
	student := &entities.StudentInfo{
		StudentId:   id,
		StudentPwd:  pwd,
		StudentName: name,
	}
	return student
}

/*
* Insert student information to the db
 */
func (*StudentInfoService) SaveAInfo(std *entities.StudentInfo) error {
	_, err := entities.MasterEngine.InsertOne(std)

	return err
}

/*
* Get a student information by ID
 */
func (*StudentInfoService) FindInfoById(id int) (*entities.StudentInfo, error) {
	std := new(entities.StudentInfo)
	_, err := entities.SlaveEngine.Id(id).Get(std)

	return std, err
}

/*
* Find all student information
 */
func (*StudentInfoService) FindAllInfo() ([]entities.StudentInfo, error) {
	stds := make([]entities.StudentInfo, 0)
	err := entities.SlaveEngine.Find(&stds)

	return stds, err
}

/*
* Update student information, just for password
 */
func (*StudentInfoService) UpdateInfo(id int, pwd string) error {
	std := new(entities.StudentInfo)
	std.StudentPwd = pwd

	_, err := entities.MasterEngine.Id(id).Update(std)

	return err
}

/*
* Delete a piece of student information
 */
func (*StudentInfoService) DeleteInfo(id int) error {
	std := new(entities.StudentInfo)
	_, err := entities.MasterEngine.Id(id).Delete(std)

	return err
}
