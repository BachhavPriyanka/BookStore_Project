package bookstorecontroller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	repository "github.com/BachhavPriyanka/BookStore_Project/storage"
	"github.com/BachhavPriyanka/BookStore_Project/types"
	tokenutil "github.com/BachhavPriyanka/BookStore_Project/util"
)

type BookStoreController struct {
	Repository *repository.BookStoreRepository
}

// GET Method to get all books
func (c *BookStoreController) GetBooks(w http.ResponseWriter, r http.Request) {

	headerToken := r.Header.Get("Authorization")
	fmt.Println(headerToken)

	userId, err := tokenutil.DecodeToken(headerToken)
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
	}
	fmt.Println("userid int int64", userId)

	books, err := c.Repository.GetBooks()
	if err != nil {
		http.Error(w, "Error getting books", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

// GET Method to get book by id
func (c *BookStoreController) GetBook(w http.ResponseWriter, r *http.Request) {
	headerToken := r.Header.Get("Authorization")
	fmt.Println(headerToken)

	userId, err := tokenutil.DecodeToken(headerToken)
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
	}
	fmt.Println("userid int int64", userId)

	book, err := c.Repository.GetBook(int(userId))
	if err != nil {
		http.Error(w, "Error getting book", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(book)

}

// POST Method to add books
func (c *BookStoreController) AddBook(w http.ResponseWriter, r *http.Request) {

	headerToken := r.Header.Get("Authorization")
	fmt.Println(headerToken)

	userId, err := tokenutil.DecodeToken(headerToken)
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
	}
	fmt.Println("userid int int64", userId)

	book := &types.Books{}
	if err := json.NewDecoder(r.Body).Decode(book); err != nil {
		http.Error(w, "Error decoding book", http.StatusBadRequest)
		return
	}

	id, err := c.Repository.AddBook(book)
	if err != nil {
		http.Error(w, "Error adding book", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(strconv.Itoa(id)))
}

// PUT Method for updating book
func (c *BookStoreController) UpdateBook(w http.ResponseWriter, r *http.Request) {

	headerToken := r.Header.Get("Authorization")
	fmt.Println(headerToken)

	userId, err := tokenutil.DecodeToken(headerToken)
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	fmt.Println("userid int int64", userId)

	book := &types.Books{}
	if err := json.NewDecoder(r.Body).Decode(book); err != nil {
		http.Error(w, "Error decoding book", http.StatusBadRequest)
		return
	}
	if err := c.Repository.UpdateBook(userId, book); err != nil {
		http.Error(w, "Error updating book", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)

}
