package redisCli

import (
	"context"
	"fmt"
	"log"

	"github.com/redis/go-redis/v9"
)

func ListPushPop() {
	opt, err := redis.ParseURL("redis://default:123456@192.168.50.100:6379/0")
	if err != nil {
		log.Fatalln(err)
	}
	client := redis.NewClient(opt)
	ctx := context.Background()
	//client.LTrim(ctx, "subjects", -1, 0)
	client.Del(ctx, "subjects")
	// 插入
	client.LPush(ctx, "subjects", "GO")
	// 获取长度
	fmt.Println(client.LLen(ctx, "subjects").Result())

	client.LPush(ctx, "subjects", "Redis", "MySQL")
	fmt.Println(client.LLen(ctx, "subjects").Result())

	client.RPush(ctx, "subjects", "Docker")
	fmt.Println(client.LLen(ctx, "subjects").Result())

	// 删除元素
	fmt.Println(client.LRem(ctx, "subjects", 1, "MySQL").Result())
	// 删除全部MySQL
	fmt.Println(client.LRem(ctx, "subjects", 0, "MySQL").Result())
	//client.RPush(ctx, "subjects", "Kubernetes", "CI/CD")

	// 插入 前提key要存在
	//fmt.Println(client.LPushX(ctx, "subjects", "GO").Result())

	// 查看
	fmt.Println(client.LRange(ctx, "subjects", 0, -1).Result())

	// 取出 Get AND Del
	//fmt.Println(client.LPop(ctx, "subjects").Result())
	//fmt.Println(client.LRange(ctx, "subjects", 0, -1).Result())
	//fmt.Println(client.RPop(ctx, "subjects").Result())
	//fmt.Println(client.LRange(ctx, "subjects", 0, -1).Result())
	//
	//fmt.Println(client.RPopCount(ctx, "subjects", 3).Result())
	//fmt.Println(client.LRange(ctx, "subjects", 0, -1).Result())
}

func ListIndex() {
	opt, err := redis.ParseURL("redis://default:123456@192.168.50.100:6379/0")
	if err != nil {
		log.Fatalln(err)
	}
	client := redis.NewClient(opt)
	ctx := context.Background()
	//client.LTrim(ctx, "subjects", -1, 0)
	client.Del(ctx, "subjects")
	// 基于索引查询
	client.LPush(ctx, "subjects", "GO", "Redis", "MySQL", "GO-Redis", "GO-Redis", "Docker", "Kubernetes", "CI/CD")
	fmt.Println(client.LRange(ctx, "subjects", 0, -1).Result())
	fmt.Println(client.LIndex(ctx, "subjects", 3).Result())

	// 基于索引设置
	fmt.Println(client.LSet(ctx, "subjects", 3, "GO-Redis").Result())
	fmt.Println(client.LRange(ctx, "subjects", 0, -1).Result())

	fmt.Println(client.LInsert(ctx, "subjects", "BEFORE", "GO-Redis", "ArgoCD").Result())
	fmt.Println(client.LInsert(ctx, "subjects", "AFTER", "GO-Redis", "ArgoCD").Result())

	fmt.Println(client.LRange(ctx, "subjects", 0, -1).Result())

	client.LTrim(ctx, "subjects", 2, 4)
	fmt.Println(client.LRange(ctx, "subjects", 0, -1).Result())

}
