package services

import (
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"

	"github.com/vschoener/auth.vincelivemix/src/config"
	"github.com/vschoener/auth.vincelivemix/src/controllers/dto"
	"github.com/vschoener/auth.vincelivemix/src/entity"
	apperrors "github.com/vschoener/auth.vincelivemix/src/errors"
)

// AuthService struct
type AuthService struct {
	securityConfig config.SecurityConfig
	userService    *UserService
}

// ProvideAuthService provide the Auth service
func ProvideAuthService(securityConfig config.SecurityConfig, userService *UserService) *AuthService {
	return &AuthService{
		securityConfig: securityConfig,
		userService:    userService,
	}
}

// HandleLoginRequest handle the user login
func (a AuthService) HandleLoginRequest(userRequest dto.UserRequest) (*dto.AuthenticatedResponse, error) {
	user, err := a.userService.GetUserByEmail(userRequest.Username)

	if err != nil {
		// We don't need to be specific, just InvalidCredential
		if errors.Is(err, apperrors.ErrUserNotFound) {
			return nil, apperrors.ErrInvalidCredential
		}

		return nil, err
	}

	if a.DoesPasswordMatch(userRequest.Password, user.Password) == false {
		return nil, apperrors.ErrInvalidCredential
	}

	token, err := a.CreateUserToken(user)

	if err != nil {
		return nil, err
	}

	return &dto.AuthenticatedResponse{
		Token: token,
	}, nil
}

// DoesPasswordMatch check if the password match
func (a AuthService) DoesPasswordMatch(plainPassword string, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword))
	if err != nil {
		fmt.Println(err)
		return false
	}

	return true
}

// CreateUserToken create the token
func (a AuthService) CreateUserToken(user *entity.User) (string, error) {
	atClaims := jwt.MapClaims{}

	currentTime := time.Now()
	atClaims["iss"] = "Vince live mix auth server"
	atClaims["exp"] = currentTime.Add(time.Duration(a.securityConfig.TokenLifeTime) * time.Minute).Unix()
	atClaims["iat"] = currentTime.Unix()
	atClaims["sub"] = "local|" + user.ID
	atClaims["roles"] = []string{"ADMIN"}

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)

	token, err := at.SignedString([]byte(a.securityConfig.AuthPrivateKey))
	if err != nil {
		return "", err
	}

	return token, nil
}
