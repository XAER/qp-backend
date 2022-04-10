package config

import (
	"github.com/gin-gonic/gin"
)

type GConfig struct {
	SERVER_NAME string
	SERVER_PORT string
	SERVER_ENV  string
	DATA_SERVER string
}

func SetGinMode(mode string) {
	switch mode {
	case "release":
		gin.SetMode(gin.ReleaseMode)
	case "test":
		gin.SetMode(gin.TestMode)
	default:
		gin.SetMode(gin.DebugMode)
	}
}
