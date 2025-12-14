package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func MiddleWare01(context *gin.Context) {
	fmt.Println("自定义中间件里面写统一的业务逻辑")
}
