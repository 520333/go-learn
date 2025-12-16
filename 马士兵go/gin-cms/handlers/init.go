package handlers

import (
	"ginCms/handlers/system"

	"github.com/gin-gonic/gin"
)

// InitEngine 初始化路由引擎
func InitEngine() *gin.Engine {
	// 1.初始化路由引擎
	r := gin.Default()

	// 2.注册不同模块的路由
	system.Router(r)

	return r
}
