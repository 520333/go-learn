package main

import (
	"fmt"
	"ginchat/models"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	var dsn = "root:123456@tcp(192.168.50.100:3306)/ginchat?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	//迁移
	if err := db.AutoMigrate(&models.UserBasic{}); err != nil {
		log.Println(err)
	}
	db.AutoMigrate(&models.Message{}, &models.Contact{}, &models.GroupBasic{})

	user := &models.UserBasic{}
	user.Name = "海绵宝宝"
	//user.PassWord = "123456"
	db.Create(user)

	// Read
	fmt.Println(db.First(user, 1))
	db.Model(user).Update("PassWord", "123456")
}
