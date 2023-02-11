package bookstorecontroller

import (
	"encoding/json"
	"net/http"
	"strconv"

	repository "github.com/BachhavPriyanka/BookStore_Project/storage"
	"github.com/BachhavPriyanka/BookStore_Project/types"
)

// type BookStoreRepository struct {
// 	DB *sql.DB
// }

type BookStoreController struct {
	Repository *repository.BookStoreRepository
}

func (c *BookStoreController) GetBooks(w http.ResponseWriter, r http.Request) {
	books, err := c.Repository.GetBooks()
	if err != nil {
		http.Error(w, "Error getting books", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

func (c *BookStoreController) GetBook(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		http.Error(w, "Invalid book ID", http.StatusBadRequest)
		return
	}
	book, err := c.Repository.GetBook(id)
	if err != nil {
		http.Error(w, "Error getting book", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(book)

}

func (c *BookStoreController) AddBook(w http.ResponseWriter, r *http.Request) {
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

func (c *BookStoreController) UpdateBook(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		http.Error(w, "Invalid book ID", http.StatusBadRequest)
		return
	}
	book := &types.Books{}
	if err := json.NewDecoder(r.Body).Decode(book); err != nil {
		http.Error(w, "Error decoding book", http.StatusBadRequest)
		return
	}
	if err := c.Repository.UpdateBook(id, book); err != nil {
		http.Error(w, "Error updating book", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)

}
