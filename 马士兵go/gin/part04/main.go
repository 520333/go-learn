package main

import (
	"net/http"
	"part04/myfunc"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	// 必须给定参数 /demo/123 否则404
	r.LoadHTMLGlob("templates/**/*")
	r.StaticFS("/s", http.Dir("static"))
	r.GET("/userindex", myfunc.Hello1)
	r.POST("/getUserInfo", myfunc.Hello2)
	r.POST("/ajaxpost", myfunc.Hello3)
	r.Run(":8080")
}
