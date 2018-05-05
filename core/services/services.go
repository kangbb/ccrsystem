package services

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
	"github.com/gorilla/sessions"
	"github.com/kangbb/ccrsystem/core/models/service"
)

var store = sessions.NewFilesystemStore("./session", []byte("ccrsystem"))

/******************************** Render File *************************/
func RenderFile(name string, w http.ResponseWriter, r *http.Request) {
	filepath := "views/" + name + ".html"
	session, _ := store.Get(r, "user")

	switch name {
	case "index":
		break
	case "studentIndex":
		if !sessionExist("Student", session, w, r) {
			return
		}
		break
	case "adminIndex":
		if !sessionExist("Admin", session, w, r) {
			return
		}
		break
	default:
		if !sessionExist("Approver", session, w, r) {
			return
		}
	}

	t := template.Must(template.ParseFiles(filepath))
	w.WriteHeader(200)
	t.Execute(w, nil)
}

/****************************** Signin and Signout *************************/
/*
* User Signin
* For student, approver, admin. If sucess, redirect to the index which belong to themselves
* If failed, redirect to the home page.
 */
func Signin(userType string, w http.ResponseWriter, r *http.Request) {
	var (
		id  int
		pwd string
	)
	store.MaxAge(86400)
	session, _ := store.Get(r, "user")

	// UserTypeEnum = []string{"Student", "Admin", "Approver"}
	defer r.Body.Close()
	js, err := simplejson.NewFromReader(r.Body)
	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
		return
	}

	id = js.Get(userType + "Id").MustInt()
	pwd = js.Get(userType + "Pwd").MustString()
	switch userType {
	case "Student":
		usr := service.StudentService.FindInfoById(id)
		if usr != nil && usr.StudentPwd == encryptPwd(id, pwd) {
			session.Values["name"] = usr.StudentName
		} else {
			type Std struct {
				StudentId  string
				StudentPwd string
			}
			w.WriteHeader(500)
			data, _ := json.Marshal(Std{"此用户不存在", ""})
			w.Write(data)
			return
		}
		break
	case "Admin":
		usr := service.AdminService.FindInfoById(id)
		if usr != nil && usr.AdminPwd == encryptPwd(id, pwd) {
			session.Values["name"] = usr.AdminName
		} else {
			type Adm struct {
				AdminId  string
				AdminPwd string
			}
			w.WriteHeader(500)
			data, _ := json.Marshal(Adm{"此用户不存在", ""})
			w.Write(data)
			return
		}
		break
	default:
		usr := service.ApproverService.FindInfoById(id)
		if usr != nil && usr.ApproverPwd == encryptPwd(id, pwd) {
			session.Values["name"] = usr.ApproverName
		} else {
			type Appr struct {
				ApproverId  string
				ApproverPwd string
			}
			w.WriteHeader(500)
			data, _ := json.Marshal(Appr{"此用户不存在", ""})
			w.Write(data)
			return
		}
	}

	session.Values["id"] = id
	session.Values["type"] = userType
	session.Values["accessTime"] = time.Now()
	err = session.Save(r, w)
	if err != nil {
		fmt.Println(err)
	}
	w.WriteHeader(200)
}

/*
* User Signout
* Just delete the sessions.
 */
func Signout(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "user")
	session.Options.MaxAge = -1
	err := session.Save(r, w)
	if err != nil {
		fmt.Println(err)
	}

	req, _ := http.NewRequest("GET", "/", r.Body)
	http.Redirect(w, req, "/", 302)
}

/************************* Data Interface For User ***********************/
/*
* Get userinfo by session
 */
func GetUserInfo(userType string, w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "user")
	if !sessionExist(userType, session, w, r) {
		return
	}
	w.WriteHeader(200)

	switch userType {
	case "Student":
		user := service.StudentService.FindInfoById(session.Values["id"].(int))
		user.StudentPwd = ""
		data, _ := json.Marshal(user)
		w.Write(data)
		break
	case "Approver":
		user := service.ApproverService.FindInfoById(session.Values["id"].(int))
		user.ApproverPwd = ""
		data, _ := json.Marshal(user)
		w.Write(data)
		break
	default:
		user := service.AdminService.FindInfoById(session.Values["id"].(int))
		user.AdminPwd = ""
		data, _ := json.Marshal(user)
		w.Write(data)
	}
}

/*
* Update userinfo by session
* Just for password
 */
func UpdateUserInfo(userType string, w http.ResponseWriter, r *http.Request) {
	var (
		status bool
		err    error
	)
	session, _ := store.Get(r, "user")
	if !sessionExist(userType, session, w, r) {
		return
	}

	//default use the parameter from body
	defer r.Body.Close()
	js, err := simplejson.NewFromReader(r.Body)
	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
		return
	}

	id := session.Values["id"].(int)
	switch userType {
	case "Student":
		pwd := encryptPwd(id, js.Get("StudentPwd").MustString())
		status, err = service.StudentService.UpdateInfo(session.Values["id"].(int), pwd)
		break
	case "Approver":
		pwd := encryptPwd(id, js.Get("ApproverPwd").MustString())
		status, err = service.ApproverService.UpdatePasswordInfo(session.Values["id"].(int), pwd)
		break
	default:
		pwd := encryptPwd(id, js.Get("AdminPwd").MustString())
		status, err = service.AdminService.UpdateInfo(session.Values["id"].(int), pwd)
	}

	if !status || err != nil {
		w.WriteHeader(500)
	} else {
		w.WriteHeader(200)
	}
}

/*
* Get userlist
* Just accessed for the admin, the password will not return
 */
func GetUserList(userType string, w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "user")
	if !sessionExist("Admin", session, w, r) {
		return
	}
	w.WriteHeader(200)

	switch userType {
	case "Student":
		user := service.StudentService.FindAllInfo()
		for k, _ := range user {
			user[k].StudentPwd = ""
		}
		data, _ := json.Marshal(user)
		w.Write(data)
		break
	case "Approver":
		user := service.ApproverService.FindAllInfo()
		for k, _ := range user {
			user[k].ApproverPwd = ""
		}

		data, _ := json.Marshal(user)
		w.Write(data)
		break
	default:
		user := service.AdminService.FindAllInfo()
		for k, _ := range user {
			user[k].AdminPwd = ""
		}
		data, _ := json.Marshal(user)
		w.Write(data)
	}

}

/*
* Add a user to db
 */
func AddUser(userType string, w http.ResponseWriter, r *http.Request) {
	var (
		state bool
		err   error
	)

	session, _ := store.Get(r, "user")
	if !sessionExist("Admin", session, w, r) {
		return
	}

	//default paremeter from r.body
	defer r.Body.Close()
	js, err := simplejson.NewFromReader(r.Body)
	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
		return
	}

	switch userType {
	case "Student":
		id := js.Get("StudentId").MustInt()
		pwd := encryptPwd(id, js.Get("StudentPwd").MustString())
		user := service.StudentService.NewStudent(id, pwd, js.Get("StudentName").MustString(),
			js.Get("Permision").MustBool())
		state, err = service.StudentService.SaveAInfo(&user)
		break
	case "Approver":
		id := js.Get("ApproverId").MustInt()
		pwd := encryptPwd(id, js.Get("ApproverPwd").MustString())
		user := service.ApproverService.NewApprover(id, pwd, js.Get("ApproverName").MustString(),
			js.Get("DepartmentId").MustInt(), js.Get("Permision").MustBool())
		state, err = service.ApproverService.SaveAInfo(&user)
		break
	case "Admin":
		id := js.Get("AdminId").MustInt()
		pwd := encryptPwd(id, js.Get("AdminPwd").MustString())
		user := service.AdminService.NewAdmin(id, pwd, js.Get("AdminName").MustString(),
			js.Get("Permision").MustBool())
		state, err = service.AdminService.SaveAInfo(&user)
		break
	}

	if !state || err != nil {
		w.WriteHeader(500)
	} else {
		w.WriteHeader(200)
	}
}

/*
* Get userinfo by id
* Just acessed for admin
 */
func GetUserById(userType string, w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "user")
	if !sessionExist("Admin", session, w, r) {
		return
	}

	//default use the parameter from r.body
	defer r.Body.Close()
	js, err := simplejson.NewFromReader(r.Body)
	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
		return
	}

	w.WriteHeader(200)

	switch userType {
	case "Student":
		user := service.StudentService.FindInfoById(js.Get("StudentId").MustInt())
		user.StudentPwd = ""
		data, _ := json.Marshal(user)
		w.Write(data)
		break
	case "Approver":
		user := service.ApproverService.FindInfoById(js.Get("ApproverId").MustInt())
		user.ApproverPwd = ""
		data, _ := json.Marshal(user)
		w.Write(data)
		break
	default:
		user := service.AdminService.FindInfoById(js.Get("AdminId").MustInt())
		user.AdminPwd = ""
		data, _ := json.Marshal(user)
		w.Write(data)
	}
}

/*
* Update userinfo by id
* just acessed for admin
 */
func UpdateUserById(userType string, w http.ResponseWriter, r *http.Request) {
	var (
		status bool
		err    error
	)
	session, _ := store.Get(r, "user")
	if !sessionExist("Admin", session, w, r) {
		return
	}

	//default use the parameter from body
	defer r.Body.Close()
	js, err := simplejson.NewFromReader(r.Body)
	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
		return
	}

	switch userType {
	case "Student":
		id := js.Get("StudentId").MustInt()
		pwd := js.Get("StudentPwd").MustString()
		status, err = service.StudentService.UpdateInfo(id, pwd)
		break
	case "Approver":
		id := js.Get("ApproverId").MustInt()
		pwd := js.Get("ApproverPwd").MustString()
		status, err = service.ApproverService.UpdatePasswordInfo(id, pwd)
		break
	default:
		id := js.Get("AdminId").MustInt()
		pwd := js.Get("AdminPwd").MustString()
		status, err = service.AdminService.UpdateInfo(id, pwd)
	}

	if !status || err != nil {
		w.WriteHeader(500)
	} else {
		w.WriteHeader(200)
	}
}

/*
* Delete userinfo by id
* Just acessed for admin
 */
func DeleteUserById(userType string, w http.ResponseWriter, r *http.Request) {
	var (
		status bool
		err    error
	)

	session, _ := store.Get(r, "user")
	if !sessionExist(userType, session, w, r) {
		return
	}

	//default use the parameter from r.body
	defer r.Body.Close()
	js, err := simplejson.NewFromReader(r.Body)
	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
		return
	}

	switch userType {
	case "Student":
		status, err = service.StudentService.DeleteInfo(js.Get("StudentId").MustInt())
		break
	case "Approver":
		status, err = service.ApproverService.DeleteInfo(js.Get("ApproverId").MustInt())
		break
	default:
		//admin can't delete count which belongs to himself
		if js.Get("AdminId").MustInt() == session.Values["id"].(int) {
			w.WriteHeader(403)
			w.Write([]byte("<h1>Forbidden</h1>"))
			return
		}
		status, err = service.AdminService.DeleteInfo(js.Get("AdminId").MustInt())
	}

	if !status || err != nil {
		w.WriteHeader(500)
	} else {
		w.WriteHeader(200)
	}
}

/************************* Data Interface For Classroom ***********************/
/*
* Get classroomlist
* Just accessed for Admin
 */
func GetClassroomList(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "user")
	if !sessionExist("Admin", session, w, r) {
		return
	}

	classrooms := service.ClassroomService.FindAllInfo()
	data, _ := json.Marshal(classrooms)

	w.WriteHeader(200)
	w.Write(data)
}

/*
* Add a classroom information to db
 */
func AddClassroom(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "user")
	if !sessionExist("Admin", session, w, r) {
		return
	}

	//default use the parameter from r.body
	defer r.Body.Close()
	js, err := simplejson.NewFromReader(r.Body)
	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
		return
	}

	classroomCampus := js.Get("ClassroomCampus").MustString()
	classroomBuilding := js.Get("ClassroomBuilding").MustString()
	classroomNum := js.Get("ClassroomNum").MustString()
	cap := js.Get("Capicity").MustInt()
	if classroomCampus == "" || classroomBuilding == "" || classroomNum == "" || cap == 0 {
		w.WriteHeader(500)
		w.Write([]byte("必填项不能为空"))
		return
	}
	classroom := service.ClassroomService.NewClassroom(classroomCampus, classroomBuilding, classroomNum, cap)
	state, err := service.ClassroomService.SaveAInfo(&classroom)

	if !state || err == nil {
		w.WriteHeader(500)
	} else {
		w.WriteHeader(200)
	}
}

/*
* Query classroom by some conditions
* Just accessed for student
 */
func QueryClassroom(w http.ResponseWriter, r *http.Request) {
	type reservation struct {
		ResId     int
		StartTime string
		EndTime   string
		ResState  string
		ResReason string
	}
	type queryResult struct {
		ClassroomId       int
		ClassroomCampus   string
		ClassroomBuilding string
		Capicity          int
		Res               map[int]reservation
	}

	session, _ := store.Get(r, "user")
	if !sessionExist("Student", session, w, r) {
		return
	}

	//default use the parameter from r.body
	defer r.Body.Close()
	js, err := simplejson.NewFromReader(r.Body)
	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
		return
	}
	index := 1
	classroomInfo := make(map[int]interface{})
	classroomInfo[0] = js.Get("ClassroomCampus").MustString()
	classroomInfo[1] = js.Get("ClassroomBuilding").MustString()
	if tmp := js.Get("Capicity").MustInt(); tmp != 0 {
		index += 1
		classroomInfo[index] = tmp
	}

	//query classrooms which meet the require
	classrooms := service.ClassroomService.NewClassroomSlice()
	switch index {
	case 1:
		classrooms = service.ClassroomService.GetClassroomBySomeCond(classroomInfo[0],
			classroomInfo[1])
		break
	case 2:
		classrooms = service.ClassroomService.GetClassroomBySomeCond(classroomInfo[0],
			classroomInfo[1], classroomInfo[2])
	}

	result := make([]queryResult, 0)
	for k, v := range classrooms {
		result[k] = queryResult{
			ClassroomId:       v.ClassroomId,
			ClassroomCampus:   v.ClassroomCampus,
			ClassroomBuilding: v.ClassroomBuilding,
			Capicity:          v.Capicity,
			Res:               make(map[int]reservation),
		}
		index = -1
		resInfo := make(map[int]interface{})
		if tmp := js.Get("StartTime").MustString(); tmp != "" {
			index += 1
			resInfo[index] = lessonToTime(tmp)
		}
		if tmp := js.Get("EndTime").MustString(); tmp != "" {
			index += 1
			resInfo[index] = lessonToTime(tmp)
		}

		//query the state of the each classroom
		res := service.ReservationService.NewReservationSlice()
		switch index {
		case -1:
			res = service.ReservationService.GetReservationBySomeCond(v.ClassroomId, time.Now())
		case 0:
			res = service.ReservationService.GetReservationBySomeCond(v.ClassroomId, lessonToTime(resInfo[0].(string)))
		case 1:
			res = service.ReservationService.GetReservationBySomeCond(v.ClassroomId, lessonToTime(resInfo[0].(string)),
				lessonToTime(resInfo[1].(string)))
		}

		for i, val := range res {
			result[k].Res[i] = reservation{
				ResId:     val.ResId,
				StartTime: timeToLesson(val.StartTime),
				EndTime:   timeToLesson(val.EndTime),
				ResState:  val.ResState,
				ResReason: val.ResReason,
			}
		}
	}

	data, _ := json.Marshal(result)

	w.WriteHeader(200)
	w.Write(data)
}

/*
* Get classroom by id
 */
func GetClassroomById(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "user")
	if !sessionExist("Admin", session, w, r) {
		return
	}
	//default use the parameter from r.body
	defer r.Body.Close()
	js, err := simplejson.NewFromReader(r.Body)
	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
		return
	}

	classroom := service.ClassroomService.FindInfoById(js.Get("ClassroomId").MustInt())
	data, _ := json.Marshal(classroom)

	w.WriteHeader(200)
	w.Write(data)
}

/*
* Update classroom by id
 */
func UpdateClassroomById(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "user")
	if !sessionExist("Admin", session, w, r) {
		return
	}
	//default use the parameter from r.body
	defer r.Body.Close()
	js, err := simplejson.NewFromReader(r.Body)
	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
		return
	}
	//default return all the parameter for the classroom
	//include the parameters which don't need to update
	id := js.Get("ClassroomId").MustInt()
	cap := js.Get("Capicity").MustInt()
	num := js.Get("ClassroomNum").MustString()
	building := js.Get("ClassroomBuilding").MustString()
	campus := js.Get("ClassroomCampus").MustString()
	status, err := service.ClassroomService.UpdateInfo(id, cap, num, building, campus)

	if !status || err != nil {
		w.WriteHeader(500)
	} else {
		w.WriteHeader(200)
	}
}

/*
* Delete classroom by id
 */
func DeleteClassroomById(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "user")
	if !sessionExist("Admin", session, w, r) {
		return
	}

	//default use the parameter from r.body
	defer r.Body.Close()
	js, err := simplejson.NewFromReader(r.Body)
	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
		return
	}

	status, err := service.ClassroomService.DeleteInfo(js.Get("ClassroomId").MustInt())

	if !status || err != nil {
		w.WriteHeader(500)
	} else {
		w.WriteHeader(200)
	}
}

/************************* Data Interface For Reservation ***********************/
/*
* Get reservation by id
 */
func GetResById(w http.ResponseWriter, r *http.Request) {
	type Reservation struct {
		ResId        int
		ResState     string
		StartTime    string
		EndTime      string
		ResReason    string
		ApprovalNote string
		DepartmentId int
		StudentId    int
		ApproverId   int
		ClassroomId  int
	}

	session, _ := store.Get(r, "user")
	if !sessionExist(session.Values["type"].(string), session, w, r) {
		return
	}

	//default use the parameter from r.body
	defer r.Body.Close()
	js, err := simplejson.NewFromReader(r.Body)
	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
		return
	}
	_, res := service.ReservationService.FindInfoById(js.Get("ResId").MustInt())

	result := Reservation{}
	// need not to load the reservation ResState = "预定失败"
	if session.Values["type"].(string) == "Approver" && res.ResState == "预定失败" {
		return
	}

	result = Reservation{
		ResId:        res.ResId,
		ResState:     res.ResState,
		StartTime:    timeToLesson(res.StartTime),
		EndTime:      timeToLesson(res.EndTime),
		ResReason:    res.ResReason,
		ApprovalNote: res.ApprovalNote,
		DepartmentId: res.DepartmentId,
		StudentId:    res.StudentId,
		ApproverId:   res.ApproverId,
		ClassroomId:  res.ClassroomId,
	}
	data, _ := json.Marshal(result)

	w.WriteHeader(200)
	w.Write(data)
}

/*
* Update reservation by id
 */
func UpdateResById(w http.ResponseWriter, r *http.Request) {
	var (
		id           int
		departmentId int
		approverId   int
		state        string
		approvalNote string
	)

	session, _ := store.Get(r, "user")
	if !sessionExist(session.Values["type"].(string), session, w, r) {
		return
	}

	//default use the parameter from r.body
	defer r.Body.Close()
	js, err := simplejson.NewFromReader(r.Body)
	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
		return
	}

	id = js.Get("ResId").MustInt()
	//student update reservation
	if session.Values["type"].(string) == "Student" {
		var reason string
		if tmp := js.Get("ResReason").MustString(); tmp != "" {
			reason = tmp
		} else {
			w.WriteHeader(500)
			w.Write([]byte("该项不能为空"))
		}
		department := service.DepartmentService.FindInfoByNote("initial")
		departmentId = department.DepartmentId
		approvers := service.ApproverService.FindInfoByDepartmentId(departmentId)
		//random select a approver from the approvers who belong to the same department
		approverId = approvers[rand.Intn(len(approvers))].ApproverId
		state = "审批中"
		approvalNote = ""
		state, err := service.ReservationService.UpdateInfo(id, approverId, departmentId,
			approvalNote, state, reason)
		if !state || err != nil {
			w.WriteHeader(500)
		} else {
			w.WriteHeader(200)
		} //approver approve the reservation
	} else if session.Values["type"].(string) == "Approver" {
		//check if reservation exist
		//if not exist, return
		has, _ := service.ReservationService.FindInfoById(id)
		if !has {
			w.WriteHeader(200)
			return
		}
		//if exist, update resinfo
		state = js.Get("ResState").MustString()
		departmentId = js.Get("DepartmentId").MustInt()
		approverId = js.Get("ApproverId").MustInt()
		approvalNote = js.Get("ApproverNote").MustString()
		if state == "等待下一个部门审批" {
			department := service.DepartmentService.FindInfoById(departmentId)
			if department.Note == "final" {
				state = "预订成功"
			} else {
				state = "审批中"
				department := service.DepartmentService.FindInfoById(departmentId)
				department = service.DepartmentService.FindInfoByOrder(department.Order + 1)
				departmentId = department.DepartmentId
				approvers := service.ApproverService.FindInfoByDepartmentId(departmentId)
				//random select a approver from the approvers who belong to the same department
				approverId = approvers[rand.Intn(len(approvers))].ApproverId
				approvalNote = ""
			}
		} else if state == "审批未通过" {
			state = "预订失败"
		}

		state, err := service.ReservationService.UpdateInfo(id, approverId, departmentId,
			approvalNote, state)

		if !state || err != nil {
			w.WriteHeader(500)
		} else {
			w.WriteHeader(200)
		}
	} else {
		w.WriteHeader(403)
		w.Write([]byte("<h1>Forbidden</h1>"))
	}
}

/*
* Delete reservation by id
 */
func DeleteResById(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "user")
	if !sessionExist("Student", session, w, r) {
		return
	}

	defer r.Body.Close()
	js, err := simplejson.NewFromReader(r.Body)
	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
		return
	}
	state, err := service.ReservationService.DeleteInfo(js.Get("ResId").MustInt())

	if !state || err != nil {
		w.WriteHeader(500)
	} else {
		w.WriteHeader(200)
	}
}

/*
* Get a student reservation list
 */
func GetStudentResList(w http.ResponseWriter, r *http.Request) {
	type Reservation struct {
		ResId        int
		ResState     string
		StartTime    string
		EndTime      string
		ResReason    string
		ApprovalNote string
		DepartmentId int
		StudentId    int
		ApproverId   int
		ClassroomId  int
	}

	session, _ := store.Get(r, "user")
	if !sessionExist("Student", session, w, r) {
		return
	}

	result := make([]Reservation, 0)
	reservations := service.ReservationService.FindInfoByStudentId(session.Values["id"].(int))
	for k, v := range reservations {
		result[k] = Reservation{
			ResId:        v.ResId,
			ResState:     v.ResState,
			StartTime:    timeToLesson(v.StartTime),
			EndTime:      timeToLesson(v.EndTime),
			ResReason:    v.ResReason,
			ApprovalNote: v.ApprovalNote,
			DepartmentId: v.DepartmentId,
			StudentId:    v.StudentId,
			ApproverId:   v.ApproverId,
			ClassroomId:  v.ClassroomId,
		}
	}
	data, _ := json.Marshal(result)

	w.WriteHeader(200)
	w.Write(data)
}

/*
* Get a approver reservation list
 */
func GetApproverResList(w http.ResponseWriter, r *http.Request) {
	type Reservation struct {
		ResId        int
		ResState     string
		StartTime    string
		EndTime      string
		ResReason    string
		ApprovalNote string
		DepartmentId int
		StudentId    int
		ApproverId   int
		ClassroomId  int
	}
	result := make([]Reservation, 0)

	session, _ := store.Get(r, "user")
	if !sessionExist("Approver", session, w, r) {
		return
	}

	reservations := service.ReservationService.FindInfoByApproverId(session.Values["id"].(int))
	for k, v := range reservations {
		if v.ResState == "预定失败" {
			continue
		}
		result[k] = Reservation{
			ResId:        v.ResId,
			ResState:     v.ResState,
			StartTime:    timeToLesson(v.StartTime),
			EndTime:      timeToLesson(v.EndTime),
			ResReason:    v.ResReason,
			ApprovalNote: v.ApprovalNote,
			DepartmentId: v.DepartmentId,
			StudentId:    v.StudentId,
			ApproverId:   v.ApproverId,
			ClassroomId:  v.ClassroomId,
		}
	}
	data, _ := json.Marshal(result)

	w.WriteHeader(200)
	w.Write(data)
}

/*
* Add a reservation to db
 */
func AddRes(w http.ResponseWriter, r *http.Request) {
	var (
		start       time.Time
		end         time.Time
		reason      string
		classroomId int
	)
	session, _ := store.Get(r, "user")
	if !sessionExist("Student", session, w, r) {
		return
	}

	//default use the parameter from r.body
	defer r.Body.Close()
	js, err := simplejson.NewFromReader(r.Body)
	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
		return
	}

	state := "审批中"
	if tmp := js.Get("StartTime").MustString(); tmp != "" {
		start = lessonToTime(tmp)
	} else {
		w.WriteHeader(500)
		w.Write([]byte("开始时间不能为空"))
		return
	}
	if tmp := js.Get("EndTime").MustString(); tmp != "" {
		end = lessonToTime(tmp)
	} else {
		w.WriteHeader(500)
		w.Write([]byte("结束时间不能为空"))
		return
	}
	if tmp := js.Get("ResReason").MustString(); tmp != "" {
		reason = tmp
	} else {
		w.WriteHeader(500)
		w.Write([]byte("申请原因不能为空"))
		return
	}
	if tmp := js.Get("ClassroomId").MustInt(); tmp != 0 {
		classroomId = tmp
	} else {
		w.WriteHeader(500)
		w.Write([]byte("申请教室不能为空"))
		return
	}

	approvalNote := ""

	department := service.DepartmentService.FindInfoByNote("initial")
	departmentId := department.DepartmentId

	studentId := session.Values["StudentId"].(int)

	approvers := service.ApproverService.FindInfoByDepartmentId(departmentId)
	//random select a approver from the approvers who belong to the same department
	approverId := approvers[rand.Intn(len(approvers))].ApproverId

	res := service.ReservationService.NewReservation(state, start, end, departmentId, reason,
		approvalNote, studentId, approverId, classroomId)
	result, err := service.ReservationService.SaveAInfo(&res)

	if !result || err != nil {
		w.WriteHeader(500)
	} else {
		w.WriteHeader(200)
	}
}

/************************* Data Interface For Department ***********************/
/*
* Get department list
 */
func GetDepartmentList(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "user")
	if !sessionExist("Admin", session, w, r) {
		return
	}

	departments := service.DepartmentService.FindAllInfo()
	data, _ := json.Marshal(departments)

	w.WriteHeader(200)
	w.Write(data)
}

/*
* Add a department information to db
 */
func AddDepartment(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "user")
	if !sessionExist("Admin", session, w, r) {
		return
	}

	defer r.Body.Close()
	js, err := simplejson.NewFromReader(r.Body)
	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
		return
	}

	name := js.Get("DepartmentName").MustString()
	order := js.Get("Order").MustInt()
	if name == "" || order == 0 {
		w.WriteHeader(500)
		w.Write([]byte("必填项不能为空"))
	}

	//update the order of the approval department
	departments := service.DepartmentService.FindAllInfo()
	for _, v := range departments {
		if v.Order >= order {
			state, err := service.DepartmentService.UpdateInfo(v.DepartmentId, v.DepartmentName,
				v.Order+1, v.Note)
			if !state || err != nil {
				fmt.Println(err)
				w.WriteHeader(500)
			} else {
				w.WriteHeader(200)
			}
		}
	}

	var note string
	if order > len(departments) {
		note = "final"
	} else if order == 1 {
		note = "initial"
	} else {
		note = "middle"
	}

	dep := service.DepartmentService.NewDeparment(name, order, note)
	state, err := service.DepartmentService.SaveAInfo(&dep)
	if !state || err != nil {
		w.WriteHeader(500)
	} else {
		w.WriteHeader(200)
	}

}

/*
* Get deparment by id
 */
func GetDepartmentById(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "user")
	if !sessionExist("Admin", session, w, r) {
		return
	}

	defer r.Body.Close()
	js, err := simplejson.NewFromReader(r.Body)
	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
		return
	}

	department := service.DepartmentService.FindInfoById(js.Get("DepartmentId").MustInt())
	data, _ := json.Marshal(department)

	w.WriteHeader(200)
	w.Write(data)
}

/*
* Update department by Id
* The result is exchange the order of the two departments
 */
func UpdateDepartmentById(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "user")
	if !sessionExist("Admin", session, w, r) {
		return
	}

	defer r.Body.Close()
	js, err := simplejson.NewFromReader(r.Body)
	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
		return
	}

	id := js.Get("DepartmentId").MustInt()
	name := js.Get("DepartmentName").MustString()
	order := js.Get("Order").MustInt()
	department_now := service.DepartmentService.FindInfoById(id)
	department := service.DepartmentService.FindInfoByOrder(order)
	if id == department.DepartmentId {
		state, err := service.DepartmentService.UpdateInfo(id, name, order, department.Note)
		if !state || err != nil {
			w.WriteHeader(500)
		}
	} else {
		state, err := service.DepartmentService.UpdateInfo(id, name, order, department.Note)
		if !state || err != nil {
			w.WriteHeader(500)
		}
		state, err = service.DepartmentService.UpdateInfo(department.DepartmentId, department.DepartmentName,
			department_now.Order, department_now.Note)
		if !state || err != nil {
			w.WriteHeader(500)
		}
	}

	w.WriteHeader(200)
}

/*
* Delete department by Id
* The result will make some departments information change
 */
func DeleteDepartmentById(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "user")
	if !sessionExist("Admin", session, w, r) {
		return
	}

	defer r.Body.Close()
	js, err := simplejson.NewFromReader(r.Body)
	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
		return
	}
	id := js.Get("DepartmentId").MustInt()
	del_dep := service.DepartmentService.FindInfoById(id)
	departments := service.DepartmentService.FindAllInfo()
	order := len(departments)
	for _, v := range departments {
		if v.DepartmentId == id {
			state, err := service.DepartmentService.DeleteInfo(id)
			if !state || err != nil {
				w.WriteHeader(500)
				break
			}
			order = v.Order
		}
		if v.Order > order {
			state, err := service.DepartmentService.UpdateInfo(v.DepartmentId, v.DepartmentName, order,
				v.Note)
			if !state || err != nil {
				w.WriteHeader(500)
				break
			}
			order += 1
		}
	}

	//delete the approvers which belong to the department
	approvers := service.ApproverService.FindInfoByDepartmentId(id)
	for _, v := range approvers {
		state, err := service.ApproverService.DeleteInfo(v.ApproverId)
		if !state || err != nil {
			fmt.Println(err)
			w.WriteHeader(500)
		}
	}

	//update reservatons
	new_dep := service.DepartmentService.FindInfoByOrder(del_dep.Order)
	approvers = service.ApproverService.FindInfoByDepartmentId(new_dep.DepartmentId)
	reservations := service.ReservationService.FindInfoByDepartmentId(id)
	for _, v := range reservations {
		v.ApproverId = approvers[rand.Intn(len(approvers))].ApproverId
		v.DepartmentId = new_dep.DepartmentId
		state, err := service.ReservationService.UpdateInfo(v.ResId, v.ApproverId, v.DepartmentId)
		if !state || err != nil {
			w.WriteHeader(500)
		}
	}

	w.WriteHeader(200)
}

/***********************************SOME SUBFUNCTION*******************************/
func lessonToTime(lessonTime string) time.Time {
	mappingRegular := map[string]int{
		"一": 1, "二": 2, "三": 3, "四": 4, "五": 5, "六": 6, "七": 7, "八": 8, "九": 9, "十": 1,
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
		lesson = mappingRegular[string(hanstr[1])]*10 + mappingRegular[string(hanstr[2])]
		break
	}
	startTime, _ := time.Parse("2006-01-02 15:04:05", date+" "+"08:00:00")
	duraTime := strconv.Itoa((lesson-1)*(45+10)) + "m"
	addTime, _ := time.ParseDuration(duraTime)
	return startTime.Add(addTime)
}

func timeToLesson(trueTime time.Time) string {
	mappingRegular := []string{"一", "二", "三", "四", "五", "六", "七", "八", "九", "十"}
	var lesson string
	date := trueTime.Format("2006-01-02")
	startTime, _ := time.Parse("2006-01-02 15:04:05", date+" "+"08:00:00")
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

func sessionExist(userType string, session *sessions.Session, w http.ResponseWriter, r *http.Request) bool {
	if session.Values["id"] == nil || session.Values["type"] == nil || session.Values["accessTime"] == nil {
		req, _ := http.NewRequest("GET", "/", r.Body)
		http.Redirect(w, req, "/", 302)
		return false
	}
	//session doesn't exist
	//session expired
	now := time.Now()
	sub := now.Sub(session.Values["accessTime"].(time.Time)).Seconds()
	if int(sub)-session.Options.MaxAge >= 0 || session.Values["type"] == nil {
		session.Options.MaxAge = -1
		session.Save(r, w)
		//redirect the utl to login layout
		req, _ := http.NewRequest("GET", "/", r.Body)
		http.Redirect(w, req, "/", 302)
		return false
	}
	//have no permission accessed
	if userType != session.Values["type"].(string) {
		w.WriteHeader(403)
		w.Write([]byte("<h1>Forbidden</h1>"))
		return false
	}
	session.Values["accessTime"] = time.Now()
	err := session.Save(r, w)
	if err != nil {
		fmt.Println(err)
	}
	return true
}
