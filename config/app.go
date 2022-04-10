package config

import "github.com/gin-gonic/gin"

var App *Services

type Services struct {
	R    *gin.Engine
	C    *GConfig
	MODE string
}

func NewServices(R *gin.Engine, C *GConfig, MODE string) *Services {
	return &Services{
		R:    R,
		C:    C,
		MODE: MODE,
	}
}

func (s *Services) GetMode() string {
	return s.MODE
}
