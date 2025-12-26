package utils

import (
	"context"
	"fmt"
	"os"
	"time"

	"log"

	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	DB  *gorm.DB
	Red *redis.Client
)

func InitConfig() {
	viper.SetConfigName("app")
	viper.SetConfigType("yml")
	viper.AddConfigPath("./config")
	if err := viper.ReadInConfig(); err != nil {
		log.Println(err)
	}
	fmt.Println("config app initialized")
}
func InitMySQL() {
	// 自定义日志打印模板 打印SQL语句
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second, // 慢sql阈值
			LogLevel:      logger.Info, //级别
			Colorful:      true,        //彩色
		},
	)
	dsn := viper.GetString("mysql.dsn")
	DB, _ = gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: newLogger})
	fmt.Println("config mysql initialized")
}

func InitRedis() {
	Red = redis.NewClient(&redis.Options{
		Addr:         viper.GetString("redis.addr"),
		Password:     viper.GetString("redis.password"),
		DB:           viper.GetInt("redis.DB"),
		PoolSize:     viper.GetInt("redis.poolSize"),
		MinIdleConns: viper.GetInt("redis.minIdleConn"),
	})
	pong, err := Red.Ping(context.Background()).Result()
	if err != nil {
		fmt.Println("init redis...", err)
	} else {
		fmt.Println("config redis initialized", pong)
	}
}

const (
	PublishKey = "websocket"
)

// Publish 发布消息到redis
func Publish(ctx context.Context, channel string, msg string) error {
	var err error
	fmt.Println("publish...", msg)
	err = Red.Publish(ctx, channel, msg).Err()
	return err
}

// Subscribe 从redis订阅消息
func Subscribe(ctx context.Context, channel string) (string, error) {
	sub := Red.Subscribe(ctx, channel)
	fmt.Println("subscribe...", ctx)
	msg, err := sub.ReceiveMessage(ctx)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	fmt.Println("subscribe...", msg.Payload)
	return msg.Payload, err
}
