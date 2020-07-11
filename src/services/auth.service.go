package services

import (
	"time"

	"github.com/dgrijalva/jwt-go"

	"github.com/vschoener/auth.vincelivemix/src/controllers/dto"
	"github.com/vschoener/auth.vincelivemix/src/errors"
)

type AuthService struct{}

func ProvideAuthService() AuthService {
	return AuthService{}
}

// HandleLoginRequest handle the user login
func (a AuthService) HandleLoginRequest(userRequest dto.UserRequest) (dto.AuthenticatedResponse, error) {
	if userRequest.Password != "admin" || userRequest.Username != "admin" {
		return dto.AuthenticatedResponse{}, errors.ErrInvalidCredential
	}

	token, err := a.CreateUserToken(userRequest)

	if err != nil {
		return dto.AuthenticatedResponse{}, err
	}

	return dto.AuthenticatedResponse{
		Token: token,
	}, nil
}

// CreateUserToken create the token
func (AuthService) CreateUserToken(user dto.UserRequest) (string, error) {
	atClaims := jwt.MapClaims{}

	atClaims["iss"] = "Vince live mix auth server"
	atClaims["iat"] = time.Now().Add(time.Minute * 15).Unix()
	atClaims["sub"] = "local|1"
	atClaims["roles"] = []string{"ADMIN"}

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)

	token, err := at.SignedString([]byte("SECRET"))
	if err != nil {
		return "", err
	}

	return token, nil
}
