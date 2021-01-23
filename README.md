# Moments

A system refers to `Wechat Momemts`.

Check this [link](https://github.com/golang-standards/project-layout) 

Gin + MongoDB + Docker + RocketMQ

[gin](https://github.com/gin-gonic/gin)

[mongo-go-driver](https://github.com/mongodb/mongo-go-driver)

[rocketmq-client-go](https://github.com/apache/rocketmq-client-go/blob/master/docs/Introduction.md)

[golang mock](https://github.com/golang/mock)


参考文章[go gin example](https://github.com/EDDYCJY/go-gin-example/blob/master/README_ZH.md)
[作者系列博客](https://eddycjy.com/posts/go/gin/2018-02-14-jwt/)


# 需求

-[ ] swagger生成接口文档
-[ ] 鉴权middlewares，jwt
-[ ] logger的配置化，debug选项可配置，全局开启debug日志，ini包比较易用
-[ ] 数据值合法性判断，各种id的取值区间，存在与否等
-[ ] 透传context，日志追加tracking-id
-[ ] 参数验证使用validator，参考[链接](https://blog.xizhibei.me/2019/06/16/an-introduction-to-golang-validator/)


# 重构
-[ ] 数据库传参统一成传结构体，精简掉map和结构体的冗余
-[ ] service层和model层要确认互相调用的形式，是否需要receiver，做好风格统一
-[ ] 单测调整，重点在service层，简单代码不考虑单测
-[x] 重构json tag名，是全部用驼峰风格，不要引入下划线
-[x] 数据库封装得不好
-[ ] 统一调整日志，日志规范化，抛错规范化
-[ ] 重构函数命名，注意service层和model层的区别，model层应该尽量通用，service层尽量贴合业务场景取取名
-[x] 文件结构简化
-[x] 消息队列代码封装
-[ ] 引入Makefile管理常用命令
-[ ] 引入gmock和sqlmock，见文章 https://draveness.me/golang-101/