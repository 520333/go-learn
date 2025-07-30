package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	// 使用路由
	goodsGroup := router.Group("/goods")
	{
		goodsGroup.GET("/list", goodsList)
		goodsGroup.GET("/:id/:action", goodsDetail)
		goodsGroup.POST("/add", createGoods)
	}

	// 不使用路由
	// router.GET("/goods/list", goodsList)
	// router.GET("/goods/1", goodsDetail)
	// router.GET("/goods/add", createGoods)
	router.Run(":8082")
}

func goodsList(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"name": "goodsList",
	})
}

func goodsDetail(c *gin.Context) {
	id := c.Param("id")         // 传参
	action := c.Param("action") // 传参
	c.JSON(http.StatusOK, gin.H{
		"id":     id,
		"action": action,
	})
}
func createGoods(c *gin.Context) {}
