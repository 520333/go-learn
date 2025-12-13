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

func Hello4(context *gin.Context) {
	var arr []Student = []Student{
		Student{"lili", 20},
		Student{"dawn", 21},
		Student{"yoda", 22},
	}
	context.HTML(http.StatusOK, "demo01/hello01.html", arr)
}

func Hello5(context *gin.Context) {
	var a map[string]int = make(map[string]int, 3)
	a["丽丽"] = 18
	a["菲菲"] = 16
	a["明明"] = 21
	context.HTML(http.StatusOK, "demo01/map.html", a)
}

func Hello6(context *gin.Context) {
	var a map[string]Student = make(map[string]Student, 3)
	a["NO1"] = Student{"丽丽", 18}
	a["NO2"] = Student{"菲菲", 21}
	a["NO3"] = Student{"明明", 22}
	context.HTML(http.StatusOK, "demo01/map-struct.html", a)
}

func Hello7(context *gin.Context) {
	var slice = []int{1, 2, 3, 4, 5, 6}
	context.HTML(http.StatusOK, "demo01/slice.html", slice)
}
