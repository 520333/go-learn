package router

import (
	"ginchat/docs"
	"ginchat/service"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Router() *gin.Engine {
	r := gin.Default()
	// swagger
	docs.SwaggerInfo.BasePath = ""
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	// 静态资源
	r.Static("/asset", "asset/")

	r.LoadHTMLGlob("view/**/*")

	// 首页
	r.GET("/", service.GetIndex)
	r.GET("/index", service.GetIndex)
	r.GET("/toRegister", service.ToRegister)
	r.GET("/toChat", service.ToChat)
	r.GET("/chat", service.Chat)

	r.POST("/searchFriends", service.SearchFriends) // 查询好友列表
	// 用户模块
	r.POST("/user/getUserList", service.GetUserList)
	r.POST("/user/createUser", service.CreateUser)
	r.POST("/user/deleteUser", service.DeleteUser)

	// 发送消息
	r.POST("/user/updateUser", service.UpdateUser)
	r.POST("/user/findUserByNameAndPwd", service.FindUserByNameAndPwd)

	// 发送消息
	r.GET("/user/SendMsg", service.SendMsg)
	r.GET("/user/SendUserMsg", service.SendUserMsg)
	r.POST("/attach/upload", service.Upload)                    // 发送图片
	r.POST("/contact/addfriend", service.AddFriends)            // 添加好友
	r.POST("/contact/createCommunity", service.CreateCommunity) // 创建群

	return r
}
