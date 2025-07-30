package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type User struct {
	gorm.Model
	CreditCards []CreditCards `gorm:"foreignKey:UserRefer"`
}

type CreditCards struct {
	gorm.Model
	Number    string
	UserRefer uint
}

func main() {
	// 1.数据源云原生操练环境
	dsn := "root:123456@tcp(192.168.50.100:3306)/gorm_test?charset=utf8mb4&parseTime=True&loc=Local"
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second, //慢SQL阈值
			LogLevel:      logger.Info, //日志级别
			Colorful:      true,        //禁用彩色打印
		},
	)
	// 全局logger 打印执行的sql语句
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&CreditCards{})
	// user := User{}
	// db.Create(&user)
	// db.Create(&CreditCards{
	// 	Number:    "12",
	// 	UserRefer: user.ID,
	// })
	// db.Create(&CreditCards{
	// 	Number:    "34",
	// 	UserRefer: user.ID,
	// })
	var user User
	db.Preload("CreditCards").First(&user)
	for _, card := range user.CreditCards {
		fmt.Println(card.Number)
	}
}
