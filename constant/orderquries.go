package constant

const (
	DeleteOrderQuery  = "Delete from orders WHERE orderId = ?"
	GetOrderQuery     = "SELECT * FROM orders WHERE orderId = ?"
	GetAllOrdersQuery = "SELECT * FROM orders"
	PostInsertQuery   = "INSERT INTO orders (orderId, userId, bookId, quantity ,orderDate, price ,orderStatus) VALUES (?,?,?,?,?,?,?)"
)
