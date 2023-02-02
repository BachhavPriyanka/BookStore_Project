package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

type Users struct {
	Id       int    `json:"id"`
	UserName string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func main() {
	var err error
	//"UserName:Password@tcp(portNumber)/databaseName"
	db, err = sql.Open("mysql", "root:root@tcp(localhost:6603)/BookStore")
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected")

	Operation()
}

func Operation() {
	http.HandleFunc("/api/userservice", handleUserService)
	http.HandleFunc("/api/userservice/get/", handleUserServiceWithId)
	http.HandleFunc("/api/userservice/register", handleUserRegistration)
	http.HandleFunc("/api/userservice/update/", handleUpdateUser)
	http.HandleFunc("/api/userservice/delete/", handleUserDelete)
	http.HandleFunc("/api/userservice/login", handleUserLogin)
	http.ListenAndServe(":8000", nil)
}

func handleUserLogin(writer http.ResponseWriter, request *http.Request) {
	if request.Method == http.MethodPost {
		var dto LoginDTO
		json.NewDecoder(request.Body).Decode(&dto)

		rows, err := db.Query("SELECT email,password from users where email = ? and password = ?", dto.Email, dto.Password)
		if err != nil {
			http.Error(writer, "User not found", http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		loginData := []LoginDTO{}

		for rows.Next() {
			var tempData LoginDTO
			err = rows.Scan(&tempData.Email, &tempData.Password)
			if err != nil {
				log.Fatal(err)
			}
			loginData = append(loginData, tempData)
		}
		if len(loginData) == 1 {
			writer.Write([]byte("Login Successful"))
		} else {
			writer.Write([]byte("Login Unsuccessful"))
		}
	}
}

func handleUserDelete(writer http.ResponseWriter, request *http.Request) {
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

func handleUpdateUser(writer http.ResponseWriter, request *http.Request) {
	if request.Method == http.MethodPut {

		var dto Users
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

func handleUserRegistration(writer http.ResponseWriter, request *http.Request) {
	if request.Method == http.MethodPost {
		var user Users
		json.NewDecoder(request.Body).Decode(&user)
		_, err := db.Exec("Insert into users (id, username, email, password) values (?,?,?,?)", user.Id, user.UserName, user.Email, user.Password)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}
		writer.Write([]byte("Record Added"))
	}
}

func handleUserServiceWithId(writer http.ResponseWriter, request *http.Request) {
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

		users := []Users{}

		for rows.Next() {
			var user Users
			err = rows.Scan(&user.Id, &user.UserName, &user.Email, &user.Password)
			if err != nil {
				http.Error(writer, err.Error(), http.StatusInternalServerError)
			}

			users = append(users, user)
		}
		json.NewEncoder(writer).Encode(users)
	}
}

func handleUserService(writer http.ResponseWriter, request *http.Request) {
	if request.Method == http.MethodGet {
		rows, err := db.Query("SELECT id,username,email,password from users")
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		users := []Users{}

		for rows.Next() {
			var user Users
			err = rows.Scan(&user.Id, &user.UserName, &user.Email, &user.Password)
			if err != nil {
				http.Error(writer, err.Error(), http.StatusInternalServerError)
			}
			users = append(users, user)
		}
		json.NewEncoder(writer).Encode(users)
	}
}
