package redisCli

import (
	"context"
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
