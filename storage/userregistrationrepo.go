package repository

import (
	"database/sql"
	"fmt"

	"github.com/BachhavPriyanka/BookStore_Project/types"
)

type UserRegistrationrepository struct {
	DB *sql.DB
}

func (r *UserRegistrationrepository) ValidateUserCredentials(username string, password string) (*types.Users, error) {
	var user types.Users

	err := r.DB.QueryRow("SELECT * from users where email = ? and password = ?", username, password).Scan(&user.Id, &user.UserName, &user.Email, &user.Password)
	if err != nil {
		return nil, fmt.Errorf("Invalid credentials")

	}
	return &user, nil

}
