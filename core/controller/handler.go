package controller

import (
	"net/http"

	"github.com/kangbb/class-reservation/core/services"
)

//Render file
func getIndex(w http.ResponseWriter, r *http.Request) {
	services.RenderFile("index", w, r)
}
func getStudentIndex(w http.ResponseWriter, r *http.Request) {
	services.RenderFile("studentIndex", w, r)
}
func getAdminIndex(w http.ResponseWriter, r *http.Request) {
	services.RenderFile("adminIndex", w, r)
}
func getApproverIndex(w http.ResponseWriter, r *http.Request) {
	services.RenderFile("approverIndex", w, r)
}

//Signin and signout
func studentSignin(w http.ResponseWriter, r *http.Request) {
	services.Signin("Student", w, r)
}
func adminSignin(w http.ResponseWriter, r *http.Request) {
	services.Signin("Admin", w, r)
}
func approverSignin(w http.ResponseWriter, r *http.Request) {
	services.Signin("Approver", w, r)
}
func signout(w http.ResponseWriter, r *http.Request) {
	services.Signout(w, r)
}

//data interface for student
func getStudentInfo(w http.ResponseWriter, r *http.Request) {
	services.GetUserInfo("Student", w, r)
}
func updateStudentInfo(w http.ResponseWriter, r *http.Request) {
	services.UpdateUserInfo("Student", w, r)
}
func getStudentList(w http.ResponseWriter, r *http.Request) {
	services.GetUserList("Student", w, r)
}
func addStudent(w http.ResponseWriter, r *http.Request) {
	services.AddUser("Student", w, r)
}
func getStudentById(w http.ResponseWriter, r *http.Request) {
	services.GetUserById("Student", w, r)
}
func updateStudentById(w http.ResponseWriter, r *http.Request) {
	services.UpdateUserById("Student", w, r)
}
func deleteStudentById(w http.ResponseWriter, r *http.Request) {
	services.DeleteUserById("Student", w, r)
}

//data interface for approver
func getApproverInfo(w http.ResponseWriter, r *http.Request) {
	services.GetUserInfo("Approver", w, r)
}
func updateApproverInfo(w http.ResponseWriter, r *http.Request) {
	services.UpdateUserInfo("Approver", w, r)
}
func getApproverList(w http.ResponseWriter, r *http.Request) {
	services.GetUserList("Approver", w, r)
}
func addApprover(w http.ResponseWriter, r *http.Request) {
	services.AddUser("Approver", w, r)
}
func getApproverById(w http.ResponseWriter, r *http.Request) {
	services.GetUserById("Approver", w, r)
}
func updateApproverById(w http.ResponseWriter, r *http.Request) {
	services.UpdateUserById("Approver", w, r)
}
func deleteApproverById(w http.ResponseWriter, r *http.Request) {
	services.DeleteUserById("Approver", w, r)
}

//data interface for admin
func getAdminInfo(w http.ResponseWriter, r *http.Request) {
	services.GetUserInfo("Admin", w, r)
}
func updateAdminInfo(w http.ResponseWriter, r *http.Request) {
	services.UpdateUserInfo("Admin", w, r)
}
func getAdminList(w http.ResponseWriter, r *http.Request) {
	services.GetUserList("Admin", w, r)
}
func addAdmin(w http.ResponseWriter, r *http.Request) {
	services.AddUser("Admin", w, r)
}
func getAdminById(w http.ResponseWriter, r *http.Request) {
	services.GetUserById("Admin", w, r)
}
func updateAdminById(w http.ResponseWriter, r *http.Request) {
	services.UpdateUserById("Admin", w, r)
}
func deleteAdminById(w http.ResponseWriter, r *http.Request) {
	services.DeleteUserById("Admin", w, r)
}

//data interface for classroom
func getClassroomList(w http.ResponseWriter, r *http.Request) {
	services.GetClassroomList(w, r)
}
func addClassroom(w http.ResponseWriter, r *http.Request) {
	services.AddClassroom(w, r)
}
func queryClassroom(w http.ResponseWriter, r *http.Request) {
	services.QueryClassroom(w, r)
}
func getClassroomById(w http.ResponseWriter, r *http.Request) {
	services.GetClassroomById(w, r)
}
func updateClassroomById(w http.ResponseWriter, r *http.Request) {
	services.UpdateClassroomById(w, r)
}
func deleteClassroomById(w http.ResponseWriter, r *http.Request) {
	services.DeleteClassroomById(w, r)
}

//data interface for reservation
func getResById(w http.ResponseWriter, r *http.Request) {
	services.GetResById(w, r)
}
func updateResById(w http.ResponseWriter, r *http.Request) {
	services.UpdateResById(w, r)
}
func deleteResById(w http.ResponseWriter, r *http.Request) {
	services.DeleteResById(w, r)
}
func getStudentResList(w http.ResponseWriter, r *http.Request) {
	services.GetStudentResList(w, r)
}
func addRes(w http.ResponseWriter, r *http.Request) {
	services.AddRes(w, r)
}
func getApproverResList(w http.ResponseWriter, r *http.Request) {
	services.GetApproverResList(w, r)
}

//data interface for department
func getDepartmentList(w http.ResponseWriter, r *http.Request) {
	services.GetDepartmentList(w, r)
}
func addDepartment(w http.ResponseWriter, r *http.Request) {
	services.AddDepartment(w, r)
}
func getDepartmentById(w http.ResponseWriter, r *http.Request) {
	services.GetDepartmentById(w, r)
}
func updateDepartmentById(w http.ResponseWriter, r *http.Request) {
	services.UpdateDepartmentById(w, r)
}
func deleteDepartmentById(w http.ResponseWriter, r *http.Request) {
	services.DeleteDepartmentById(w, r)
}
