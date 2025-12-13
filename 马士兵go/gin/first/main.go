package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func Hello(context *gin.Context) {
	context.String(200, "第一个Gin项目")
}
func main() {
	//r := gin.New()
	r := gin.Default()
	r.GET("/", func(context *gin.Context) {
		context.String(200, "first gin")

	})
	r.GET("/hello", Hello)
	err := r.Run(":9999")
	if err != nil {
		log.Fatalln(err)
	}
}
