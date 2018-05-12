package controller

import (
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

// GetServer return web server
func GetServer() *negroni.Negroni {
	webRoot := os.Getenv("WEBROOT")
	if len(webRoot) == 0 {
		if root, err := os.Getwd(); err != nil {
			panic("Could not retrive working directory")
		} else {
			webRoot = root
		}
	}
	r := mux.NewRouter()
	//load static template
	r.HandleFunc("/", getIndex).Methods("GET")
	r.HandleFunc("/student", getStudentIndex).Methods("GET")
	r.HandleFunc("/approver", getApproverIndex).Methods("GET")
	r.HandleFunc("/admin", getAdminIndex).Methods("GET")
	// function interface
	r.HandleFunc("/signout", signout).Methods("POST")
	r.HandleFunc("/student/signin", studentSignin).Methods("POST")
	r.HandleFunc("/approver/signin", approverSignin).Methods("POST")
	r.HandleFunc("/admin/signin", adminSignin).Methods("POST")
	// data interface
	//for student
	r.HandleFunc("/api/users/student", getStudentInfo).Methods("GET")
	r.HandleFunc("/api/users/student", updateStudentInfo).Methods("PUT")
	r.HandleFunc("/api/users/students", getStudentList).Methods("GET")
	r.HandleFunc("/api/users/students", addStudent).Methods("POST")
	r.HandleFunc("/api/users/students/{id}", getStudentById).Methods("GET")
	r.HandleFunc("/api/users/students/{id}", updateStudentById).Methods("PUT")
	r.HandleFunc("/api/users/students/{id}", deleteStudentById).Methods("DELETE")
	//for approver
	r.HandleFunc("/api/users/approver", getApproverInfo).Methods("GET")
	r.HandleFunc("/api/users/approver", updateApproverInfo).Methods("PUT")
	r.HandleFunc("/api/users/approvers", getApproverList).Methods("GET")
	r.HandleFunc("/api/users/approvers", addApprover).Methods("POST")
	r.HandleFunc("/api/users/approvers/{id}", getApproverById).Methods("GET")
	r.HandleFunc("/api/users/approvers/{id}", updateApproverById).Methods("PUT")
	r.HandleFunc("/api/users/approvers/{id}", deleteApproverById).Methods("DELETE")
	//for admin
	r.HandleFunc("/api/users/admin", getAdminInfo).Methods("GET")
	r.HandleFunc("/api/users/admin", updateAdminInfo).Methods("PUT")
	r.HandleFunc("/api/users/admins", getAdminList).Methods("GET")
	r.HandleFunc("/api/users/admins", addAdmin).Methods("POST")
	r.HandleFunc("/api/users/admins/{id}", getAdminById).Methods("GET")
	r.HandleFunc("/api/users/admins/{id}", updateAdminById).Methods("PUT")
	r.HandleFunc("/api/users/admins/{id}", deleteAdminById).Methods("DELETE")
	//for classroom
	r.HandleFunc("/api/classrooms", getClassroomList).Methods("GET")
	r.HandleFunc("/api/classrooms", addClassroom).Methods("POST")
	r.HandleFunc("/api/classrooms/state", queryClassroom)
	r.HandleFunc("/api/classrooms/{id}", getClassroomById).Methods("GET")
	r.HandleFunc("/api/classrooms/{id}", updateClassroomById).Methods("PUT")
	r.HandleFunc("/api/classrooms/{id}", deleteClassroomById).Methods("DELETE")
	//for reservation
	r.HandleFunc("/api/reservations/{id}", getResById).Methods("GET")
	r.HandleFunc("/api/reservations/{id}", updateResById).Methods("PUT")
	r.HandleFunc("/api/reservations/{id}", deleteResById).Methods("DELETE")
	r.HandleFunc("/api/users/student/reservations", addRes).Methods("POST")
	r.HandleFunc("/api/users/student/reservations", getStudentResList).Methods("GET")
	r.HandleFunc("/api/users/approver/reservations", getApproverResList)
	//for department
	r.HandleFunc("/api/departments", getDepartmentList).Methods("GET")
	r.HandleFunc("/api/departments", addDepartment).Methods("POST")
	r.HandleFunc("/api/departments/{id}", getDepartmentById).Methods("GET")
	r.HandleFunc("/api/departments/{id}", updateDepartmentById).Methods("PUT")
	r.HandleFunc("/api/departments/{id}", deleteDepartmentById).Methods("DELETE")

	// 跨域控制， 暂定
	r.HandleFunc("/", accessControl).Methods("OPTIONS")
	r.HandleFunc("/student", accessControl).Methods("OPTIONS")
	r.HandleFunc("/approver", accessControl).Methods("OPTIONS")
	r.HandleFunc("/admin", accessControl).Methods("OPTIONS")
	// function interface
	r.HandleFunc("/signout", accessControl).Methods("OPTIONS")
	r.HandleFunc("/student/signin", accessControl).Methods("OPTIONS")
	r.HandleFunc("/approver/signin", accessControl).Methods("OPTIONS")
	r.HandleFunc("/admin/signin", accessControl).Methods("OPTIONS")
	// data interface
	//for student
	r.HandleFunc("/api/users/student", accessControl).Methods("OPTIONS")
	r.HandleFunc("/api/users/student", accessControl).Methods("OPTIONS")
	r.HandleFunc("/api/users/students", accessControl).Methods("OPTIONS")
	r.HandleFunc("/api/users/students", accessControl).Methods("OPTIONS")
	r.HandleFunc("/api/users/students/{id}", accessControl).Methods("OPTIONS")
	r.HandleFunc("/api/users/students/{id}", accessControl).Methods("OPTIONS")
	r.HandleFunc("/api/users/students/{id}", accessControl).Methods("OPTIONS")
	//for approver
	r.HandleFunc("/api/users/approver", accessControl).Methods("OPTIONS")
	r.HandleFunc("/api/users/approver", accessControl).Methods("OPTIONS")
	r.HandleFunc("/api/users/approvers", accessControl).Methods("OPTIONS")
	r.HandleFunc("/api/users/approvers", accessControl).Methods("OPTIONS")
	r.HandleFunc("/api/users/approvers/{id}", accessControl).Methods("OPTIONS")
	r.HandleFunc("/api/users/approvers/{id}", accessControl).Methods("OPTIONS")
	r.HandleFunc("/api/users/approvers/{id}", accessControl).Methods("OPTIONS")
	//for admin
	r.HandleFunc("/api/users/admin", accessControl).Methods("OPTIONS")
	r.HandleFunc("/api/users/admin", accessControl).Methods("OPTIONS")
	r.HandleFunc("/api/users/admins", accessControl).Methods("OPTIONS")
	r.HandleFunc("/api/users/admins", addAdmin).Methods("OPTIONS")
	r.HandleFunc("/api/users/admins/{id}", accessControl).Methods("OPTIONS")
	r.HandleFunc("/api/users/admins/{id}", accessControl).Methods("OPTIONS")
	r.HandleFunc("/api/users/admins/{id}", accessControl).Methods("OPTIONS")
	//for classroom
	r.HandleFunc("/api/classrooms", accessControl).Methods("OPTIONS")
	r.HandleFunc("/api/classrooms", accessControl).Methods("OPTIONS")
	r.HandleFunc("/api/classrooms/state", accessControl)
	r.HandleFunc("/api/classrooms/{id}", accessControl).Methods("OPTIONS")
	r.HandleFunc("/api/classrooms/{id}", accessControl).Methods("OPTIONS")
	r.HandleFunc("/api/classrooms/{id}", accessControl).Methods("OPTIONS")
	//for reservation
	r.HandleFunc("/api/reservations/{id}", accessControl).Methods("OPTIONS")
	r.HandleFunc("/api/reservations/{id}", accessControl).Methods("OPTIONS")
	r.HandleFunc("/api/reservations/{id}", accessControl).Methods("OPTIONS")
	r.HandleFunc("/api/users/student/reservations", accessControl).Methods("OPTIONS")
	r.HandleFunc("/api/users/student/reservations", accessControl).Methods("OPTIONS")
	r.HandleFunc("/api/users/approver/reservations", accessControl)
	//for department
	r.HandleFunc("/api/departments", accessControl).Methods("OPTIONS")
	r.HandleFunc("/api/departments", addDepartment).Methods("OPTIONS")
	r.HandleFunc("/api/departments/{id}", accessControl).Methods("OPTIONS")
	r.HandleFunc("/api/departments/{id}", accessControl).Methods("OPTIONS")
	r.HandleFunc("/api/departments/{id}", accessControl).Methods("OPTIONS")

	static := "/static/"
	r.PathPrefix("/static").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir(webRoot+static))))

	s := negroni.Classic()
	s.UseHandler(r)
	return s
}
