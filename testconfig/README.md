# 校园课室预定系统测试说明

本项目采用docker容器部署测试，如有需要进行本地化测试，请先安装如下工具：
>docker<br/>
>docker-compose

具体步骤请参考官方网站：

[Docker - Build, Ship, and Run Any App, Anywhere](https://www.docker.com/)
## 测试配置

在进行测试前，需要进行简单的配置，具体包括以下步骤：
>1. 配置mysql数据库
>2. 修改部分文件，适应测试环境
>3. 启动测试程序，分析测试结果


### 配置mysql数据库
**1. 构建镜像**

由于使用docker部署测试，您需要先构建合适的mysql镜像。具体的配置和`Dockerfile`文档已经编写完成，只需要在项目根目录执行以下命令即可：
```bash
$ cd testconfig/mysql/dbmaster
$ docker run build -t registry.cn-shenzhen.aliyuncs.com/selfmysql/master:latest .
$ cd ../dbslave
$ docker run build -t registry.cn-shenzhen.aliyuncs.com/selfmysql/slave .
```
运行命令后，你将的到如下两个镜像：
```basj
$ docker images
registry.cn-shenzhen.aliyuncs.com/selfmysql/slave     latest              ad0a4971c1a5        4 weeks ago         371 MB
registry.cn-shenzhen.aliyuncs.com/selfmysql/master    latest              cfa626cc975d        4 weeks ago         371 MB
```

**2. 修改docker-compose文件配置**

对`testconfig/mysql`目录下的`docker-compose.yml`文件进行修改。 主要修改`volumes`配置项，根据自己的情况挂载数据卷。例如，文件默认的配置为：
> 将本地目录`$HOME/Work/data/dbslave`挂载到镜像的`/var/lib/mysql`目录下。<br />
> 将本地目录`$HOME/Work/data/dbmaster`挂载到镜像的`/var/lib/mysql`目录下。

**3. 启动**

在`testconfig/mysql`目录下运行如下命令：
```bash
$ docker-compose up -d
```
在`testconfig/mysql`目录下使用如下命令查看是否启动成功:
```
$ docker-compose logs
```

**4. 主从配置**

参考网址：[https://blog.csdn.net/kiloveyousmile/article/details/79833043](https://blog.csdn.net/kiloveyousmile/article/details/79833043)

**5. 备注**

如果你并喜欢自己动手构建镜像，也可以直接在`testconfig/mysql`目录下运行如下命令：
```bash
$ sudo chmod +x update.sh
$ ./update.sh
```
之后，可以直接跳到第4步，进行主从配置即可。

### 修改部分文件

分别按照如下方式修改指定文件：

<font color="gray">models/entities/init.go</font>
```
// Just for test
MasterEngine, err = xorm.NewEngine("mysql", "root:master@tcp(localhost:3307)/ccrsystem?charset=utf8&parseTime=true")
// True connection for mysql
// MasterEngine, err = xorm.NewEngine("mysql", "root:master@tcp(dbmaster:3306)/ccrsystem?charset=utf8&parseTime=true")

// Just for test
SlaveEngine, err = xorm.NewEngine("mysql", "root:slave@tcp(localhost:3308)/ccrsystem?charset=utf8&parseTime=true")
// True connection for mysql
// SlaveEngine, err = xorm.NewEngine("mysql", "root:slave@tcp(dbslave:3306)/ccrsystem?charset=utf8&parseTime=true")

```

### 在指定目录启动测试

由于时间原因，本系统编写的测试较为简单，主要分布在根目录下。进入根目录，执行如下命令即可：

```bash
$ go test .
```