package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	v1 := router.Group("/v1")
	{
		v1.POST("/login")
		v1.POST("/submit")
		v1.POST("/read")
	}
	v2 := router.Group("/v2")
	{
		v2.POST("/login")
		v2.POST("/submit")
		v2.POST("/read")
	}
	router.Run("8082")
}
