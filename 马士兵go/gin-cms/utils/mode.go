package utils

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func SetMode() {
	switch strings.ToLower(viper.GetString("app.mode")) {
	case "release":
		gin.SetMode(gin.ReleaseMode)
	case "test":
		gin.SetMode(gin.TestMode)
	case "debug":
		fallthrough
	default:
		gin.SetMode(gin.DebugMode)
	}
}
