package utils

import (
	"log"

	"github.com/spf13/viper"
)

// 默认配置
func defaultConfig() {
	viper.SetDefault("app.mode", "debug")
	viper.SetDefault("app.addr", ":8080")
	viper.SetDefault("app.log.path", "./logs")
}

// ParseConfig 解析配置
func ParseConfig() {
	// 1.默认配置
	defaultConfig()

	//2.配置解析参数
	viper.AddConfigPath(".")       //从哪些目录搜索配置文件
	viper.SetConfigName("configs") //配置文件名字
	viper.SetConfigType("yaml")    //配置文件类型

	//3.执行解析
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalln(err)
	}
}
