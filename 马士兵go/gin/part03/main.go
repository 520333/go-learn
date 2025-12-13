package main

import (
	"gin/part03/myfunc"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	// 必须给定参数 /demo/123 否则404
	r.LoadHTMLGlob("part03/templates/**/*")
	r.StaticFS("/s", http.Dir("part03/static"))
	r.GET("/userindex", myfunc.Hello1)
	r.POST("/getUserInfo", myfunc.Hello2)
	r.Run(":8080")
}
