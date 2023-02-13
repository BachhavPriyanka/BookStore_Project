package bookstorecontroller

import (
	"encoding/json"
	"net/http"
	"strconv"

	repository "github.com/BachhavPriyanka/BookStore_Project/storage"
	"github.com/BachhavPriyanka/BookStore_Project/types"

	_ "github.com/go-sql-driver/mysql"
)

type CartController struct {
	Repository *repository.CartRepository
}

// GET Method to decrease Quantity of item
func (c *CartController) DecreaseQuantity(writer http.ResponseWriter, request *http.Request) {
	id, err := strconv.Atoi(request.URL.Query().Get("id"))
	if err != nil {
		http.Error(writer, "id not found", http.StatusUnauthorized)
	}

	decreaseItem, err := c.Repository.DecreaseQuantity(int(id))
	if err != nil {
		http.Error(writer, "Error scanning book", http.StatusInternalServerError)
		return
	}
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(decreaseItem)
}

// GET Method to increase Quantity of item
func (c *CartController) IncreaseQuantity(writer http.ResponseWriter, request *http.Request) {
	id, err := strconv.Atoi(request.URL.Query().Get("id"))
	if err != nil {
		http.Error(writer, "id not found", http.StatusUnauthorized)
	}

	increaseItem, err := c.Repository.IncreaseQuantity(int(id))
	if err != nil {
		http.Error(writer, "Error", http.StatusInternalServerError)
		return
	}
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(increaseItem)
}

// Delete Method to delete the cart
func (c *CartController) Delete(writer http.ResponseWriter, request *http.Request) {
	id, err := strconv.Atoi(request.URL.Query().Get("id"))
	if err != nil {
		http.Error(writer, "id not found", http.StatusUnauthorized)
	}

	deletedId, err := c.Repository.Delete(int(id))
	if err != nil {
		http.Error(writer, "Error", http.StatusInternalServerError)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(deletedId)
}

// PUT Method to Update the record
func (c *CartController) UpdateById(writer http.ResponseWriter, request *http.Request) {
	id, err := strconv.Atoi(request.URL.Query().Get("id"))
	if err != nil {
		http.Error(writer, "id not found", http.StatusUnauthorized)
	}

	var dataStore types.Cart
	json.NewDecoder(request.Body).Decode(&dataStore)

	updateData, err := c.Repository.UpdateById(id, &dataStore)
	if err != nil {
		http.Error(writer, "Id is not present", http.StatusBadRequest)
	}
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(updateData)
}

// GET Method to get the record
func (c *CartController) GetById(writer http.ResponseWriter, request *http.Request) {
	id, err := strconv.Atoi(request.URL.Query().Get("id"))
	if err != nil {
		http.Error(writer, "id not found", http.StatusUnauthorized)
	}

	data, err := c.Repository.GetById(int(id))
	if err != nil {
		http.Error(writer, "Data retrieve error", http.StatusBadRequest)
	}
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(data)
}

// POST Method to add the record
func (c *CartController) Create(writer http.ResponseWriter, request *http.Request) {
	var data types.Cart
	json.NewDecoder(request.Body).Decode(&data)

	newData := c.Repository.Create(&data)
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(newData)

}
