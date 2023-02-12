package main

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

// var db *sql.DB

// func main() {
// 	var err error
// 	//"UserName:Password@tcp(portNumber)/databaseName"
// 	db, err = sql.Open("mysql", "root:root@tcp(localhost:6603)/BookStore")
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	pingErr := db.Ping()
// 	if pingErr != nil {
// 		log.Fatal(pingErr)
// 	}
// 	fmt.Println("Connected")

// 	Operation()
// }

type UserRegister struct {
	Repository *repository.UserRegistrationrepository
}

// func Operation() {
// 	http.HandleFunc("/api/userservice", handleUserService)
// 	http.HandleFunc("/api/userservice/get/", handleUserServiceWithId)
// 	http.HandleFunc("/api/userservice/register", handleUserRegistration)
// 	http.HandleFunc("/api/userservice/update/", handleUpdateUser)
// 	http.HandleFunc("/api/userservice/delete/", handleUserDelete)
// 	http.HandleFunc("/api/userservice/login", handleUserLogin)
// 	http.HandleFunc("/api/userservice/verifytoken", verifyToken)
// 	http.ListenAndServe(":8000", nil)
// }

func verifyToken(w http.ResponseWriter, r *http.Request) {

	jwtToken := r.FormValue("jwttoken")

	useridInt64, _ := tokenutil.DecodeToken(jwtToken)

	userid := int(useridInt64)
	fmt.Println(userid)
	useridString := strconv.Itoa(userid)
	fmt.Println(useridString)
	w.Write([]byte(useridString))

}

func UserLogin(writer http.ResponseWriter, request *http.Request) {
	var loginDetails types.LoginDTO
	json.NewDecoder(request.Body).Decode(&loginDetails)
	userName := loginDetails.Email
	passWord := loginDetails.Password
	fmt.Println("Username", userName)
	fmt.Println("password", passWord)

	user := &repository.UserRegistrationrepository{
		DB: db,
	}

	userValidation, err := user.ValidateUserCredentials(userName, passWord)
	if err != nil {
		http.Error(writer, "User not Valid", http.StatusInternalServerError)
		return
	}
	fmt.Println(userValidation)

	token, err := tokenutil.CreateToken(int64(userValidation.Id))
	if err != nil {
		http.Error(writer, "Error creating token", http.StatusInternalServerError)
	}
	writer.Write([]byte(token))

}

func UserDelete(writer http.ResponseWriter, request *http.Request) {
	if request.Method == http.MethodDelete {
		id, err := strconv.Atoi(request.URL.Path[len("/api/userservice/delete/"):])
		if err != nil {
			http.Error(writer, "Error", http.StatusBadRequest)
			return
		}

		_, err = db.Exec("DELETE from users WHERE id = ?", id)
		if err != nil {
			panic(err.Error())
		}

		writer.Write([]byte("Record Deleted"))
	}
}

func UpdateUser(writer http.ResponseWriter, request *http.Request) {
	if request.Method == http.MethodPut {

		var dto types.Users
		json.NewDecoder(request.Body).Decode(&dto)

		id, err := strconv.Atoi(request.URL.Path[len("/api/userservice/update/"):])
		if err != nil {
			http.Error(writer, "Error", http.StatusBadRequest)
			return
		}
		_, err = db.Exec("UPDATE users SET id=?, username=?, email=?, password=? WHERE id=?", dto.Id, dto.UserName, dto.Email, dto.Password, id)
		if err != nil {
			http.Error(writer, "Error updating user", http.StatusInternalServerError)
			return
		}
		writer.Write([]byte("Record Updated"))
		writer.WriteHeader(http.StatusOK)
	}
}

func UserRegistration(writer http.ResponseWriter, request *http.Request) {
	if request.Method == http.MethodPost {
		var user types.Users
		json.NewDecoder(request.Body).Decode(&user)
		_, err := db.Exec("Insert into users (id, username, email, password) values (?,?,?,?)", user.Id, user.UserName, user.Email, user.Password)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}
		writer.Write([]byte("Record Added"))
	}
}

func UserServiceWithId(writer http.ResponseWriter, request *http.Request) {
	if request.Method == http.MethodGet {
		id, err := strconv.Atoi(request.URL.Path[len("/api/userservice/get/"):])
		if err != nil {
			http.Error(writer, "Error", http.StatusBadRequest)
			return
		}
		rows, err := db.Query("SELECT id, username, email, password FROM users WHERE id = ?", id)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		users := []types.Users{}

		for rows.Next() {
			var user types.Users
			err = rows.Scan(&user.Id, &user.UserName, &user.Email, &user.Password)
			if err != nil {
				http.Error(writer, err.Error(), http.StatusInternalServerError)
			}

			users = append(users, user)
		}
		json.NewEncoder(writer).Encode(users)
	}
}

func UserService(writer http.ResponseWriter, request *http.Request) {
	if request.Method == http.MethodGet {
		rows, err := db.Query("SELECT id,username,email,password from users")
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		users := []types.Users{}

		for rows.Next() {
			var user types.Users
			err = rows.Scan(&user.Id, &user.UserName, &user.Email, &user.Password)
			if err != nil {
				http.Error(writer, err.Error(), http.StatusInternalServerError)
			}
			users = append(users, user)
		}
		json.NewEncoder(writer).Encode(users)
	}
}
