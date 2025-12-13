package myfunc

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Hello1(context *gin.Context) {
	context.HTML(http.StatusOK, "demo01/hello.html", nil)
}
func Hello2(context *gin.Context) {
	var uname = context.PostForm("username")
	var pwd = context.PostForm("pwd")
	fmt.Println(uname, pwd)
}
