package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

/*
优雅退出gin程序:让程序处理完再关闭
*/

func main() {
	router := gin.Default()
	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"msg": "pong",
		})
	})
	go func() {
		router.Run(":8083")
	}()
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	// 处理关闭逻辑
	time.Sleep(2 * time.Second)
	fmt.Println("关闭server中...")
	fmt.Println("注销服务...")
}
