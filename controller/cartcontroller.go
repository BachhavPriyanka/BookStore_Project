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

type CartController struct {
	Repository *repository.CartRepository
}

// GET Method to decrease Quantity of item
func (c *CartController) DecreaseQuantity(writer http.ResponseWriter, request *http.Request) {

	headerToken := request.Header.Get("Authorization")
	fmt.Println(headerToken)

	userId, err := tokenutil.DecodeToken(headerToken)
	if err != nil {
		http.Error(writer, "Unauthorized", http.StatusUnauthorized)
	}
	fmt.Println("userid int int64", userId)

	decreaseItem, err := c.Repository.DecreaseQuantity(int(userId))
	if err != nil {
		http.Error(writer, "Error scanning book", http.StatusInternalServerError)
		return
	}
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(decreaseItem)
}

func (c *CartController) IncreaseQuantity(writer http.ResponseWriter, request *http.Request) {

	headerToken := request.Header.Get("Authorization")
	fmt.Println(headerToken)

	userId, err := tokenutil.DecodeToken(headerToken)
	if err != nil {
		http.Error(writer, "Unauthorized", http.StatusUnauthorized)
	}
	fmt.Println("userid int int64", userId)

	increaseItem, err := c.Repository.IncreaseQuantity(int(userId))
	if err != nil {
		http.Error(writer, "Error", http.StatusInternalServerError)
		return
	}
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(increaseItem)
}

func (c *CartController) Delete(writer http.ResponseWriter, request *http.Request) {

	headerToken := request.Header.Get("Authorization")
	fmt.Println(headerToken)

	userId, err := tokenutil.DecodeToken(headerToken)
	if err != nil {
		http.Error(writer, "Unauthorized", http.StatusUnauthorized)
	}
	fmt.Println("userid int int64", userId)

	deletedId, err := c.Repository.Delete(int(userId))
	if err != nil {
		http.Error(writer, "Error", http.StatusInternalServerError)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(deletedId)
}

func (c *CartController) UpdateById(writer http.ResponseWriter, request *http.Request) {
	headerToken := request.Header.Get("Authorization")
	fmt.Println(headerToken)

	userId, err := tokenutil.DecodeToken(headerToken)
	if err != nil {
		http.Error(writer, "Unauthorized", http.StatusUnauthorized)
	}
	fmt.Println("userid int int64", userId)

	var dataStore types.Cart
	json.NewDecoder(request.Body).Decode(&dataStore)

	id := int(userId)
	updateData, err := c.Repository.UpdateById(id, &dataStore)
	if err != nil {
		http.Error(writer, "Id is not present", http.StatusBadRequest)
	}
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(updateData)

}

func (c *CartController) GetById(writer http.ResponseWriter, request *http.Request) {
	headerToken := request.Header.Get("Authorization")
	fmt.Println(headerToken)

	userId, err := tokenutil.DecodeToken(headerToken)
	if err != nil {
		http.Error(writer, "Unauthorized", http.StatusUnauthorized)
		return
	}
	fmt.Println("userid int int64", userId)

	data, err := c.Repository.GetById(int(userId))
	if err != nil {
		http.Error(writer, "Data retrieve error", http.StatusBadRequest)
	}
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(data)
}

func (c *CartController) Create(writer http.ResponseWriter, request *http.Request) {

	headerToken := request.Header.Get("Authorization")
	fmt.Println(headerToken)

	userId, err := tokenutil.DecodeToken(headerToken)
	if err != nil {
		http.Error(writer, "Unauthorized", http.StatusUnauthorized)
		return
	}
	fmt.Println("userid int int64", userId)

	var data types.Cart
	json.NewDecoder(request.Body).Decode(&data)

	xyz := c.Repository.Create(&data)
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(xyz)

}
