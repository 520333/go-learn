package main

import (
	"net/http"
	"part01/myfunc"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/**/*")
	r.StaticFS("/s", http.Dir("static"))
	r.GET("/demo", myfunc.Hello)
	r.Run(":8080")
}
