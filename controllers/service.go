package controllers

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"html/template"
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
	"github.com/kangbb/ccrsystem/middlewares"
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

/******************************** Render File *************************/
func renderFile(name string, w http.ResponseWriter, r *http.Request) {
	filepath := "views/" + name + ".html"
	t := template.Must(template.ParseFiles(filepath))
	w.Header().Set("Content-Type", "text/html; charset=UTF-8")
	w.WriteHeader(200)
	t.Execute(w, nil)
}

/****************************** Signin and Signout *************************/
/*
* User Signin
* For student, approver, admin. If sucess, redirect to the index which belong to themselves
* If failed, redirect to the home page.
 */
func signin(userType string, w http.ResponseWriter, r *http.Request) {
	var (
		id   int
		name string
		pwd  string
	)

	// UserTypeEnum = []string{"Student", "Admin", "Approver"}
	defer r.Body.Close()
	js, err := simplejson.NewFromReader(r.Body)
	if logs.NormalError(err, w) {
		return
	}

	// The format of the id and password are wrong.
	stringId := fmt.Sprintf("%d", js.Get(userType+"Id").MustInt())
	if !UserIfoFormatValidate(stringId, js.Get(userType+"Pwd").MustString(), w) {
		return
	}

	id = js.Get(userType + "Id").MustInt()
	pwd = js.Get(userType + "Pwd").MustString()
	name = js.Get(userType + "Name").MustString()
	switch userType {
	case "Student":
		usr, err := services.StudentService.FindInfoById(id)
		if logs.SqlError(err, w, usr.StudentName != "") {
			return
		}
		if !UserPasswordValidate(encryptPwd(id, pwd), usr.StudentPwd, w) {
			return
		}
		break
	case "Admin":
		usr, err := services.AdminService.FindInfoById(id)
		if logs.SqlError(err, w, usr.AdminName != "") {
			return
		}
		if !UserPasswordValidate(encryptPwd(id, pwd), usr.AdminPwd, w) {
			return
		}
		break
	case "Approver":
		usr, err := services.ApproverService.FindInfoById(id)
		if logs.SqlError(err, w, usr.ApproverName != "") {
			return
		}
		if !UserPasswordValidate(encryptPwd(id, pwd), usr.ApproverPwd, w) {
			return
		}
		break
	default:
		logs.NormalError(logs.SWITCH_BRANCH_ERROR, w)
		return
	}

	if middlewares.SessionProcess(w, r, stringId, userType, name) {
		w.WriteHeader(200)
	}
}

/*
* User Signout
* Just delete the sessions.
 */
func signout(w http.ResponseWriter, r *http.Request) {
	if !middlewares.SessionProcess(w, r) {
		return
	}
	req, _ := http.NewRequest("GET", "/", r.Body)
	http.Redirect(w, req, "/", 302)
}

/************************* Data Interface For User ***********************/
/*
* Get userinfo by session
 */
func getUserInfo(userType string, w http.ResponseWriter, r *http.Request) {
	var data []byte
	var id int
	id, _ = strconv.Atoi(r.Header.Get("userId"))
	switch userType {
	case "Student":
		user, err := services.StudentService.FindInfoById(id)
		if logs.SqlError(err, w, user.StudentName != "") {
			return
		}
		user.StudentPwd = ""
		data, _ = json.Marshal(*user)
		break
	case "Approver":
		user, err := services.ApproverService.FindInfoById(id)
		if logs.SqlError(err, w, user.ApproverName != "") {
			return
		}
		user.ApproverPwd = ""
		data, _ = json.Marshal(*user)
		break
	case "Admin":
		user, err := services.AdminService.FindInfoById(id)
		if logs.SqlError(err, w, user.AdminName != "") {
			return
		}
		user.AdminPwd = ""
		data, _ = json.Marshal(*user)
		break
	default:
		logs.NormalError(logs.SWITCH_BRANCH_ERROR, w)
		return
	}
	w.WriteHeader(200)
	w.Write(data)
}

/*
* Update userinfo by session
* Just for password
 */
func updateUserInfo(userType string, w http.ResponseWriter, r *http.Request) {
	var err error
	//default use the parameter from body
	defer r.Body.Close()
	js, err := simplejson.NewFromReader(r.Body)
	if logs.NormalError(err, w) {
		return
	}

	var id int
	id, _ = strconv.Atoi(r.Header.Get("userId"))
	pwd := encryptPwd(id, js.Get(userType+"Pwd").MustString())
	switch userType {
	case "Student":
		err = services.StudentService.UpdateInfo(id, pwd)
		break
	case "Approver":
		err = services.ApproverService.UpdatePasswordInfo(id, pwd)
		break
	case "Admin":
		err = services.AdminService.UpdateInfo(id, pwd)
		break
	default:
		logs.NormalError(logs.SWITCH_BRANCH_ERROR, w)
		return
	}

	if logs.SqlError(err, w, true) {
		return
	} else {
		w.WriteHeader(200)
	}
}

/*
* Get userlist
* Just accessed for the admin, the password will not return
 */
func getUserList(userType string, w http.ResponseWriter, r *http.Request) {
	var data []byte
	switch userType {
	case "Student":
		users, err := services.StudentService.FindAllInfo()
		if logs.SqlError(err, w, len(users) != 0) {
			return
		}
		for k, _ := range users {
			users[k].StudentPwd = ""
		}
		data, _ = json.Marshal(users)
		break
	case "Approver":
		users, err := services.ApproverService.FindAllInfo()
		if logs.SqlError(err, w, len(users) != 0) {
			return
		}
		for k, _ := range users {
			users[k].ApproverPwd = ""
		}

		data, _ = json.Marshal(users)
		break
	case "Admin":
		users, err := services.AdminService.FindAllInfo()
		if logs.SqlError(err, w, len(users) != 0) {
			return
		}
		for k, _ := range users {
			users[k].AdminPwd = ""
		}
		data, _ = json.Marshal(users)
		break
	default:
		logs.NormalError(logs.SWITCH_BRANCH_ERROR, w)
		return
	}
	w.WriteHeader(200)
	w.Write(data)
}

/*
* Add a user to db
 */
func addUser(userType string, w http.ResponseWriter, r *http.Request) {
	var err error
	//default paremeter from r.body
	defer r.Body.Close()
	js, err := simplejson.NewFromReader(r.Body)
	if logs.NormalError(err, w) {
		return
	}

	id := js.Get(userType + "Id").MustInt()
	pwd := encryptPwd(id, js.Get(userType+"Pwd").MustString())
	switch userType {
	case "Student":
		user := services.StudentService.NewStudent(id, pwd, js.Get("StudentName").MustString())
		err = services.StudentService.SaveAInfo(user)
		break
	case "Approver":
		// Validating whether departmentId is exiting.
		departmentId := js.Get("DepartmentId").MustInt()
		dep, err := services.DepartmentService.FindInfoById(departmentId)
		if logs.SqlError(err, w, dep.DepartmentName != "") {
			return
		}
		user := services.ApproverService.NewApprover(id, pwd, js.Get("ApproverName").MustString(),
			departmentId)
		err = services.ApproverService.SaveAInfo(user)
		break
	case "Admin":
		user := services.AdminService.NewAdmin(id, pwd, js.Get("AdminName").MustString())
		err = services.AdminService.SaveAInfo(user)
		break
	default:
		logs.NormalError(logs.SWITCH_BRANCH_ERROR, w)
		return
	}

	if logs.SqlError(err, w, true) {
		return
	} else {
		w.WriteHeader(200)
	}
}

/*
* Get userinfo by id
* Just acessed for admin or the student, approver who possess the account.
 */
func getUserById(userType string, w http.ResponseWriter, r *http.Request) {
	var data []byte

	id := getIdFromPath(r)
	switch userType {
	case "Student":
		user, err := services.StudentService.FindInfoById(id)
		if logs.SqlError(err, w, user.StudentName != "") {
			return
		}
		user.StudentPwd = ""
		data, _ = json.Marshal(*user)
		break
	case "Approver":
		user, err := services.ApproverService.FindInfoById(id)
		if logs.SqlError(err, w, user.ApproverName != "") {
			return
		}
		user.ApproverPwd = ""
		data, _ = json.Marshal(*user)
		break
	case "Admin":
		user, err := services.AdminService.FindInfoById(id)
		if logs.SqlError(err, w, user.AdminName != "") {
			return
		}
		user.AdminPwd = ""
		data, _ = json.Marshal(*user)
		break
	default:
		logs.NormalError(logs.SWITCH_BRANCH_ERROR, w)
		return
	}
	w.WriteHeader(200)
	w.Write(data)
}

/*
* Update userinfo by id
* just accessed for admin or the student, approver who possess the account.
 */
func updateUserById(userType string, w http.ResponseWriter, r *http.Request) {
	var err error
	//default use the parameter from body
	defer r.Body.Close()
	js, err := simplejson.NewFromReader(r.Body)
	if logs.NormalError(err, w) {
		return
	}

	id := getIdFromPath(r)
	pwd := encryptPwd(id, js.Get(userType+"Pwd").MustString())
	switch userType {
	case "Student":
		err = services.StudentService.UpdateInfo(id, pwd)
		break
	case "Approver":
		err = services.ApproverService.UpdatePasswordInfo(id, pwd)
		break
	case "Admin":
		err = services.AdminService.UpdateInfo(id, pwd)
		break
	default:
		logs.NormalError(logs.SWITCH_BRANCH_ERROR, w)
		return
	}

	if logs.SqlError(err, w, true) {
		return
	} else {
		w.WriteHeader(200)
	}
}

/*
* Delete userinfo by id
* Just acessed for admin
 */
func deleteUserById(userType string, w http.ResponseWriter, r *http.Request) {
	var err error
	//default use the parameter from r.body
	defer r.Body.Close()

	id := getIdFromPath(r)
	switch userType {
	case "Student":
		err = services.StudentService.DeleteInfo(id)
		if logs.SqlError(err, w, true) {
			return
		}

		// Delete the reservations related.
		reservations, err := services.ReservationService.FindInfoByStudentId(id)
		if logs.SqlError(err, w, true) {
			return
		}
		for _, v := range reservations {
			err = services.ReservationService.DeleteInfo(v.ResId)
			if logs.SqlError(err, w, true) {
				return
			}
		}
		break
	case "Approver":
		approver, err := services.ApproverService.FindInfoById(id)
		if logs.SqlError(err, w, approver.ApproverName != "") {
			return
		}
		//Delete the approver
		err = services.ApproverService.DeleteInfo(id)

		//Get the approvers who are in the same department
		approvers, err := services.ApproverService.FindInfoByDepartmentId(approver.DepartmentId)
		if logs.SqlError(err, w, true) {
			return
		}
		// No approver in this department, then delete the department.
		if len(approvers) == 0 {
			del_dep, err := services.DepartmentService.FindInfoById(approver.DepartmentId)
			if logs.SqlError(err, w, del_dep.DepartmentName != "") {
				return
			}
			err = services.DepartmentService.DeleteInfo(id)
			if logs.SqlError(err, w, true) {
				return
			}

			// Update order
			departments, err := services.DepartmentService.FindAllInfo()
			if logs.SqlError(err, w, len(departments) != 0) {
				return
			}
			order := del_dep.Order
			for _, v := range departments {
				if v.Order > order {
					err = services.DepartmentService.UpdateInfo(v.DepartmentId, v.DepartmentName, v.Introduction, v.Order-1,
						v.Note)
					if logs.SqlError(err, w, true) {
						return
					}
				}
			}

			// Update reservation
			//Define some variable for updating reservatons
			new_dep, err := services.DepartmentService.FindInfoByOrder(del_dep.Order)
			if logs.SqlError(err, w, new_dep.DepartmentName != "") {
				return
			}
			new_approvers, err := services.ApproverService.FindInfoByDepartmentId(new_dep.DepartmentId)
			if logs.SqlError(err, w, len(new_approvers) != 0) {
				return
			}
			reservations, err := services.ReservationService.FindInfoByApproverId(id)
			if logs.SqlError(err, w, true) {
				return
			}
			for _, res := range reservations {
				res.ApproverId = approvers[rand.Intn(len(new_approvers))].ApproverId
				err := services.ReservationService.UpdateInfo(res.ResId, res.ResReason, res.OrganizationName, res.ResState, res.ApprovalNote, res.ApproverId)
				if logs.SqlError(err, w, true) {
					return
				}
			}
		} else {
			// Has approvers in this department, then delete the department.
			reservations, err := services.ReservationService.FindInfoByApproverId(id)
			if logs.SqlError(err, w, true) {
				return
			}
			for _, res := range reservations {
				res.ApproverId = approvers[rand.Intn(len(approvers))].ApproverId
				err := services.ReservationService.UpdateInfo(res.ResId, res.ResReason, res.OrganizationName, res.ResState, res.ApprovalNote, res.ApproverId)
				if logs.SqlError(err, w, true) {
					return
				}
			}
		}
		break
	case "Admin":
		//admin can't delete count which belongs to himself
		if fmt.Sprintf("%d", id) == r.Header.Get("userId") {
			logs.NormalError(logs.PERMISSION_DENY, w)
			return
		}
		err = services.AdminService.DeleteInfo(id)
		break
	default:
		logs.NormalError(logs.SWITCH_BRANCH_ERROR, w)
		return
	}

	if logs.SqlError(err, w, true) {
		return
	} else {
		w.WriteHeader(200)
	}
}

/************************* Data Interface For Classroom ***********************/
/*
* Get classroomlist
* Accessed for all user, which are student, approver, and admin
 */
func getClassroomList(w http.ResponseWriter, r *http.Request) {

	classrooms, err := services.ClassroomService.FindAllInfo()
	if logs.SqlError(err, w, len(classrooms) != 0) {
		return
	}

	data, _ := json.Marshal(classrooms)
	w.WriteHeader(200)
	w.Write(data)
}

/*
* Add a classroom information to db
* Just accessed for admin.
 */
func addClassroom(w http.ResponseWriter, r *http.Request) {

	//default use the parameter from r.body
	defer r.Body.Close()
	js, err := simplejson.NewFromReader(r.Body)
	if logs.NormalError(err, w) {
		return
	}

	classroomCampus := js.Get("ClassroomCampus").MustString()
	classroomBuilding := js.Get("ClassroomBuilding").MustString()
	classroomNum := js.Get("ClassroomNum").MustString()
	cap := js.Get("Capacity").MustInt()
	if classroomCampus == "" || classroomBuilding == "" || classroomNum == "" || cap == 0 {
		w.WriteHeader(500)
		data, _ := json.Marshal(logs.ErrorMsg{Msg: "必填项不能为空"})
		w.Write(data)
		return
	}
	classroom := services.ClassroomService.NewClassroom(classroomCampus, classroomBuilding, classroomNum, cap)
	err = services.ClassroomService.SaveAInfo(classroom)

	if logs.SqlError(err, w, true) {
		return
	} else {
		w.WriteHeader(200)
	}
}

/*
* Query classroom by some conditions
* Accessed for all users, which are student, approver, admin
 */
func queryClassroom(w http.ResponseWriter, r *http.Request) {
	var capacity int
	//Get request, parameters in headers
	r.ParseForm()
	campuse := r.Form["ClassroomCampus"][0]
	building := r.Form["ClassroomBuilding"][0]
	capacity, _ = strconv.Atoi(r.Form["Capacity"][0])
	start_time := r.Form["StartTime"][0]
	end_time := r.Form["EndTime"][0]
	if campuse == "" || building == "" || capacity == 0 || start_time == "" || end_time == "" {
		logs.RequestError(500, logs.ErrorMsg{Msg: "必填项不能为空"}, w)
		return
	}
	start := lessonToTime(start_time)
	end := lessonToTime(end_time)

	//query classrooms which meet the require
	classrooms, err := services.ClassroomService.GetClassroomBySomeCond(campuse, building, capacity)
	if logs.SqlError(err, w, len(classrooms) != 0) {
		return
	}

	// Can't be 0 in make
	// If do, then will result in error: index out of range
	result := services.ClassroomService.NewClassroomSlice()
	for _, v := range classrooms {
		//query the state of the each classroom
		res, err := services.ReservationService.GetReservationBySomeCond(v.ClassroomId, start, end)
		if logs.SqlError(err, w, true) {
			return
		}

		// If no reservation for the classroom, then the classrrom is idle.
		if len(res) == 0 {
			result = append(result, v)
		}
	}

	data, _ := json.Marshal(result)

	w.WriteHeader(200)
	w.Write(data)
}

/*
* Get classroom by id
* Accessed for all users, which are student, approver, admin
 */
func getClassroomById(w http.ResponseWriter, r *http.Request) {

	id := getIdFromPath(r)
	classroom, err := services.ClassroomService.FindInfoById(id)
	if logs.SqlError(err, w, classroom.ClassroomCampus != "") {
		return
	}

	data, _ := json.Marshal(*classroom)
	w.WriteHeader(200)
	w.Write(data)
}

/*
* Update classroom by id
* Just accessed for admin.
 */
func updateClassroomById(w http.ResponseWriter, r *http.Request) {
	//default use the parameter from r.body
	defer r.Body.Close()
	js, err := simplejson.NewFromReader(r.Body)
	if logs.NormalError(err, w) {
		return
	}

	//default return all the parameter for the classroom
	//include the parameters which don't need to update
	id := getIdFromPath(r)
	cap := js.Get("Capacity").MustInt()
	num := js.Get("ClassroomNum").MustString()
	building := js.Get("ClassroomBuilding").MustString()
	campus := js.Get("ClassroomCampus").MustString()
	if cap == 0 || num == "" || building == "" || campus == "" {
		logs.RequestError(500, logs.ErrorMsg{Msg: "必填字段不能为空"}, w)
		return
	}
	err = services.ClassroomService.UpdateInfo(id, campus, building, num, cap)

	if logs.SqlError(err, w, true) {
		return
	} else {
		w.WriteHeader(200)
	}
}

/*
* Delete classroom by id
* Just accessed for admin.
 */
func deleteClassroomById(w http.ResponseWriter, r *http.Request) {
	//default use the parameter from r.body
	id := getIdFromPath(r)
	err := services.ClassroomService.DeleteInfo(id)

	if logs.SqlError(err, w, true) {
		return
	}
	// Delete the reservations related.
	reservations, err := services.ReservationService.FindInfoByClassroomId(id)
	if logs.SqlError(err, w, true) {
		return
	}
	for _, v := range reservations {
		err = services.ReservationService.DeleteInfo(v.ResId)
		if logs.SqlError(err, w, true) {
			return
		}
	}

	w.WriteHeader(200)
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
				state = 3
			} else {
				state = 1

				department, err = services.DepartmentService.FindInfoByOrder(department.Order + 1)
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

/************************* Data Interface For Department ***********************/
/*
* Get department list
 */
func getDepartmentList(w http.ResponseWriter, r *http.Request) {

	departments, err := services.DepartmentService.FindAllInfo()
	if logs.SqlError(err, w, len(departments) != 0) {
		return
	}

	data, _ := json.Marshal(departments)

	w.WriteHeader(200)
	w.Write(data)
}

/*
* Add a department information to db
* At the same time, the order should update. Default the order be latest; if define, than set custom the order.
 */
func addDepartment(w http.ResponseWriter, r *http.Request) {

	defer r.Body.Close()
	js, err := simplejson.NewFromReader(r.Body)
	if logs.NormalError(err, w) {
		return
	}

	name := js.Get("DepartmentName").MustString()
	order := js.Get("Order").MustInt()
	introduction := js.Get("Introduction").MustString()
	if name == "" || introduction == "" || order == 0 {
		w.WriteHeader(500)
		w.Write([]byte("必填项不能为空"))
	}

	//update the order of the approval department
	departments, err := services.DepartmentService.FindAllInfo()
	if logs.SqlError(err, w, true) {
		return
	}

	val := departments[0]
	if len(departments) == 1 && order > departments[0].Order {
		val.Note = "initial"
		err = services.DepartmentService.UpdateInfo(val.DepartmentId, val.DepartmentName, val.Introduction,
			val.Order+1, val.Note)
		if logs.SqlError(err, w, true) {
			return
		}
	} else if len(departments) == 1 && order == departments[0].Order {
		val.Note = "final"
		err = services.DepartmentService.UpdateInfo(val.DepartmentId, val.DepartmentName, val.Introduction,
			val.Order+1, val.Note)
		if logs.SqlError(err, w, true) {
			return
		}
	} else {
		for _, v := range departments {
			if v.Order >= order {
				err = services.DepartmentService.UpdateInfo(v.DepartmentId, v.DepartmentName, v.Introduction,
					v.Order+1, v.Note)
				if logs.SqlError(err, w, true) {
					return
				}
			}
		}
	}

	var note string
	if len(departments) == 0 {
		note = "initial,final"
	} else if order > len(departments) {
		note = "final"
	} else if order == 1 {
		note = "initial"
	} else {
		note = "middle"
	}

	dep := services.DepartmentService.NewDeparment(name, introduction, order, note)
	err = services.DepartmentService.SaveAInfo(dep)
	if logs.SqlError(err, w, true) {
		return
	}
	w.WriteHeader(200)
}

/*
* Get deparment by id
 */
func getDepartmentById(w http.ResponseWriter, r *http.Request) {
	id := getIdFromPath(r)

	department, err := services.DepartmentService.FindInfoById(id)
	if logs.SqlError(err, w, department.DepartmentName != "") {
		return
	}

	data, _ := json.Marshal(department)
	w.WriteHeader(200)
	w.Write(data)
}

/*
* Update department by Id
* The result is exchange the order of the two departments
 */
func updateDepartmentById(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	js, err := simplejson.NewFromReader(r.Body)
	if logs.NormalError(err, w) {
		return
	}

	id := getIdFromPath(r)
	name := js.Get("DepartmentName").MustString()
	intro := js.Get("Introduction").MustString()
	order := js.Get("Order").MustInt()
	if name == "" || intro == "" || order == 0 {
		logs.RequestError(500, logs.ErrorMsg{Msg: "必填字段不能为空"}, w)
		return
	}

	// Get department by id and order.
	departmentUpdate, err := services.DepartmentService.FindInfoById(id)
	if logs.SqlError(err, w, departmentUpdate.DepartmentName != "") {
		return
	}
	departmentNoUpdate, err := services.DepartmentService.FindInfoByOrder(order)
	if logs.SqlError(err, w, departmentNoUpdate.DepartmentName != "") {
		return
	}

	// If order no update, then update other information.
	if id == departmentNoUpdate.DepartmentId {
		err = services.DepartmentService.UpdateInfo(id, name, intro, order, departmentUpdate.Note)
		if logs.SqlError(err, w, true) {
			return
		}
	} else { // If order update, then change the two information.
		err = services.DepartmentService.UpdateInfo(id, name, intro, order, departmentUpdate.Note)
		if logs.SqlError(err, w, true) {
			return
		}

		err = services.DepartmentService.UpdateInfo(departmentNoUpdate.DepartmentId, departmentNoUpdate.DepartmentName, departmentNoUpdate.Introduction,
			departmentUpdate.Order, departmentNoUpdate.Note)
		if logs.SqlError(err, w, true) {
			return
		}
	}

	w.WriteHeader(200)
}

/*
* Delete department by Id
* The result will make some departments information change
 */
func deleteDepartmentById(w http.ResponseWriter, r *http.Request) {

	id := getIdFromPath(r)
	del_dep, err := services.DepartmentService.FindInfoById(id)
	if logs.SqlError(err, w, del_dep.DepartmentName != "") {
		return
	}
	if del_dep.Note == "initial,final" {
		logs.RequestError(500, logs.ErrorMsg{Msg: "系统正常运行，至少需要保留一个部门"}, w)
		return
	}

	// Update order
	departments, err := services.DepartmentService.FindAllInfo()
	if logs.SqlError(err, w, len(departments) != 0) {
		return
	}
	order := del_dep.Order
	isUPdateNoteInitial := false
	isUPdateNoteFinal := false
	if del_dep.Note == "initial" && len(departments) == 2 {
		departments[1].Note = "initial,final"
	} else if del_dep.Note == "final" && len(departments) == 2 {
		departments[0].Note = "initial,final"
	} else if del_dep.Note == "initial" {
		isUPdateNoteInitial = true
	} else if del_dep.Note == "final" {
		isUPdateNoteFinal = true
	}

	for _, v := range departments {
		if isUPdateNoteFinal && v.Order == (order-1) {
			err = services.DepartmentService.UpdateInfo(v.DepartmentId, v.DepartmentName, v.Introduction, v.Order,
				"final")
			if logs.SqlError(err, w, true) {
				return
			}
		}
		if v.Order > order {
			if isUPdateNoteInitial && v.Order == (order+1) {
				v.Note = "initial"
			}
			err = services.DepartmentService.UpdateInfo(v.DepartmentId, v.DepartmentName, v.Introduction, v.Order-1,
				v.Note)
			if logs.SqlError(err, w, true) {
				return
			}
		}
	}
	// Delete the department.
	err = services.DepartmentService.DeleteInfo(id)
	if logs.SqlError(err, w, true) {
		return
	}

	//Define some variable for updating reservatons
	new_dep, err := services.DepartmentService.FindInfoByOrder(del_dep.Order)
	if logs.SqlError(err, w, new_dep.DepartmentName != "") {
		return
	}
	new_approvers, err := services.ApproverService.FindInfoByDepartmentId(new_dep.DepartmentId)
	if logs.SqlError(err, w, len(new_approvers) != 0) {
		return
	}

	//Delete the approvers which belong to the department
	approvers, err := services.ApproverService.FindInfoByDepartmentId(id)
	if logs.SqlError(err, w, true) {
		return
	}
	for _, v := range approvers {
		err := services.ApproverService.DeleteInfo(v.ApproverId)
		if logs.SqlError(err, w, true) {
			return
		}

		//Update reservations.
		reservations, err := services.ReservationService.FindInfoByApproverId(v.ApproverId)
		if logs.SqlError(err, w, true) {
			return
		}
		for _, res := range reservations {
			res.ApproverId = approvers[rand.Intn(len(new_approvers))].ApproverId
			err := services.ReservationService.UpdateInfo(res.ResId, res.ResReason, res.OrganizationName, res.ResState, res.ApprovalNote, res.ApproverId)
			if logs.SqlError(err, w, true) {
				return
			}
		}
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
