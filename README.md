## 基于 *beego* 的 CMS 系统

> @authro Zhengqian <zhuzhengqian@gmail.com>



## 依赖安装

```
dep ensure

```

## 数据库配置

```
cd conf && cp app.conf.example app.conf

## 应用名称
appname = cms

## 应用访问端口
httpport = 8080

## 允许模式
runmode = dev

## 数据库配置
mysqluser = root
mysqlpass = root
mysqlport = 3307
mysqlhost = 127.0.0.1
mysqldb   = cms

```

## 启用

```
go run main.go

```