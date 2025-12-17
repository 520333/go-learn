package role

import (
	"fmt"
	"ginCms/models"
	"ginCms/utils"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetList(ctx *gin.Context) {
	// 1.解析请求消息
	req := GetListReq{}
	if err := ctx.ShouldBindQuery(&req); err != nil {
		utils.Logger().Error(err.Error()) //记录日志
		ctx.JSON(http.StatusOK, gin.H{
			"code":    100,
			"message": err.Error(),
		})
		return
	}
	// 2.整理请求参数
	req.Clean()
	log.Println(req.Keyword, req.SortMethod, req.SortField, req.PageNum, req.PageSize)
	log.Println(*req.Keyword, *req.SortMethod, *req.SortField, *req.PageNum, *req.PageSize)

	// 3.基于model查询
	rows, err := models.RoleFetchList(false, req.RoleFilter, req.Sorter, req.Pager)
	if err != nil {
		utils.Logger().Error(err.Error()) //记录日志
		ctx.JSON(http.StatusOK, gin.H{
			"code":    100,
			"message": fmt.Sprintf("数据查询错误"),
		})
		return
	}
	// 4.响应
	ctx.JSON(http.StatusOK, gin.H{
		"code": 0,
		"data": rows,
	})
}

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
	row, err := models.RoleFetch(req.ID, false)
	if err != nil {
		utils.Logger().Error(err.Error()) //记录日志
		ctx.JSON(http.StatusOK, gin.H{
			"code":    100,
			"message": fmt.Sprintf("数据查询错误:%s", err.Error()),
		})
		return
	}
	// 3.响应
	ctx.JSON(http.StatusOK, gin.H{
		"code": 0,
		"data": row,
	})
}

func Delete(ctx *gin.Context) {
	req := DeleteReq{}
	if err := ctx.ShouldBindQuery(&req); err != nil {
		utils.Logger().Error(err.Error()) //记录日志
		ctx.JSON(http.StatusOK, gin.H{
			"code":    100,
			"message": err.Error(),
		})
		return
	}

	rowNum, err := models.RoleDelete(req.IDList)
	if err != nil {
		utils.Logger().Error(err.Error())
		ctx.JSON(http.StatusOK, gin.H{
			"code":    100,
			"message": "数据删除错误",
		})
		return
	}

	// 响应
	ctx.JSON(http.StatusOK, gin.H{
		"code": 0,
		"data": rowNum,
	})
}

func Add(ctx *gin.Context) {
	// 1.解析请求数据
	req := AddReq{}
	if err := ctx.ShouldBind(&req); err != nil {
		utils.Logger().Error(err.Error()) //记录日志
		ctx.JSON(http.StatusOK, gin.H{
			"code":    100,
			"message": err.Error(),
		})
		return
	}
	// 2.利用模型完成插入
	role := req.ToRole()
	if err := models.RoleInsert(role); err != nil {
		utils.Logger().Error(err.Error())
		ctx.JSON(http.StatusOK, gin.H{
			"code":    100,
			"message": "数据查询错误",
		})
		return
	}
	// 响应
	row, err := models.RoleFetch(role.ID, false)
	if err != nil {
		utils.Logger().Error(err.Error()) //记录日志
		ctx.JSON(http.StatusOK, gin.H{
			"code":    100,
			"message": "数据查询错误",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": 0,
		"data": row,
	})
}
