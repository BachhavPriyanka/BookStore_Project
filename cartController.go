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

type ResponseDTO struct {
	Message string `json:"message"`
}

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
	http.HandleFunc("/api/cart/create", handleCreate)
	http.HandleFunc("/api/cart/getById/", handleGetById)
	http.HandleFunc("/api/cart/updateById/", handleUpdateById)
	http.HandleFunc("/api/cart/delete/", handleDelete)
	http.HandleFunc("/api/cart/increaseQuantity/", handleIncreaseQuantity)
	http.HandleFunc("/api/cart/decreaseQuantity/", handleDecreaseQuantity)
	http.ListenAndServe(":8000", nil)

}

func handleDecreaseQuantity(writer http.ResponseWriter, request *http.Request) {
	if request.Method == http.MethodGet {
		var cart types.Cart
		var book types.Books

		cartId, err := strconv.Atoi(request.URL.Path[len("/api/cart/decreaseQuantity/"):])
		if err != nil {
			http.Error(writer, "Id is not present", http.StatusBadRequest)
		}
		// Find the cart record
		rows, err := db.Query("SELECT * FROM cart WHERE cartId = ?", cartId)
		if err != nil {
			log.Fatalf("Error finding cart: %v", err)
		}
		defer rows.Close()

		for rows.Next() {
			err = rows.Scan(&cart.BookName, &cart.BookID, &cart.CartID, &cart.Quantity)
			if err != nil {
				log.Fatalf("Error scanning cart: %v", err)
			}
		}

		// Find the book record
		rows, err = db.Query("SELECT * FROM books WHERE Id = ?", cart.BookID)
		if err != nil {
			log.Fatalf("Error finding book: %v", err)
		}
		defer rows.Close()

		for rows.Next() {
			err = rows.Scan(&book.Id, &book.Title, &book.Author, &book.BookQuantity)
			if err != nil {
				log.Fatalf("Error scanning book: %v", err)
			}
		}

		if book.BookQuantity >= 1 {
			cart.Quantity -= 1

			_, err = db.Exec("UPDATE cart SET quantity = ? WHERE cartId = ?", cart.Quantity, cart.CartID)
			if err != nil {
				log.Fatalf("Error saving cart: %v", err)
			}

			log.Println("Quantity in cart record updated successfully")

			book.BookQuantity += 1

			// Update the book record
			_, err = db.Exec("UPDATE books SET bookQuantity = ? WHERE Id = ?", book.BookQuantity, book.Id)
			if err != nil {
				log.Fatalf("Error saving book: %v", err)
			}
		} else {
			log.Fatalf("Requested quantity is not available")
		}
	}
}

func handleIncreaseQuantity(writer http.ResponseWriter, request *http.Request) {
	if request.Method == http.MethodGet {
		var cart types.Cart
		var book types.Books

		cartId, err := strconv.Atoi(request.URL.Path[len("/api/cart/increaseQuantity/"):])
		if err != nil {
			http.Error(writer, "Id is not present", http.StatusBadRequest)
		}

		// Find the cart record
		rows, err := db.Query("SELECT * FROM cart WHERE cartId = ?", cartId)
		if err != nil {
			log.Fatalf("Error finding cart: %v", err)
		}
		defer rows.Close()

		for rows.Next() {
			err = rows.Scan(&cart.BookName, &cart.BookID, &cart.CartID, &cart.Quantity)
			if err != nil {
				log.Fatalf("Error scanning cart: %v", err)
			}
		}

		// Find the book record
		rows, err = db.Query("SELECT * FROM books WHERE Id = ?", cart.BookID)
		if err != nil {
			log.Fatalf("Error finding book: %v", err)
		}
		defer rows.Close()

		for rows.Next() {
			err = rows.Scan(&book.Id, &book.Title, &book.Author, &book.BookQuantity)
			if err != nil {
				log.Fatalf("Error scanning book: %v", err)
			}
		}

		if book.BookQuantity >= 1 {
			cart.Quantity += 1

			_, err = db.Exec("UPDATE cart SET quantity = ? WHERE cartId = ?", cart.Quantity, cart.CartID)
			if err != nil {
				log.Fatalf("Error saving cart: %v", err)
			}
			log.Println("Quantity in cart record updated successfully")

			book.BookQuantity -= 1

			// Update the book record
			_, err = db.Exec("UPDATE books SET bookQuantity = ? WHERE Id = ?", book.BookQuantity, book.Id)
			if err != nil {
				log.Fatalf("Error saving book: %v", err)
			}
		} else {
			log.Fatalf("Requested quantity is not available")
		}
	}
}

func handleDelete(writer http.ResponseWriter, request *http.Request) {
	if request.Method == http.MethodDelete {
		id, err := strconv.Atoi(request.URL.Path[len("/api/cart/delete/"):])
		if err != nil {
			http.Error(writer, "Id is not present", http.StatusBadRequest)
		}
		_, err = db.Exec("delete from cart where cartId =?", id)
		if err != nil {
			http.Error(writer, "Id is not present", http.StatusInternalServerError)
		}
		writer.Write([]byte("Record Deleted"))
	}
}

func handleUpdateById(writer http.ResponseWriter, request *http.Request) {
	if request.Method == http.MethodPut {
		var dataStore types.Cart
		json.NewDecoder(request.Body).Decode(&dataStore)

		id, err := strconv.Atoi(request.URL.Path[len("/api/cart/updateById/"):])
		if err != nil {
			http.Error(writer, "Error", http.StatusBadRequest)
			return
		}
		_, err = db.Exec("update cart set bookName = ?, bookId = ?, cartId = ?, quantity = ? where cartId = ?", dataStore.BookName, dataStore.BookID, dataStore.CartID, dataStore.Quantity, id)
		if err != nil {
			http.Error(writer, "sql query error", http.StatusInternalServerError)
		}
		writer.Write([]byte("Record Updated"))
	}
}

func handleGetById(writer http.ResponseWriter, request *http.Request) {
	if request.Method == http.MethodGet {
		cartID, err := strconv.Atoi(request.URL.Path[len("/api/cart/getById/"):])
		if err != nil {
			http.Error(writer, "Error", http.StatusBadRequest)
			return
		}

		rows, err := db.Query("select bookName, bookId, cartId, quantity from cart where cartId = ?", cartID)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
		}

		cartDetails := []types.Cart{}
		for rows.Next() {
			var cartDetail types.Cart
			rows.Scan(&cartDetail.BookName, &cartDetail.BookID, &cartDetail.CartID, &cartDetail.Quantity)
			if err != nil {
				http.Error(writer, err.Error(), http.StatusInternalServerError)
			}
			cartDetails = append(cartDetails, cartDetail)
		}
		json.NewEncoder(writer).Encode(cartDetails)
	}
}

func handleCreate(writer http.ResponseWriter, request *http.Request) {
	if request.Method == http.MethodPost {
		var data types.Cart
		json.NewDecoder(request.Body).Decode(&data)

		_, err := db.Exec("insert into cart (BookName,BookID, CartID, Quantity) values (?,?,?,?)", data.BookName, data.BookID, data.CartID, data.Quantity)
		if err != nil {
			http.Error(writer, "Invalid Insertion", http.StatusInternalServerError)
		}
		response := ResponseDTO{
			Message: "Item added to cart successfully",
		}
		json.NewEncoder(writer).Encode(response)
	}
}
