package role

import "github.com/gin-gonic/gin"

func Router(r *gin.Engine) {
	g := r.Group("/role")
	g.GET("", GetRow)         //查询一条 GET /role?id=21
	g.GET("list", GetList)    //查询多条 GET /role/list?pageSize=1&pageNum=2&sortMethod=desc
	g.POST("", Add)           //添加 GET /role/
	g.DELETE("", Delete)      //删除 DELETE /role?id=22&id=33&id=44
	g.GET("recycle", Recycle) //查询回收站 GET /role/recycle?keyword=超人
	g.PUT("restore", Restore) //还原 PUT /role?id=1&id=2
}
