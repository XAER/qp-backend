package routes

import (
	"qp/back/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.SetTrustedProxies(nil)

	v1 := r.Group("/v1")
	{
		v1.GET("/ping", controllers.PingHandler)
		v1.POST("/searchLineData", controllers.SearchLineDataHandler)
	}

	return r
}
