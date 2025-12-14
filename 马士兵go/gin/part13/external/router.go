package external

import (
	"gin/part13/middleware"

	"github.com/gin-gonic/gin"
)

func Router(r *gin.RouterGroup) {
	r.GET("/userindex3", Hello4, middleware.MiddleWare01)
	r.POST("/toajax", Hello5)
	r.GET("/userindex4/:uname/:age", Hello6)
}
