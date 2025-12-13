package myfunc

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Hello(context *gin.Context) {
	name := "我是你爹不不不"
	context.HTML(http.StatusOK, "demo01/hello01.html", name)
}

type Student struct {
	Name string
	Age  int
}

func Hello2(context *gin.Context) {
	var s = Student{
		Name: "hello",
		Age:  18,
	}
	context.HTML(http.StatusOK, "demo01/hello01.html", s)
}

func Hello3(context *gin.Context) {
	var array = [3]int{10, 20, 30}
	context.HTML(http.StatusOK, "demo01/hello01.html", array)
}
