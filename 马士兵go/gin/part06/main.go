package main

import (
	"gin/part06/myfunc"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	// 必须给定参数 /demo/123 否则404
	r.LoadHTMLGlob("part06/templates/**/*")
	r.StaticFS("/s", http.Dir("part06/static"))
	r.GET("/userindex", myfunc.Hello1)
	r.POST("/savefile", myfunc.Hello4)
	r.Run(":8080")
}
