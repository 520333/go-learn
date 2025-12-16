package main

import (
	"ginCms/handlers"
	"ginCms/utils"

	"github.com/spf13/viper"
)

func main() {
	// 解析配置
	utils.ParseConfig()

	r := handlers.InitEngine()
	r.Run(viper.GetString("app.addr"))
}
