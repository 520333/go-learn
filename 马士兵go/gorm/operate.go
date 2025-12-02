package gorm

import (
	"fmt"
	"log"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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
	Likes       uint  `gorm:"default:99"`
	Views       *uint `gorm:"default:99"`
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

func CreateMulti() {
	DB.AutoMigrate(&Content{})
	// model
	cs := []Content{
		{Subject: "标题1"},
		{Subject: "标题2"},
		{Subject: "标题3"},
	}
	result := DB.Create(&cs)
	if result.Error != nil {
		log.Fatal(result.Error)
	}
	fmt.Println("RowsAffected:", result.RowsAffected)
	for _, c := range cs {
		fmt.Println("ID:", c.ID)
	}
	vs := []map[string]any{
		{"Subject": "标题4"},
		{"Subject": "标题5"},
		{"Subject": "标题6"},
	}
	result2 := DB.Model(&Content{}).Create(vs)
	if result2.Error != nil {
		log.Fatal(result2.Error)
	}
	fmt.Println("RowsAffected:", result2.RowsAffected)
}

func CreateBatch() {
	DB.AutoMigrate(&Content{})

	// model
	cs := []Content{
		{Subject: "标题1"},
		{Subject: "标题2"},
		{Subject: "标题3"},
		{Subject: "标题4"},
		{Subject: "标题5"},
	}
	result1 := DB.CreateInBatches(&cs, 2)
	if result1.Error != nil {
		log.Fatal(result1.Error)
	}
	fmt.Println(result1.RowsAffected)
	for _, c := range cs {
		fmt.Println(c.ID)
	}

	// map
	vs := []map[string]any{
		{"Subject": "标题6"},
		{"Subject": "标题7"},
		{"Subject": "标题8"},
		{"Subject": "标题9"},
		{"Subject": "标题10"},
	}
	result2 := DB.Model(&Content{}).CreateInBatches(vs, 2)
	if result2.Error != nil {
		log.Fatal(result2.Error)
	}
	fmt.Println(result2.RowsAffected)
}

func UpSert() {
	DB.AutoMigrate(&Content{})
	c1 := Content{}
	c1.Likes = 10
	c1.Subject = "标题"
	DB.Create(&c1)

	// 主键冲突
	//c2 := Content{}
	//c2.ID = c1.ID
	//c2.Subject = "新标题"
	//c2.Likes = 20
	//if err := DB.Create(&c2).Error; err != nil {
	//	log.Fatal(err)
	//}

	c3 := Content{}
	c3.ID = c1.ID
	c3.Subject = "新标题"
	c3.Likes = 20
	if err := DB.
		Clauses(clause.OnConflict{UpdateAll: true}).
		Create(&c3).Error; err != nil {
		log.Fatal(err)
	}
	c4 := Content{}
	c4.ID = c1.ID
	c4.Subject = "新标题4"
	c4.Likes = 40
	if err := DB.Clauses(clause.
		OnConflict{DoUpdates: clause.
		AssignmentColumns([]string{"likes", "subject"})}).
		Create(&c4).Error; err != nil {
		log.Fatal(err)
	}
}

func DefaultValue() {
	DB.AutoMigrate(&Content{})
	c1 := Content{}
	c1.Subject = "标题"
	c1.Likes = 0
	views := uint(0)
	c1.Views = &views
	DB.Create(&c1)
	fmt.Println(c1.Likes, *c1.Views)
}
