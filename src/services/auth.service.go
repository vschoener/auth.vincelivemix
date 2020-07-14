package services

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"

	"github.com/vschoener/auth.vincelivemix/src/controllers/dto"
	"github.com/vschoener/auth.vincelivemix/src/entity"
	apperrors "github.com/vschoener/auth.vincelivemix/src/errors"
)

// AuthService struct
type AuthService struct {
	userService *UserService
}

// ProvideAuthService provide the Auth service
func ProvideAuthService(userService *UserService) *AuthService {
	return &AuthService{
		userService: userService,
	}
}

// HandleLoginRequest handle the user login
func (a AuthService) HandleLoginRequest(userRequest dto.UserRequest) (*dto.AuthenticatedResponse, error) {
	user, err := a.userService.GetUserByEmailAndPassword(userRequest.Username, userRequest.Password)

	if err != nil {
		// We don't need to be specific, just InvalidCredential
		if errors.Is(err, apperrors.ErrUserNotFound) {
			return nil, apperrors.ErrInvalidCredential
		}

		return nil, err
	}

	token, err := a.CreateUserToken(user)

	if err != nil {
		return nil, err
	}

	return &dto.AuthenticatedResponse{
		Token: token,
	}, nil
}

// CreateUserToken create the token
func (AuthService) CreateUserToken(user *entity.User) (string, error) {
	atClaims := jwt.MapClaims{}

	atClaims["iss"] = "Vince live mix auth server"
	atClaims["iat"] = time.Now().Add(time.Minute * 15).Unix()
	atClaims["sub"] = "local|" + user.ID
	atClaims["roles"] = []string{"ADMIN"}

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)

	token, err := at.SignedString([]byte("SECRET"))
	if err != nil {
		return "", err
	}

	return token, nil
}
