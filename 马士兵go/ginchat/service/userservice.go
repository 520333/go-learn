package service

import (
	"fmt"
	"ginchat/models"
	"ginchat/utils"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
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
		"code":    0,
		"message": data,
		"data":    data,
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
	salt := fmt.Sprintf("%06d", rand.Int31())

	username := models.FindUserByName(user.Name)
	if username.Name != "" {
		c.JSON(-1, gin.H{
			"code":    -1,
			"message": "用户名已注册！",
			"data":    user,
		})
		return
	}

	if passWord != rePassWord {
		c.JSON(-1, gin.H{
			"code":    -1,
			"message": "两次密码不一致！",
			"data":    user,
		})
		return
	}
	//user.PassWord = passWord
	user.PassWord = utils.MakePassword(passWord, salt)
	user.Salt = salt
	fmt.Println(user.PassWord)
	models.CreateUser(user)
	c.JSON(200, gin.H{
		"code":    0,
		"message": "新增用户成功！",
		"data":    user,
	})
}

// FindUserByNameAndPwd
// @Summary 用户登录
// @Tags 用户模块
// @Param name query string true "用户名"
// @Param password query string true "密码"
// @Produce json
// @Success 200 {string} json {"code","message"}
// @Router /user/findUserByNameAndPwd [post]
func FindUserByNameAndPwd(c *gin.Context) {
	data := models.UserBasic{}
	name := c.Query("name")
	password := c.Query("password")
	user := models.FindUserByName(name)
	if user.Name == "" {
		c.JSON(200, gin.H{
			"code":    -1,
			"message": "该用户不存在",
			"data":    data,
		})
		return
	}
	fmt.Println(user)

	flag := utils.ValidPassword(password, user.Salt, user.PassWord)
	if !flag {
		c.JSON(200, gin.H{
			"code":    -1,
			"message": "密码不正确",
			"data":    data,
		})
		return
	}
	pwd := utils.MakePassword(password, user.Salt)
	data = models.FindUserByNameAndPwd(name, pwd)
	c.JSON(200, gin.H{
		"code":    0, // 0成功 -1失败
		"message": "登录成功",
		"data":    data,
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
		"message": "删除用户成功!",
		"data":    user,
	})
	return
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
		"data":    user,
	})
}

// 防止跨域站点伪造请求
var upGrade = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func SendMsg(c *gin.Context) {
	ws, err := upGrade.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func(ws *websocket.Conn) {
		err := ws.Close()
		if err != nil {
			log.Fatalln(err)
		}
	}(ws)
	MsgHandler(ws, c)
}

func MsgHandler(ws *websocket.Conn, c *gin.Context) {
	for {
		msg, err := utils.Subscribe(c, utils.PublishKey)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("发送消息：")
		tm := time.Now().Format("2006-01-02 15:04:05")
		m := fmt.Sprintf("[ws][%s]:%s", tm, msg)
		err = ws.WriteMessage(1, []byte(m))
		if err != nil {
			fmt.Println(err)
		}
	}
}
