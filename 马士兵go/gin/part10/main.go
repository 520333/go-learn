package main

import (
	"gin/part10/myfunc"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("part10/templates/**/*")
	r.GET("/userindex", myfunc.Hello1)
	r.GET("/toFormBind", myfunc.Hello2)

	r.GET("/userindex2", myfunc.Hello3)
	r.Run(":8080")
}
