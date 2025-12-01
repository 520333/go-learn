package gorm

import (
	"log"

	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
}
type Category struct {
	gorm.Model
}
type PostCategory struct {
	gorm.Model
}
type Box struct {
	gorm.Model
}

func (Box) TableName() string { //某张表特殊表名
	return "my_box"
}
func Migrate() {
	if err := DB.AutoMigrate(&Post{}, &Category{}, &PostCategory{}, &Box{}); err != nil {
		log.Fatal(err)
	}
}
