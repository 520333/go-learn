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
	if err := DB.AutoMigrate(&IAndC{}, &FieldTag{}, &TypeMap{}, &Post{}, &Category{}, &PostCategory{}, &Box{}); err != nil {
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

type IAndC struct {
	ID    uint   `gorm:"primaryKey"`                                //主键索引
	Email string `gorm:"type:varchar(255);uniqueIndex"`             // 唯一索引
	Age   int8   `gorm:"index;check:age>=18 AND email is not null"` // 普通索引
	// 复合索引
	FirstName string `gorm:"index:name"`
	LastName  string `gorm:"index:name"`

	// 顺序关键索引 默认的priority: 10
	FirstName1 string `gorm:"index:name1;priority:2"`
	LastName1  string `gorm:"index:name1;priority:1"`

	// 索引选项,前缀长度，排序方式,comment
	Height      float32 `gorm:"index:,sort:desc"` //降序遍历
	AddressHash string  `gorm:"index:,length:12,comment:前12个字符作为索引关键字"`
}

func IAndCreate() {
	iac := &IAndC{}
	iac.Age = 21
	if err := DB.Create(iac).Error; err != nil {
		log.Fatal(err)
	}
	fmt.Printf(": %+v\n", iac)
}
