package services

import (
	"github.com/kangbb/ccrsystem/models/entities"
)

type ClassroomInfoService struct{}

var ClassroomService = ClassroomInfoService{}

//Create a new Classroom
func (*ClassroomInfoService) NewClassroom(campuse string, building string, num string, cap int) *entities.ClassroomInfo {
	classroom := &entities.ClassroomInfo{
		ClassroomCampus:   campuse,
		ClassroomBuilding: building,
		ClassroomNum:      num,
		Capacity:          cap,
	}
	return classroom
}

//Create a slice of Classroom
func (*ClassroomInfoService) NewClassroomSlice() []entities.ClassroomInfo {
	return make([]entities.ClassroomInfo, 0)
}

//Insert classroom information to the db
func (*ClassroomInfoService) SaveAInfo(classroom *entities.ClassroomInfo) error {
	_, err := entities.MasterEngine.InsertOne(classroom)

	return err
}

//Get a classroom information by ID
func (*ClassroomInfoService) FindInfoById(id int) (*entities.ClassroomInfo, error) {
	classroom := new(entities.ClassroomInfo)
	_, err := entities.SlaveEngine.Id(id).Get(classroom)

	return classroom, err
}

//Find all classroom information
func (*ClassroomInfoService) FindAllInfo() ([]entities.ClassroomInfo, error) {
	classroom := make([]entities.ClassroomInfo, 0)
	err := entities.SlaveEngine.Find(&classroom)

	return classroom, err
}

//Update classroom information
func (*ClassroomInfoService) UpdateInfo(id int, arg ...interface{}) error {
	classroom := new(entities.ClassroomInfo)
	switch len(arg) {
	case 4:
		classroom.Capacity = arg[3].(int)
		classroom.ClassroomNum = arg[2].(string)
		classroom.ClassroomBuilding = arg[1].(string)
		classroom.ClassroomCampus = arg[0].(string)
	case 3:
		classroom.ClassroomNum = arg[2].(string)
		classroom.ClassroomBuilding = arg[1].(string)
		classroom.ClassroomCampus = arg[0].(string)
	case 2:
		classroom.ClassroomBuilding = arg[1].(string)
		classroom.ClassroomCampus = arg[0].(string)
	case 1:
		classroom.ClassroomCampus = arg[0].(string)
	default:
	}
	_, err := entities.MasterEngine.Id(id).Update(classroom)

	return err
}

//Delete classroom information
func (*ClassroomInfoService) DeleteInfo(id int) error {
	classroom := new(entities.ClassroomInfo)
	_, err := entities.MasterEngine.Id(id).Delete(classroom)

	return err
}

//Get the classroom by campus, building, Capticity
func (*ClassroomInfoService) GetClassroomBySomeCond(arg ...interface{}) ([]entities.ClassroomInfo, error) {
	var err error
	classrooms := make([]entities.ClassroomInfo, 0)
	switch len(arg) {
	case 1:
		err = entities.SlaveEngine.Where("classroom_campus = ?", arg[0]).Find(&classrooms)
		break
	case 2:
		err = entities.SlaveEngine.Where("classroom_campus = ? AND  classroom_building = ?", arg[0],
			arg[1]).Find(&classrooms)
		break
	case 3:
		// err = entities.SlaveEngine.Where("classroom_campus = ? AND  classroom_building = ? and capacity > ?", arg[0],
		// 	arg[1], arg[2]).Find(&classrooms)
		err = entities.SlaveEngine.Where("classroom_campus = ? AND  classroom_building = ? and capacity = ?", arg[0],
			arg[1], arg[2]).Find(&classrooms)
	}

	return classrooms, err
}
