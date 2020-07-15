package repositories

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v4"
	"github.com/vschoener/auth.vincelivemix/src/database"
	"github.com/vschoener/auth.vincelivemix/src/entity"
)

// UserRepository structure
type UserRepository struct {
	database *database.Database
}

// ProvideUserRepository provide the user repository
func ProvideUserRepository(database *database.Database) *UserRepository {
	fmt.Println("Provide User Repository", database)
	return &UserRepository{
		database: database,
	}
}

// FindUserByEmail will query database to find the user by its email
func (u *UserRepository) FindUserByEmail(email string) (*entity.User, error) {
	user := &entity.User{}

	if err := u.database.Connection.
		QueryRow(context.Background(),
			"SELECT user_id, email, password FROM users WHERE email=$1",
			email).
		Scan(&user.ID, &user.Email, &user.Password); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}

		return nil, err
	}

	return user, nil
}
