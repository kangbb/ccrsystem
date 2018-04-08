# class-reservation
## 路径设计
```
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
1. 节数<----->时间
2. 默认从一节课的开始计算
3. 节数展示给前端，时间存储在数据库中
格式：
1. 节数 eg:"2016-10-01 第一节"
2. 时间 eg:"2016-10-01 08:00"