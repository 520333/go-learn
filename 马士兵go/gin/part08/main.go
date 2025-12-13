package main

import (
	"gin/part08/myfunc"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/red1", myfunc.Red1)
	r.GET("/red2", myfunc.Red2)
	r.Run(":8080")
}
