package gorm

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
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
	if err := DB.AutoMigrate(&FieldTag{}, &TypeMap{}, &Post{}, &Category{}, &PostCategory{}, &Box{}); err != nil {
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

type CustomTypeModel struct {
	gorm.Model
	FTime       time.Time
	FNullTime   sql.NullTime
	FString     string
	FNullString sql.NullString
	FUuid       uuid.UUID
	FNullUUID   uuid.NullUUID
}

func CustomType() {
	//id:=uuid.UUID{}
	//id.Scan()
	//id.Value()
	// 初始化模型
	ctm := &CustomTypeModel{}
	// 迁移数据表
	DB.AutoMigrate(ctm)
	// 创建
	ctm.FTime = time.Now()             //当前时间
	ctm.FNullTime = sql.NullTime{}     //nil零值
	ctm.FString = ""                   //零值
	ctm.FNullString = sql.NullString{} //nil零值
	ctm.FUuid = uuid.New()             //零值
	ctm.FNullUUID = uuid.NullUUID{}    //nil零值
	DB.Create(ctm)
	// 查询
	DB.First(ctm, ctm.ID)
	// 判断字段是否为Null
	if ctm.FString == "" {
		fmt.Println("FString is NULL")
	} else {
		fmt.Println("FString is NOT NULL ")
	}
	if ctm.FNullString.Valid == false {
		fmt.Println("FNullString is NULL")
	} else {
		fmt.Println("FNullString is NOT NULL")
	}
}

type FieldTag struct {
	gorm.Model
	FStringDefault string `gorm:""`
	FTypeChar      string `gorm:"type:char(32)"`
	FTypeVarChar   string `gorm:"type:varchar(255)"`
	FTypeText      string `gorm:"type:text"`
	FTypeBlob      []byte `gorm:"type:blob"`
	FTypeEnum      string `gorm:"type:enum('Go','GORM','MYSQL')"`
	FTypeSet       string `gorm:"type:set('Go','GORM','MYSQL')"`
	FColNum        string `gorm:"column:custom_column_name"`
	FColNotNull    string `gorm:"type:varchar(255);not null"`
	FColDefault    string `gorm:"type:varchar(255);not null;default:'gorm middle ware'"`
	FColComment    string `gorm:"type:varchar(255);comment:带有注释的字段"`
}
