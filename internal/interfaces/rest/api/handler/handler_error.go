package handler

import (
	"auth-system/internal/core/user"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func handlerError(c *gin.Context, err error) {
	switch {
	case errors.Is(err, user.ErrEmptyName):
		c.JSON(http.StatusBadRequest, gin.H{
			"error":  "name cannot be empty",
			"status": http.StatusBadRequest,
			"data":   nil,
		})

	default:
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":  "something went wrong",
			"status": http.StatusInternalServerError,
			"data":   nil,
		})
	}
}
