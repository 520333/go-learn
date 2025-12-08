package redisCli

import (
	"context"
	"fmt"
	"log"

	"github.com/redis/go-redis/v9"
)

func BitMapSetGet() {
	opt, err := redis.ParseURL("redis://default:123456@192.168.50.100:6379/0")
	if err != nil {
		log.Fatalln(err)
	}
	client := redis.NewClient(opt)
	ctx := context.Background()
	client.Set(ctx, "userLog", "go", 0)
	client.Del(ctx, "userLog")
	// 10011
	client.SetBit(ctx, "userLog", 0, 1)
	client.SetBit(ctx, "userLog", 1, 0)
	client.SetBit(ctx, "userLog", 2, 0)
	client.SetBit(ctx, "userLog", 3, 1)
	client.SetBit(ctx, "userLog", 4, 1)
	client.SetBit(ctx, "userLog", 7, 1)
	client.SetBit(ctx, "userLog", 8, 1)
	client.SetBit(ctx, "userLog", 15, 1)
	client.SetBit(ctx, "userLog", 16, 1)

	// 获取
	for i := 0; i < 5; i++ {
		fmt.Println(client.GetBit(ctx, "userLog", int64(i)).Result())
	}
	// 未设置的位 为 0
	fmt.Println(client.GetBit(ctx, "userLog", 5).Result())
	fmt.Println("==================")
	fmt.Println(client.BitCount(ctx, "userLog", &redis.BitCount{
		// 1byte== 8bits
		//Start: 0,
		// 默认第一个字节
		//End: 0,
		Start: 0,
		End:   -1,
	}).Result())
	// 大offset 2^32-1
	//client.SetBit(ctx, "big", 1<<32-1, 1)
	//fmt.Println(client.GetBit(ctx, "big", 1<<32-1).Result())
	//fmt.Println(client.GetBit(ctx, "big", 1<<32-2).Result())
	//127.0.0.1:6379> memory usage big
	//(integer) 671,088,688

}
