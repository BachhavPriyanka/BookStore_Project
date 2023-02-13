package bookstorecontroller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	repository "github.com/BachhavPriyanka/BookStore_Project/storage"
	"github.com/BachhavPriyanka/BookStore_Project/types"
	tokenutil "github.com/BachhavPriyanka/BookStore_Project/util"

	_ "github.com/go-sql-driver/mysql"
)

type UserRegister struct {
	Repository *repository.UserRegistrationrepository
}

// Get Method to verify the tokens
func (c *UserRegister) VerifyToken(writer http.ResponseWriter, request *http.Request) {
	jwtToken := request.FormValue("jwttoken")

	useridInt64, _ := tokenutil.DecodeToken(jwtToken)
	userid := int(useridInt64)
	fmt.Println(userid)

	useridString := strconv.Itoa(userid)

	fmt.Println(useridString)

	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(useridString)
}

// POST Method for user login
func (c *UserRegister) UserLogin(writer http.ResponseWriter, request *http.Request) {
	var loginDetails types.LoginDTO
	json.NewDecoder(request.Body).Decode(&loginDetails)
	userName := loginDetails.Email
	passWord := loginDetails.Password
	fmt.Println("Username", userName)
	fmt.Println("password", passWord)

	userValidation, err := c.Repository.ValidateUserCredentials(userName, passWord)
	if err != nil {
		http.Error(writer, "User not Valid", http.StatusInternalServerError)
		return
	}
	fmt.Println(userValidation)

	token, err := tokenutil.CreateToken(int64(userValidation.Id))
	if err != nil {
		http.Error(writer, "Error creating token", http.StatusInternalServerError)
	}

	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(token)

}

// DELETE Method to delete the user record
func (c *UserRegister) UserDelete(writer http.ResponseWriter, request *http.Request) {
	userId, err := strconv.Atoi(request.URL.Path[len("/api/userservice/update/"):])
	if err != nil {
		http.Error(writer, "Id is not present", http.StatusBadRequest)
	}
	deleteUser, err := c.Repository.UserDelete(int(userId))
	if err != nil {
		http.Error(writer, "Error", http.StatusBadRequest)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(deleteUser)
}

// PUT Method to update the user
func (c *UserRegister) UpdateUser(writer http.ResponseWriter, request *http.Request) {
	userId, err := strconv.Atoi(request.URL.Path[len("/api/userservice/update/"):])
	if err != nil {
		http.Error(writer, "Id is not present", http.StatusBadRequest)
	}
	var userData types.Users
	json.NewDecoder(request.Body).Decode(&userData)

	updatedData, err := c.Repository.UpdateUser(int(userId), &userData)
	if err != nil {
		http.Error(writer, "Error", http.StatusBadRequest)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(updatedData)
}

// POST Method for user registration
func (c *UserRegister) UserRegistration(writer http.ResponseWriter, request *http.Request) {
	var userData types.Users
	json.NewDecoder(request.Body).Decode(&userData)

	response, err := c.Repository.UserRegistration(&userData)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(response)
}

// GET Method to retrieve single user record
func (c *UserRegister) UserServiceWithId(writer http.ResponseWriter, request *http.Request) {
	headerToken := request.Header.Get("Authorization")

	userId, err := tokenutil.DecodeToken(headerToken)
	if err != nil {
		http.Error(writer, "Unauthorized", http.StatusUnauthorized)
	}
	fmt.Println("userid int int64", userId)

	retrieveData, err := c.Repository.UserServiceWithId(int(userId))
	if err != nil {
		http.Error(writer, "Error", http.StatusBadRequest)
		return
	}
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(retrieveData)
}

// GET Method to retrieve all user records
func (c *UserRegister) UserService(writer http.ResponseWriter, request *http.Request) {
	retrieveData, err := c.Repository.UserService()
	if err != nil {
		http.Error(writer, "Error", http.StatusBadRequest)
		return
	}
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(retrieveData)
}
