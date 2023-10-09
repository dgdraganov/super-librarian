package repo

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/dgdraganov/super-librarian/internal/model"

	_ "github.com/go-sql-driver/mysql"
)

// libraryRepository represents a mysql books repository
type libraryRepository struct {
	db *sql.DB
}

// NewLibraryRepository is a constructor function for the libraryRepository type
func NewLibraryRepository(config *model.MySqlConfig) *libraryRepository {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s",
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
func (repo *libraryRepository) Ping() error {
	err := repo.db.Ping()
	if err != nil {
		return fmt.Errorf("db ping: %w", err)
	}
	return nil
}

// GetBook retrieves a single book from the db
func (repo *libraryRepository) GetBook(ctx context.Context, id int) (model.Book, error) {

	stmt, err := repo.db.Prepare(`select id, title, author, bookCover, publishedAt 
								  from Books
								  where id = ?`)
	if err != nil {
		return model.Book{}, fmt.Errorf("db prepare stmt: %w", err)
	}
	defer stmt.Close()
	row := stmt.QueryRowContext(ctx, id)
	book := model.Book{}
	if err := row.Scan(&book); err != nil {
		return model.Book{}, fmt.Errorf("row scan: %w", err)
	}

	return book, nil
}

// GetBooks retrieves paginated books
func (repo *libraryRepository) GetBooks(skip, take int) ([]model.Book, error) {
	return []model.Book{}, nil
}
