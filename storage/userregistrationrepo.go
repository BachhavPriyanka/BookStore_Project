package repository

import (
	"database/sql"
	"fmt"

	"github.com/BachhavPriyanka/BookStore_Project/constant"
	"github.com/BachhavPriyanka/BookStore_Project/types"
	_ "github.com/go-sql-driver/mysql"
)

type UserRegistrationrepository struct {
	DB *sql.DB
}

// Method for Validatation
func (r *UserRegistrationrepository) ValidateUserCredentials(username string, password string) (*types.Users, error) {
	var user types.Users

	err := r.DB.QueryRow("SELECT * from users where email = ? and password = ?", username, password).Scan(&user.Id, &user.UserName, &user.Email, &user.Password)
	if err != nil {
		return nil, fmt.Errorf("invalid credentials")

	}
	return &user, nil
}

// DELETE Method to delete user record
func (r *UserRegistrationrepository) UserDelete(id int) (string, error) {
	_, err := r.DB.Exec(constant.DeleteRecord, id)
	if err != nil {
		return "invalid User Id", err
	}
	return "Successfully Deleted", err
}

// PUT Method to update the user record
func (r *UserRegistrationrepository) UpdateUser(id int, updateData *types.Users) (string, error) {
	_, err := r.DB.Exec(constant.PutUpdateQuery, updateData.Id, updateData.UserName, updateData.Email, updateData.Password, id)
	if err != nil {
		return "unsuccessfull to update", err
	}
	return "successfully updated", nil
}

// POST Method to add user details
func (r *UserRegistrationrepository) UserRegistration(newUserData *types.Users) (string, error) {
	_, err := r.DB.Exec(constant.PostDetailsQuery, newUserData.Id, newUserData.UserName, newUserData.Email, newUserData.Password)
	if err != nil {
		return "unsuccessfull to Insert", err
	}
	return "successfully updated", nil
}

// GET Method to retrieve single user record
func (r *UserRegistrationrepository) UserServiceWithId(id int) (*[]types.Users, error) {
	rows, err := r.DB.Query(constant.GetRecordQuery, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := []types.Users{}

	for rows.Next() {
		var user types.Users
		err = rows.Scan(&user.Id, &user.UserName, &user.Email, &user.Password)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return &users, err
}

// GET Method to retrieve all user records
func (r *UserRegistrationrepository) UserService() (*[]types.Users, error) {
	rows, err := r.DB.Query(constant.GetAllRecordQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := []types.Users{}

	for rows.Next() {
		var user types.Users
		err = rows.Scan(&user.Id, &user.UserName, &user.Email, &user.Password)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return &users, err
}
