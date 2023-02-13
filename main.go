package main

import (
	"fmt"
	"net/http"

	bookstorecontroller "github.com/BachhavPriyanka/BookStore_Project/controller"
	"github.com/BachhavPriyanka/BookStore_Project/middleware"
	repository "github.com/BachhavPriyanka/BookStore_Project/storage"
	"github.com/BachhavPriyanka/BookStore_Project/util"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	var err error

	db, err := util.DBConnector()
	if err != nil {
		fmt.Println("Error in Connecting")
		return
	}
	defer db.Close()
	fmt.Println("Connected")

	Bookcontroller := &bookstorecontroller.BookStoreController{
		Repository: &repository.BookStoreRepository{
			DB: db,
		},
	}

	UserRegistrationcontroller := &bookstorecontroller.UserRegister{
		Repository: &repository.UserRegistrationrepository{
			DB: db,
		},
	}

	Ordercontroller := &bookstorecontroller.OrderController{
		Repository: &repository.OrderRepository{
			DB: db,
		},
	}

	Cartcontroller := &bookstorecontroller.CartController{
		Repository: &repository.CartRepository{
			DB: db,
		},
	}

	http.HandleFunc("/api/userservice", UserRegistrationcontroller.UserService)
	http.Handle("/api/userservice/get/", middleware.LoggerMiddleware(middleware.AuthMiddleware(http.HandlerFunc(UserRegistrationcontroller.UserServiceWithId))))
	http.HandleFunc("/api/userservice/register", UserRegistrationcontroller.UserRegistration)
	http.HandleFunc("/api/userservice/update/", UserRegistrationcontroller.UpdateUser)
	http.HandleFunc("/api/userservice/delete/", UserRegistrationcontroller.UserDelete)
	http.HandleFunc("/api/userservice/login", UserRegistrationcontroller.UserLogin)
	http.HandleFunc("/api/userservice/verifytoken", UserRegistrationcontroller.VerifyToken)

	http.Handle("/books", middleware.LoggerMiddleware(middleware.AuthMiddleware(http.HandlerFunc(Bookcontroller.GetBooks))))

	http.HandleFunc("/getBooks", Bookcontroller.GetBooks)
	http.Handle("/getBookByName/", middleware.LoggerMiddleware(middleware.AuthMiddleware(http.HandlerFunc(Bookcontroller.HandleBookByName))))
	http.Handle("/getBook/", middleware.LoggerMiddleware(middleware.AuthMiddleware(http.HandlerFunc(Bookcontroller.GetBook))))
	http.Handle("/addBook", middleware.LoggerMiddleware(middleware.AuthMiddleware(http.HandlerFunc(Bookcontroller.AddBook))))
	http.Handle("/update/", middleware.LoggerMiddleware(middleware.AuthMiddleware(http.HandlerFunc(Bookcontroller.UpdateBook))))
	http.Handle("/delete/", middleware.LoggerMiddleware(middleware.AuthMiddleware(http.HandlerFunc(Bookcontroller.Delete))))

	http.Handle("/api/order/insert", middleware.LoggerMiddleware(http.HandlerFunc(Ordercontroller.Insertion)))
	http.HandleFunc("/api/order/retrieveAllOrder", Ordercontroller.RetrieveAllOrders)
	http.Handle("/api/order/retrieveOrder/", middleware.LoggerMiddleware(http.HandlerFunc(Ordercontroller.RetrieveOrderByID)))
	http.Handle("/api/order/cancelOrder/", middleware.LoggerMiddleware(middleware.AuthMiddleware(http.HandlerFunc(Ordercontroller.CancelOrderByID))))

	http.HandleFunc("/api/cart/create", Cartcontroller.Create)
	http.Handle("/api/cart/getById/", middleware.LoggerMiddleware(middleware.AuthMiddleware(http.HandlerFunc(Cartcontroller.GetById))))
	http.Handle("/api/cart/updateById/", middleware.LoggerMiddleware(middleware.AuthMiddleware(http.HandlerFunc(Cartcontroller.UpdateById))))
	http.Handle("/api/cart/delete/", middleware.LoggerMiddleware(middleware.AuthMiddleware(http.HandlerFunc(Cartcontroller.Delete))))
	http.Handle("/api/cart/increaseQuantity/", middleware.LoggerMiddleware(middleware.AuthMiddleware(http.HandlerFunc(Cartcontroller.IncreaseQuantity))))
	http.Handle("/api/cart/decreaseQuantity/", middleware.LoggerMiddleware(middleware.AuthMiddleware(http.HandlerFunc(Cartcontroller.DecreaseQuantity))))
	http.ListenAndServe(":8000", nil)
}
