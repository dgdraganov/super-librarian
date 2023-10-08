package librarianapi

import (
	"context"
	"log"

	librarian "github.com/dgdraganov/super-librarian/gen/librarian"
)

// librarian service example implementation.
// The example methods log the requests and return zero values.
type librariansrvc struct {
	logger *log.Logger
}

// NewLibrarian returns the librarian service implementation.
func NewLibrarian(logger *log.Logger) librarian.Service {
	return &librariansrvc{logger}
}

// Retrieve a book by id.
func (s *librariansrvc) GetBook(ctx context.Context, p *librarian.GetBookPayload) (res *librarian.Getbookresponse, err error) {
	res = &librarian.Getbookresponse{}
	s.logger.Print("librarian.get-book")
	return
}

// Get paginated books by specifying the number of books to skip and take.
func (s *librariansrvc) GetBooks(ctx context.Context, p *librarian.GetBooksPayload) (res *librarian.Getbooksresponse, err error) {
	res = &librarian.Getbooksresponse{}
	s.logger.Print("librarian.get-books")
	return
}

// Create a single book.
func (s *librariansrvc) CreateBook(ctx context.Context, p *librarian.CreateBookPayload) (res *librarian.Createbookresponse, err error) {
	res = &librarian.Createbookresponse{}
	s.logger.Print("librarian.create-book")
	return
}

// Updates a book by the given id.
func (s *librariansrvc) UpdateBook(ctx context.Context, p *librarian.UpdateBookPayload) (res *librarian.Updatebookresponse, err error) {
	res = &librarian.Updatebookresponse{}
	s.logger.Print("librarian.update-book")
	return
}

// Delete a single book.
func (s *librariansrvc) DeleteBook(ctx context.Context, p *librarian.DeleteBookPayload) (err error) {
	s.logger.Print("librarian.delete-book")
	return
}
