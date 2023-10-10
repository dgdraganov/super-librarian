package librarianapi

import (
	"context"

	"github.com/dgdraganov/super-librarian/internal/model"
)

type BooksRepo interface {
	Ping(ctx context.Context) error
	GetBook(ctx context.Context, id int) (model.Book, error)
	GetBooks(ctx context.Context, skip, take int) ([]model.Book, error)
	CreateBook(context.Context, model.Book) (model.Book, error)
}

type Logger interface {
	Errorw(msg string, keysAndValues ...interface{})
	Irrorw(msg string, keysAndValues ...interface{})
	Errorf(template string, args ...interface{})
	Infof(template string, args ...interface{})
}
