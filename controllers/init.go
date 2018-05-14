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
	user := services.StudentService.NewStudent(15331124, pwd, "王芳", 1)
	services.StudentService.SaveAInfo(user)

	pwd = encryptPwd(15331125, "123456")
	user = services.StudentService.NewStudent(15331125, pwd, "赵红", 2)
	services.StudentService.SaveAInfo(user)

	pwd = encryptPwd(16331124, "123456")
	user = services.StudentService.NewStudent(16331124, pwd, "王芸", 3)
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

	pwd = encryptPwd(10031124, "123456")
	user = services.ApproverService.NewApprover(10031124, pwd, "李芸", 3)
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

	pwd = encryptPwd(11031124, "123456")
	user = services.AdminService.NewAdmin(11031124, pwd, "万余里")
	services.AdminService.SaveAInfo(user)
}

/*
* Add some classroom information.
 */
func initClassroom() {
	classroom := services.ClassroomService.NewClassroom("东校园", "公教楼", "A204", 180)
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
* Add some organization information.
 */
func initOrg() {
	org := services.OrganizationService.NewOrganization("个人", "这是所有特立独行的人的组织")
	services.OrganizationService.SaveAInfo(org)

	org = services.OrganizationService.NewOrganization("中山大学东校园学生会", "致力于为学生提供优质服务，方便学生们的学习生活")
	services.OrganizationService.SaveAInfo(org)

	org = services.OrganizationService.NewOrganization("中大青年", "汇聚中大资讯，树洞欢迎你")
	services.OrganizationService.SaveAInfo(org)

}

/*
* Add some reservation information.
 */
func initRes() {
	start := lessonToTime("2018-05-23 第三节课")
	end := lessonToTime("2018-05-23 第五节课")
	approvalNote := "批准"
	res := services.ReservationService.NewReservation("审批中", start, end, 1, "社团面试",
		approvalNote, 15331124, 10331124, 1, 1)
	services.ReservationService.SaveAInfo(res)

	start = lessonToTime("2018-05-24 第八节课")
	end = lessonToTime("2018-05-24 第十节课")
	approvalNote = "不允许的活动"

	res = services.ReservationService.NewReservation("预定失败", start, end, 1, "生日party",
		approvalNote, 15331125, 10331124, 1, 2)
	services.ReservationService.SaveAInfo(res)
}

// func init() {
// 	initStudent()
// 	initAdmin()
// 	initApprover()
// 	initClassroom()
// 	initDep()
// 	initRes()
// }
