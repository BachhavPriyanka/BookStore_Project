package repository

import (
	"database/sql"

	"github.com/BachhavPriyanka/BookStore_Project/constant"
	"github.com/BachhavPriyanka/BookStore_Project/types"
	_ "github.com/go-sql-driver/mysql"
)

type BookStoreRepository struct {
	DB *sql.DB
}

func (r *BookStoreRepository) GetBooks() ([]*types.Books, error) {
	rows, err := r.DB.Query(constant.GetBooksQuery)
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
	row := r.DB.QueryRow(constant.GetBookQuery, id)

	book := &types.Books{}
	if err := row.Scan(&book.Id, &book.Title, &book.Author, &book.BookQuantity); err != nil {
		return nil, err
	}

	return book, nil
}

func (r *BookStoreRepository) AddBook(book *types.Books) (int, error) {
	result, err := r.DB.Exec(constant.AddBookQuery, book.Id, book.Title, book.Author, book.BookQuantity)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (r *BookStoreRepository) UpdateBook(id int, book *types.Books) (string, error) {
	_, err := r.DB.Exec(constant.UpdateBookQuery, book.Title, book.Author, id)
	return "Successfully Updated", err
}

func (r *BookStoreRepository) DeleteBook(id int) (string, error) {
	_, err := r.DB.Exec(constant.DeleteBookQuery, id)
	return "Successfully Deleted", err
}

func (r *BookStoreRepository) HandleBookByName(bookName string) (*[]types.Books, error) {
	rows, err := r.DB.Query("select Id, Title, Author , bookQuantity from books where Title = ?", bookName)
	if err != nil {
		return nil, err
	}

	bookDetails := []types.Books{}
	for rows.Next() {
		var bookDetail types.Books
		rows.Scan(&bookDetail.Id, &bookDetail.Title, &bookDetail.Author, &bookDetail.BookQuantity)
		if err != nil {
			return nil, err
		}
		bookDetails = append(bookDetails, bookDetail)
	}
	return &bookDetails, err
}
