package librarianapi

import (
	"context"

	"github.com/dgdraganov/super-librarian/internal/model"
)

type BooksRepo interface {
	Ping(ctx context.Context) error
	GetBook(ctx context.Context, id int) (model.Book, error)
	GetBooks(ctx context.Context, skip, take int) ([]model.Book, error)
	DeleteBook(ctx context.Context, id int) error
	CreateBook(ctx context.Context, book model.Book) (model.Book, error)
	UpdateBook(ctx context.Context, book model.Book) (model.Book, error)
}
