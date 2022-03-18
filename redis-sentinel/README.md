## 使用docker-compose实现redis哨兵模式
> Sentinel（哨岗、哨兵）是Redis的高可用性（high availability）解决方案：由一个或多个Sentinel实例（instance）组成的Sentinel系统（system）可以监视任意多个主服务器，以及这些主服务器属下的所有从服务器，并在被监视的主服务器进入下线状态时，自动将下线主服务器属下的某个从服务器升级为新的主服务器，然后由新的主服务器代替已下线的主服务器继续处理命令请求。

- [参考文章](https://www.jianshu.com/p/d8c5f5d33a29)
## 目录结构
```shell
redis-sentinel/
├── sentinel
│   ├── docker-compose.yml
│   └── redis-sentinel-1.conf
│   └── redis-sentinel-2.conf
│   └── redis-sentinel-3.conf
├── server
    ├── docker-compose.yml
    └── redis-master.conf
    └── redis-slave1.conf
    └── redis-slave2.conf
```
## 操作
- 启动server，在server目录下执行
```shell
docker-compose up
```
- 启动sentinel，在sentinel目录下执行
```shell
docker-compose up
```
- 观察sentinel的日志
![观察sentinel的日志](https://cdn.jsdelivr.net/gh/mouweng/FigureBed/img/20220319000440.jpg)
可以看到三个sentinel节点都监视了master节点
- 下线master节点
```shell
docker stop redis-server-master
```
- 观察sentinel日志
![](https://cdn.jsdelivr.net/gh/mouweng/FigureBed/img/20220319002218.jpg)
主节点已经变成redis-server-slave-1

- 观察server日志
![观察server日志](https://cdn.jsdelivr.net/gh/mouweng/FigureBed/img/20220319002141.jpg)
主节点已经变成redis-server-slave-1

- 操作新生成的master节点
```shell
docker exec -it redis-server-slave-1 /bin/bash
redis-cli -p 6380 -a 123456

127.0.0.1:6380> set hello "hi"
OK
```
发现可以插入，说明此时salve-1节点已经升级为master节点

- 重启redis-server-master节点
![重启redis-server-master节点](https://cdn.jsdelivr.net/gh/mouweng/FigureBed/img/20220319001821.jpg)
原先的redis-server-master变为新的master节点（redis-server-slave-1）的slave节点。
- 操作redis-server-master节点
```shell
docker exec -it redis-server-master /bin/bash
redis-cli -p 6379 -a 123456

127.0.0.1:6379> set hello "hihihi"
(error) READONLY You can't write against a read only replica.
127.0.0.1:6379> get hello
"hi"
```
## Sentinel命令
- **PING** ：返回 PONG 。
- **SENTINEL masters [name]**：列出所有被监视的主服务器，以及这些主服务器的当前状态。
- **SENTINEL slaves [name]**：列出给定主服务器的所有从服务器，以及这些从服务器的当前状态。