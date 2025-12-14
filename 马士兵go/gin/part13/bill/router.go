package bill

import (
	"github.com/gin-gonic/gin"
)

func Router(r *gin.RouterGroup) {
	r.GET("/userindex", Hello1)
	r.GET("/toFormBind", Hello2)
	r.GET("/userindex2", Hello3)
}
