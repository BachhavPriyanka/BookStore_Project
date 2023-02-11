package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/BachhavPriyanka/BookStore_Project/types"
	tokenutil "github.com/BachhavPriyanka/BookStore_Project/util"

	_ "github.com/go-sql-driver/mysql"
)

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
	http.HandleFunc("/getBooks", handleBooks)
	http.HandleFunc("/getBookByName/", handleBookByName)
	http.HandleFunc("/getBook/", handleBookById)
	http.HandleFunc("/addBook", handleAddBook)
	http.HandleFunc("/update/", handleUpdate)
	http.HandleFunc("/delete/", handleDelete)
	http.ListenAndServe(":8000", nil)
}

func handleBooks(w http.ResponseWriter, r *http.Request) {
	// Geting the token from the request header
	headerToken := r.Header.Get("Authorization")

	// Verifying the token
	userID, err := tokenutil.DecodeToken(headerToken)
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	fmt.Println(userID)
	if r.Method == http.MethodGet {
		rows, err := db.Query("SELECT * FROM books")
		if err != nil {
			http.Error(w, "Not found", http.StatusInternalServerError)
			return
		}

		bookDetails := []types.Books{}

		for rows.Next() {
			var book types.Books
			err = rows.Scan(&book.Id, &book.Title, &book.Author, &book.BookQuantity)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			bookDetails = append(bookDetails, book)
		}
		json.NewEncoder(w).Encode(bookDetails)
	}
}
func handleDelete(writer http.ResponseWriter, request *http.Request) {
	if request.Method == http.MethodDelete {
		id, err := strconv.Atoi(request.URL.Path[len("/delete/"):])
		if err != nil {
			http.Error(writer, "Id is not present", http.StatusBadRequest)
		}
		_, err = db.Exec("delete from books where id =?", id)
		if err != nil {
			http.Error(writer, "Id is not present", http.StatusInternalServerError)
		}
		writer.Write([]byte("Record Deleted"))
	}
}

func handleUpdate(writer http.ResponseWriter, request *http.Request) {
	if request.Method == http.MethodPut {
		var dataStore types.Books
		json.NewDecoder(request.Body).Decode(&dataStore)

		id, err := strconv.Atoi(request.URL.Path[len("/update/"):])
		if err != nil {
			http.Error(writer, "Error", http.StatusBadRequest)
			return
		}
		_, err = db.Exec("update books set Id = ?, Title = ?, Author=? , bookQuantity = ? where Id = ?", dataStore.Id, dataStore.Title, dataStore.Author, dataStore.BookQuantity, id)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
		}
		writer.Write([]byte("Record Updated"))
	}
}

func handleAddBook(writer http.ResponseWriter, request *http.Request) {
	if request.Method == http.MethodPost {
		var dataStore types.Books
		json.NewDecoder(request.Body).Decode(&dataStore)

		_, err := db.Exec("Insert into books (Id,Title,Author,bookQuantity) values (?, ?, ?, ?)", dataStore.Id, dataStore.Title, dataStore.Author, dataStore.BookQuantity)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
		}
		writer.Write([]byte("Record Added"))
	}
}

func handleBookById(writer http.ResponseWriter, request *http.Request) {
	if request.Method == http.MethodGet {
		bookID, err := strconv.Atoi(request.URL.Path[len("/getBook/"):])
		if err != nil {
			http.Error(writer, "Error", http.StatusBadRequest)
			return
		}

		rows, err := db.Query("select Id, Title, Author, bookQuantity from books where Id = ?", bookID)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
		}

		bookDetails := []types.Books{}
		for rows.Next() {
			var bookDetail types.Books
			rows.Scan(&bookDetail.Id, &bookDetail.Title, &bookDetail.Author, &bookDetail.BookQuantity)
			if err != nil {
				http.Error(writer, err.Error(), http.StatusInternalServerError)
			}
			bookDetails = append(bookDetails, bookDetail)
		}
		json.NewEncoder(writer).Encode(bookDetails)
	}
}

func handleBookByName(writer http.ResponseWriter, request *http.Request) {
	if request.Method == http.MethodGet {
		bookName := request.URL.Path[len("/getBookByName/"):]

		rows, err := db.Query("select Id, Title, Author , bookQuantity from books where Title = ?", bookName)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
		}

		bookDetails := []types.Books{}
		for rows.Next() {
			var bookDetail types.Books
			rows.Scan(&bookDetail.Id, &bookDetail.Title, &bookDetail.Author, &bookDetail.BookQuantity)
			if err != nil {
				http.Error(writer, err.Error(), http.StatusInternalServerError)
			}
			bookDetails = append(bookDetails, bookDetail)
		}
		json.NewEncoder(writer).Encode(bookDetails)
	}
}

func handleAllBooks(writer http.ResponseWriter, request *http.Request) {
	if request.Method == http.MethodGet {

		rows, err := db.Query("select Id, Title, Author, bookQuantity from books")
		if err != nil {
			http.Error(writer, "Not found", http.StatusInternalServerError)
			return
		}

		bookDetails := []types.Books{}

		for rows.Next() {
			var book types.Books
			err = rows.Scan(&book.Id, &book.Title, &book.Author, &book.BookQuantity)
			if err != nil {
				http.Error(writer, err.Error(), http.StatusInternalServerError)
			}
			bookDetails = append(bookDetails, book)
		}
		json.NewEncoder(writer).Encode(bookDetails)

	}
}
