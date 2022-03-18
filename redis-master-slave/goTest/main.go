package main

import (
	"fmt"

	redis "github.com/go-redis/redis"
)

// 初始化连接
func main(){
	rdb_master := redis.NewClient(&redis.Options{
		Addr:     "localhost:87",
		Password: "redispwd", // no password set
		DB:       0,  // use default DB
	});
	defer rdb_master.Close()

	rdb_slave1 := redis.NewClient(&redis.Options{
		Addr:     "localhost:88",
		Password: "redispwd", // no password set
		DB:       0,  // use default DB
	});
	defer rdb_slave1.Close()

	rdb_slave2 := redis.NewClient(&redis.Options{
		Addr:     "localhost:89",
		Password: "redispwd", // no password set
		DB:       0,  // use default DB
	});
	defer rdb_slave2.Close()

	// 连接测试
	pong, err := rdb_master.Ping().Result()
	if err != nil {
		fmt.Println(err)
		return;
	}
	fmt.Printf("connection success : %s\n", pong)

	// 往master添加数据
	err = rdb_master.Set("hello", "hi", 0).Err();
	if err != nil {
		fmt.Println(err)
		return;
	}
	fmt.Println("add data success!")

	// 从master获取数据
	val1, err := rdb_master.Get("hello").Result()
	if err != nil {
		fmt.Println(err)
		return;
	}
	fmt.Println("get master success : ", val1)

	// 从slave1获取数据
	val2, err := rdb_slave1.Get("hello").Result()
	if err != nil {
		fmt.Println(err)
		return;
	}
	fmt.Println("get slave1 success : ", val2)

	// 从slave2获取数据
	val3, err := rdb_slave2.Get("hello").Result()
	if err != nil {
		fmt.Println(err)
		return;
	}
	fmt.Println("get slave2 success : ", val3)

}