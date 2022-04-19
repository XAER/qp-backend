package routes

import (
	"qp/back/controllers"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.SetTrustedProxies(nil)

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization", "accept", "origin", "Cache-Control", "X-Requested-With"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return true
		},
		MaxAge: 15 * time.Second,
	}))

	v1 := r.Group("/v1")
	{
		v1.GET("/ping", controllers.PingHandler)
		v1.POST("/searchLineData", controllers.SearchLineDataHandler)
		v1.GET("/testDBConnection", controllers.TestDBConnectionHandler)
		v1.GET("/getStopRealTimeData/:languageName/:stopCode", controllers.GetStopRealTimeDataHandler)
	}

	return r
}
