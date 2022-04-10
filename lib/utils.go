package lib

import (
	"qp/back/models"

	"github.com/gin-gonic/gin"
)

func RespondWithError(c *gin.Context, code int, message string) {
	c.JSON(code, models.ErrorResponse{
		Message: message,
	})
}
