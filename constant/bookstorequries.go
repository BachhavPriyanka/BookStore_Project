package constant

const (
	GetBooksQuery   = "SELECT * FROM books"
	GetBookQuery    = "SELECT * FROM books WHERE id = ?"
	AddBookQuery    = "INSERT INTO books (id, title, author, bookQuantity) VALUES (?, ?,?,?)"
	UpdateBookQuery = "UPDATE books SET title = ?, author = ? WHERE id = ?"
	DeleteBookQuery = "DELETE FROM books WHERE id = ?"
)
