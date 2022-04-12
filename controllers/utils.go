package controllers

import (
	"errors"
	"net/http"
	"qp/back/config"
	"qp/back/lib"
	"qp/back/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func PingHandler(c *gin.Context) {
	c.JSON(http.StatusOK, models.PingResponse{
		Message: "Pong da " + config.App.C.SERVER_NAME,
		Mode:    config.App.GetMode(),
	})
}

func TestDBConnectionHandler(c *gin.Context) {
	Logs := models.Logs{}
	result := config.App.DB.First(&Logs)

	errors.Is(result.Error, gorm.ErrRecordNotFound)
	if result.Error != nil {
		lib.RespondWithError(c, http.StatusInternalServerError, result.Error.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":          "Successfully connected to the database",
		"connectionStatus": "Connected",
	})

}
