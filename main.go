package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

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

	http.HandleFunc("/api/userservice", UserService)
	http.HandleFunc("/api/userservice/get/", UserServiceWithId)
	http.HandleFunc("/api/userservice/register", UserRegistration)
	http.HandleFunc("/api/userservice/update/", UpdateUser)
	http.HandleFunc("/api/userservice/delete/", UserDelete)
	http.HandleFunc("/api/userservice/login", UserLogin)
	http.HandleFunc("/api/userservice/verifytoken", verifyToken)

	http.HandleFunc("/getBooks", GetBooks)
	http.HandleFunc("/getBookByName/", handleBookByName)
	http.HandleFunc("/getBook/", GetBook)
	http.HandleFunc("/addBook", AddBook)
	http.HandleFunc("/update/", UpdateBook)
	http.HandleFunc("/delete/", Delete)

	http.HandleFunc("/api/order/insert", Insertion)
	http.HandleFunc("/api/order/retrieveAllOrder", RetrieveAllOrders)
	http.HandleFunc("/api/order/retrieveOrder/", RetrieveOrderByID)
	http.HandleFunc("/api/order/cancelOrder/", CancelOrderByID)

	http.HandleFunc("/api/cart/create", Create)
	http.HandleFunc("/api/cart/getById/", GetById)
	http.HandleFunc("/api/cart/updateById/", UpdateById)
	http.HandleFunc("/api/cart/delete/", Delete)
	http.HandleFunc("/api/cart/increaseQuantity/", IncreaseQuantity)
	http.HandleFunc("/api/cart/decreaseQuantity/", DecreaseQuantity)
	http.ListenAndServe(":8000", nil)
}
