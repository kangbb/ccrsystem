package controller

import (
	"net/http"

	"github.com/kangbb/ccrsystem/core/services"
)

//Render file
func getIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	services.RenderFile("index", w, r)
}
func getStudentIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	services.RenderFile("studentIndex", w, r)
}
func getAdminIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	services.RenderFile("adminIndex", w, r)
}
func getApproverIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	services.RenderFile("approverIndex", w, r)
}

//Signin and signout
func studentSignin(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	services.Signin("Student", w, r)
}
func adminSignin(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	services.Signin("Admin", w, r)
}
func approverSignin(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	services.Signin("Approver", w, r)
}
func signout(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	services.Signout(w, r)
}

//data interface for student
func getStudentInfo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	services.GetUserInfo("Student", w, r)
}
func updateStudentInfo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	services.UpdateUserInfo("Student", w, r)
}
func getStudentList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	services.GetUserList("Student", w, r)
}
func addStudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	services.AddUser("Student", w, r)
}
func getStudentById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	services.GetUserById("Student", w, r)
}
func updateStudentById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	services.UpdateUserById("Student", w, r)
}
func deleteStudentById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	services.DeleteUserById("Student", w, r)
}

//data interface for approver
func getApproverInfo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	services.GetUserInfo("Approver", w, r)
}
func updateApproverInfo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	services.UpdateUserInfo("Approver", w, r)
}
func getApproverList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	services.GetUserList("Approver", w, r)
}
func addApprover(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	services.AddUser("Approver", w, r)
}
func getApproverById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	services.GetUserById("Approver", w, r)
}
func updateApproverById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	services.UpdateUserById("Approver", w, r)
}
func deleteApproverById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	services.DeleteUserById("Approver", w, r)
}

//data interface for admin
func getAdminInfo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	services.GetUserInfo("Admin", w, r)
}
func updateAdminInfo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	services.UpdateUserInfo("Admin", w, r)
}
func getAdminList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	services.GetUserList("Admin", w, r)
}
func addAdmin(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	services.AddUser("Admin", w, r)
}
func getAdminById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	services.GetUserById("Admin", w, r)
}
func updateAdminById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	services.UpdateUserById("Admin", w, r)
}
func deleteAdminById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	services.DeleteUserById("Admin", w, r)
}

//data interface for classroom
func getClassroomList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	services.GetClassroomList(w, r)
}
func addClassroom(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	services.AddClassroom(w, r)
}
func queryClassroom(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	services.QueryClassroom(w, r)
}
func getClassroomById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	services.GetClassroomById(w, r)
}
func updateClassroomById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	services.UpdateClassroomById(w, r)
}
func deleteClassroomById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	services.DeleteClassroomById(w, r)
}

//data interface for reservation
func getResById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	services.GetResById(w, r)
}
func updateResById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	services.UpdateResById(w, r)
}
func deleteResById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	services.DeleteResById(w, r)
}
func getStudentResList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	services.GetStudentResList(w, r)
}
func addRes(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	services.AddRes(w, r)
}
func getApproverResList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	services.GetApproverResList(w, r)
}

//data interface for department
func getDepartmentList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	services.GetDepartmentList(w, r)
}
func addDepartment(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	services.AddDepartment(w, r)
}
func getDepartmentById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	services.GetDepartmentById(w, r)
}
func updateDepartmentById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	services.UpdateDepartmentById(w, r)
}
func deleteDepartmentById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	services.DeleteDepartmentById(w, r)
}
func accessControl(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", r.Header.Get("Origin"))
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Add("Access-Control-Allow-Method", "POST, OPTIONS, GET, HEAD, PUT, PATCH, DELETE")

	w.Header().Add("Access-Control-Allow-Headers", "Origin, X-Requested-With, X-HTTP-Method-Override,accept-charset,accept-encoding , Content-Type, Accept, Cookie")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write([]byte("跨域已经可以"))
}
