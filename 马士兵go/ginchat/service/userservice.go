package service

import (
	"ginchat/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetUserList
// @Summary 用户列表
// @Tags 用户模块
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

// CreateUser
// @Summary 新增用户
// @Tags 用户模块
// @Param name query string true "用户名"
// @Param password query string true "密码"
// @Param repassword query string true "确认密码"
// @Produce json
// @Success 200 {string} json {"code","message"}
// @Router /user/createUser [get]
func CreateUser(c *gin.Context) {
	user := models.UserBasic{}
	user.Name = c.Query("name")
	passWord := c.Query("password")
	rePassWord := c.Query("repassword")

	username := models.FindUserByName(user.Name)
	if username.Name != "" {
		c.JSON(-1, gin.H{
			"code":    -1,
			"message": "用户名已注册！",
		})
		return
	}

	if passWord != rePassWord {
		c.JSON(-1, gin.H{
			"code":    -1,
			"message": "两次密码不一致！",
		})
		return
	}
	user.PassWord = passWord
	models.CreateUser(user)
	c.JSON(200, gin.H{
		"code":    0,
		"message": "新增用户成功！",
	})
}

// DeleteUser
// @Summary 删除用户
// @Tags 用户模块
// @Param id query string true "id"
// @Produce json
// @Success 200 {string} json {"code","message"}
// @Router /user/deleteUser [get]
func DeleteUser(c *gin.Context) {
	user := models.UserBasic{}
	id, _ := strconv.Atoi(c.Query("id"))
	user.ID = uint(id)

	models.DeleteUser(user)
	c.JSON(200, gin.H{
		"code":    0,
		"message": "删除用户成功！",
	})
}

// UpdateUser
// @Summary 修改用户
// @Tags 用户模块
// @Param id formData string true "id"
// @Param name formData string true "name"
// @Param password formData string true "password"
// @Produce json
// @Success 200 {string} json {"code","message"}
// @Router /user/updateUser [post]
func UpdateUser(c *gin.Context) {
	user := models.UserBasic{}
	id, _ := strconv.Atoi(c.PostForm("id"))
	user.ID = uint(id)
	user.Name = c.PostForm("name")
	user.PassWord = c.PostForm("password")

	models.UpdateUser(user)
	c.JSON(200, gin.H{
		"code":    0,
		"message": "更新用户成功！",
	})
}
