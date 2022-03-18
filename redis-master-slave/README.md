# redis主从分布模式

## 部署
- 使用docker-compose部署集群
```shell
docker-compose up -d
```

## 测试
- 往master集群写入数据
```shell
docker exec -it redis-master /bin/bash
redis-cli -h 127.0.0.1 -p 6379 -a redispwd
set name yifanweng
keys *
```
- 从slave1读取数据
```shell
docker exec -it redis-slave-1 /bin/bash
redis-cli -h 127.0.0.1 -p 6379 -a redispwd
keys *
get name
```
- 从slave2读取数据
```shell
docker exec -it redis-slave-2 /bin/bash
redis-cli -h 127.0.0.1 -p 6379 -a redispwd
keys *
get name
```

## go-redis
```
cd redis-master-slave/goTest
go run main.go
```