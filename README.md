# class-reservation
## 路径设计
```
https://example.com/         主页，登录页面
https://example.com/student  学生主页
https://example.com/admin    管理员主页
https://example.com/approver 审批者主页
https://example.com/student/login         //学生登录
https://example.com/student/logout        //学生退出登录
https://example.com/student/classroom     //学生查询课室
https://example.com/student/reservation   //学生预定订单，查询订单，修改订单；第一个参数为操作方式
https://example.com/admin/login          //管理员登录
https://example.com/admin/logout       //管理员退出
https://example.com/admin//student      //管理员查询、修改、删除学生信息
https://example.com/admin/aprover       //管理员查询、修改、删除审批者信息
https://example.com/admin/classroom     //管理员查询、修改、删除教室信息
https://example.com/approver/login      //审批者登录
https://example.com/approver/logout     //审批者退出
https://example.com/student/reservation  //审批者审批预定
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
	ResId       int `xorm:"autoincr pk 'id'"`
	ResState    string
	StartTime   string
	EndTime     string
	ResReason   string
	StudentId   int
	ApproverId  int
	ClassroomId int `xorm:"'classroom_id'"`
}
```
注：<br />
- 前端的数据结构字段名称需与上面一致，方便处理
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