package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/BachhavPriyanka/BookStore_Project/types"
	_ "github.com/go-sql-driver/mysql"
)

var err error
var db *sql.DB

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
	http.HandleFunc("/api/order/insert", handleInsertion)
	http.HandleFunc("/api/order/retrieveAllOrder", handleRetrieveAllOrders)
	http.HandleFunc("/api/order/retrieveOrder/", handleRetrieveOrderByID)
	http.HandleFunc("/api/order/cancelOrder/", handleCancelOrderByID)
	http.ListenAndServe(":8000", nil)
}

func handleCancelOrderByID(writer http.ResponseWriter, request *http.Request) {
	if request.Method == http.MethodPut {
		writer.Write([]byte("1 to keep Order Active AND 0 to Cancel Order"))

		var data types.Orders
		json.NewDecoder(request.Body).Decode(&data)

		id, err := strconv.Atoi(request.URL.Path[len("/api/order/cancelOrder/"):])
		if err != nil {
			http.Error(writer, "Invalid order ID", http.StatusBadRequest)
			return
		}

		_, err = db.Exec("UPDATE orders SET orderStatus=? WHERE orderId = ?", data.OrderStatus, id)
		if err != nil {
			http.Error(writer, "Some Error Occured", http.StatusInternalServerError)
		}
		writer.Write([]byte("Successfully Updated"))
	}
}

func handleRetrieveOrderByID(writer http.ResponseWriter, request *http.Request) {
	if request.Method == http.MethodGet {
		id, err := strconv.Atoi(request.URL.Path[len("/api/order/retrieveOrder/"):])
		if err != nil {
			http.Error(writer, "Invalid order ID", http.StatusBadRequest)
			return
		}

		rows, err := db.Query("SELECT * FROM orders WHERE orderId = ?", id)
		defer rows.Close()

		orderGet := []types.Orders{}

		for rows.Next() {
			var order types.Orders
			if err := rows.Scan(&order.OrderId, &order.UserId, &order.BookId, &order.Quantity, &order.OrderDate, &order.PriceOfOrder, &order.OrderStatus); err != nil {
				writer.Write([]byte("ORDER-ID NOT FOUND"))
				http.Error(writer, err.Error(), http.StatusInternalServerError)
			}
			orderGet = append(orderGet, order)
		}
		json.NewEncoder(writer).Encode(orderGet)
	}

}

func handleRetrieveAllOrders(writer http.ResponseWriter, request *http.Request) {
	if request.Method == http.MethodGet {
		var order []types.Orders
		rows, err := db.Query("SELECT * FROM orders;")
		if err != nil {
			fmt.Printf("error in quering all orders: %v", err)
			http.Error(writer, err.Error(), http.StatusInternalServerError)
		}
		defer rows.Close()

		for rows.Next() {
			var orderData types.Orders
			if err := rows.Scan(&orderData.OrderId, &orderData.UserId, &orderData.BookId, &orderData.Quantity, &orderData.OrderDate, &orderData.PriceOfOrder, &orderData.OrderStatus); err != nil {
				fmt.Printf("error in query all orders: %v", err)
				http.Error(writer, err.Error(), http.StatusInternalServerError)
			}
			order = append(order, orderData)
		}
		json.NewEncoder(writer).Encode(order)
	}
}

func handleInsertion(writer http.ResponseWriter, request *http.Request) {
	if request.Method == http.MethodPost {
		var orderPost types.Orders
		json.NewDecoder(request.Body).Decode(&orderPost)

		_, err = db.Exec("INSERT INTO orders (orderId, userId, bookId, quantity ,orderDate, price ,orderStatus) VALUES (?,?,?,?,?,?,?)", orderPost.OrderId, orderPost.UserId, orderPost.BookId, orderPost.Quantity, orderPost.OrderDate, orderPost.PriceOfOrder, orderPost.OrderStatus)
		if err != nil {
			fmt.Printf("Error in reading %v", err)
			http.Error(writer, err.Error(), http.StatusInternalServerError)
		}
		writer.Write([]byte("Order Placed Sucessfully..!"))
		fmt.Println("Order Placed Sucessfully..!")
	}
}
