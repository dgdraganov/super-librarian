package repo

import (
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
		config.Database,
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
