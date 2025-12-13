package main

import (
	"gin/part07/myfunc"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	// 必须给定参数 /demo/123 否则404
	r.LoadHTMLGlob("part07/templates/**/*")
	r.StaticFS("/s", http.Dir("part07/static"))
	r.GET("/userindex", myfunc.Hello1)
	r.POST("/savefile", myfunc.Hello2)
	r.Run(":8080")
}
