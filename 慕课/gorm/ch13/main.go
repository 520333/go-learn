package main

// 自定义表名 统一前缀
import (
	"database/sql"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

type Language struct {
	gorm.Model
	Name string
	// AddTime time.Time // 每个记录创建的时候自动加上当前时间
	AddTime sql.NullTime
}

// 创建之前add_time字段加上当前时间
// func (l *Language) BeforeCreate(tx *gorm.DB) (err error) {
// 	l.AddTime = time.Now()
// 	return
// }

// 1.自定义表名
// func (Language) TableName() string {
// 	return "my_language"
// }

func main() {
	// 1.数据源云原生操练环境
	dsn := "root:123456@tcp(192.168.50.100:3306)/gorm_test?charset=utf8mb4&parseTime=True&loc=Local"
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second, //慢SQL阈值
			LogLevel:      logger.Info, //日志级别
			Colorful:      true,        //禁用彩色打印
		},
	)
	// 全局logger 打印执行的sql语句
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{TablePrefix: "baoge_"}, // 2.表名统一前缀  不能和自定义表名一起存在否则不生效
		Logger:         newLogger,
	})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Language{})
	db.Create(&Language{
		Name: "golang",
	})

}
