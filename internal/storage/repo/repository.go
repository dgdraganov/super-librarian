package repo

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/dgdraganov/super-librarian/internal/model"

	_ "github.com/go-sql-driver/mysql"
)

var ErrNotFound error = errors.New("book not found")

// libraryRepository represents a mysql books repository
type libraryRepository struct {
	db *sql.DB
}

// NewLibraryRepository is a constructor function for the libraryRepository type
func NewLibraryRepository(config model.MySqlConfig) *libraryRepository {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?parseTime=true",
		config.Username,
		config.Password,
		config.Host,
		config.Port,
		config.DbName,
	)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(fmt.Sprintf("cannot create connection to database (%s:%s)", config.Host, config.Port))
	}
	return &libraryRepository{
		db: db,
	}
}

// Ping the database
func (repo *libraryRepository) Ping(ctx context.Context) error {
	err := repo.db.Ping()
	if err != nil {
		return fmt.Errorf("db ping: %w", err)
	}

	return nil
}

// GetBook retrieves a single book from the db
func (repo *libraryRepository) GetBook(ctx context.Context, id int) (model.Book, error) {

	stmt, err := repo.db.Prepare(`SELECT id, author, title, bookCover, publishedAt 
								  FROM Books
								  WHERE id = ?
								  LIMIT 1`)
	if err != nil {
		return model.Book{}, fmt.Errorf("db prepare statement: %w", err)
	}
	defer stmt.Close()
	row := stmt.QueryRowContext(ctx, id)

	book := model.Book{}
	err = row.Scan(&book.Id, &book.Author, &book.Title, &book.BookCover, &book.PublishedAt)

	if err == sql.ErrNoRows {
		return model.Book{}, ErrNotFound
	}
	if err != nil {
		return model.Book{}, fmt.Errorf("row scan: %w", err)
	}

	return book, nil
}

// CreateBook adds a new book to the db
func (repo *libraryRepository) CreateBook(ctx context.Context, book model.Book) (model.Book, error) {
	stmt, err := repo.db.Prepare("INSERT INTO Books(title, author, bookCover, publishedAt) VALUES(?, ?, ?, ?)")
	if err != nil {
		return model.Book{}, fmt.Errorf("db prepare statement: %w", err)
	}
	defer stmt.Close()

	res, err := stmt.Exec(book.Title, book.Author, book.BookCover, book.PublishedAt)
	if err != nil {
		return model.Book{}, fmt.Errorf("execute statement: %w", err)
	}

	lastId, err := res.LastInsertId()
	if err != nil {
		return model.Book{}, fmt.Errorf("get last inserted id: %w", err)
	}

	return model.Book{
		Id:          int(lastId),
		Title:       book.Title,
		Author:      book.Author,
		BookCover:   book.BookCover,
		PublishedAt: book.PublishedAt,
	}, nil
}

// GetBooks retrieves paginated books
func (repo *libraryRepository) GetBooks(ctx context.Context, skip, take int) ([]model.Book, error) {

	stmt, err := repo.db.Prepare(`SELECT id, author, title, bookCover, publishedAt 
								  FROM Books
								  ORDER BY id
								  LIMIT ?
								  OFFSET ?`)
	if err != nil {
		return nil, fmt.Errorf("db prepare statement: %w", err)
	}
	defer stmt.Close()

	rows, err := stmt.QueryContext(ctx, take, skip)
	if err != nil {
		return nil, fmt.Errorf("execute query: %w", err)
	}

	defer rows.Close()

	result := make([]model.Book, 0, take)
	for rows.Next() {
		book := model.Book{}

		err := rows.Scan(&book.Id, &book.Author, &book.Title, &book.BookCover, &book.PublishedAt)
		if err != nil {
			return nil, fmt.Errorf("rows scan: %w", err)
		}
		result = append(result, book)
	}

	return result, nil
}

func (repo *libraryRepository) DeleteBook(ctx context.Context, id int) error {
	stmt, err := repo.db.Prepare(`DELETE FROM Books WHERE id = ?;`)
	if err != nil {
		return fmt.Errorf("db prepare statement: %w", err)
	}

	_, err = stmt.Exec(id)
	if err != nil {
		return fmt.Errorf("execute statement: %w", err)
	}
	return nil
}

func (repo *libraryRepository) UpdateBook(ctx context.Context, book model.Book) (model.Book, error) {
	stmt, err := repo.db.Prepare(`UPDATE Books 
								  SET title = ?, author = ?, publishedAt = ?, bookCover = ?
								  WHERE id = ?;`)
	if err != nil {
		return model.Book{}, fmt.Errorf("db prepare statement: %w", err)
	}

	defer stmt.Close()
	_, err = stmt.Exec(book.Title, book.Author, book.PublishedAt, book.BookCover, book.Id)
	if err != nil {
		return model.Book{}, fmt.Errorf("execute statement: %w", err)
	}
	return book, nil
}
