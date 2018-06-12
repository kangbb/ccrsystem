package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	simplejson "github.com/bitly/go-simplejson"
	"github.com/kangbb/ccrsystem/logs"
	"github.com/kangbb/ccrsystem/models/services"
)

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
