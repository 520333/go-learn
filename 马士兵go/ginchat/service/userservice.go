package service

import (
	"ginchat/models"

	"github.com/gin-gonic/gin"
)

// GetUserList
// @Summary 获取用户列表
// @Tags 用户列表
// @Produce json
// @Success 200 {string} json {"code","message"}
// @Router /user/getUserList [get]
func GetUserList(c *gin.Context) {
	data := make([]*models.UserBasic, 10)
	data = models.GetUserList()
	c.JSON(200, gin.H{
		"message": data,
	})
}
