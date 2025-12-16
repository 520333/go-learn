package main

import (
	"ginCms/handlers"
	"ginCms/utils"

	"github.com/spf13/viper"
)

func main() {
	// 解析配置
	utils.ParseConfig()
	// 设置应用模式
	utils.SetMode()
	// 设置日志
	utils.SetLogger()
	// 初始化数据库连接
	utils.InitDB()

	r := handlers.InitEngine()
	utils.Logger().Info("service is listening", "addr", viper.GetString("app.addr")) //输出应用日志
	r.Run(viper.GetString("app.addr"))
}
