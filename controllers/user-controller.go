package controllers

import (
	"encoding/json"
	"fmt"
	"html/template"
	"math/rand"
	"net/http"
	"strconv"

	simplejson "github.com/bitly/go-simplejson"
	"github.com/kangbb/ccrsystem/logs"
	"github.com/kangbb/ccrsystem/middlewares"
	"github.com/kangbb/ccrsystem/models/services"
)

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
	var (
		err error
		id  int
		pwd string
	)
	//default use the parameter from body
	defer r.Body.Close()
	js, err := simplejson.NewFromReader(r.Body)
	if logs.NormalError(err, w) {
		return
	}

	id, _ = strconv.Atoi(r.Header.Get("userId"))

	pwd = js.Get(userType + "Pwd").MustString()
	if !UserPasswordFormatValidate(pwd, w) {
		return
	}
  pwd = encryptPwd(id, pwd)
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
	pwd := encryptPwd(id, "123456")
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
			order := del_dep.DepartmentOrder
			for _, v := range departments {
				if v.DepartmentOrder > order {
					err = services.DepartmentService.UpdateInfo(v.DepartmentId, v.DepartmentName, v.Introduction, v.DepartmentOrder-1,
						v.Note)
					if logs.SqlError(err, w, true) {
						return
					}
				}
			}

			// Update reservation
			//Define some variable for updating reservatons
			new_dep, err := services.DepartmentService.FindInfoByOrder(del_dep.DepartmentOrder)
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

/******************************** Render File *************************/
func renderFile(name string, w http.ResponseWriter, r *http.Request) {
	filepath := "views/" + name + ".html"
	t := template.Must(template.ParseFiles(filepath))
	w.Header().Set("Content-Type", "text/html; charset=UTF-8")
	w.WriteHeader(200)
	t.Execute(w, nil)
}
