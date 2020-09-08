package services

import (
	"errors"
	"fmt"
	"strings"
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

type jwtMetadata struct {
	UserLocation string
	UserID       string
	Roles        []string
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

func (a AuthService) VerifyToken(requestToken string) (*jwt.Token, error) {
	token, err := jwt.Parse(requestToken, func(token *jwt.Token) (interface{}, error) {
		//Make sure that the token method conform to "SigningMethodHMAC"
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(a.securityConfig.AuthPrivateKey), nil
	})

	if err != nil {
		return nil, err
	}

	return token, nil
}

func (a AuthService) ExtractTokenMetadata(requestToken string) (*jwtMetadata, error) {
	token, err := a.VerifyToken(requestToken)
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		sub, ok := claims["sub"].(string)
		if !ok {
			return nil, err
		}

		rolesSlice, ok := claims["roles"].([]interface{})
		if !ok {
			return nil, err
		}

		// Need conversion from concrete type
		roles := make([]string, len(rolesSlice))
		for i, v := range rolesSlice {
			roles[i] = fmt.Sprint(v)
		}

		userInfo := strings.Split(sub, "|")

		return &jwtMetadata{
			UserLocation: userInfo[0],
			UserID:       userInfo[1],
			Roles:        roles,
		}, nil
	}

	return nil, err
}

func (a AuthService) GetUserFromToken(requestToken string) (*dto.UserFetchResponse, error) {
	jwtMetadata, err := a.ExtractTokenMetadata(requestToken)
	if err != nil {
		return nil, err
	}

	return &dto.UserFetchResponse{
		UserID: jwtMetadata.UserID,
		Roles:  jwtMetadata.Roles,
	}, nil
}
