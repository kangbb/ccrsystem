package service

import (
	"github.com/kangbb/class-reservation/core/models/entities"
)

type ClassroomInfoService struct{}

var ClassroomService = ClassroomInfoService{}

//Create a new Classroom
func (*ClassroomInfoService) NewClassroom(campuse string, building string, num string, cap int) entities.ClassroomInfo {
	classroom := entities.ClassroomInfo{
		ClassroomCampus:   campuse,
		ClassroomBuilding: building,
		ClassroomNum:      num,
		Capicity:          cap,
	}
	return classroom
}

//Create a slice of Classroom
func (*ClassroomInfoService) NewClassroomSlice() []entities.ClassroomInfo {
	return make([]entities.ClassroomInfo, 0)
}

//Insert classroom information to the db
func (*ClassroomInfoService) SaveAInfo(classroom *entities.ClassroomInfo) (bool, error) {
	_, err := entities.MasterEngine.InsertOne(classroom)
	if err != nil {
		return false, err
	}
	return true, nil
}

//Get a classroom information by ID
func (*ClassroomInfoService) FindInfoById(id int) entities.ClassroomInfo {
	classroom := new(entities.ClassroomInfo)
	_, err := entities.SlaveEngine.Id(id).Get(classroom)
	if err != nil {
		panic(err)
	}
	return *classroom
}

//Find all classroom information
func (*ClassroomInfoService) FindAllInfo() []entities.ClassroomInfo {
	classroom := make([]entities.ClassroomInfo, 0)
	err := entities.SlaveEngine.Find(&classroom)
	if err != nil {
		panic(err)
	}
	return classroom
}

//Update classroom information
func (*ClassroomInfoService) UpdateInfo(id int, arg ...interface{}) (bool, error) {
	classroom := new(entities.ClassroomInfo)
	switch len(arg) {
	case 4:
		classroom.Capicity = arg[3].(int)
	case 3:
		classroom.ClassroomNum = arg[2].(string)
	case 2:
		classroom.ClassroomBuilding = arg[1].(string)
	case 1:
		classroom.ClassroomCampus = arg[0].(string)
		break
	default:
	}
	_, err := entities.MasterEngine.Id(id).Update(classroom)
	if err != nil {
		panic(err)
	}
	return true, nil
}

//Delete classroom information
func (*ClassroomInfoService) DeleteInfo(id int) (bool, error) {
	classroom := new(entities.ClassroomInfo)
	_, err := entities.MasterEngine.Id(id).Delete(classroom)
	if err != nil {
		panic(err)
	}
	return false, nil
}

//Get the classroom by campus, building, Capticity
func (*ClassroomInfoService) GetClassroomBySomeCond(arg ...interface{}) []entities.ClassroomInfo {
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
		err = entities.SlaveEngine.Where("classroom_campus = ? AND  classroom_building = ? and capicity = ?", arg[0],
			arg[1], arg[2]).Find(&classrooms)
	}
	if err != nil {
		panic(err)
	}
	return classrooms
}
