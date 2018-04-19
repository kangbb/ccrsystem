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
	r.HandleFunc("/", getIndex)
	r.HandleFunc("/student", getStudentIndex)
	r.HandleFunc("/approver", getApproverIndex)
	r.HandleFunc("/admin", getAdminIndex)
	// function interface
	r.HandleFunc("/signout", signout)
	r.HandleFunc("/student/signin", studentSignin)
	r.HandleFunc("/approver/signin", approverSignin)
	r.HandleFunc("/admin/signin", adminSignin)
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

	static := "static"
	r.PathPrefix("/static").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir(webRoot+static))))

	s := negroni.Classic()
	s.UseHandler(r)
	return s
}
