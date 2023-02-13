package repository

import (
	"database/sql"
	"fmt"

	"github.com/BachhavPriyanka/BookStore_Project/constant"
	"github.com/BachhavPriyanka/BookStore_Project/types"
	_ "github.com/go-sql-driver/mysql"
)

type OrderRepository struct {
	DB *sql.DB
}

// Delete Method for canceling order
func (r *OrderRepository) CancelOrderByID(id int) (string, error) {
	_, err := r.DB.Exec(constant.DeleteOrderQuery, id)
	if err != nil {
		return "Id not present", err
	}
	return "Successfully Deleted", err

}

// GET Method to retrieve order by ID
func (r *OrderRepository) RetrieveOrderByID(id int) (*types.Orders, error) {
	var dataStore types.Orders
	fmt.Println(id)
	err := r.DB.QueryRow(constant.GetOrderQuery, id).Scan(&dataStore.OrderId, &dataStore.UserId, &dataStore.BookId, &dataStore.Quantity, &dataStore.OrderDate, &dataStore.PriceOfOrder, &dataStore.OrderStatus)
	if err != nil {
		return nil, err
	}

	return &dataStore, nil
}

// GET Method to retrieve all orders
func (r *OrderRepository) RetrieveAllOrders() (*[]types.Orders, error) {
	rows, err := r.DB.Query(constant.GetAllOrdersQuery)
	if err != nil {
		return nil, err
	}
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

// POST Method to post orders
func (r *OrderRepository) Insertion(orderDetails *types.Orders) (int, error) {
	result, err := r.DB.Exec(constant.PostInsertQuery, orderDetails.OrderId, orderDetails.UserId, orderDetails.BookId, orderDetails.Quantity, orderDetails.OrderDate, orderDetails.PriceOfOrder, orderDetails.OrderStatus)
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
