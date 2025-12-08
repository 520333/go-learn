package redisCli

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
)

func Client() {
	client := redis.NewClient(&redis.Options{
		Addr:                "192.168.50.100:6379",
		Username:            "default",
		Password:            "123456",
		CredentialsProvider: nil,
		DB:                  0,
		DialTimeout:         1 * time.Second,
		ReadTimeout:         1 * time.Second,
	})
	status := client.Ping(context.Background())
	fmt.Println(status.Result())
	err := client.Set(context.Background(), "name", "dawn", time.Second*5).Err()
	if err != nil {
		return
	}

	//另一种连接方式
	client2, err := redis.ParseURL("redis://default:123456@192.168.50.100:6379/0?dial_timeout=1")
	if err != nil {
		log.Fatalln(err)
	}
	rdb := redis.NewClient(client2)
	fmt.Println(rdb.Ping(context.Background()).Result())

}

func CmdString() {
	opt, err := redis.ParseURL("redis://default:123456@192.168.50.100:6379/0")
	if err != nil {
		log.Fatalln(err)
	}
	client := redis.NewClient(opt)
	ctx := context.Background()
	// 设置
	//status, err := client.Set(ctx, "name", "", 3*time.Second).Result()
	//status, err := client.Set(ctx, "name", "yoda", redis.KeepTTL).Result()

	// 设置条件 NX 不存在时 XX 存在时
	status := client.SetArgs(ctx, "name", "kubernetes", redis.SetArgs{
		Mode:     "NX",        // 不存在设置
		TTL:      0,           // 有效期 时间周期
		ExpireAt: time.Time{}, // 有效期时间点
		Get:      false,       // 是否返回原有值
		KeepTTL:  false,       // 是否保持原有有效期
	})

	fmt.Printf("写入: %v %v\n", status, err)

	// 获取
	result := client.Get(ctx, "name")
	val, err := result.Result()
	if errors.Is(err, redis.Nil) {
		fmt.Println("key not exists")
	} else if err != nil {
		fmt.Println(err)
	} else if val == "" {
		fmt.Println("value is empty")
	}
	fmt.Println(val)
}

func CmdStringAppendIncrDecr() {
	opt, err := redis.ParseURL("redis://default:123456@192.168.50.100:6379/0")
	if err != nil {
		log.Fatalln(err)
	}
	client := redis.NewClient(opt)
	ctx := context.Background()
	status := client.Set(ctx, "name", "kubernetes", 0)
	client.Append(ctx, "name", "go")
	client.Append(ctx, "name", "python")
	fmt.Println(status.Result())
	result := client.Get(ctx, "name")
	fmt.Println(result)

	client.Set(ctx, "counter", "0", 0)
	for i := 0; i <= 4; i++ {
		client.Incr(ctx, "counter")
	}
	client.IncrBy(ctx, "counter", 10)
	fmt.Println(client.Get(ctx, "counter"))

	client.DecrBy(ctx, "counter", 10)
	fmt.Println(client.Get(ctx, "counter"))
}

func CmdStringSub() {
	opt, err := redis.ParseURL("redis://default:123456@192.168.50.100:6379/0")
	if err != nil {
		log.Fatalln(err)
	}
	client := redis.NewClient(opt)
	ctx := context.Background()
	client.Set(ctx, "name", "kubernetes", 0)

	fmt.Println(client.GetRange(ctx, "name", 0, 4).Result())
	fmt.Println(client.GetRange(ctx, "name", -4, -1).Result())

	client.SetRange(ctx, "name", 0, "Go")
}
