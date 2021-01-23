# 新手看我

这个简单的微信朋友圈后台（功能有阉割）是基于Gin框架做的小demo，本身是作者用于学习Gin框架的一个简单产物。

作者是个没有Web开发经验的菜鸟，整体项目结构极大参考了[go gin example](https://github.com/EDDYCJY/go-gin-example/blob/master/README_ZH.md)。

整体流程遵循MVC结构：`routers -> service -> models`。

先设计好功能点，再根据功能点抽象出接口。


# Docker环境创建

## 手动创建环境

1. docker image pull ubuntu:20.04   # 从docker hub拉取ubuntu带标签20.04的镜像到本地
2. docker run --name myExample -itd -p 8000-9000:8000-9000 ubuntu:20.04 /bin/bash   # 基于ubuntu:20.04镜像运行叫myExample的容器，容器中运行/bin/bash，对外暴露端口8000-9000
3. docker attach {your container_id}    # 进入你的容器中
4. 


## 基于Dockerfile创建环境（推荐）




```shell
#! /bin/bash
# mongo run.sh
db="/usr/local/mongodb/"
${db}"bin/mongod" -f ${db}"db.conf"
```

```yaml
# mongodb config
net:
  port: 8080
  bindIpAll: true
systemLog:
  destination: file
  path: "/usr/local/mongodb/log/m.log"
  logAppend: true
storage:
  journal:
    enabled: true
  dbPath: "/usr/local/mongodb/data"
processManagement:
   fork: true
setParameter:
  enableLocalhostAuthBypass: false
```

```shell
#! /bin/bash
mqpath="/usr/local/rocketmq"
export PATH=${PATH}:${mqpath}/bin

start(){
    nohup mqnamesrv -c ${mqpath}/namesrv.properties 2>&1 &
    sleep 1
    nohup mqbroker -c ${mqpath}/broker.properties 2>&1 &
}

stop(){
    mqshutdown namesrv
    sleep 1
    mqshutdown broker
}

if [ "$1" == "start" ];
then
    start
elif [ "$1" == "stop" ];
then
    stop
fi
```

# rocketmq broker
```
namesrvAddr=127.0.0.1:8000
brokerIP1=127.0.0.1
brokerClusterName=DefaultCluster
brokerName=broker-a
brokerId=0
deleteWhen=04
fileReservedTime=48
brokerRole=SYNC_MASTER
flushDiskType=ASYNC_FLUSH
listenPort=8100
storePathRootDir=/usr/local/rocketmq/hxh/
```
# rocketmq namesrv
```
listenPort=8000
```
