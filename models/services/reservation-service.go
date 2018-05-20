package services

import (
	"time"

	"github.com/kangbb/ccrsystem/models/entities"
)

type ReservationInfoService struct{}

var ReservationService = ReservationInfoService{}

func (*ReservationInfoService) NewReservation(reason string, start time.Time, end time.Time, classroomId int,
	studentId int, organizationName string, approverId int, note string, state int) *entities.ReservationInfo {
	reservation := &entities.ReservationInfo{
		ResReason:        reason,
		StartTime:        start,
		EndTime:          end,
		ClassroomId:      classroomId,
		StudentId:        studentId,
		OrganizationName: organizationName,
		ApproverId:       approverId,
		ApprovalNote:     note,
		ResState:         state,
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

//Get reservation information by classroom Id
func (*ReservationInfoService) FindInfoByClassroomId(classroomId int) ([]entities.ReservationInfo, error) {
	reservations := make([]entities.ReservationInfo, 0)
	err := entities.SlaveEngine.Where("classroom_id = ?", classroomId).Find(&reservations)

	return reservations, err
}

//Find all reservation information
func (*ReservationInfoService) FindAllInfo() ([]entities.ReservationInfo, error) {
	res := make([]entities.ReservationInfo, 0)
	err := entities.SlaveEngine.Find(&res)

	return res, err
}

//Update reservation information
//Just for resReason, organizationName, resState, approvalNote, and approverId.
func (*ReservationInfoService) UpdateInfo(id int, arg ...interface{}) error {
	reservation := new(entities.ReservationInfo)
	switch len(arg) {
	case 5:
		reservation.ApproverId = arg[4].(int)
	case 4:
		reservation.ApprovalNote = arg[3].(string)
		reservation.ResState = arg[2].(int)
	case 2:
		reservation.OrganizationName = arg[1].(string)
		reservation.ResReason = arg[0].(string)
		break
	default:
	}
	_, err := entities.MasterEngine.Cols("res_reason", "organization_name", "res_state", "approval_note", "approver_id").Id(id).Update(reservation)

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
	// To implement the fuzzy query(模糊查询)
	// switch len(arg) {
	// case 0:
	// 	err = entities.SlaveEngine.Where("classroom_id = ?", classroomId).Find(&reservations)
	// 	break
	// case 1:
	// 	err = entities.SlaveEngine.Where("classroom_id = ? AND (start_time > ? OR  end_time < ?)",
	// 		classroomId, arg[0].(time.Time), arg[0].(time.Time)).Find(&reservations)
	// 	break
	// default:
	// 	err = entities.SlaveEngine.Where("classroom_id = ? AND (start_time > ? OR  end_time < ?)",
	// 		classroomId, arg[1].(time.Time), arg[0].(time.Time)).Find(&reservations)
	// }

	// To implement the accurate query
	err = entities.SlaveEngine.Where("classroom_id = ? AND ((start_time <= ? AND  end_time >= ?) OR (start_time <= ? AND  end_time >= ?))",
		classroomId, arg[0].(time.Time), arg[0].(time.Time), arg[1].(time.Time), arg[1].(time.Time)).Find(&reservations)

	return reservations, err
}
