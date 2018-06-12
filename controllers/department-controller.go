package controllers

import (
	"encoding/json"
	"math/rand"
	"net/http"

	simplejson "github.com/bitly/go-simplejson"
	"github.com/kangbb/ccrsystem/logs"
	"github.com/kangbb/ccrsystem/models/services"
)

/************************* Data Interface For Department ***********************/
/*
* Get department list
 */
func getDepartmentList(w http.ResponseWriter, r *http.Request) {

	departments, err := services.DepartmentService.FindAllInfo()
	if logs.SqlError(err, w, len(departments) != 0) {
		return
	}

	data, _ := json.Marshal(departments)

	w.WriteHeader(200)
	w.Write(data)
}

/*
* Add a department information to db
* At the same time, the order should update. Default the order be latest; if define, than set custom the order.
 */
func addDepartment(w http.ResponseWriter, r *http.Request) {

	defer r.Body.Close()
	js, err := simplejson.NewFromReader(r.Body)
	if logs.NormalError(err, w) {
		return
	}

	name := js.Get("DepartmentName").MustString()
	order := js.Get("DepartmentOrder").MustInt()
	introduction := js.Get("Introduction").MustString()
	if name == "" || introduction == "" || order == 0 {
		w.WriteHeader(500)
		w.Write([]byte("必填项不能为空"))
	}

	//update the order of the approval department
	departments, err := services.DepartmentService.FindAllInfo()
	if logs.SqlError(err, w, true) {
		return
	}

	val := departments[0]
	if len(departments) == 1 && order > departments[0].DepartmentOrder {
		val.Note = "initial"
		err = services.DepartmentService.UpdateInfo(val.DepartmentId, val.DepartmentName, val.Introduction,
			val.DepartmentOrder+1, val.Note)
		if logs.SqlError(err, w, true) {
			return
		}
	} else if len(departments) == 1 && order == departments[0].DepartmentOrder {
		val.Note = "final"
		err = services.DepartmentService.UpdateInfo(val.DepartmentId, val.DepartmentName, val.Introduction,
			val.DepartmentOrder+1, val.Note)
		if logs.SqlError(err, w, true) {
			return
		}
	} else {
		for _, v := range departments {
			if v.DepartmentOrder >= order {
				err = services.DepartmentService.UpdateInfo(v.DepartmentId, v.DepartmentName, v.Introduction,
					v.DepartmentOrder+1, v.Note)
				if logs.SqlError(err, w, true) {
					return
				}
			}
		}
	}

	var note string
	if len(departments) == 0 {
		note = "initial,final"
	} else if order > len(departments) {
		note = "final"
	} else if order == 1 {
		note = "initial"
	} else {
		note = "middle"
	}

	dep := services.DepartmentService.NewDeparment(name, introduction, order, note)
	err = services.DepartmentService.SaveAInfo(dep)
	if logs.SqlError(err, w, true) {
		return
	}
	w.WriteHeader(200)
}

/*
* Get deparment by id
 */
func getDepartmentById(w http.ResponseWriter, r *http.Request) {
	id := getIdFromPath(r)

	department, err := services.DepartmentService.FindInfoById(id)
	if logs.SqlError(err, w, department.DepartmentName != "") {
		return
	}

	data, _ := json.Marshal(department)
	w.WriteHeader(200)
	w.Write(data)
}

/*
* Update department by Id
* The result is exchange the order of the two departments
 */
func updateDepartmentById(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	js, err := simplejson.NewFromReader(r.Body)
	if logs.NormalError(err, w) {
		return
	}

	id := getIdFromPath(r)
	name := js.Get("DepartmentName").MustString()
	intro := js.Get("Introduction").MustString()
	order := js.Get("DepartmentOrder").MustInt()
	if name == "" || intro == "" || order == 0 {
		logs.RequestError(500, logs.ErrorMsg{Msg: "必填字段不能为空"}, w)
		return
	}

	// Get department by id and order.
	departmentUpdate, err := services.DepartmentService.FindInfoById(id)
	if logs.SqlError(err, w, departmentUpdate.DepartmentName != "") {
		return
	}
	departmentNoUpdate, err := services.DepartmentService.FindInfoByOrder(order)
	if logs.SqlError(err, w, departmentNoUpdate.DepartmentName != "") {
		return
	}

	// If order no update, then update other information.
	if id == departmentNoUpdate.DepartmentId {
		err = services.DepartmentService.UpdateInfo(id, name, intro, order, departmentUpdate.Note)
		if logs.SqlError(err, w, true) {
			return
		}
	} else { // If order update, then change the two information.
		err = services.DepartmentService.UpdateInfo(id, name, intro, order, departmentUpdate.Note)
		if logs.SqlError(err, w, true) {
			return
		}

		err = services.DepartmentService.UpdateInfo(departmentNoUpdate.DepartmentId, departmentNoUpdate.DepartmentName, departmentNoUpdate.Introduction,
			departmentUpdate.DepartmentOrder, departmentNoUpdate.Note)
		if logs.SqlError(err, w, true) {
			return
		}
	}

	w.WriteHeader(200)
}

/*
* Delete department by Id
* The result will make some departments information change
 */
func deleteDepartmentById(w http.ResponseWriter, r *http.Request) {

	id := getIdFromPath(r)
	del_dep, err := services.DepartmentService.FindInfoById(id)
	if logs.SqlError(err, w, del_dep.DepartmentName != "") {
		return
	}
	// if del_dep.Note == "initial,final" {
	// 	logs.RequestError(500, logs.ErrorMsg{Msg: "系统正常运行，至少需要保留一个部门"}, w)
	// 	return
	// }
	// Update order
	departments, err := services.DepartmentService.FindAllInfo()
	if logs.SqlError(err, w, len(departments) != 0) {
		return
	}
	order := del_dep.DepartmentOrder
	isUPdateNoteInitial := false
	isUPdateNoteFinal := false
	if del_dep.Note == "initial" && len(departments) == 2 {
		departments[1].Note = "initial,final"
	} else if del_dep.Note == "final" && len(departments) == 2 {
		departments[0].Note = "initial,final"
	} else if del_dep.Note == "initial" {
		isUPdateNoteInitial = true
	} else if del_dep.Note == "final" {
		isUPdateNoteFinal = true
	}

	for _, v := range departments {
		if isUPdateNoteFinal && v.DepartmentOrder == (order-1) {
			err = services.DepartmentService.UpdateInfo(v.DepartmentId, v.DepartmentName, v.Introduction, v.DepartmentOrder,
				"final")
			if logs.SqlError(err, w, true) {
				return
			}
		}
		if v.DepartmentOrder > order {
			if isUPdateNoteInitial && v.DepartmentOrder == (order+1) {
				v.Note = "initial"
			}
			err = services.DepartmentService.UpdateInfo(v.DepartmentId, v.DepartmentName, v.Introduction, v.DepartmentOrder-1,
				v.Note)
			if logs.SqlError(err, w, true) {
				return
			}
		}
	}
	// Delete the department.
	err = services.DepartmentService.DeleteInfo(id)
	if logs.SqlError(err, w, true) {
		return
	}

	//Define some variable for updating reservatons
	new_dep, err := services.DepartmentService.FindInfoByOrder(del_dep.DepartmentOrder)
	if logs.SqlError(err, w, new_dep.DepartmentName != "") {
		return
	}
	new_approvers, err := services.ApproverService.FindInfoByDepartmentId(new_dep.DepartmentId)
	if logs.SqlError(err, w, len(new_approvers) != 0) {
		return
	}

	//Delete the approvers which belong to the department
	approvers, err := services.ApproverService.FindInfoByDepartmentId(id)
	if logs.SqlError(err, w, true) {
		return
	}
	for _, v := range approvers {
		err := services.ApproverService.DeleteInfo(v.ApproverId)
		if logs.SqlError(err, w, true) {
			return
		}

		//Update reservations.
		reservations, err := services.ReservationService.FindInfoByApproverId(v.ApproverId)
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
	}

	w.WriteHeader(200)
}
