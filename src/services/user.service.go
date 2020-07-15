package services

import (
	"github.com/vschoener/auth.vincelivemix/src/entity"
	apperrors "github.com/vschoener/auth.vincelivemix/src/errors"
	"github.com/vschoener/auth.vincelivemix/src/repositories"
)

// UserService structure
type UserService struct {
	userRepository *repositories.UserRepository
}

// ProvideUserService provides the user service
func ProvideUserService(userRepository *repositories.UserRepository) *UserService {
	return &UserService{
		userRepository: userRepository,
	}
}

// GetUserByEmailAndPassword retrieves the user by email and its password
func (u *UserService) GetUserByEmail(email string) (*entity.User, error) {
	user, err := u.userRepository.FindUserByEmail(email)

	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, apperrors.ErrUserNotFound
	}

	return user, err
}
