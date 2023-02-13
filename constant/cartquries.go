package constant

const (
	GetDecreaseQuery      = "SELECT * FROM cart WHERE cartId = ?"
	GetDecBookRecordQuery = "SELECT * FROM books WHERE Id = ?"
	GetDecCartQuery       = "UPDATE cart SET quantity = ? WHERE cartId = ?"
	GetDecBookQuery       = "UPDATE books SET bookQuantity = ? WHERE Id = ?"

	GetIncreaseQuery      = "SELECT * FROM cart WHERE cartId = ?"
	GetIncBookRecordQuery = "SELECT * FROM books WHERE Id = ?"
	GetIncCartQuery       = "UPDATE cart SET quantity = ? WHERE cartId = ?"
	GetIncBookQuery       = "UPDATE books SET bookQuantity = ? WHERE Id = ?"
	DeleteCartQuery       = "delete from cart where cartId =?"
	PutUpdateCartQuery    = "update cart set bookName = ?, bookId = ?, cartId = ?, quantity = ? where cartId = ?"
	GetCartDataById       = "select bookName, bookId, cartId, quantity from cart where cartId = ?"
	PostCartQuery         = "insert into cart (BookName,BookID, CartID, Quantity) values (?,?,?,?)"
)
