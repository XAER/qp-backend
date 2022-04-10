package controllers

import (
	"net/http"
	"qp/back/config"
	"qp/back/models"

	"github.com/gin-gonic/gin"
)

func PingHandler(c *gin.Context) {
	c.JSON(http.StatusOK, models.PingResponse{
		Message: "Pong da " + config.App.C.SERVER_NAME,
		Mode:    config.App.GetMode(),
	})
}
