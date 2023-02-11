package repository

import (
	"database/sql"

	"github.com/BachhavPriyanka/BookStore_Project/types"
	_ "github.com/go-sql-driver/mysql"
)

type BookStoreRepository struct {
	DB *sql.DB
}

func (r *BookStoreRepository) GetBooks() ([]*types.Books, error) {
	rows, err := r.DB.Query("SELECT * FROM books")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var books []*types.Books
	for rows.Next() {
		book := &types.Books{}
		if err := rows.Scan(&book.Id, &book.Title, &book.Author, &book.BookQuantity); err != nil {
			return nil, err
		}
		books = append(books, book)
	}
	return books, nil
}

func (r *BookStoreRepository) GetBook(id int) (*types.Books, error) {
	row := r.DB.QueryRow("SELECT * FROM books WHERE id = ?", id)

	book := &types.Books{}
	if err := row.Scan(&book.Id, &book.Title, &book.Author, &book.BookQuantity); err != nil {
		return nil, err
	}

	return book, nil
}

func (r *BookStoreRepository) AddBook(book *types.Books) (int, error) {
	result, err := r.DB.Exec("INSERT INTO books (title, author) VALUES (?, ?)", book.Title, book.Author)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (r *BookStoreRepository) UpdateBook(id int, book *types.Books) error {
	_, err := r.DB.Exec("UPDATE books SET title = ?, author = ? WHERE id = ?", book.Title, book.Author, id)
	return err
}

func (r *BookStoreRepository) DeleteBook(id int) error {
	_, err := r.DB.Exec("DELETE FROM books WHERE id = ?", id)
	return err
}
