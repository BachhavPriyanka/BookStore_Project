package bookstorecontroller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	repository "github.com/BachhavPriyanka/BookStore_Project/storage"
	"github.com/BachhavPriyanka/BookStore_Project/types"
	_ "github.com/go-sql-driver/mysql"
)

// BookStoreController is a controller for book store operations
type BookStoreController struct {
	Repository *repository.BookStoreRepository
}

// GET book By name
func (c *BookStoreController) HandleBookByName(writer http.ResponseWriter, request *http.Request) {
	bookName := request.URL.Path[len("/getBookByName/"):]

	bookData, _ := c.Repository.HandleBookByName(bookName)
	json.NewEncoder(writer).Encode(bookData)
}

// GetBooks returns a list of all books
func (c *BookStoreController) GetBooks(writer http.ResponseWriter, request *http.Request) {
	books, err := c.Repository.GetBooks()
	if err != nil {
		http.Error(writer, "Error getting books", http.StatusInternalServerError)
		return
	}
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(books)
}

// GetBook returns a book with a specific ID
func (c *BookStoreController) GetBook(writer http.ResponseWriter, request *http.Request) {
	id, err := strconv.Atoi(request.URL.Query().Get("id"))
	fmt.Println(id)
	if err != nil {
		http.Error(writer, "Invalid book ID", http.StatusBadRequest)
		return
	}

	book, err := c.Repository.GetBook(id)
	if err != nil {
		http.Error(writer, "Error getting book", http.StatusInternalServerError)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(book)
}

// AddBook adds a new book to the store
func (c *BookStoreController) AddBook(writer http.ResponseWriter, request *http.Request) {
	book := &types.Books{}
	if err := json.NewDecoder(request.Body).Decode(book); err != nil {
		http.Error(writer, "Error decoding book", http.StatusBadRequest)
		return
	}

	id, err := c.Repository.AddBook(book)
	if err != nil {
		http.Error(writer, "Error adding book", http.StatusInternalServerError)
		return
	}

	writer.WriteHeader(http.StatusCreated)
	writer.Write([]byte(strconv.Itoa(id)))

}

// UpdateBook updates a book with a specific ID
func (c *BookStoreController) UpdateBook(writer http.ResponseWriter, request *http.Request) {
	id, err := strconv.Atoi(request.URL.Query().Get("id"))
	if err != nil {
		http.Error(writer, "id not found", http.StatusUnauthorized)
	}
	book := &types.Books{}
	if err := json.NewDecoder(request.Body).Decode(book); err != nil {
		http.Error(writer, "Error decoding book", http.StatusBadRequest)
		return
	}

	data, err := c.Repository.UpdateBook(id, book)
	if err != nil {
		http.Error(writer, "Error updating book", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(writer).Encode(data)
}

// Delete Method for deleating book
func (c *BookStoreController) Delete(writer http.ResponseWriter, request *http.Request) {
	id, err := strconv.Atoi(request.URL.Query().Get("id"))
	if err != nil {
		http.Error(writer, "id not found", http.StatusUnauthorized)
	}
	deleteStr, err := c.Repository.DeleteBook(id)
	if err != nil {
		http.Error(writer, "Error getting books", http.StatusInternalServerError)
		return
	}
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(deleteStr)
}
