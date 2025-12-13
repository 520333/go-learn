package main

import (
	"log"
	"net/http"
	"test01/myfunc"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	//r.LoadHTMLFiles("templates/hello01.html", "templates/hello02.html") // 不推荐
	r.LoadHTMLGlob("templates/**/*") // 加载全部 多级两个星号
	//r.Static("/s", "static") // 指定静态文件：指定css文件
	r.StaticFS("/s", http.Dir("static"))

	r.GET("/hello", myfunc.Hello2)
	err := r.Run(":9999")
	if err != nil {
		log.Fatalln(err)
	}
}
