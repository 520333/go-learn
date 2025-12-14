package myfunc

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Hello1(context *gin.Context) {
	context.HTML(http.StatusOK, "demo01/hello.html", nil)
}

type User struct {
	Username string `form:"username"`
	Pwd      string `form:"pwd"`
}

func Hello2(context *gin.Context) {
	var user User
	err := context.ShouldBind(&user)
	fmt.Println(user)
	if err != nil {
		context.String(404, "绑定失败")
	} else {
		context.String(200, "绑定成功")
	}
}

func Hello3(context *gin.Context) {
	var user User
	err := context.ShouldBind(&user)
	fmt.Println(user)
	if err != nil {
		context.String(404, "绑定失败")
	} else {
		context.String(200, "绑定成功")
	}
}
