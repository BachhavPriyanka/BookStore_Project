package repository

import (
	"database/sql"
	"fmt"

	"github.com/BachhavPriyanka/BookStore_Project/types"
)

type OrderRepository struct {
	DB *sql.DB
}

// Delete Method for canceling order
func (r *OrderRepository) CancelOrderByID(id int) (string, error) {
	_, err := r.DB.Exec("Delete from orders WHERE orderId = ?", id)
	if err != nil {
		return "Id not present", err
	}
	return "Successfully Deleted", err

}

// GET Method to retrieve order by ID
func (r *OrderRepository) RetrieveOrderByID(id int) (*[]types.Orders, error) {

	rows, _ := r.DB.Query("SELECT * FROM orders WHERE orderId = ?", id)
	defer rows.Close()

	orderGet := []types.Orders{}

	for rows.Next() {
		var order types.Orders
		if err := rows.Scan(&order.OrderId, &order.UserId, &order.BookId, &order.Quantity, &order.OrderDate, &order.PriceOfOrder, &order.OrderStatus); err != nil {
			return nil, nil
		}
		orderGet = append(orderGet, order)
	}
	return &orderGet, nil
}

// GET Method to retrieve all orders
func (r *OrderRepository) RetrieveAllOrders() (*types.Orders, error) {
	var orderData types.Orders
	err := r.DB.QueryRow("SELECT * FROM orders").Scan(&orderData.OrderId, &orderData.UserId, &orderData.BookId, &orderData.Quantity, &orderData.OrderDate, &orderData.PriceOfOrder, &orderData.OrderStatus)
	if err != nil {
		return nil, fmt.Errorf("not readable")
	}

	return &orderData, nil
}

// POST Method to post orders
func (r *OrderRepository) Insertion(orderDetails *types.Orders) (int, error) {
	result, err := r.DB.Exec("INSERT INTO orders (orderId, userId, bookId, quantity ,orderDate, price ,orderStatus) VALUES (?,?,?,?,?,?,?)", orderDetails.OrderId, orderDetails.UserId, orderDetails.BookId, orderDetails.Quantity, orderDetails.OrderDate, orderDetails.PriceOfOrder, orderDetails.OrderStatus)
	if err != nil {
		return 0, fmt.Errorf("not readable")
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	fmt.Println("Order Placed Sucessfully..!")

	return int(id), nil
}
