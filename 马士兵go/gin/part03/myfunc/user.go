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
	var age = context.DefaultPostForm("age", "18")
	var loveLange = context.PostFormArray("loveLangue")
	var user = context.PostFormMap("user")
	context.String(http.StatusOK, "hello %s %s %s 兴趣爱好:%s  地址邮箱:%s", uname, pwd, age, loveLange, user)
	fmt.Println(uname, pwd, age, loveLange, user)
}
