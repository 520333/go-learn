package myfunc

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Student struct {
	Name string
	Age  int
}

func Hello1(context *gin.Context) {
	var age = 19
	var arr = []int{33, 66, 99}
	flag := true
	var username = "丽丽"

	var stu = Student{
		Name: "丽丽",
		Age:  18,
	}
	var nowTime = time.Now()
	var mapData = map[string]interface{}{
		"age":      age,
		"arr":      arr,
		"flag":     flag,
		"username": username,
		"stu":      stu,
		"nowTime":  nowTime,
	}
	context.HTML(http.StatusOK, "demo01/hello.html", mapData)
}
