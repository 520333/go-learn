package main

import (
	"gin/part01/myfunc"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("part01/templates/**/*")
	r.StaticFS("/s", http.Dir("part01/static"))
	r.GET("/demo", myfunc.Hello)
	r.Run(":8080")
}
