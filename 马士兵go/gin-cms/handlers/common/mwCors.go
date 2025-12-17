package common

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func UseCors(engine *gin.Engine) {
	//1.设置中间件
	cfg := cors.DefaultConfig()
	cfg.AllowAllOrigins = true
	cfg.AllowCredentials = true
	cfg.AddAllowHeaders("Authorization")
	//2.初始化 并使用中间件
	engine.Use(cors.New(cfg))
}
