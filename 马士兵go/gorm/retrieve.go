package gorm

import (
	"fmt"
	"log"
	"time"
)

func GetByPK() {
	DB.AutoMigrate(&Content{}, &ContentStrPK{})
	// 查询单条
	c := Content{}
	if err := DB.First(&c, 10).Error; err != nil {
		log.Println(err)
	}
	// 字符串类型主键
	cStr := ContentStrPK{}
	if err := DB.First(&cStr, "id = ?", "some pk").Error; err != nil {
		log.Println(err)
	}
	// 查询多条
	var cs []Content
	if err := DB.Find(&cs, []uint{10, 11, 12}).Error; err != nil {
		log.Println(err)
	}

	// 字符串类型的主键
	var cStrs []ContentStrPK
	if err := DB.Find(&cStrs, "id IN ?", []string{"some", "pk", "item"}).Error; err != nil {
		log.Println(err)
	}
}

func GeOne() {
	c := Content{}
	if err := DB.First(&c, "id > ?", 42).Error; err != nil {
		log.Println(err)
	}
	o := Content{}
	if err := DB.Last(&o, "id > ?", 42).Error; err != nil {
		log.Println(err)
	}
	n := Content{}
	if err := DB.Take(&n, "id > ?", 42).Error; err != nil {
		log.Println(err)
	}
	f := Content{}
	if err := DB.Limit(1).Find(&f, "id > ?", 42).Error; err != nil {
		log.Println(err)
	}
	fs := Content{}
	if err := DB.Find(&fs, "id > ?", 42).Error; err != nil {
		log.Println(err)
	}
}

func GetToMap() {
	c := map[string]any{}
	if err := DB.Model(&Content{}).First(&c, 13).Error; err != nil {
		log.Println(err)
	}
	fmt.Println(c, c["id"] == 13)
	if c["id"].(uint) == 13 {
		fmt.Println("id bingo")
	}
	// time类型处理
	fmt.Println(c["created_at"])
	t, _ := time.Parse("2006-01-02 15:04:05.00 -0700 CST", "2025-12-02 23:23:34.22 +0800 CST")
	if c["created_at"].(time.Time) == t {
		fmt.Println("created_at bingo")
	}

	// 多条
	var cs []map[string]any
	if err := DB.Model(&Content{}).Find(&cs, []uint{13, 14, 15}).Error; err != nil {
		log.Println(err)
	}
	for _, c := range cs {
		fmt.Println(c["id"], c["subject"].(string), c["created_at"].(time.Time))
	}
}
