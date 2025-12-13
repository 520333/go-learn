package myfunc

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Red1(context *gin.Context) {
	fmt.Println("Red1")
	context.Redirect(http.StatusFound, "/red2")
}

func Red2(context *gin.Context) {
	fmt.Println("Red2")
	context.String(http.StatusOK, "重定向成功 red2页面")

}
