package main

import (
	"net/http"
	"part05/myfunc"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	// 必须给定参数 /demo/123 否则404
	r.LoadHTMLGlob("templates/**/*")
	r.StaticFS("/s", http.Dir("static"))
	r.GET("/userindex", myfunc.Hello1)
	r.POST("/savefile", myfunc.Hello2)
	r.Run(":8080")
}
