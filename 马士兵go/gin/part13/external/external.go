package external

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Hello4(context *gin.Context) {
	context.HTML(http.StatusOK, "demo01/hello3.html", nil)
}

type User2 struct {
	Uname string `json:"uname" uri:"uname"`
	Age   int    `json:"age" uri:"age"`
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

func Hello6(context *gin.Context) {
	var user User2
	err := context.ShouldBindUri(&user)
	fmt.Println(user)
	if err != nil {
		context.String(404, "绑定失败")
	} else {
		context.String(200, "绑定成功")
	}
}
