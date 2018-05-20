/*
* The file used to initiale mysql database.
* It will create some data which can be used to test the system.
* Please make sure that you have config a right environment to run this program before everything start.
 */
package controllers

import (
	"github.com/kangbb/ccrsystem/models/services"
)

/*
* Add some student information
 */
func initStudent() {
	pwd := encryptPwd(15331124, "123456")
	user := services.StudentService.NewStudent(15331124, pwd, "王芳")
	services.StudentService.SaveAInfo(user)

	pwd = encryptPwd(15331125, "123456")
	user = services.StudentService.NewStudent(15331125, pwd, "赵红")
	services.StudentService.SaveAInfo(user)

	pwd = encryptPwd(16331124, "123456")
	user = services.StudentService.NewStudent(16331124, pwd, "王芸")
	services.StudentService.SaveAInfo(user)

	pwd = encryptPwd(16331125, "123456")
	user = services.StudentService.NewStudent(16331125, pwd, "张晓")
	services.StudentService.SaveAInfo(user)
}

/*
* Add some approver information.
 */
func initApprover() {
	pwd := encryptPwd(10331124, "123456")
	user := services.ApproverService.NewApprover(10331124, pwd, "李怡", 1)
	services.ApproverService.SaveAInfo(user)

	pwd = encryptPwd(10331125, "123456")
	user = services.ApproverService.NewApprover(10331125, pwd, "宋晓", 2)
	services.ApproverService.SaveAInfo(user)

	pwd = encryptPwd(10331126, "123456")
	user = services.ApproverService.NewApprover(10331126, pwd, "李芸", 3)
	services.ApproverService.SaveAInfo(user)

	pwd = encryptPwd(10331127, "123456")
	user = services.ApproverService.NewApprover(10331127, pwd, "张丽", 1)
	services.ApproverService.SaveAInfo(user)
}

/*
* Add some admin information.
 */
func initAdmin() {
	pwd := encryptPwd(11331124, "123456")
	user := services.AdminService.NewAdmin(11331124, pwd, "赵毅")
	services.AdminService.SaveAInfo(user)

	pwd = encryptPwd(11331125, "123456")
	user = services.AdminService.NewAdmin(11331125, pwd, "李阳")
	services.AdminService.SaveAInfo(user)

	pwd = encryptPwd(11331126, "123456")
	user = services.AdminService.NewAdmin(11331126, pwd, "万余里")
	services.AdminService.SaveAInfo(user)

	pwd = encryptPwd(11331127, "123456")
	user = services.AdminService.NewAdmin(11331127, pwd, "陈向阳")
	services.AdminService.SaveAInfo(user)
}

/*
* Add some classroom information.
 */
func initClassroom() {
	classroom := services.ClassroomService.NewClassroom("东校园", "公教楼", "A201", 50)
	services.ClassroomService.SaveAInfo(classroom)

	classroom = services.ClassroomService.NewClassroom("东校园", "公教楼", "A202", 50)
	services.ClassroomService.SaveAInfo(classroom)

	classroom = services.ClassroomService.NewClassroom("东校园", "公教楼", "C102", 100)
	services.ClassroomService.SaveAInfo(classroom)

	classroom = services.ClassroomService.NewClassroom("东校园", "公教楼", "C101", 200)
	services.ClassroomService.SaveAInfo(classroom)

	classroom = services.ClassroomService.NewClassroom("南校园", "第一教学楼", "A104", 100)
	services.ClassroomService.SaveAInfo(classroom)
}

/*
* Add some department information.
 */
func initDep() {
	dep := services.DepartmentService.NewDeparment("团委", "位于东校明德园6号，负责学生活动审批", 1, "initial")
	services.DepartmentService.SaveAInfo(dep)

	dep = services.DepartmentService.NewDeparment("学院团委", "该部门主要负责本院活动的审批", 2, "middle")
	services.DepartmentService.SaveAInfo(dep)

	dep = services.DepartmentService.NewDeparment("保卫办", "该部门主要负责对于活动和场所的审查", 3, "final")
	services.DepartmentService.SaveAInfo(dep)
}

/*
* Add some reservation information.
 */
func initRes() {
	start := lessonToTime("2018-05-23 第三节课")
	end := lessonToTime("2018-05-23 第五节课")
	res := services.ReservationService.NewReservation("社团面试", start, end, 1, 15331124, "东校园学生会", 10331124, "", 0)
	services.ReservationService.SaveAInfo(res)

	start = lessonToTime("2018-05-24 第八节课")
	end = lessonToTime("2018-05-24 第十节课")
	res = services.ReservationService.NewReservation("生日party", start, end, 2, 15331124, "软工4班", 10331124, "不允许教学区进行该活动", 3)
	services.ReservationService.SaveAInfo(res)

	start = lessonToTime("2018-06-01 第十三节课")
	end = lessonToTime("2018-06-01 第十四节课")
	res = services.ReservationService.NewReservation("班级会议", start, end, 3, 15331124, "软工4班", 10331126, "", 2)
	services.ReservationService.SaveAInfo(res)

}

func init() {
	initStudent()
	initAdmin()
	initApprover()
	initClassroom()
	initDep()
	initRes()
}
