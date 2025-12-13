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
	r.GET("/demo2/*id", myfunc.Hello2)

	// 路径中以键值对传入参数 在路由规则中不用做文章
	r.GET("/demo3", myfunc.Hello3)
	r.GET("/demo4", myfunc.Hello4)
	r.Run(":8080")
}
