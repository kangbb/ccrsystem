package controllers

import (
	"encoding/json"
	"net/http"
	"regexp"

	"github.com/kangbb/ccrsystem/logs"
)

//Render file
func GetIndex(w http.ResponseWriter, r *http.Request) {
	renderFile("index", w, r)
}
func GetStudentIndex(w http.ResponseWriter, r *http.Request) {
	if !validatePerm([]string{"Student"}, false, w, r) {
		return
	}
	renderFile("studentIndex", w, r)
}
func GetAdminIndex(w http.ResponseWriter, r *http.Request) {
	if !validatePerm([]string{"Admin"}, false, w, r) {
		return
	}
	renderFile("adminIndex", w, r)
}
func GetApproverIndex(w http.ResponseWriter, r *http.Request) {
	if !validatePerm([]string{"Approver"}, false, w, r) {
		return
	}
	renderFile("approverIndex", w, r)
}

//Signin and signout
func StudentSignin(w http.ResponseWriter, r *http.Request) {
	signin("Student", w, r)
}
func AdminSignin(w http.ResponseWriter, r *http.Request) {
	signin("Admin", w, r)
}
func ApproverSignin(w http.ResponseWriter, r *http.Request) {
	signin("Approver", w, r)
}
func Signout(w http.ResponseWriter, r *http.Request) {
	signout(w, r)
}

//data interface for student
func GetStudentInfo(w http.ResponseWriter, r *http.Request) {
	if !validatePerm([]string{"Student"}, false, w, r) {
		return
	}
	getUserInfo("Student", w, r)
}
func UpdateStudentInfo(w http.ResponseWriter, r *http.Request) {
	if !validatePerm([]string{"Student"}, false, w, r) {
		return
	}
	updateUserInfo("Student", w, r)
}
func GetStudentList(w http.ResponseWriter, r *http.Request) {
	if !validatePerm([]string{"Admin"}, false, w, r) {
		return
	}
	getUserList("Student", w, r)
}
func AddStudent(w http.ResponseWriter, r *http.Request) {
	if !validatePerm([]string{"Admin"}, false, w, r) {
		return
	}
	addUser("Student", w, r)
}
func GetStudentById(w http.ResponseWriter, r *http.Request) {
	if !validatePerm([]string{"Student", "Admin"}, true, w, r) {
		return
	}
	getUserById("Student", w, r)
}
func UpdateStudentById(w http.ResponseWriter, r *http.Request) {
	if !validatePerm([]string{"Student", "Admin"}, true, w, r) {
		return
	}
	updateUserById("Student", w, r)
}
func DeleteStudentById(w http.ResponseWriter, r *http.Request) {
	if !validatePerm([]string{"Adimin"}, false, w, r) {
		return
	}
	deleteUserById("Student", w, r)
}

//data interface for approver
func GetApproverInfo(w http.ResponseWriter, r *http.Request) {
	if !validatePerm([]string{"Approver"}, false, w, r) {
		return
	}
	getUserInfo("Approver", w, r)
}
func UpdateApproverInfo(w http.ResponseWriter, r *http.Request) {
	if !validatePerm([]string{"Approver"}, false, w, r) {
		return
	}
	updateUserInfo("Approver", w, r)
}
func GetApproverList(w http.ResponseWriter, r *http.Request) {
	if !validatePerm([]string{"Admin"}, false, w, r) {
		return
	}
	getUserList("Approver", w, r)
}
func AddApprover(w http.ResponseWriter, r *http.Request) {
	if !validatePerm([]string{"Admin"}, false, w, r) {
		return
	}
	addUser("Approver", w, r)
}
func GetApproverById(w http.ResponseWriter, r *http.Request) {
	if !validatePerm([]string{"Approver", "Admin"}, true, w, r) {
		return
	}
	getUserById("Approver", w, r)
}
func UpdateApproverById(w http.ResponseWriter, r *http.Request) {
	if !validatePerm([]string{"Approver", "Admin"}, true, w, r) {
		return
	}
	updateUserById("Approver", w, r)
}
func DeleteApproverById(w http.ResponseWriter, r *http.Request) {
	if !validatePerm([]string{"Admin"}, false, w, r) {
		return
	}
	deleteUserById("Approver", w, r)
}

//data interface for admin
func GetAdminInfo(w http.ResponseWriter, r *http.Request) {
	if !validatePerm([]string{"Admin"}, false, w, r) {
		return
	}
	getUserInfo("Admin", w, r)
}
func UpdateAdminInfo(w http.ResponseWriter, r *http.Request) {
	if !validatePerm([]string{"Admin"}, false, w, r) {
		return
	}
	updateUserInfo("Admin", w, r)
}
func GetAdminList(w http.ResponseWriter, r *http.Request) {
	if !validatePerm([]string{"Admin"}, false, w, r) {
		return
	}
	getUserList("Admin", w, r)
}
func AddAdmin(w http.ResponseWriter, r *http.Request) {
	if !validatePerm([]string{"Admin"}, false, w, r) {
		return
	}
	addUser("Admin", w, r)
}
func getAdminById(w http.ResponseWriter, r *http.Request) {
	if !validatePerm([]string{"Admin"}, false, w, r) {
		return
	}
	getUserById("Admin", w, r)
}
func UpdateAdminById(w http.ResponseWriter, r *http.Request) {
	if !validatePerm([]string{"Admin"}, false, w, r) {
		return
	}
	updateUserById("Admin", w, r)
}
func DeleteAdminById(w http.ResponseWriter, r *http.Request) {
	if !validatePerm([]string{"Admin"}, false, w, r) {
		return
	}
	deleteUserById("Admin", w, r)
}

//data interface for classroom
func GetClassroomList(w http.ResponseWriter, r *http.Request) {
	getClassroomList(w, r)
}
func AddClassroom(w http.ResponseWriter, r *http.Request) {
	if !validatePerm([]string{"Admin"}, false, w, r) {
		return
	}
	addClassroom(w, r)
}
func QueryClassroom(w http.ResponseWriter, r *http.Request) {
	queryClassroom(w, r)
}
func GetClassroomById(w http.ResponseWriter, r *http.Request) {
	getClassroomById(w, r)
}
func UpdateClassroomById(w http.ResponseWriter, r *http.Request) {
	if !validatePerm([]string{"Admin"}, false, w, r) {
		return
	}
	updateClassroomById(w, r)
}
func DeleteClassroomById(w http.ResponseWriter, r *http.Request) {
	if !validatePerm([]string{"Admin"}, false, w, r) {
		return
	}
	deleteClassroomById(w, r)
}

//data interface for reservation
func GetResById(w http.ResponseWriter, r *http.Request) {
	if !validatePerm([]string{"Student", "Approver"}, false, w, r) {
		return
	}
	getResById(w, r)
}
func UpdateResById(w http.ResponseWriter, r *http.Request) {
	if !validatePerm([]string{"Student", "Approver"}, false, w, r) {
		return
	}
	updateResById(w, r)
}
func DeleteResById(w http.ResponseWriter, r *http.Request) {
	if !validatePerm([]string{"Student"}, false, w, r) {
		return
	}
	deleteResById(w, r)
}
func GetStudentResList(w http.ResponseWriter, r *http.Request) {
	if !validatePerm([]string{"Student"}, false, w, r) {
		return
	}
	getStudentResList(w, r)
}
func AddRes(w http.ResponseWriter, r *http.Request) {
	if !validatePerm([]string{"Student"}, false, w, r) {
		return
	}
	addRes(w, r)
}
func GetApproverResList(w http.ResponseWriter, r *http.Request) {
	if !validatePerm([]string{"Approver"}, false, w, r) {
		return
	}
	getApproverResList(w, r)
}

//data interface for department
func GetDepartmentList(w http.ResponseWriter, r *http.Request) {
	if !validatePerm([]string{"Admin"}, false, w, r) {
		return
	}
	getDepartmentList(w, r)
}
func AddDepartment(w http.ResponseWriter, r *http.Request) {
	if !validatePerm([]string{"Admin"}, false, w, r) {
		return
	}
	addDepartment(w, r)
}
func GetDepartmentById(w http.ResponseWriter, r *http.Request) {
	getDepartmentById(w, r)
}
func UpdateDepartmentById(w http.ResponseWriter, r *http.Request) {
	if !validatePerm([]string{"Admin"}, false, w, r) {
		return
	}
	updateDepartmentById(w, r)
}
func DeleteDepartmentById(w http.ResponseWriter, r *http.Request) {
	if !validatePerm([]string{"Admin"}, false, w, r) {
		return
	}
	deleteDepartmentById(w, r)
}

/***********************************SOME SUBFUNCTION*******************************/
func validatePerm(userType []string, isNeedValidateId bool, w http.ResponseWriter, r *http.Request) bool {
	if isNeedValidateId {
		url := r.URL.Path
		reg := regexp.MustCompile(`\d{8}`)
		Id := reg.FindString(url)
		for _, v := range userType {
			if v == r.Header.Get("userType") {
				if v == "Admin" {
					return true
				} else if (v == "Student" || v == "Approver") && Id == r.Header.Get("userId") {
					return true
				}
			}
		}
	} else {
		for _, v := range userType {
			if v == r.Header.Get("userType") {
				return true
			}
		}
	}
	// Permission deny
	w.WriteHeader(401)
	data, _ := json.Marshal(logs.ErrorMsg{"Permission deny."})
	w.Write(data)
	return false
}
