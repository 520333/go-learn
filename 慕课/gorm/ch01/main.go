package main

import (
	"database/sql"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Product struct {
	gorm.Model
	Code  sql.NullString
	Price uint
}

func main() {
	// 1.数据源
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

	// 2.定义表结构 自动创建mysql表  表明是product小写
	db.AutoMigrate(&Product{})

	// 新增
	db.Create(&Product{Code: sql.NullString{String: "D42", Valid: true}, Price: 100})

	// 查询
	var product Product
	db.First(&product, 1)                 // 根据整型主键查询
	db.First(&product, "code = ?", "D42") //查看code字段值位D42的记录

	// 更新
	db.Model(&product).Update("Price", 200) //更新单个字段
	// 更新多个字段
	db.Model(&product).Updates(Product{Price: 200, Code: sql.NullString{String: "", Valid: true}})

	// 删除
	// db.Delete(&product, 1) //逻辑删除 将delete_at字段更新
}
