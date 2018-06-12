package controllers

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"regexp"
	"strconv"
	"time"

	simplejson "github.com/bitly/go-simplejson"
	"github.com/gorilla/mux"
	"github.com/kangbb/ccrsystem/logs"
	"github.com/kangbb/ccrsystem/models/services"
)

type Reservation struct {
	ResId            int
	ResReason        string
	StartTime        string
	EndTime          string
	StudentId        int
	ClassroomId      int
	ClassroomNum     string
	Capacity         int
	OrganizationName string
	ApproverId       int
	ResState         int
	ApprovalNote     string
}

/************************* Data Interface For Reservation ***********************/
/*
* Get reservation by id
 */
func getResById(w http.ResponseWriter, r *http.Request) {
	//default use the parameter from r.body
	id := getIdFromPath(r)
	res, err := services.ReservationService.FindInfoById(id)
	if logs.SqlError(err, w, res.ApproverId != 0) {
		return
	}

	userType := r.Header.Get("userType")
	var userId int
	userId, _ = strconv.Atoi(r.Header.Get("userId"))
	// If the user have no permission accessed the reservation, return
	if userType == "Student" && userId != res.StudentId ||
		userType == "Approver" && userId != res.ApproverId {
		logs.RequestError(401, logs.ErrorMsg{Msg: "Permission deny."}, w)
		return
	}
	// need not to load the reservation ResState = 3 or 2
	if userType == "Approver" && (res.ResState == 3 || res.ResState == 2) {
		logs.RequestError(404, logs.ErrorMsg{Msg: "您查询的信息不存在."}, w)
		return
	}

	classroom, err := services.ClassroomService.FindInfoById(res.ClassroomId)
	if logs.SqlError(err, w, classroom.ClassroomNum != "") {
		return
	}

	result := Reservation{
		ResId:            res.ResId,
		ResReason:        res.ResReason,
		StartTime:        timeToLesson(res.StartTime),
		EndTime:          timeToLesson(res.EndTime),
		StudentId:        res.StudentId,
		ClassroomId:      res.ClassroomId,
		ClassroomNum:     classroom.ClassroomNum,
		Capacity:         classroom.Capacity,
		OrganizationName: res.OrganizationName,
		ApproverId:       res.ApproverId,
		ResState:         res.ResState,
		ApprovalNote:     res.ApprovalNote,
	}
	data, _ := json.Marshal(result)

	w.WriteHeader(200)
	w.Write(data)
}

/*
* Update reservation by id
* For student, only can update reservation reason.
* For approver, only can update reservation state, approvalNote.
 */
func updateResById(w http.ResponseWriter, r *http.Request) {
	var (
		id           int
		departmentId int
		approverId   int
		state        int
		approvalNote string
	)

	//default use the parameter from r.body
	defer r.Body.Close()
	js, err := simplejson.NewFromReader(r.Body)
	if logs.NormalError(err, w) {
		return
	}

	id = getIdFromPath(r)
	userType := r.Header.Get("userType")
	//student update reservation
	if userType == "Student" {
		reason := js.Get("ResReason").MustString()
		organizationName := js.Get("OrganizationName").MustString()
		if reason == "" || organizationName == "" {
			logs.RequestError(500, logs.ErrorMsg{Msg: "必填项不能为空"}, w)
			return
		}

		//If left departments more than one.
		department, err := services.DepartmentService.FindInfoByNote("initial")
		//If just left a department.
		department_one, err_one := services.DepartmentService.FindInfoByNote("initial,final")
		if logs.SqlError(err, w, department.DepartmentName != "") &&
			logs.SqlError(err_one, w, department_one.DepartmentName != "") {
			return
		}
		if department.DepartmentName != "" {
			departmentId = department.DepartmentId
		} else {
			departmentId = department_one.DepartmentId
		}

		approvers, err := services.ApproverService.FindInfoByDepartmentId(departmentId)
		if logs.SqlError(err, w, len(approvers) != 0) {
			return
		}
		//random select a approver from the approvers who belong to the same department
		approverId = approvers[rand.Intn(len(approvers))].ApproverId
		state = 0
		approvalNote = ""
		err = services.ReservationService.UpdateInfo(id, reason, organizationName,
			state, approvalNote, approverId)
		if logs.SqlError(err, w, true) {
			return
		} //approver approve the reservation
	} else if userType == "Approver" {
		//check if reservation exist
		//if not exist, return
		res, err := services.ReservationService.FindInfoById(id)
		if logs.SqlError(err, w, res.ResId != 0) {
			return
		}
		//if exist, update resinfo
		state = js.Get("ResState").MustInt()
		approvalNote = js.Get("ApproverNote").MustString()
		// Get the departmentId
		approver, err := services.ApproverService.FindInfoById(res.ApproverId)
		if logs.SqlError(err, w, approver.ApproverName != "") {
			return
		}
		departmentId = approver.DepartmentId

		if state == 2 {
			department, err := services.DepartmentService.FindInfoById(departmentId)
			if logs.SqlError(err, w, department.DepartmentName != "") {
				return
			}
			if department.Note == "final" || department.Note == "initial,final" {
				state = 2
			} else {
				state = 1

				department, err = services.DepartmentService.FindInfoByOrder(department.DepartmentOrder + 1)
				if logs.SqlError(err, w, department.DepartmentName != "") {
					return
				}
				departmentId = department.DepartmentId
				approvers, err := services.ApproverService.FindInfoByDepartmentId(departmentId)
				if logs.SqlError(err, w, len(approvers) != 0) {
					return
				}
				//random select a approver from the approvers who belong to the same department
				approverId = approvers[rand.Intn(len(approvers))].ApproverId
				approvalNote = ""
			}
		}
		err = services.ReservationService.UpdateInfo(id, res.ResReason, res.OrganizationName,
			state, approvalNote, approverId)

		if logs.SqlError(err, w, true) {
			return
		}

		w.WriteHeader(200)
	} else {
		logs.NormalError(logs.SWITCH_BRANCH_ERROR, w)
	}
}

/*
* Delete reservation by id
 */
func deleteResById(w http.ResponseWriter, r *http.Request) {

	id := getIdFromPath(r)
	err := services.ReservationService.DeleteInfo(id)
	if logs.SqlError(err, w, true) {
		return
	}
	w.WriteHeader(200)
}

/*
* Get a student reservation list
 */
func getStudentResList(w http.ResponseWriter, r *http.Request) {

	var userId int
	userId, _ = strconv.Atoi(r.Header.Get("userId"))
	reservations, err := services.ReservationService.FindInfoByStudentId(userId)
	if logs.SqlError(err, w, len(reservations) != 0) {
		return
	}

	result := make([]Reservation, len(reservations))
	for k, v := range reservations {
		classroom, err := services.ClassroomService.FindInfoById(v.ClassroomId)
		if logs.SqlError(err, w, classroom.ClassroomNum != "") {
			return
		}

		result[k] = Reservation{
			ResId:            v.ResId,
			ResReason:        v.ResReason,
			StartTime:        timeToLesson(v.StartTime),
			EndTime:          timeToLesson(v.EndTime),
			StudentId:        v.StudentId,
			ClassroomId:      v.ClassroomId,
			ClassroomNum:     classroom.ClassroomNum,
			Capacity:         classroom.Capacity,
			OrganizationName: v.OrganizationName,
			ApproverId:       v.ApproverId,
			ResState:         v.ResState,
			ApprovalNote:     v.ApprovalNote,
		}
	}

	data, _ := json.Marshal(result)

	w.WriteHeader(200)
	w.Write(data)
}

/*
* Get a approver reservation list
 */
func getApproverResList(w http.ResponseWriter, r *http.Request) {
	var userId int
	userId, _ = strconv.Atoi(r.Header.Get("userId"))

	reservations, err := services.ReservationService.FindInfoByApproverId(userId)
	if logs.SqlError(err, w, len(reservations) != 0) {
		return
	}

	result := make([]Reservation, 0)
	for _, v := range reservations {
		if v.ResState == 2 || v.ResState == 3 {
			continue
		}
		classroom, err := services.ClassroomService.FindInfoById(v.ClassroomId)
		if logs.SqlError(err, w, classroom.ClassroomNum != "") {
			return
		}
		result = append(result, Reservation{
			ResId:            v.ResId,
			ResReason:        v.ResReason,
			StartTime:        timeToLesson(v.StartTime),
			EndTime:          timeToLesson(v.EndTime),
			StudentId:        v.StudentId,
			ClassroomId:      v.ClassroomId,
			ClassroomNum:     classroom.ClassroomNum,
			Capacity:         classroom.Capacity,
			OrganizationName: v.OrganizationName,
			ApproverId:       v.ApproverId,
			ResState:         v.ResState,
			ApprovalNote:     v.ApprovalNote,
		})
	}

	data, _ := json.Marshal(result)

	w.WriteHeader(200)
	w.Write(data)
}

/*
* Add a reservation to db
* Just accessed for the student.
 */
func addRes(w http.ResponseWriter, r *http.Request) {

	//default use the parameter from r.body
	defer r.Body.Close()
	js, err := simplejson.NewFromReader(r.Body)
	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
		return
	}

	start_time := js.Get("StartTime").MustString()
	end_time := js.Get("EndTime").MustString()
	reason := js.Get("ResReason").MustString()
	classroomId := js.Get("ClassroomId").MustInt()
	organizationName := js.Get("OrganizationName").MustString()
	if start_time == "" || end_time == "" || reason == "" || classroomId == 0 || organizationName == "" {
		logs.RequestError(500, logs.ErrorMsg{Msg: "必填项不能为空"}, w)
		return
	}
	start := lessonToTime(start_time)
	end := lessonToTime(end_time)

	approvalNote := ""

	var studentId int
	studentId, _ = strconv.Atoi(r.Header.Get("userId"))

	department, err := services.DepartmentService.FindInfoByNote("initial")
	if logs.SqlError(err, w, department.DepartmentName != "") {
		return
	}
	departmentId := department.DepartmentId
	approvers, err := services.ApproverService.FindInfoByDepartmentId(departmentId)
	if logs.SqlError(err, w, len(approvers) != 0) {
		return
	}
	//random select a approver from the approvers who belong to the same department
	approverId := approvers[rand.Intn(len(approvers))].ApproverId

	res := services.ReservationService.NewReservation(reason, start, end, classroomId, studentId,
		organizationName, approverId, approvalNote, 0)
	err = services.ReservationService.SaveAInfo(res)

	if logs.SqlError(err, w, true) {
		return
	}
	w.WriteHeader(200)
}

/***********************************SOME SUBFUNCTION*******************************/
func lessonToTime(lessonTime string) time.Time {
	mappingRegular := map[string]int{
		"一": 1, "二": 2, "三": 3, "四": 4, "五": 5, "六": 6, "七": 7, "八": 8, "九": 9, "十": 10,
	}
	var lesson int
	reg := regexp.MustCompile(`[0-9\-]+`)
	date := reg.FindAllString(lessonTime, -1)[0]
	reg = regexp.MustCompile(`[\p{Han}]+`)
	hanstr := []rune(reg.FindAllString(lessonTime, -1)[0])
	switch len(hanstr) {
	case 4:
		lesson = mappingRegular[string(hanstr[1])]
		break
	case 5:
		lesson = mappingRegular[string(hanstr[1])] + mappingRegular[string(hanstr[2])]
		break
	}
	//The database store time as CST
	//So we need parse time in Location.
	loc, _ := time.LoadLocation("Local")
	startTime, _ := time.ParseInLocation("2006-01-02 15:04:05", date+" "+"08:00:00", loc)
	duraTime := strconv.Itoa((lesson-1)*(45+10)) + "m"
	addTime, _ := time.ParseDuration(duraTime)
	return startTime.Add(addTime)
}
func timeToLesson(trueTime time.Time) string {
	mappingRegular := []string{"一", "二", "三", "四", "五", "六", "七", "八", "九", "十"}
	var lesson string
	date := trueTime.Format("2006-01-02")
	loc, _ := time.LoadLocation("Local")
	startTime, _ := time.ParseInLocation("2006-01-02 15:04:05", date+" "+"08:00:00", loc)
	lessonInNum := int(trueTime.Sub(startTime).Minutes())/55 + 1
	if lessonInNum > 10 {
		lesson = mappingRegular[9] + mappingRegular[lessonInNum%10-1]
	} else if lessonInNum == 10 {
		lesson = mappingRegular[9]
	} else {
		lesson = mappingRegular[lessonInNum%10-1]
	}

	return date + " " + "第" + lesson + "节"
}

func encryptPwd(id int, pwd string) string {
	h := md5.New()
	io.WriteString(h, pwd)

	id_str := fmt.Sprintf("%d", id)
	pwmd5 := fmt.Sprintf("%x", h.Sum(nil))

	//指定两个 salt
	salt1 := "@#$**"
	salt2 := "^&*(("

	//salt1+用户名+salt2+MD5拼接
	io.WriteString(h, salt1)
	io.WriteString(h, id_str)
	io.WriteString(h, salt2)
	io.WriteString(h, pwmd5)

	encrypt := fmt.Sprintf("%x", h.Sum(nil))
	return encrypt
}

func getIdFromPath(r *http.Request) int {
	var id int
	id, _ = strconv.Atoi(mux.Vars(r)["id"])
	return id
}
