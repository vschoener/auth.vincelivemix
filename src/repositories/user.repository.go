package repositories

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v4"
	"github.com/vschoener/auth.vincelivemix/src/database"
	"github.com/vschoener/auth.vincelivemix/src/entity"
)

type UserRepository struct {
	database *database.Database
}

func ProvideUserRepository(database *database.Database) *UserRepository {
	fmt.Println("Provide User Repository", database)
	return &UserRepository{
		database: database,
	}
}

func (u *UserRepository) FindUserWithEmailAndPassword(email string, password string) (*entity.User, error) {
	user := &entity.User{}

	if err := u.database.Connection.
		QueryRow(context.Background(),
			"SELECT user_id, email FROM users WHERE email=$1 AND password=$2",
			email, password).
		Scan(&user.ID, &user.Email); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}

		return nil, err
	}

	return user, nil
}
