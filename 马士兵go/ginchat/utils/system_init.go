package utils

import (
	"fmt"

	"log"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
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
	dsn := viper.GetString("mysql.dsn")
	DB, _ = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	fmt.Println("config mysql initialized")
}
