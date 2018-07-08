package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gorilla/mux"
	"github.com/kangbb/ccrsystem/controllers"
	"github.com/kangbb/ccrsystem/logs"
	"github.com/kangbb/ccrsystem/models/entities"
	"github.com/kangbb/ccrsystem/models/services"
)

var cookies string

/*
* A simple error logs print function.
 */
func errStdout(t *testing.T, resp *http.Response) {
	t.Logf("[Response] StatusCode: %v", resp.StatusCode)
	body, _ := ioutil.ReadAll(resp.Body)
	data := &logs.ErrorMsg{}
	err := json.Unmarshal(body, data)
	if err != nil {
		t.Logf("[Error]json unmarshal error: %s", err)
		t.Fail()
	}
	t.Logf("[Response]Error from logs: %+v", data)

	new_data := &controllers.UserErrorMsg{}
	err = json.Unmarshal(body, new_data)
	if err != nil {
		t.Logf("[Error]json unmarshal error: %s", err)
		t.Fail()
	}
	t.Logf("[Response]Error from validate: %+v", new_data)
}

/*------------------------Student Data Interface Test--------------------*/
/*
* Test the html template render function
 */
func TestGetIndex(t *testing.T) {
	req := httptest.NewRequest(
		http.MethodGet,
		"/",
		nil,
	)

	w := httptest.NewRecorder()
	controllers.GetIndex(w, req)
	resp := w.Result()
	if resp.StatusCode != 200 {
		t.Fail()
	}
}

/*
* Test the Signin function
 */
func TestStudentSignin(t *testing.T) {
	reqData := struct {
		StudentId  int
		StudentPwd string
	}{15331124, "123456"}

	reqBody, _ := json.Marshal(reqData)
	req := httptest.NewRequest(
		http.MethodPost,
		"/student/sigin",
		bytes.NewReader(reqBody),
	)

	w := httptest.NewRecorder()
	controllers.StudentSignin(w, req)

	resp := w.Result()
	if resp.StatusCode != 200 {
		t.Fail()
		errStdout(t, resp)
	}
	if resp.Header.Get("Set-Cookie") == "" {
		t.Fail()
	}

	cookies = resp.Header.Get("Set-Cookie")
}

/*
* Test the get student information function
 */
func TestGetStudentInfo(t *testing.T) {

	req := httptest.NewRequest(
		http.MethodGet,
		"/api/users/student",
		nil,
	)
	req.Header.Set("userType", "Student")
	req.Header.Set("userId", "15331124")

	w := httptest.NewRecorder()
	controllers.GetStudentInfo(w, req)

	resp := w.Result()
	if resp.StatusCode != 200 {
		t.Fail()
		errStdout(t, resp)
	}
}

/*
* Test the get student information list
 */
func TestGetStudentList(t *testing.T) {
	req := httptest.NewRequest(
		http.MethodGet,
		"/api/users/students",
		nil,
	)

	req.Header.Set("userType", "Admin")
	req.Header.Set("userId", "11331124")

	w := httptest.NewRecorder()
	controllers.GetStudentList(w, req)

	resp := w.Result()
	if resp.StatusCode != 200 {
		t.Fail()
		errStdout(t, resp)
	}
}

/*
* Test the update student information by Id
 */

func TestUpdateStudentById(t *testing.T) {
	reqData := struct {
		StudentPwd string
	}{"example123"}
	reqBody, _ := json.Marshal(reqData)
	req := httptest.NewRequest(
		http.MethodPut,
		"/api/users/students/id",
		bytes.NewReader(reqBody),
	)
	req = mux.SetURLVars(req, map[string]string{"id": "15331124"})
	req.Header.Set("userType", "Admin")
	req.Header.Set("userId", "11331124")

	w := httptest.NewRecorder()
	controllers.UpdateStudentById(w, req)

	resp := w.Result()
	if resp.StatusCode != 200 {
		t.Fail()
		errStdout(t, resp)
	}
}

/*
* Test the delete and add student information.
 */
func TestDeleteStudentById(t *testing.T) {

	req := httptest.NewRequest(
		http.MethodDelete,
		"/api/users/students/id",
		nil,
	)
	req = mux.SetURLVars(req, map[string]string{"id": "15331124"})
	req.Header.Set("userType", "Admin")
	req.Header.Set("userId", "11331124")

	w := httptest.NewRecorder()
	controllers.DeleteStudentById(w, req)

	resp := w.Result()
	if resp.StatusCode != 200 {
		t.Fail()
		errStdout(t, resp)
	}
}

/*
* Test the add student information by Id
 */
func TestAddStudent(t *testing.T) {
	reqData := entities.StudentInfo{
		StudentId:   15331124,
		StudentPwd:  "123456",
		StudentName: "王芳",
	}
	reqBody, _ := json.Marshal(reqData)
	req := httptest.NewRequest(
		http.MethodPost,
		"/api/users/students",
		bytes.NewReader(reqBody),
	)

	req.Header.Set("userType", "Admin")
	req.Header.Set("userId", "11331124")

	w := httptest.NewRecorder()
	controllers.AddStudent(w, req)

	resp := w.Result()
	if resp.StatusCode != 200 {
		t.Fail()
		errStdout(t, resp)
	}
}

/*
* Test the get student information by Id.
 */
func TestGetStudentById(t *testing.T) {
	req := httptest.NewRequest(
		http.MethodGet,
		"/api/users/students/id",
		nil,
	)
	req = mux.SetURLVars(req, map[string]string{"id": "15331124"})

	req.Header.Set("userType", "Admin")
	req.Header.Set("userId", "11331124")

	w := httptest.NewRecorder()
	controllers.GetStudentById(w, req)

	resp := w.Result()
	if resp.StatusCode != 200 {
		t.Fail()
		errStdout(t, resp)
	}
}

func TestSigout(t *testing.T) {
	req := httptest.NewRequest(
		http.MethodPost,
		"/signout",
		nil,
	)

	req.Header.Set("userType", "Student")
	req.Header.Set("userId", "15331124")
	req.Header.Add("Cookie", cookies)
	w := httptest.NewRecorder()
	controllers.Signout(w, req)

	resp := w.Result()
	if resp.StatusCode != 302 {
		t.Fail()
		errStdout(t, resp)
	}
}

/*------------------------Classroom Data Interface Test--------------------*/
/*
* Test get classroom list function.
 */
func TestGetClassroomList(t *testing.T) {
	req := httptest.NewRequest(
		http.MethodGet,
		"/api/classrooms",
		nil,
	)

	w := httptest.NewRecorder()
	controllers.GetClassroomList(w, req)

	resp := w.Result()
	if resp.StatusCode != 200 {
		t.Fail()
		errStdout(t, resp)
	}
}

/*
* Test add classroom function.
 */
func TestAddClassroom(t *testing.T) {
	reqData := services.ClassroomService.NewClassroom("北校园", "医学院院楼", "D104", 100)
	reqBody, _ := json.Marshal(*reqData)
	req := httptest.NewRequest(
		http.MethodPost,
		"/api/classrooms",
		bytes.NewReader(reqBody),
	)
	req.Header.Set("userType", "Admin")
	req.Header.Set("userId", "11331124")

	w := httptest.NewRecorder()
	controllers.AddClassroom(w, req)

	resp := w.Result()
	if resp.StatusCode != 200 {
		t.Fail()
		errStdout(t, resp)
	}
}

/*
* Test get classroom by id function. Use the data we add lastest.
 */
func TestGetClassroomById(t *testing.T) {
	req := httptest.NewRequest(
		http.MethodGet,
		"/api/classrooms/id",
		nil,
	)

	req = mux.SetURLVars(req, map[string]string{"id": "3"})
	w := httptest.NewRecorder()
	controllers.GetClassroomById(w, req)

	resp := w.Result()
	if resp.StatusCode != 200 {
		t.Fail()
		errStdout(t, resp)
	}
}

/*
* Test update classrom inforamtion function.
 */
func TestUpdateClassroomById(t *testing.T) {
	reqData := services.ClassroomService.NewClassroom("珠海校园", "医学院院楼", "D104", 80)
	reqBody, _ := json.Marshal(*reqData)
	req := httptest.NewRequest(
		http.MethodPut,
		"/api/classrooms/id",
		bytes.NewReader(reqBody),
	)
	req = mux.SetURLVars(req, map[string]string{"id": "3"})
	req.Header.Set("userType", "Admin")
	req.Header.Set("userId", "11331124")

	w := httptest.NewRecorder()
	controllers.UpdateClassroomById(w, req)

	resp := w.Result()
	if resp.StatusCode != 200 {
		t.Fail()
		errStdout(t, resp)
	}
}

/*
* Test delete classroom information function.
 */
func TestDeleteClassroomById(t *testing.T) {
	req := httptest.NewRequest(
		http.MethodPut,
		"/api/classrooms/id",
		nil,
	)

	req = mux.SetURLVars(req, map[string]string{"id": "3"})
	req.Header.Set("userType", "Admin")
	req.Header.Set("userId", "11331124")

	w := httptest.NewRecorder()
	controllers.DeleteClassroomById(w, req)

	resp := w.Result()
	if resp.StatusCode != 200 {
		t.Fail()
		errStdout(t, resp)
	}
}

/*
* Test query classroom information function.
 */
// func TestQueryClassroom(t *testing.T) {
// 	reqData := struct {
// 		ClassroomCampus   string
// 		ClassroomBuilding string
// 		StartTime         string
// 		EndTime           string
// 		Capicity          int
// 	}{"东校园", "公教楼", "2018-05-23 第三节课", "2018-05-23 第五节课", 50}
// 	reqBody, _ := json.Marshal(reqData)
// 	req := httptest.NewRequest(
// 		http.MethodGet,
// 		"/api/classrooms/state",
// 		bytes.NewReader(reqBody),
// 	)

// 	w := httptest.NewRecorder()
// 	controllers.QueryClassroom(w, req)

// 	resp := w.Result()
// 	if resp.StatusCode != 200 {
// 		t.Fail()
// 		errStdout(t, resp)
// 	}
// }

/*------------------------Reservation Data Interface Test--------------------*/
/*
* Test get reservation list function
 */
func TestGetResById(t *testing.T) {
	req := httptest.NewRequest(
		http.MethodGet,
		"/api/reservations/id",
		nil,
	)
	req.Header.Set("userType", "Student")
	req.Header.Set("userId", "15331124")
	req = mux.SetURLVars(req, map[string]string{"id": "1"})

	w := httptest.NewRecorder()
	controllers.GetResById(w, req)

	resp := w.Result()
	if resp.StatusCode != 200 {
		t.Fail()
		errStdout(t, resp)
	}
}

/*
* Test update reservation list function
 */
func TestUpdateResById(t *testing.T) {
	reqData := struct {
		ResReason        string
		OrganizationName string
	}{"考试复习", "东校园学生会"}
	reqBody, _ := json.Marshal(reqData)
	req := httptest.NewRequest(
		http.MethodGet,
		"/api/reservations/id",
		bytes.NewReader(reqBody),
	)
	req.Header.Set("userType", "Student")
	req.Header.Set("userId", "15331124")
	req = mux.SetURLVars(req, map[string]string{"id": "1"})

	w := httptest.NewRecorder()
	controllers.UpdateResById(w, req)

	resp := w.Result()
	if resp.StatusCode != 200 {
		t.Fail()
		errStdout(t, resp)
	}
}

/*
* Test delete reservation list function
 */
func TestDeleteResById(t *testing.T) {
	req := httptest.NewRequest(
		http.MethodGet,
		"/api/reservations/id",
		nil,
	)
	req.Header.Set("userType", "Student")
	req.Header.Set("userId", "15331124")
	req = mux.SetURLVars(req, map[string]string{"id": "1"})

	w := httptest.NewRecorder()
	controllers.DeleteResById(w, req)

	resp := w.Result()
	if resp.StatusCode != 200 {
		t.Fail()
		errStdout(t, resp)
	}
}

/*
* Test student get his/her reservation list function
 */
func TestGetStudentResList(t *testing.T) {
	req := httptest.NewRequest(
		http.MethodGet,
		"/api/users/student/reservations",
		nil,
	)
	req.Header.Set("userType", "Student")
	req.Header.Set("userId", "15331124")

	w := httptest.NewRecorder()
	controllers.GetStudentResList(w, req)

	resp := w.Result()
	if resp.StatusCode != 404 {
		t.Fail()
		errStdout(t, resp)
	}
}

/*
* Test student add his/her reservation list function
 */
func TestAddRes(t *testing.T) {
	reqData := struct {
		StartTime   string
		EndTime     string
		ResReason   string
		ClassroomId int
	}{"2018-06-13 第三节课", "2018-06-13 第五节课", "习题辅导课", 2}
	reqBody, _ := json.Marshal(reqData)
	req := httptest.NewRequest(
		http.MethodPost,
		"/api/users/student/reservations",
		bytes.NewReader(reqBody),
	)
	req.Header.Set("userType", "Student")
	req.Header.Set("userId", "15331124")

	w := httptest.NewRecorder()
	controllers.AddRes(w, req)

	resp := w.Result()
	if resp.StatusCode != 200 {
		t.Fail()
		errStdout(t, resp)
	}

	//Test wether the data has in databases
	time.Sleep(10)
	req = httptest.NewRequest(
		http.MethodGet,
		"/api/users/student/reservations",
		nil,
	)
	req.Header.Set("userType", "Student")
	req.Header.Set("userId", "15331124")

	w = httptest.NewRecorder()
	controllers.GetStudentResList(w, req)

	resp = w.Result()
	if resp.StatusCode != 200 {
		t.Fail()
		errStdout(t, resp)
	}
}

/*
* Test approver get his/her reservation list function
 */
func TestGetApproverResList(t *testing.T) {
	req := httptest.NewRequest(
		http.MethodGet,
		"/api/users/approver/reservations",
		nil,
	)
	req.Header.Set("userType", "Approver")
	req.Header.Set("userId", "10331124")

	w := httptest.NewRecorder()
	controllers.GetApproverResList(w, req)

	resp := w.Result()
	if resp.StatusCode != 200 {
		t.Fail()
		errStdout(t, resp)
	}
}

/*------------------------Department Data Interface Test--------------------*/
/*
* These tests are very important
* Some function need to implement the cascade delete.
 */

/*
* Test the get department list function.
 */
func TestGetDepartmentList(t *testing.T) {
	req := httptest.NewRequest(
		http.MethodGet,
		"/api/departments",
		nil,
	)

	req.Header.Set("userType", "Admin")
	req.Header.Set("userId", "11331124")

	w := httptest.NewRecorder()
	controllers.GetDepartmentList(w, req)

	resp := w.Result()
	if resp.StatusCode != 200 {
		t.Fail()
		errStdout(t, resp)
	}
}

/*
* Test the add department function.
 */
func TestAddDepartment(t *testing.T) {
	reqData := struct {
		DepartmentName string
		Order          int
		Introduction   string
	}{"学生处", 2, "活动审批的第二个部门"}
	reqBody, _ := json.Marshal(reqData)
	req := httptest.NewRequest(
		http.MethodPost,
		"/api/departments",
		bytes.NewReader(reqBody),
	)

	req.Header.Set("userType", "Admin")
	req.Header.Set("userId", "11331124")

	w := httptest.NewRecorder()
	controllers.AddDepartment(w, req)

	resp := w.Result()
	if resp.StatusCode != 200 {
		t.Fail()
		errStdout(t, resp)
	}
}

/*
* Test get department by id function.
* And validate the add information result.
 */
func TestGetDepartmentById(t *testing.T) {
	req := httptest.NewRequest(
		http.MethodPost,
		"/api/departments/id",
		nil,
	)

	req.Header.Set("userType", "Admin")
	req.Header.Set("userId", "11331124")
	req = mux.SetURLVars(req, map[string]string{"id": "4"})
	w := httptest.NewRecorder()
	controllers.GetDepartmentById(w, req)

	resp := w.Result()
	if resp.StatusCode != 200 {
		t.Fail()
		errStdout(t, resp)
	}

	body, _ := ioutil.ReadAll(resp.Body)
	data := &entities.DepartmentInfo{}
	err := json.Unmarshal(body, data)
	if err != nil {
		t.Fail()
	}
	if data.DepartmentName != "学生处" {
		t.Fail()
		t.Log(data)
	}
}
