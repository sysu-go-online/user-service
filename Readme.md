# go-online service end

## 简介

## 维护者

## 依赖的环境变量

- DATABASE_ADDRESS 数据库地址，默认为localhost
- DATABASE_PORT 数据库端口，默认为3306
- DEVELOP 是否为开发环境，默认为false
- DOCKER_ADDRESS 容器服务地址，默认为localhost
- REDIS_ADDRESS redis地址，默认为localhost
- REDIS_PORT redis端口，默认为6379

## 依赖的外部软件

- mysql

  需要包含有`mydb`数据库，具体说明参见技术文档

- redis

  用来存放失效的jwt，具体说明见技术文档

## 运行方式

`go run main.go`