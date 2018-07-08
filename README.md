# class-reservation

![test](https://www.travis-ci.org/kangbb/ccrsystem.svg?branch=master)

一个简单的校园课室预定系统。采用前后端分离的方式开发，前端采用`vuex`框架，后端采用`golang`编写。

## 小组分工与贡献率
|学号|姓名|分工|贡献率|
|----|----|---|------|
|15331134|亢辈辈(组长)|项目管理、后端代码、软件设计文档、软件测试文档|30%|
|15331090|郭比仪|项目UI设计、UI交互设计、软件设计文档|17%|
|15331346|严晓珊|前端Student模块、Approver模块|18%|
|15331083|符淼|前端Admin模块、软件设计文档|18%|
|15331003|敖津聪|安装部署说明、用户收测、软件需求规格说明书|17%|

## 制品与贡献率
|   |制品|亢辈辈|郭比仪|严晓珊|符淼|敖津聪|
|---|---|:------:|:------:|:-----:|:----:|:-----:|
|源码|前端Student模块|-|-|100%|-|-|
|源码|前端Approver模块|-|-|100%|-|-|
|源码|前端Admin模块|-|-|-|100%|-|
|源码|后端及测试代码|100%|-|-|-|-|
|文档|Install Instruction|-|-|-|-|100%|
|文档|SD|35%|30%|-|35%|-|
|文档|SRS|-|-|-|-|100%|
|文档|User Handset|-|-|-|-|100%|
|文档|SoftWare Test|100%|-|-|-|-|
## 文档

### API文档

API使用`showdoc`进行管理，地址如下：

[https://www.showdoc.cc/web/#/69795491818399](https://www.showdoc.cc/web/#/69795491818399)

<font color="#0366d6">密码：123456</font>

### 数据字典

定义了前后端数据交互的数据表格式以及数据库中对应的数据表格式。

[https://www.showdoc.cc/web/#/69795491818399](https://www.showdoc.cc/web/#/69795491818399)

<font color="#0366d6">密码：123456</font>

### 技术文档

描述了数据库设计相关内容。

[https://www.showdoc.cc/web/#/69795491818399](https://www.showdoc.cc/web/#/69795491818399)

<font color="#0366d6">密码：123456</font>

## 后端

后端采用golang开发，如果需要本地化测试，可以按照`testconfig`文件夹下的`README.md`文档进行配置，测试。

## 前端

进入文件夹
```
$ cd front_end
```

安装依赖
```
$ npm install
```

运行
```
$ npm run dev
```

编译
```
$ npm build .
```
静态文件部署测试
```
$ node server.js
```

## 遇到的问题及解决

### 解跨域问题

[跨域资源共享 CORS 详解](http://www.ruanyifeng.com/blog/2016/04/cors.html)

[MDN HTTP访问控制(CORS)](https://developer.mozilla.org/zh-CN/docs/Web/HTTP/Access_control_CORS)

### 日志

#### 日志管理策略之一——日志轮转

[日志轮转](https://www.cnblogs.com/guo-xiang/p/5806563.html)


[logrotate简单实现各种日志自动轮转](https://blog.csdn.net/nerissa/article/details/17707913)


