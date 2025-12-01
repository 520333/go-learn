package gorm

import (
	"io"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var logWriter io.Writer

func init() {
	// 定义DSN
	//const dsn = "root:123456@tcp(192.168.50.100:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local"
	// 连接服务器
	//db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
	//	Logger: logger.Default.LogMode(logger.Info), // 全局Info级别日志
	//})
	//if err != nil {
	//	log.Fatal(err)
	//}
	//DB = db

	// 自定义日志
	logWriter, _ = os.OpenFile("./sql.log", os.O_CREATE|os.O_APPEND, 0644)
	const dsn = "root:123456@tcp(192.168.50.100:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local"
	customLogger := logger.New(log.New(logWriter, "", log.LstdFlags), logger.Config{
		SlowThreshold:             200 * time.Millisecond,
		LogLevel:                  logger.Info,
		IgnoreRecordNotFoundError: false,
		Colorful:                  false, //不显示彩色输出
	})
	// 连接服务器
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: customLogger, // 自定义日志类型
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "go_", // 表名前缀
			SingularTable: true,  // 表名使用单数形式
			NameReplacer:  nil,   // 替换命名中某些字符
			NoLowerCase:   true,  // 不将名称转为小写
		}, // 设置默认的命名策略选项
	})
	if err != nil {
		log.Fatal(err)
	}
	DB = db
}
