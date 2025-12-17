package handlers

import (
	"ginCms/handlers/common"
	"ginCms/handlers/role"
	"ginCms/handlers/system"

	"github.com/gin-gonic/gin"
)

// InitEngine 初始化路由引擎
func InitEngine() *gin.Engine {
	// 1.初始化路由引擎
	r := gin.Default()
	// 设置中间件
	common.UseCors(r)
	// 2.注册不同模块的路由
	system.Router(r)
	role.Router(r)

	return r
}
