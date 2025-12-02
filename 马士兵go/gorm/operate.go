package gorm

import (
	"fmt"
	"log"
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model

	Username string
	Name     string
	Email    string
	Birthday *time.Time
}

func OperatorType() {
	DB.AutoMigrate(&User{})

	var users []User

	// 一步操作
	//DB.Where("birthday IS NOT NULL").Where("email like ?", "@163.com%").Order("name DESC").Find(&users)

	// 分步操作
	query := DB.Where("birthday IS NOT NULL")
	query.Where("email like ?", "@163.com%")
	query.Order("name DESC")
	query.Find(&users)
}

type Content struct {
	gorm.Model

	Subject     string
	Likes       uint
	PublishTime *time.Time
}

func CreateBasic() {
	DB.AutoMigrate(&Content{})

	c1 := Content{}
	c1.Subject = "GORM的使用"
	c1.Likes = 1000

	result1 := DB.Create(&c1)
	if result1.Error != nil {
		log.Fatal(result1.Error)
	}
	fmt.Println(c1.ID, result1.RowsAffected)

	values := map[string]interface{}{
		"Subject":     "Map指定值",
		"Likes":       999,
		"PublishTime": time.Now(),
	}
	result := DB.Model(&Content{}).Create(values)
	if result.Error != nil {
		log.Fatal(result.Error)
	}
	fmt.Println(result.RowsAffected)

}
