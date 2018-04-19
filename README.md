# class-reservation
## 路径设计
```bash
# 静态文件控制 #
http://www.example.com/           //主页,登录页
http://www.example.com/student    //学生主页
http://www.example.com/admin      //管理员主页
http://www.example.com/approver   //审批者主页
# 功能接口 #
http://www.example.com/signout           //用户退出
http://www.example.com/student/signin    //学生登录
http://www.example.com/admin/signin      //管理员主页
http://www.example.com/approver/signin   //审批者主页
# 数据api接口 #
http://www.example.com/api/users/student        //获取(GET)、修改(PUT)单个学生信息(通过session); 当前登录用户
http://www.example.com/api/users/students       //获取学生信息列表, 增加(POST)学生信息; 管理员权限 
http://www.example.com/api/users/students/{id}  //获取(GET)、修改(PUT)、删除(DELETE)单个学生信息(通过id); 管理员权限
http://www.example.com/api/users/approver       //获取(GET)、修改(PUT)单个审批者信息(通过session); 当前登录用户
http://www.example.com/api/users/approvers      //获取审批者信息列表，增加(POST)审批者信息; 管理员权限
http://www.example.com/api/users/approvers/{id} //获取(GET)、修改(PUT)、删除(DELETE)单个审批者信息(通过id); 管理员权限
http://www.example.com/api/user/admin           //获取(GET)、修改(PUT)单个管理员信息(通过session); 当前登录用户
http://www.example.com/api/users/admins         //获取管理员信息列表, 增加(POST)管理员信息; 管理员权限
http://www.example.com/api/users/admins/{id}    //获取(GET)、修改(PUT)、删除(DELETE)单个管理员信息(通过id); 管理员权限
http://www.example.com/api/classrooms           //获取课室信息列表; 管理员权限
http://www.example.com/api/classrooms/{id}      //获取(GET)、修改(PUT)、删除(DELETE)、增加(POST)单个课室信息;获取为所有人，增删改为管理员权限
http://www.example.com/api/classrooms/state     //获取教室的状态(学生根据各种条件查询可用教室); 学生权限
http://www.example.com/api/reservations/{id}                //获取(GET)、修改(PUT)、删除(DELETE)单个预定信息(通过id); 当前登录用户
http://www.example.com/api/users/student/reservations       //获取学生查询的课室预定信息,增加(POST)课室预定; 当前登录用户
http://www.example.com/api/users/approver/reservations      //获取审批者查询的课室预定信息; 当前登录用户
http://www.example.com/api/departments                       //获取部门信息列表,增加(POST)单个部门信息; 管理员权限
http://www.example.com/api/departments/{id}                  //获取(GET)、修改(PUT)、删除(DELETE)单个课室信息;管理员权限
```

注意：<br />
|业务需求|请求方法|数据位置|备注|
|:-:|:-:|:-:|:-:|
|更新数据|PUT|body|
|创建数据|POST|body|
|获取数据|GET|body|
|获取数据|POST|path|表格提交，例外
|删除数据|DELETE|body|

## 数据结构约定
```go
// StudentInfo store student information
type StudentInfo struct {
	StudentId   int `xorm:"pk 'id'"`
	StudentPwd  string
	StudentName string
	Permission  bool
}

//AdminInfo store admin information
type AdminInfo struct {
	AdminId    int `xorm:"pk 'id'"`
	AdminPwd   string
	AdminName  string
	Permission bool
}

//ApproverInfo store approver information
type ApproverInfo struct {
	ApproverId   int `xorm:"pk 'id'"`
	ApproverPwd  string
	ApproverName string
	DepartmentId int
	Permission   bool
}

//ClassroomInfo store approver information
type ClassroomInfo struct {
	ClassroomId       int `xorm:"pk autoincr 'id'"`
	ClassroomCampus   string
	ClassroomBuilding string
	ClassroomNum      string
	Capicity          int
}

//ReservationInfo store reservation information
type ReservationInfo struct {
	ResId        int `xorm:"autoincr pk 'id'"`
	ResState     string
	StartTime    time.Time
	EndTime      time.Time
	ResReason    string
	ApprovalNote string
	DepartmentId int
	StudentId    int
	ApproverId   int
	ClassroomId  int `xorm:"'classroom_id'"`
}
/*
* ResState的值包括：
* 预定成功： 审批完成后的标志
* 审批中： 等待审批的标志
* 预定失败： 审批未通过的标志
* 等待下一个部门审批： 审批人员提交审批时的内容.
*
* ApprovalNote:
* 审批人员对于预定的审批备注
*/

//DepartmentInfo store approval department information
type DepartmentInfo struct {
	DepartmentId   int `xorm:"'id' autoincr pk"`
	DepartmentName string
	Order          int
	Note           string //note whether it is Initial approval department or final approval department
	//value: initial, middle, final
}
```
注：<br />
- 前端的数据结构字段名称需与上面一致，方便处理
## 关于session
后端处理，文件方式存储

## 查询课室使用情况
校区|楼名|教室|预定开始时间|预定结束时间|课室容量
:-:|:-:|:-:|:-:|:-:|:-:|
必填|必填|可查询 | 选填| 选填| 选填|

#### 第一步——课室查询
1. 校区+楼区
2. 校区+楼区+课室
3. 校区+楼区+容量

#### 第二步——预定查询：
课室ID+时间
1. 开始时间存在： StartTime < ResStartTime or StartTime > ResEndtTime
2. 结束时间存在： EndTime < ResStartTime or EndTime > ResEndtTime

默认开始时间为当前系统时间。所以至少有一个开始时间<br />
其他展示逻辑等，前段负责；这里负责返回原始数据

## 一些必需的业务
时间转换： 
1. 节数<=====>时间
2. 默认从一节课的开始计算
3. 节数展示给前端，时间存储在数据库中
格式：
1. 节数 eg:"2016-10-01 第一节"
2. 时间 eg:"2016-10-01 08:00"