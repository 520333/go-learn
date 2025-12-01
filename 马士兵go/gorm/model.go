package gorm

import (
	"fmt"
	"log"
	"time"

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

type TypeMap struct {
	gorm.Model
	FInt       int
	FUint      uint
	Float32    float32
	Float64    float64
	FString    string
	FTime      time.Time
	FByteSlice []byte

	FIntP     *int
	FUintP    *uint
	FFloat32P *float32
	FFloat64P *float64
	FStringP  *string
	FTimeP    *time.Time
}

func Migrate() {
	if err := DB.AutoMigrate(&TypeMap{}, &Post{}, &Category{}, &PostCategory{}, &Box{}); err != nil {
		log.Fatal(err)
	}
}

// PointerDiff 指针类型和非指针类型区别
func PointerDiff() {
	// 模型的零值
	typeMap := &TypeMap{}
	fmt.Printf("%+v\n", typeMap)
	fmt.Println("==============================")

	// 查询数据 数据库中NULL对应指针类型nil值
	DB.First(typeMap, 1)
	fmt.Printf("%+v\n", typeMap)
}
