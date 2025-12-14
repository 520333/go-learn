package router

import (
	"gin/part13/bill"
	"gin/part13/external"
	"gin/part13/middleware"

	"github.com/gin-gonic/gin"
)

func Router(r *gin.Engine) {
	var b = r.Group("/bill")       //支票模块
	b.Use(middleware.MiddleWare01) //
	var e = r.Group("/external")   //三方工具模块
	// 模块分组：
	bill.Router(b)
	external.Router(e)
}
