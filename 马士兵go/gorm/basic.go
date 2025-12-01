package gorm

import (
	"fmt"
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Article 创建模型
type Article struct {
	gorm.Model  //内置模型
	Subject     string
	Likes       uint
	Published   bool
	PublishTime time.Time
	AuthorID    uint
}

func BasicUsage() {
	// 定义DSN
	var dsn = "root:123456@tcp(192.168.50.100:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local"
	// 连接服务器
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	// 连接成功
	fmt.Println(db)

	// 基于表模型完成表结构（设计）的迁移
	if err := db.AutoMigrate(&Article{}); err != nil {
		log.Fatal(err)
	}
	fmt.Println("migrate success")
}

var DB *gorm.DB

// Create 增加数据
func Create() {
	article := &Article{
		Subject:     "GORM的CRUD基础操作",
		Likes:       0,
		Published:   true,
		PublishTime: time.Now(),
		AuthorID:    42,
	}
	if err := DB.Create(article).Error; err != nil {
		log.Fatal(err)
	}
	fmt.Println(article)
}

// Retrieve 查询
func Retrieve(id uint) {
	// 初始化Article模型，零值
	article := &Article{}
	if err := DB.First(article, id).Error; err != nil {
		log.Fatal(err)
	}
	fmt.Println(article)
}

// Update 更新操作
func Update() {
	article := &Article{}
	if err := DB.First(article, 1).Error; err != nil {
		log.Fatal(err)
	}
	article.AuthorID = 23
	article.Likes = 101
	article.Subject = "新文章的标题"
	if err := DB.Save(article).Error; err != nil {
		log.Fatal(err)
	}
}

// Delete 删除操作
func Delete(id uint) {
	article := &Article{}
	if err := DB.First(article, id).Error; err != nil {
		log.Fatal(err)
	}
	if err := DB.Delete(article, id).Error; err != nil {
		log.Fatal(err)
	}
	fmt.Println("article delete success")
}

// Debug 日志级别
func Debug() {
	article := &Article{
		Subject:     "Article Subject",
		PublishTime: time.Now(),
	}
	if err := DB.Debug().Create(article).Error; err != nil {
		log.Fatal(err)
	}
	if err := DB.Debug().First(article, article.ID).Error; err != nil {
		log.Fatal(err)
	}
}

func Log() {
	Create()
}
