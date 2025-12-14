package main

import (
	"gin/part09/myfunc"
	"html/template"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.SetFuncMap(template.FuncMap{ //自定义模板函数
		"add": myfunc.Add,
	})
	r.LoadHTMLGlob("part09/templates/**/*")
	r.GET("/userindex", myfunc.Hello1)
	r.Run(":8080")
}
