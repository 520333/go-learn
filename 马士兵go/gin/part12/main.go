package main

import (
	"gin/part12/router"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("part12/templates/**/*")
	router.Router(r)

	r.Run(":8080")
}
