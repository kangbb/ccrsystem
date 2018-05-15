package services

import (
	"time"

	"github.com/kangbb/ccrsystem/models/entities"
)

type ReservationInfoService struct{}

var ReservationService = ReservationInfoService{}

func (*ReservationInfoService) NewReservation(state string, start time.Time, end time.Time, departmentId int,
	reason string, note string, studentId int, approverId int, classroomId int, OrganizationId int) *entities.ReservationInfo {
	reservation := &entities.ReservationInfo{
		ResState:       state,
		StartTime:      start,
		EndTime:        end,
		ResReason:      reason,
		ApprovalNote:   note,
		DepartmentId:   departmentId,
		StudentId:      studentId,
		ApproverId:     approverId,
		ClassroomId:    classroomId,
		OrganizationId: OrganizationId,
	}
	return reservation
}

func (*ReservationInfoService) NewReservationSlice() []entities.ReservationInfo {
	return make([]entities.ReservationInfo, 0)
}

//Insert reservation information to the db
func (*ReservationInfoService) SaveAInfo(res *entities.ReservationInfo) error {
	_, err := entities.MasterEngine.InsertOne(res)

	return err
}

//Get a reservation information by ID
func (*ReservationInfoService) FindInfoById(id int) (*entities.ReservationInfo, error) {
	res := new(entities.ReservationInfo)
	_, err := entities.SlaveEngine.Id(id).Get(res)

	return res, err
}

//Get reservation information by approver Id
func (*ReservationInfoService) FindInfoByApproverId(approverId int) ([]entities.ReservationInfo, error) {
	reservations := make([]entities.ReservationInfo, 0)
	err := entities.SlaveEngine.Where("approver_id = ?", approverId).Find(&reservations)

	return reservations, err
}

//Get reservation information by student Id
func (*ReservationInfoService) FindInfoByStudentId(studentId int) ([]entities.ReservationInfo, error) {
	reservations := make([]entities.ReservationInfo, 0)
	err := entities.SlaveEngine.Where("student_id = ?", studentId).Find(&reservations)

	return reservations, err
}

//Get resvation information by department Id
func (*ReservationInfoService) FindInfoByDepartmentId(departmentId int) ([]entities.ReservationInfo, error) {
	reservations := make([]entities.ReservationInfo, 0)
	err := entities.SlaveEngine.Where("department_id = ?", departmentId).Find(&reservations)

	return reservations, err
}

//Find all reservation information
func (*ReservationInfoService) FindAllInfo() ([]entities.ReservationInfo, error) {
	res := make([]entities.ReservationInfo, 0)
	err := entities.SlaveEngine.Find(&res)

	return res, err
}

//Update reservation information
//Just for ResState, Department, ApproverId and ApproverNote
func (*ReservationInfoService) UpdateInfo(id int, arg ...interface{}) error {
	reservation := new(entities.ReservationInfo)
	switch len(arg) {
	case 5:
		reservation.ResReason = arg[4].(string)
	case 4:
		reservation.ApprovalNote = arg[2].(string)
		reservation.ResState = arg[3].(string)
	case 2:
		reservation.ApproverId = arg[0].(int)
		reservation.DepartmentId = arg[1].(int)
		break
	default:
	}
	_, err := entities.MasterEngine.AllCols().Id(id).Update(reservation)

	return err
}

//Delete reservation information
func (*ReservationInfoService) DeleteInfo(id int) error {
	res := new(entities.ReservationInfo)
	_, err := entities.MasterEngine.Id(id).Delete(res)

	return err
}

//Get Reservation by classroomId and time
//default argument sequence: classroomId, startTime, endTime
func (*ReservationInfoService) GetReservationBySomeCond(classroomId int, arg ...interface{}) ([]entities.ReservationInfo, error) {
	var err error
	reservations := make([]entities.ReservationInfo, 0)
	switch len(arg) {
	case 0:
		err = entities.SlaveEngine.Where("classroom_id = ?", classroomId).Find(&reservations)
		break
	case 1:
		err = entities.SlaveEngine.Where("classroom_id = ? AND (start_time > ? OR  end_time < ?)",
			classroomId, arg[0].(time.Time), arg[0].(time.Time)).Find(&reservations)
		break
	default:
		err = entities.SlaveEngine.Where("classroom_id = ? AND (start_time > ? OR  end_time < ?)",
			classroomId, arg[1].(time.Time), arg[0].(time.Time)).Find(&reservations)
	}

	return reservations, err
}
