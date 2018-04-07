# class-reservation
## 路径设计
https://example.com/student   //获取、修改学生信息
https://example.com/admin   //获取、修改教师信息
## 数据结构
```go
type ClassroomInfo struct {
  ClassroomId   int `xorm:"pk autoincr index 'id'"`  //唯一编号，自增即可
  ClassroomCampus string `xorm:""`  //教室校区
  ClassroomBuilding string   //教室楼名
  ClassroomNum    string   //教室编号
	Capicity      int //教室容量
}

//ReservationInfo store reservation information
type ReservationInfo struct {
	ResId       int    `xorm:"autoincr pk 'id'"`
	ResState    string `xorm:"pk index(res)"`
	StartTime   string `xorm:"pk index(res)"`
	EndTIme     string `xorm:"pk index(res)"`
	ResReason   string
	ResCapicity int `xorm:"pk index(res)"`
	StudentId   int
	ApproverId  int
	ClassroomId int `xorm:"pk index(res)"`
}
```
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

## 一些必须的业务
时间转换： 
1. 节数<----->时间
2. 默认从一节课的开始计算
3. 节数展示给前端，时间存储在数据库中
