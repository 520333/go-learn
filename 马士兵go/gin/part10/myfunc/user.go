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

func Hello4(context *gin.Context) {
	context.HTML(http.StatusOK, "demo01/hello3.html", nil)
}

type User2 struct {
	Uname string `json:"uname"`
	Age   int    `json:"age"`
}

func Hello5(context *gin.Context) {
	var user User2
	err := context.ShouldBind(&user)
	if err != nil {
		context.JSON(404, gin.H{
			"msg": "绑定失败",
		})
	} else {
		context.JSON(200, gin.H{
			"msg": "绑定成功",
		})
	}
	fmt.Println(user)

}
