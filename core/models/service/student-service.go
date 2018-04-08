import (
	"log"

	"github.com/kangbb/class-reservation/core/models/entities"
)

type StudentInfoService struct{}

var StudentServie = StudentInfoService{}

//Create a new student
func (*StudentService) NewStudent(id int, pwd string, name string, perm bool) entities.StudentInfo {
	student := entities.StudentInfo{
		StudentId:   id,
		StudentPwd:  pwd,
		StudentName: name,
		Permission:  perm,
	}
	return student
}

//Insert student information to the db
func (*StudentService) SaveAInfo(std *entities.StudentInfo) (bool, error) {
	_, err := entities.MasterEngine.InsertOne(std)
	if err != nil {
		return false, err
	}
	return true, nil
}

//Get a student information by ID
func (*StudentService) FindInfoById(id int) entities.StudentInfo {
	std := new(entities.StudentInfo)
	_, err := entities.SlaveEngine.Id(id).Get(std)
	if err != nil {
		log.Println(err)
	}
	return std
}

//Find all student information
func (*StudentService) FindAllInfo() []entities.StudentInfo {
	std := make([]entities.StudentInfo, 0)
	err := entities.SlaveEngine.Find(&std)
	if err != nil {
		panic(err)
	}
	return std
}

//Update student information
//Just for password
func (*StudentService) UpdateInfo(id int, pwd string) (bool, error) {
	std := new(entities.StudentInfo)
	std.StudentPwd = pwd

	_, err := entities.MasterEngine.Id(id).Update(std)
	if err != nil {
		panic(err)
		return false, err
	}
	return true, nil
}

//Delete student information
func (*StudentService) DeleteInfo(id int) (bool, error) {
	std := new(entities.StudentInfo)
	_, err := entities.MasterEngine.Id(id).Delete(std)
	if err != nil {
		panic(err)
		return false, err
	}
	return false, nil
}