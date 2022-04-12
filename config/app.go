package config

import (
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

var App *Services

type Services struct {
	DB   *gorm.DB
	R    *gin.Engine
	C    *GConfig
	MODE string
}

func NewServices(DB *gorm.DB, R *gin.Engine, C *GConfig, MODE string) *Services {
	return &Services{
		DB:   DB,
		R:    R,
		C:    C,
		MODE: MODE,
	}
}

func (s *Services) GetMode() string {
	return s.MODE
}
