package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vschoener/auth.vincelivemix/src/controllers/dto"
	"github.com/vschoener/auth.vincelivemix/src/errors"
	"github.com/vschoener/auth.vincelivemix/src/services"
)

// Login catch the login request
func Login(c *gin.Context) {
	var userRequest dto.UserRequest
	if err := c.ShouldBindJSON(&userRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":  err.Error(),
			"status": http.StatusBadRequest,
		})
		return
	}

	body, err := services.HandleLoginRequest(userRequest)
	if err != nil {
		if err == errors.ErrInvalidCredential {
			c.JSON(http.StatusForbidden, gin.H{
				"message": err.Error(),
				"status":  http.StatusForbidden,
			})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"status":  http.StatusInternalServerError,
		})
		return
	}

	c.JSON(http.StatusOK, body)
}
