package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vschoener/auth.vincelivemix/src/controllers/dto"
	apperrors "github.com/vschoener/auth.vincelivemix/src/errors"
	"github.com/vschoener/auth.vincelivemix/src/services"
)

// AuthController structure
type AuthController struct {
	authService *services.AuthService
}

// ProvideAuthController provide the AuthController
func ProvideAuthController(authService *services.AuthService) AuthController {
	return AuthController{
		authService: authService,
	}
}

// Login route
func (a AuthController) Login(c *gin.Context) {
	var userRequest dto.UserRequest
	if err := c.ShouldBindJSON(&userRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":  err.Error(),
			"status": http.StatusBadRequest,
		})
		return
	}

	body, err := a.authService.HandleLoginRequest(userRequest)
	if err != nil {
		if err == apperrors.ErrInvalidCredential {
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
