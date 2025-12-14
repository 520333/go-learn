package main

import (
	"gin/part13/middleware"
	"gin/part13/router"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("part13/templates/**/*")
	// 使用中间件
	r.Use(middleware.MiddleWare01)
	r.Use(middleware.MiddleWare02())
	router.Router(r)

	r.Run(":8080")
}
