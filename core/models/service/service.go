package service

import (
	"log"

	"github.com/kangbb/class-reservation/core/models/entities"
)

//Create a new student
func NewStudent(id int, pwd string, name string, perm bool) entities.StudentInfo {
	student := entities.StudentInfo{
		StudentId:   id,
		StudentPwd:  pwd,
		StudentName: name,
		Permission:  perm,
	}
	return student
}
func NewAdmin(id int, pwd string, name string, perm bool) entities.AdminInfo {
	admin := entities.AdminInfo{
		AdminId:    id,
		AdminPwd:   pwd,
		AdminName:  name,
		Permission: perm,
	}
	return admin
}
func NewApprover(id int, pwd string, name string, perm bool) entities.ApproverInfo {
	approver := entities.ApproverInfo{
		ApproverId:   id,
		ApproverPwd:  pwd,
		ApproverName: name,
		Permission:   perm,
	}
	return approver
}
func NewClassroom(campuse string, building string, num string, cap int) entities.ClassroomInfo {
	classroom := entities.ClassroomInfo{
		ClassroomCampus:   campuse,
		ClassroomBuilding: building,
		ClassroomNum:      num,
		Capicity:          cap,
	}
	return classroom
}
func NewReservation(state string, start string, end string, reason string, studentId int,
	approverId int, classroomId int) entities.ReservationInfo {
	reservation := entities.ReservationInfo{
		ResState:    state,
		StartTime:   start,
		EndTime:     end,
		ResReason:   reason,
		StudentId:   studentId,
		ApproverId:  approverId,
		ClassroomId: classroomId,
	}
	return reservation
}

/***********************************CREATE*******************************/

// Create a item and insert it to the db
// The item maybe student, admin, approver, classroom, reservation
func CreateAItem(itemType string, arg ...interface{}) (bool, error) {
	var item interface{}
	switch itemType {
	case "StudentInfo":
		item = NewStudent(arg[0].(int), arg[1].(string), arg[2].(string), arg[3].(bool))
		break
	case "AdminInfo":
		item = NewAdmin(arg[0].(int), arg[1].(string), arg[2].(string), arg[3].(bool))
		break
	case "ApproverInfo":
		item = NewApprover(arg[0].(int), arg[1].(string), arg[2].(string), arg[3].(bool))
		break
	case "ClassroomInfo":
		item = NewClassroom(arg[0].(string), arg[1].(string), arg[2].(string), arg[3].(int))
		break
	default:
		item = NewReservation(arg[0].(string), arg[1].(string), arg[2].(string), arg[3].(string),
			arg[4].(int), arg[5].(int), arg[6].(int))
	}

	_, err := entities.MasterEngine.InsertOne(&item)
	if err != nil {
		return false, err
	}
	return true, nil
}

/***********************************RETRIEVE*******************************/
//Get a user information by ID
//Could be student, admin, approver
//Return is a slice
func GetAUser(userType string, id int) interface{} {
	var user interface{}
	switch userType {
	case "Student":
		user = new(entities.StudentInfo)
		break
	case "Admin":
		user = new(entities.AdminInfo)
		break
	case "Approver":
		user = new(entities.ApproverInfo)
		break
	}
	_, err := entities.SlaveEngine.Id(id).Get(user)
	if err != nil {
		log.Println(err)
	}
	return user
}

//Get all items
func GetAllItem(itemType string) interface{} {
	var items interface{}
	switch itemType {
	case "StudentInfo":
		items = make([]entities.StudentInfo, 0)
		break
	case "AdminInfo":
		items = make([]entities.AdminInfo, 0)
		break
	case "ApproverInfo":
		items = make([]entities.ApproverInfo, 0)
		break
	case "ClassroomInfo":
		items = make([]entities.ClassroomInfo, 0)
		break
	default:
		items = make([]entities.ReservationInfo, 0)
	}
	err := entities.SlaveEngine.Find(&items)
	if err != nil {
		panic(err)
	}
	return items
}

// Get the classroom by id
func GetClassroomById(id int) *entities.ClassroomInfo {
	classroom := new(entities.ClassroomInfo)
	_, err := entities.SlaveEngine.Id(id).Get(classroom)
	if err != nil {
		log.Println(err)
	}
	return classroom
}

//Get the classroom by campus, building, Capticity
func GetClassroomBySomeCon(arg ...interface{}) []entities.ClassroomInfo {
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

//Get the reservation by id
func GetReservationById(id int) entities.ReservationInfo {
	reservation := new(entities.ReservationInfo)
	_, err := entities.SlaveEngine.Id(id).Get(reservation)
	if err != nil {
		log.Println(err)
	}
	return *reservation
}

//Get Reservation by classroomID and time
//default argument sequence: classroomID, startTime, endTime
func GetReservationBySomeCon(classroomID int, arg ...interface{}) []entities.ReservationInfo {
	var err error
	reservations := make([]entities.ReservationInfo, 0)
	switch len(arg) {
	case 0:
		err = entities.SlaveEngine.Where("classroom_id = ?", classroomID).Find(&reservations)
		break
	case 1:
		err = entities.SlaveEngine.Where("classroom_id = ? AND (start_time > ? OR  end_time < ?)",
			classroomID, arg[0], arg[0]).Find(&reservations)
		break
	default:
		err = entities.SlaveEngine.Where("classroom_id = ? AND (start_time > ? OR  end_time < ?)",
			classroomID, arg[1], arg[0]).Find(&reservations)
	}
	if err != nil {
		panic(err)
	}
	return reservations
}

/***********************************UPDATE*******************************/
//Update user password
func UpdatePassword(userType string, id int, pwd string) (bool, error) {
	var user interface{}
	switch userType {
	case "Student":
		user = &entities.StudentInfo{StudentPwd: pwd}
		break
	case "Admin":
		user = &entities.AdminInfo{AdminPwd: pwd}
		break
	case "Approver":
		user = &entities.ApproverInfo{ApproverPwd: pwd}
		break
	}
	_, err := entities.MasterEngine.Id(id).Update(user)
	if err != nil {
		panic(err)
		return false, err
	}
	return true, nil
}

//Update permission;
//Just for approver
func UpdatePermission(id int, perm bool) (bool, error) {
	admin := new(entities.AdminInfo)
	admin.Permission = perm
	_, err := entities.MasterEngine.Id(id).Cols("permission").Update(admin)
	if err != nil {
		panic(err)
		return false, err
	}
	return true, nil
}

//Update class information
//Please attention the sequence
func UpdateClassroom(id int, arg ...interface{}) (bool, error) {
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
		return false, err
	}
	return true, nil
}

//Update the Reservation
//Just for ResState, ResReason and ApproverId
func UpdateReservation(id int, arg ...interface{}) (bool, error) {
	reservation := new(entities.ReservationInfo)
	switch len(arg) {
	case 3:
		reservation.ApproverId = arg[2].(int)
	case 2:
		reservation.ResReason = arg[1].(string)
	case 1:
		reservation.ResReason = arg[0].(string)
		break
	default:
	}
	_, err := entities.MasterEngine.Id(id).Update(reservation)
	if err != nil {
		panic(err)
		return false, err
	}
	return true, nil
}

/***********************************DELETE*******************************/
//Delete item
//Just for student, approver, classroom, reservation
func DeleteItem(itemType string, id int) (bool, error) {
	var item interface{}
	switch itemType {
	case "StudentInfo":
		item = new(entities.StudentInfo)
		break
	case "ApproverInfo":
		item = new(entities.ApproverInfo)
		break
	case "ClassroomInfo":
		item = new(entities.ClassroomInfo)
		break
	default:
		item = new(entities.ReservationInfo)
	}
	_, err := entities.MasterEngine.Id(id).Delete(item)
	if err != nil {
		panic(err)
		return false, err
	}
	return false, nil
}
