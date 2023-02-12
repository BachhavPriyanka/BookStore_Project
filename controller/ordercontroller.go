package bookstorecontroller

import (
	"encoding/json"
	"fmt"
	"net/http"

	repository "github.com/BachhavPriyanka/BookStore_Project/storage"
	"github.com/BachhavPriyanka/BookStore_Project/types"
	tokenutil "github.com/BachhavPriyanka/BookStore_Project/util"
	_ "github.com/go-sql-driver/mysql"
)

type OrderController struct {
	Repository *repository.OrderRepository
}

// Delete Method
func (c *OrderController) CancelOrderByID(writer http.ResponseWriter, request *http.Request) {

	headerToken := request.Header.Get("Authorization")
	fmt.Println(headerToken)

	userId, err := tokenutil.DecodeToken(headerToken)
	if err != nil {
		http.Error(writer, "Unauthorized", http.StatusUnauthorized)
		return
	}
	fmt.Println("userid int int64", userId)

	orderCancel, err := c.Repository.CancelOrderByID(int(userId))
	if err != nil {
		http.Error(writer, "Error getting order", http.StatusInternalServerError)
		return
	}
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(orderCancel)

}

// Get Method to retrieve orders by id
func (c *OrderController) RetrieveOrderByID(writer http.ResponseWriter, request *http.Request) {

	headerToken := request.Header.Get("Authorization")
	fmt.Println(headerToken)

	userId, err := tokenutil.DecodeToken(headerToken)
	if err != nil {
		http.Error(writer, "Unauthorized", http.StatusUnauthorized)
		return
	}
	fmt.Println("userid int int64", userId)

	dataRetrieve, err := c.Repository.RetrieveOrderByID(int(userId))
	if err != nil {
		http.Error(writer, "Invalid order ID", http.StatusBadRequest)
		return
	}
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(dataRetrieve)
}

// Get Method to retrieve all orders
func (c *OrderController) RetrieveAllOrders(writer http.ResponseWriter, request *http.Request) {

	headerToken := request.Header.Get("Authorization")
	fmt.Println(headerToken)

	userId, err := tokenutil.DecodeToken(headerToken)
	if err != nil {
		http.Error(writer, "Unauthorized", http.StatusUnauthorized)
		return
	}
	fmt.Println("userid int int64", userId)

	allOrders, err := c.Repository.RetrieveAllOrders()
	if err != nil {
		http.Error(writer, "Invalid", http.StatusBadRequest)
		return
	}
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(allOrders)
}

// POST Method
func (c *OrderController) Insertion(writer http.ResponseWriter, request *http.Request) {
	headerToken := request.Header.Get("Authorization")
	fmt.Println(headerToken)

	userId, err := tokenutil.DecodeToken(headerToken)
	if err != nil {
		http.Error(writer, "Unauthorized", http.StatusUnauthorized)
		return
	}
	fmt.Println("userid int int64", userId)

	var orderData types.Orders
	json.NewDecoder(request.Body).Decode(&orderData)
	orderInsertion, err := c.Repository.Insertion(&orderData)
	if err != nil {
		http.Error(writer, "Invalid", http.StatusBadRequest)
		return
	}
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(orderInsertion)

}
