package bookstorecontroller

import (
	"encoding/json"
	"net/http"
	"strconv"

	repository "github.com/BachhavPriyanka/BookStore_Project/storage"
	"github.com/BachhavPriyanka/BookStore_Project/types"
	_ "github.com/go-sql-driver/mysql"
)

type OrderController struct {
	Repository *repository.OrderRepository
}

// Delete Method for deleting order
func (c *OrderController) CancelOrderByID(writer http.ResponseWriter, request *http.Request) {
	id, err := strconv.Atoi(request.URL.Query().Get("id"))
	if err != nil {
		http.Error(writer, "id not found", http.StatusUnauthorized)
	}
	orderCancel, err := c.Repository.CancelOrderByID(int(id))
	if err != nil {
		http.Error(writer, "Error getting order", http.StatusInternalServerError)
		return
	}
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(orderCancel)

}

// Get Method to retrieve orders by id
func (c *OrderController) RetrieveOrderByID(writer http.ResponseWriter, request *http.Request) {
	id, err := strconv.Atoi(request.URL.Query().Get("id"))
	if err != nil {
		http.Error(writer, "id not found", http.StatusUnauthorized)
	}
	dataRetrieve, err := c.Repository.RetrieveOrderByID(id)
	if err != nil {
		http.Error(writer, "Invalid order ID", http.StatusBadRequest)
		return
	}
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(dataRetrieve)
}

// Get Method to retrieve all orders
func (c *OrderController) RetrieveAllOrders(writer http.ResponseWriter, request *http.Request) {

	allOrders, err := c.Repository.RetrieveAllOrders()
	if err != nil {
		http.Error(writer, "Invalid", http.StatusBadRequest)
		return
	}
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(allOrders)
}

// POST Method for insertion
func (c *OrderController) Insertion(writer http.ResponseWriter, request *http.Request) {
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
