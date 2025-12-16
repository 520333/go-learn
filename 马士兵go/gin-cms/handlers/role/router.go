package role

import "github.com/gin-gonic/gin"

func Router(r *gin.Engine) {
	g := r.Group("/role")
	g.GET("", GetRow)      //GET /role?id=21
	g.GET("list", GetList) //GET /role/list?
}
