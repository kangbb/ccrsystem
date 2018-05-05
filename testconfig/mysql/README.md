# MySql镜像配置

## 修改配置docker-compose.yml
主要修改`volumes`配置项，根据自己的情况挂载数据卷。例如，本人的配置是将本地目录`$HOME/Work/data/dbslave`挂载到镜像的`/var/lib/mysql`目录下。

## 启动
```bash
$ sudo chmod update.sh +x
$ ./update.sh
```

## 主从配置
参考网址：[https://blog.csdn.net/kiloveyousmile/article/details/79833043](https://blog.csdn.net/kiloveyousmile/article/details/79833043)