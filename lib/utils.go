package lib

import (
	"errors"
	"qp/back/config"
	"qp/back/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RespondWithError(c *gin.Context, code int, message string) {
	c.JSON(code, models.ErrorResponse{
		Message: message,
	})
}

func LogRequest(code int, event string, client string, outcome string) error {
	Logs := models.Logs{
		EVENT:   event,
		OUTCOME: outcome,
		CLIENT:  client,
	}

	result := config.App.DB.Create(&Logs)

	errors.Is(result.Error, gorm.ErrInvalidValue)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
