package myfunc

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Hello(context *gin.Context) {
	context.HTML(http.StatusOK, "demo01/hello.html", nil)
}
