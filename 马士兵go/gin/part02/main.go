package main

import (
	"part02/myfunc"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	// 必须给定参数 /demo/123 否则404
	r.GET("/demo/:id", myfunc.Hello1)
	// *占位符 无参数直接忽略 输出:/
	r.GET("/demo2/*id", myfunc.Hello1)
	r.Run(":8080")
}
