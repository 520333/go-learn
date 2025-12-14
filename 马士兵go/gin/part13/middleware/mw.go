package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func MiddleWare01(context *gin.Context) {
	fmt.Println("这是自定义中间件-方式1 开始")
	context.Next()
	fmt.Println("这是自定义中间件-方式1 结束")
}

func MiddleWare02() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("这是自定义中间件-方式2 开始")
		fmt.Println("这是自定义中间件-方式2 结束")
	}
}
func MiddleWare03() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("这是自定义中间件-方式3 开始")
		c.Next()
		fmt.Println("这是自定义中间件-方式3 结束")
	}
}
