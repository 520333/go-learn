package main

import (
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type User struct {
	gorm.Model
	Name      string
	CompanyID int // 数据库中存储的字段 company_id
	Company   Company
}

type Company struct {
	ID   int
	Name string
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
	db.AutoMigrate(&User{}) //新建user、company表 并设置外键
	// db.Create(&User{
	// 	Name: "dawn",
	// 	Company: Company{
	// 		Name: "宝哥网",
	// 	},
	// })
	db.Create(&User{
		Name: "dawn",
		Company: Company{
			ID: 1,
		},
	})
}
