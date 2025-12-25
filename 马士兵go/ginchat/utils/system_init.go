package utils

import (
	"fmt"
	"os"
	"time"

	"log"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func InitConfig() {
	viper.SetConfigName("app")
	viper.SetConfigType("yml")
	viper.AddConfigPath("./config")
	if err := viper.ReadInConfig(); err != nil {
		log.Println(err)
	}
	fmt.Println("config app initialized")
}
func InitMySQL() {
	// 自定义日志打印模板 打印SQL语句
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second, // 慢sql阈值
			LogLevel:      logger.Info, //级别
			Colorful:      true,        //彩色
		},
	)
	dsn := viper.GetString("mysql.dsn")
	DB, _ = gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: newLogger})
	fmt.Println("config mysql initialized")
}
