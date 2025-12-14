package main

import (
	"gin/part09/myfunc"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	// 必须给定参数 /demo/123 否则404
	r.LoadHTMLGlob("part09/templates/**/*")
	r.GET("/userindex", myfunc.Hello1)
	r.Run(":8080")
}
