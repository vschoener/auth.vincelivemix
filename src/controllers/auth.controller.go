package controllers

import (
	"net/http"
	"strings"

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

// GetUser fetch user
func (a AuthController) GetUser(c *gin.Context) {
	authorization := c.GetHeader("Authorization")

	authorizations := strings.Split(authorization, " ")

	if len(authorizations) != 2 {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized",
			"status":  http.StatusUnauthorized,
		})
		return
	}

	userFetch, err := a.authService.GetUserFromToken(authorizations[1])

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized",
			"status":  http.StatusUnauthorized,
		})
		return
	}

	c.JSON(http.StatusOK, userFetch)
}
