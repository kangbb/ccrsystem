package routers

import (
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/kangbb/ccrsystem/controllers"
	"github.com/kangbb/ccrsystem/middlewares"
)

/*
* Get routers. The function return a router list to the main.go
* It registes some routers which used to process business request.
 */
func GetRouters() *mux.Router {

	r := mux.NewRouter()
	// Load static template
	r.HandleFunc("/", controllers.GetIndex).Methods("GET")
	r.HandleFunc("/student", controllers.GetStudentIndex).Methods("GET")
	r.HandleFunc("/approver", controllers.GetApproverIndex).Methods("GET")
	r.HandleFunc("/admin", controllers.GetAdminIndex).Methods("GET")
	// Function interface
	r.HandleFunc("/signout", controllers.Signout).Methods("POST")
	r.HandleFunc("/student/signin", controllers.StudentSignin).Methods("POST")
	r.HandleFunc("/approver/signin", controllers.ApproverSignin).Methods("POST")
	r.HandleFunc("/admin/signin", controllers.AdminSignin).Methods("POST")
	// Data interface
	// For student
	r.HandleFunc("/api/users/student", controllers.GetStudentInfo).Methods("GET")
	r.HandleFunc("/api/users/student", controllers.UpdateStudentInfo).Methods("PUT")
	r.HandleFunc("/api/users/students", controllers.GetStudentList).Methods("GET")
	r.HandleFunc("/api/users/students", controllers.AddStudent).Methods("POST")
	r.HandleFunc("/api/users/students/{id}", controllers.GetStudentById).Methods("GET")
	r.HandleFunc("/api/users/students/{id}", controllers.UpdateStudentById).Methods("PUT")
	r.HandleFunc("/api/users/students/{id}", controllers.DeleteStudentById).Methods("DELETE")
	// For approver
	r.HandleFunc("/api/users/approver", controllers.GetApproverInfo).Methods("GET")
	r.HandleFunc("/api/users/approver", controllers.UpdateApproverInfo).Methods("PUT")
	r.HandleFunc("/api/users/approvers", controllers.GetApproverList).Methods("GET")
	r.HandleFunc("/api/users/approvers", controllers.AddApprover).Methods("POST")
	r.HandleFunc("/api/users/approvers/{id}", controllers.GetApproverById).Methods("GET")
	r.HandleFunc("/api/users/approvers/{id}", controllers.UpdateApproverById).Methods("PUT")
	r.HandleFunc("/api/users/approvers/{id}", controllers.DeleteApproverById).Methods("DELETE")
	// For admin
	r.HandleFunc("/api/users/admin", controllers.GetAdminInfo).Methods("GET")
	r.HandleFunc("/api/users/admin", controllers.UpdateAdminInfo).Methods("PUT")
	r.HandleFunc("/api/users/admins", controllers.GetAdminList).Methods("GET")
	r.HandleFunc("/api/users/admins", controllers.AddAdmin).Methods("POST")
	r.HandleFunc("/api/users/admins/{id}", controllers.GetAdminById).Methods("GET")
	r.HandleFunc("/api/users/admins/{id}", controllers.UpdateAdminById).Methods("PUT")
	r.HandleFunc("/api/users/admins/{id}", controllers.DeleteAdminById).Methods("DELETE")
	// For classroom
	r.HandleFunc("/api/classrooms", controllers.GetClassroomList).Methods("GET")
	r.HandleFunc("/api/classrooms", controllers.AddClassroom).Methods("POST")
	r.HandleFunc("/api/classrooms/state", controllers.QueryClassroom)
	r.HandleFunc("/api/classrooms/{id}", controllers.GetClassroomById).Methods("GET")
	r.HandleFunc("/api/classrooms/{id}", controllers.UpdateClassroomById).Methods("PUT")
	r.HandleFunc("/api/classrooms/{id}", controllers.DeleteClassroomById).Methods("DELETE")
	// For reservation
	r.HandleFunc("/api/reservations/{id}", controllers.GetResById).Methods("GET")
	r.HandleFunc("/api/reservations/{id}", controllers.UpdateResById).Methods("PUT")
	r.HandleFunc("/api/reservations/{id}", controllers.DeleteResById).Methods("DELETE")
	r.HandleFunc("/api/users/student/reservations", controllers.AddRes).Methods("POST")
	r.HandleFunc("/api/users/student/reservations", controllers.GetStudentResList).Methods("GET")
	r.HandleFunc("/api/users/approver/reservations", controllers.GetApproverResList)
	// For department
	r.HandleFunc("/api/departments", controllers.GetDepartmentList).Methods("GET")
	r.HandleFunc("/api/departments", controllers.AddDepartment).Methods("POST")
	r.HandleFunc("/api/departments/{id}", controllers.GetDepartmentById).Methods("GET")
	r.HandleFunc("/api/departments/{id}", controllers.UpdateDepartmentById).Methods("PUT")
	r.HandleFunc("/api/departments/{id}", controllers.DeleteDepartmentById).Methods("DELETE")
	// For organization
	r.HandleFunc("/api/organizations", controllers.GetOrganizationList).Methods("GET")
	r.HandleFunc("/api/organizations", controllers.AddOrganization).Methods("POST")
	r.HandleFunc("/api/organizations/{id}", controllers.GetOrganizationById).Methods("GET")
	r.HandleFunc("/api/organizations/{id}", controllers.UpdateOrganizationById).Methods("PUT")
	r.HandleFunc("/api/organizations/{id}", controllers.DeleteOrganizationById).Methods("DELETE")
	// For request not found
	//r.NotFoundHandler(NotFoundHandler)

	// Construct a static file sever
	webRoot := os.Getenv("WEBROOT")
	if len(webRoot) == 0 {
		if root, err := os.Getwd(); err != nil {
			panic("Could not retrive working directory")
		} else {
			webRoot = root
		}
	}
	static := "/static/"
	// or static = "/static"
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir(webRoot+static))))

	// Use the middleware to handle request.
	r.Use(middlewares.CorsHandler)
	r.Use(middlewares.Authentication)
	return r
}
