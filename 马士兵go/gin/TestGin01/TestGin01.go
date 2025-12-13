package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Hello(context *gin.Context) {
	context.HTML(http.StatusOK, "demo01/hello01.html", nil)
}
func main() {
	r := gin.Default()

	//r.LoadHTMLFiles("templates/hello01.html", "templates/hello02.html") // 不推荐
	r.LoadHTMLGlob("templates/**/*") // 加载全部 多级两个星号

	r.GET("/hello", Hello)
	err := r.Run(":9999")
	if err != nil {
		log.Fatalln(err)
	}
}
