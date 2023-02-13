package repository

import (
	"database/sql"
	"log"

	"github.com/BachhavPriyanka/BookStore_Project/constant"
	"github.com/BachhavPriyanka/BookStore_Project/types"
	_ "github.com/go-sql-driver/mysql"
)

type CartRepository struct {
	DB *sql.DB
}

// GET Method for decreasing the items in cart
func (r *CartRepository) DecreaseQuantity(cartId int) (string, error) {
	var cart types.Cart
	var book types.Books

	// Find the cart record
	rows, err := r.DB.Query(constant.GetDecreaseQuery, cartId)
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
	// Finding the book record
	rows, err = r.DB.Query(constant.GetDecBookRecordQuery, cart.BookID)
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

		_, err = r.DB.Exec(constant.GetDecCartQuery, cart.Quantity, cart.CartID)
		if err != nil {
			log.Fatalf("Error saving cart: %v", err)
		}

		log.Println("Quantity in cart record updated successfully")

		book.BookQuantity += 1

		// Updating the book record
		_, err = r.DB.Exec(constant.GetDecBookRecordQuery, book.BookQuantity, book.Id)
		if err != nil {
			log.Fatalf("Error saving book: %v", err)
		}
	} else {
		return "Requested quantity is not available", err
	}

	return "Decremented Successfully", err

}

// GET Method for increasing the items in cart
func (r *CartRepository) IncreaseQuantity(cartId int) (string, error) {
	var cart types.Cart
	var book types.Books

	// Finding the cart record
	rows, err := r.DB.Query(constant.GetIncreaseQuery, cartId)
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

	// Finding the book record
	rows, err = r.DB.Query(constant.GetIncBookRecordQuery, cart.BookID)
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

		_, err = r.DB.Exec(constant.GetIncCartQuery, cart.Quantity, cart.CartID)
		if err != nil {
			log.Fatalf("Error saving cart: %v", err)
		}
		log.Println("Quantity in cart record updated successfully")

		book.BookQuantity -= 1

		// Updating the book record
		_, err = r.DB.Exec(constant.GetIncBookQuery, book.BookQuantity, book.Id)
		if err != nil {
			log.Fatalf("Error saving book: %v", err)
		}
	} else {
		return "Requested quantity is not available", err
	}
	return "Incremented Successfully", err
}

// Delete Method to delete the item from cart
func (r *CartRepository) Delete(id int) (string, error) {
	_, err := r.DB.Exec(constant.DeleteCartQuery, id)
	return "Successfully deleted", err
}

// PUT Method for updating data using ID
func (r *CartRepository) UpdateById(id int, dataStore *types.Cart) (string, error) {
	_, err := r.DB.Exec(constant.PutUpdateCartQuery, dataStore.BookName, dataStore.BookID, dataStore.CartID, dataStore.Quantity, id)
	return "Successfully Updated", err
}

// GET Method for getting the data by using ID
func (r *CartRepository) GetById(cartID int) (*[]types.Cart, error) {
	rows, err := r.DB.Query(constant.GetCartDataById, cartID)
	if err != nil {
		return nil, err
	}
	cartDetails := []types.Cart{}
	for rows.Next() {
		var cartDetail types.Cart
		rows.Scan(&cartDetail.BookName, &cartDetail.BookID, &cartDetail.CartID, &cartDetail.Quantity)
		if err != nil {
			return nil, err
		}
		cartDetails = append(cartDetails, cartDetail)
	}
	return &cartDetails, err
}

// POST Method for inserting data
func (r *CartRepository) Create(data *types.Cart) string {

	_, err := r.DB.Exec(constant.PostCartQuery, data.BookName, data.BookID, data.CartID, data.Quantity)
	if err != nil {
		log.Fatalf("Invalid Insertion %v", err)
	}
	response := types.ResponseDTO{
		Message: "Item added to cart successfully",
	}
	return response.Message
}
