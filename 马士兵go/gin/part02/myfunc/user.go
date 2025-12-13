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
