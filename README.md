# class-reservation

一个简单的校园课室预定系统。采用前后端分离的方式开发，前端采用`vuex`框架，后端采用`golang`编写。

## 文档

### API文档

API使用`showdoc`进行管理，地址如下：

[https://www.showdoc.cc/web/#/69795491818399](https://www.showdoc.cc/web/#/69795491818399)

### 数据字典

定义了前后端数据交互的数据表格式以及数据库中对应的数据表格式。

[https://www.showdoc.cc/web/#/69795491818399](https://www.showdoc.cc/web/#/69795491818399)


## 关于后端

后端采用golang开发，如果需要本地化测试，可以按照`testconfig`文件夹下的`README.md`文档进行配置，测试。

## 关于前段

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

## 遇到的问题及解决

### 解跨域问题

[跨域资源共享 CORS 详解](http://www.ruanyifeng.com/blog/2016/04/cors.html)