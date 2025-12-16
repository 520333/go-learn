package role

import (
	"ginCms/models"
	"ginCms/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetRow(ctx *gin.Context) {
	// 1.解析请求数据（消息）
	req := GetRowReq{}
	if err := ctx.ShouldBindQuery(&req); err != nil {
		utils.Logger().Error(err.Error()) //记录日志
		ctx.JSON(http.StatusOK, gin.H{
			"code":    100,
			"message": err.Error(),
		})
		return
	}
	// 2.利用模型完成查询
	row, err := models.RoleFetchRow(false, "`id` = ?", req.ID)
	if err != nil {
		utils.Logger().Error(err.Error()) //记录日志
		ctx.JSON(http.StatusOK, gin.H{
			"code":    100,
			"message": err.Error(),
		})
		return
	}
	// 3.响应
	ctx.JSON(http.StatusOK, gin.H{
		"code": 0,
		"data": row,
	})
}
