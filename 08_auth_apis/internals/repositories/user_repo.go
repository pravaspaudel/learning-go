package repositories

import (
	"database/sql"

	"github.com/pravaspaudel/10_Auth_Apis/internals/models"
)

type UserRepositories struct {
	Db *sql.DB
}

func (u *UserRepositories) CreateUser(user models.User) (models.User, error) {
	query := `INSERT INTO users(username,email,password) 
			  VALUES ($1, $2, $3) 
			  RETURNING id,username,email,password`

	var createdUser models.User

	err := u.Db.QueryRow(query, user.Username, user.Email, user.Password).Scan(&createdUser.Id, &createdUser.Username, &createdUser.Email, &createdUser.Password)

	if err != nil {
		return models.User{}, err
	}
	return createdUser, nil
}

func (u *UserRepositories) GetUserByEmail(email string) (models.User, error) {
	var user models.User

	query := `SELECT id,username,email,password 
			  FROM users
			  WHERE email=$1`

	err := u.Db.QueryRow(query, email).Scan(&user.Id,
		&user.Username,
		&user.Email,
		&user.Password)

	if err != nil {
		return models.User{}, err
	}

	return user, err
}
