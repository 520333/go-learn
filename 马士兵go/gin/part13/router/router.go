package router

import (
	"gin/part12/bill"
	"gin/part12/external"

	"github.com/gin-gonic/gin"
)

func Router(r *gin.Engine) {
	var b = r.Group("/bill")     //支票模块
	var e = r.Group("/external") //三方工具模块
	// 模块分组：
	bill.Router(b)
	external.Router(e)
}
