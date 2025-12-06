package gorm

import (
	"fmt"
	"log"
)

func StdAssocModel() {
	if err := DB.AutoMigrate(&Author{}, &Essay{}, &Tag{}, &EssayMate{}); err != nil {
		log.Fatalln(err)
	}
	log.Println("migrate successful")
}

func AssocAppend() {
	// 一对多的关系， Author 1:n Essay
	// 创建测试数据
	var a Author
	a.Name = "一位大佬"
	if err := DB.Create(&a).Error; err != nil {
		log.Println(err)
	}
	log.Println("a:", a.ID)
	var e1, e2, e3 Essay
	e1.Subject = "师与徒"
	e2.Subject = "罪与罚"
	e3.Subject = "劳与烦"
	if err := DB.Create([]*Essay{&e1, &e2}).Error; err != nil {
		log.Println(err)
	}
	log.Println("e1:", e1.AuthorID, "e2:", e2.AuthorID)

	// 添加关联
	if err := DB.Model(&a).Association("Essay").Append([]Essay{e1, e2}); err != nil {
		log.Println(err)
	}
	fmt.Println(a.Essay)
	var t1, t2, t3, t4 Tag
	t1.Title = "曼达洛人"
	t2.Title = "波巴非特之书"
	t3.Title = "安多"
	t4.Title = "阿索卡"
	if err := DB.Create([]*Tag{&t1, &t2, &t3, &t4}).Error; err != nil {
		log.Println(err)
	}
	log.Println("t1,t2,t3,t4:", t1.ID, t2.ID, t3.ID, t4.ID)
	if err := DB.Model(&e1).Association("Tag").Append([]Tag{t1, t2}); err != nil {
		log.Println(err)
	}
	if err := DB.Model(&e2).Association("Tag").Append([]Tag{t3, t4}); err != nil {
		log.Println(err)
	}
	// To Essay N:1 Author
	if err := DB.Create([]*Essay{&e3}).Error; err != nil {
		log.Println(err)
	}
	log.Println("e3:", e3.ID)

	if err := DB.Model(&e3).Association("Author").Append(&a); err != nil {
		log.Println(err)
	}
	log.Println("e3:", e3.Author)

	var a2 Author
	a2.Name = "尤达"
	if err := DB.Create(&a2).Error; err != nil {
		log.Println(err)
	}
	log.Println("a2:", a2.ID)
	if err := DB.Model(&e3).Association("Author").Append(&a2); err != nil {
		log.Println(err)
	}
	log.Println("e3:", e3.Author.ID)

}
