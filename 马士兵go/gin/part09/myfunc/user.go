package myfunc

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Hello1(context *gin.Context) {
	var age = 19
	var arr = []int{33, 66, 99}
	var mapData = map[string]interface{}{
		"age": age,
		"arr": arr,
	}
	context.HTML(http.StatusOK, "demo01/hello.html", mapData)
}
