# redis集群Cluster模式
- [集群搭建概述](http://kmanong.top/kmn/qxw/form/article?id=74467&cate=72)
- [修复redis cluster 插槽](https://www.cnblogs.com/xizhunet/p/15346765.html)
## 部署
- 使用docker-compose部署集群
```shell
docker-compose up -d
```

## 测试
### 检查当前集群状况
```shell
127.0.0.1:6371> cluster nodes
08011ef4871d2dbf0c1bd407498539ba37909b39 :6371@16371 myself,master - 0 0 0 connected
```
可以看到目前各个节点还只能返回自己的信息，每个节点还不能感知到彼此。
### 节点握手
- 在某个节点上执行`cluster meet {ip} {port}`命令，达到两个节点间的握手，这两个节点就组成了一个真正的彼此感知的集群，之后两个节点间会定期通过`ping/pong`消息进行正常的节点通信
- 在集群中任意节点上执行`cluster meet {ip} {port}`命令，添加尚未加入集群的新节点
- 所有节点全部加入之后可以看到集群中所有节点信息

```shell
127.0.0.1:6371> cluster meet 127.0.0.1 6372
OK
127.0.0.1:6371> cluster meet 127.0.0.1 6373
OK
127.0.0.1:6371> cluster nodes
e4ea5e44833a161601a01d013ea50b084c57fc39 127.0.0.1:6373@16373 master - 0 1647663768268 0 connected
9aee3c8bd78f51507d848cf0c2cf42068ff4b64b 127.0.0.1:6372@16372 master - 0 1647663766263 2 connected
08011ef4871d2dbf0c1bd407498539ba37909b39 127.0.0.1:6371@16371 myself,master - 0 1647663765000 1 connected
```
### 分配槽
- 节点建立握手之后，集群还处于下线状态，无法执行写操作。
```shell
127.0.0.1:6371> set hello world
(error) CLUSTERDOWN Hash slot not served
```
- 进行槽修复
```shell
$ redis-cli --cluster fix 127.0.0.1:6371
```
- 请求重定向访问

在集群中某个节点读写不属于此节点的数据会返回错误`(error) MOVED 5798 172.20.0.6:6372`,为了减少手动切换的环节，在开启客户端时可以添加`-c`参数，开启请求重定向，详细命令`redis-cli -p 6371 -c`，这样以后操作跨节点时会自动跳转到相应的节点
```shell
redis-cli -p 6371 -c
```
![请求重定向访问](https://cdn.jsdelivr.net/gh/mouweng/FigureBed/img/20220319122605.jpg)
## 官方集群快速搭建工具
> 在redis5.x中更是可以直接使用redis-cli命令来直接完成集群的一键搭建
- 搭建
```
redis-cli --cluster create 127.0.0.1:6371 127.0.0.1:6372 127.0.0.1:6373
```
![集群搭建](https://cdn.jsdelivr.net/gh/mouweng/FigureBed/img/20220319121536.jpg)
以下这种方式可以搭建3主3从模式
```
redis-cli --cluster create 127.0.0.1:7000 127.0.0.1:7001 127.0.0.1:7002 127.0.0.1:7003 127.0.0.1:7004 127.0.0.1:7005 --cluster-replicas 
```
- 请求重定向
```shell
redis-cli -p 6371 -c
```
![请求重定向](https://cdn.jsdelivr.net/gh/mouweng/FigureBed/img/20220319121705.jpg)