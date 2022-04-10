package main

import (
	"fmt"
	"log"
	"qp/back/config"
	"qp/back/helper"
	"qp/back/routes"

	"github.com/gin-gonic/gin"
)

var (
	ginConfig config.GConfig
	router    *gin.Engine
)

func initApp() {
	ginConfig.SERVER_NAME = helper.GetEnv("SERVER_NAME", "QP_BUS")
	ginConfig.SERVER_PORT = helper.GetEnv("SERVER_PORT", "8081")
	ginConfig.SERVER_ENV = helper.GetEnv("SERVER_ENV", "dev")
	ginConfig.DATA_SERVER = "https://www.atac.roma.it/Service/WebService.asmx"

	config.SetGinMode(ginConfig.SERVER_ENV)
	router = routes.SetupRouter()
}

func main() {
	initApp()

	config.App = config.NewServices(router, &ginConfig, ginConfig.SERVER_ENV)
	log.Fatal(router.Run(fmt.Sprintf("0.0.0.0:%s", ginConfig.SERVER_PORT)))
}
