package services

import (
	"github.com/kangbb/ccrsystem/core/models/service"
)

func addStudent() {
	pwd := encryptPwd(15331124, "123456")
	user := service.StudentService.NewStudent(15331124, pwd, "王芳", true)
	service.StudentService.SaveAInfo(&user)
	// if !state || err != nil {
	// 	panic(err)
	// }

	pwd = encryptPwd(15331125, "123456")
	user = service.StudentService.NewStudent(15331125, pwd, "赵红", true)
	service.StudentService.SaveAInfo(&user)
	// if !state || err != nil {
	// 	panic(err)
	// }

	pwd = encryptPwd(16331124, "123456")
	user = service.StudentService.NewStudent(16331124, pwd, "王芸", true)
	service.StudentService.SaveAInfo(&user)
	// if !state || err != nil {
	// 	panic(err)
	// }
}

func addApprover() {
	pwd := encryptPwd(10331124, "123456")
	user := service.ApproverService.NewApprover(10331124, pwd, "李怡", 1, true)
	service.ApproverService.SaveAInfo(&user)
	// if !state || err != nil {
	// 	panic(err)
	// }

	pwd = encryptPwd(10331125, "123456")
	user = service.ApproverService.NewApprover(10331125, pwd, "宋晓", 2, true)
	service.ApproverService.SaveAInfo(&user)
	// if !state || err != nil {
	// 	panic(err)
	// }

	pwd = encryptPwd(10031124, "123456")
	user = service.ApproverService.NewApprover(10031124, pwd, "李芸", 3, true)
	service.ApproverService.SaveAInfo(&user)
	// if !state || err != nil {
	// 	panic(err)
	// }
}

func addAdmin() {
	pwd := encryptPwd(11331124, "123456")
	user := service.AdminService.NewAdmin(11331124, pwd, "赵毅", true)
	service.AdminService.SaveAInfo(&user)
	// if !state || err != nil {
	// 	panic(err)
	// }

	pwd = encryptPwd(11331125, "123456")
	user = service.AdminService.NewAdmin(11331125, pwd, "李阳", true)
	service.AdminService.SaveAInfo(&user)
	// if !state || err != nil {
	// 	panic(err)
	// }

	pwd = encryptPwd(11031124, "123456")
	user = service.AdminService.NewAdmin(11031124, pwd, "万余里", true)
	service.AdminService.SaveAInfo(&user)
	// if !state || err != nil {
	// 	panic(err)
	// }
}

func addClassroom() {
	classroom := service.ClassroomService.NewClassroom("东校园", "公教楼", "A204", 180)
	service.ClassroomService.SaveAInfo(&classroom)
	// if !state || err != nil {
	// 	panic(err)
	// }

	classroom = service.ClassroomService.NewClassroom("南校园", "第一教学楼", "A104", 100)
	service.ClassroomService.SaveAInfo(&classroom)
	// if !state || err != nil {
	// 	panic(err)
	// }
}
func addDep() {
	dep := service.DepartmentService.NewDeparment("团委", 1, "initial")
	service.DepartmentService.SaveAInfo(&dep)
	// if !state || err != nil {
	// 	panic(err)
	// }

	dep = service.DepartmentService.NewDeparment("学院团委", 2, "middle")
	service.DepartmentService.SaveAInfo(&dep)
	// if !state || err != nil {
	// 	panic(err)
	// }

	dep = service.DepartmentService.NewDeparment("保卫办", 3, "final")
	service.DepartmentService.SaveAInfo(&dep)
	// if !state || err != nil {
	// 	panic(err)
	// }
}
func addRes() {
	start := lessonToTime("2018-05-23 第三节课")
	end := lessonToTime("2018-05-23 第五节课")
	approvalNote := "批准"
	res := service.ReservationService.NewReservation("审批中", start, end, 1, "社团面试",
		approvalNote, 15331124, 10331124, 1)
	result, err := service.ReservationService.SaveAInfo(&res)
	if !result || err != nil {
		panic(err)
	}

	start = lessonToTime("2018-05-24 第八节课")
	end = lessonToTime("2018-05-24 第十节课")
	approvalNote = "不允许的活动"

	res = service.ReservationService.NewReservation("预定失败", start, end, 1, "生日party",
		approvalNote, 15331124, 10331124, 1)
	result, err = service.ReservationService.SaveAInfo(&res)
	if !result || err != nil {
		panic(err)
	}

}

// func init() {
// 	addStudent()
// 	addAdmin()
// 	addApprover()
// 	addClassroom()
// 	addDep()
// 	addRes()
// }
