package utils

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormLoger "gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

// InitDB 初始化数据库连接
func InitDB() {
	// 1.配置
	logLevel := gormLoger.Warn
	// 根据当前应用mode 控制日志级别
	switch gin.Mode() {
	case gin.ReleaseMode:
		logLevel = gormLoger.Warn
	case gin.DebugMode, gin.TestMode:
		fallthrough
	default:
		logLevel = gormLoger.Info
	}
	// db日志
	unionLogger := gormLoger.New(log.New(LogWriter(), "\n", log.LstdFlags), gormLoger.Config{
		SlowThreshold:             time.Second,
		Colorful:                  false, //不显示彩色输出
		IgnoreRecordNotFoundError: false,
		ParameterizedQueries:      false,
		LogLevel:                  logLevel,
	})

	// gorm配置
	config := &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 单数表名
		}, // 数据表命名策略
		Logger: unionLogger,
	}

	// 2.创建db对象
	dsn := viper.GetString("db.dsn")
	if dbNew, err := gorm.Open(mysql.Open(dsn), config); err != nil {
		log.Fatalln(err)
	} else {
		db = dbNew
	}
}

// 全局的DB对象
var db *gorm.DB

// DB 全局访问DB对象的方法
func DB() *gorm.DB {
	return db
}
