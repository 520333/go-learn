package myfunc

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Hello1(context *gin.Context) {
	// 获取路径中的参数值:
	var id = context.Param("id")
	context.String(http.StatusOK, "hello world %s", id)
}

func Hello2(context *gin.Context) {
	// 获取路径中的参数值:
	var id = context.Param("id")
	context.String(http.StatusOK, "hello world %s", id)
}

func Hello3(context *gin.Context) {
	// 获取路径中的参数值: 通过key获取对应value
	var id = context.Query("id")
	var name = context.Query("name")
	context.String(http.StatusOK, "hello world %s %s", id, name)
}

func Hello4(context *gin.Context) {
	// 获取路径中的参数值: 通过key获取对应value
	var id = context.DefaultQuery("id", "123")
	var name = context.DefaultQuery("name", "丽丽")
	context.String(http.StatusOK, "hello world %s %s", id, name)
}

func Hello5(context *gin.Context) {
	// 获取路径中的参数值: 通过key获取对应value
	var idValues = context.QueryArray("id")
	context.String(http.StatusOK, "hello world %s", idValues)
}

func Hello6(context *gin.Context) {
	// 获取路径中的参数值: 通过key获取对应value 访问/demo6?user[10001]=丽丽&user[10002]=派大星
	var idValues = context.QueryMap("user")
	context.String(http.StatusOK, "hello world %s", idValues)
}
