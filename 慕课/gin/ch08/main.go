package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func MyLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		// time.Sleep(1 * time.Second)
		c.Set("example", "123456")
		c.Next() //让原本该执行的逻辑继续执行
		end := time.Since(t)
		fmt.Printf("耗时:%V\n", end)
		status := c.Writer.Status()
		fmt.Println("状态", status)
	}

}

func TokenRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		var token string
		for k, v := range c.Request.Header {
			if k == "X-Token" {
				token = v[0]
			}
		}
		if token != "dawn" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"msg": "未登录",
			})
			c.Abort()
		}
		c.Next()
	}
}
func main() {
	// router := gin.New()
	// // 全局使用Logger中间件和Recovery中间件
	// router.Use(gin.Logger(), gin.Recovery())

	router := gin.Default()
	router.Use(TokenRequired())
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	router.Run(":8083")
}
