package librarianapi

import (
	"context"
	"log"
	"time"

	librarian "github.com/dgdraganov/super-librarian/gen/librarian"
	"github.com/dgdraganov/super-librarian/internal/model"
)

// librarian service example implementation.
// The example methods log the requests and return zero values.
type librariansrvc struct {
	logs *log.Logger
	repo BooksRepo
}

// NewLibrarian returns the librarian service implementation.
func NewLibrarian(booksRepository BooksRepo, logger *log.Logger) librarian.Service {
	return &librariansrvc{
		logs: logger,
		repo: booksRepository,
	}
}

// Retrieve a book by id.
func (s *librariansrvc) GetBook(ctx context.Context, p *librarian.GetBookPayload) (res *librarian.Getbookresponse, err error) {
	res = &librarian.Getbookresponse{}
	result, err := s.repo.GetBook(ctx, p.ID)
	if err != nil {
		s.logs.Printf("ERROR: repo get book: %s", err)
		// todo: return Invalid request or server error
		return &librarian.Getbookresponse{}, nil
	}
	return &librarian.Getbookresponse{
		ID:          result.Id,
		Title:       result.Title,
		Author:      result.Author,
		BookCover:   result.BookCover,
		PublishedAt: result.PublishedAt.Format("2006-01-02"),
	}, nil
}

// Get paginated books by specifying the number of books to skip and take.
func (s *librariansrvc) GetBooks(ctx context.Context, p *librarian.GetBooksPayload) (res *librarian.Getbooksresponse, err error) {
	res = &librarian.Getbooksresponse{}
	s.logs.Print("librarian.get-books")
	return
}

// Create a single book.
func (s *librariansrvc) CreateBook(ctx context.Context, p *librarian.CreateBookPayload) (res *librarian.Createbookresponse, err error) {
	s.logs.Print("librarian.create-book")
	published, err := time.Parse("2006-01-02", p.PublishedAt)
	if err != nil {
		// todo: return Invalid request
		s.logs.Printf("ERROR: time parse: %s", err)

		return &librarian.Createbookresponse{}, nil
	}

	book := model.Book{
		Title:       p.Title,
		Author:      p.Author,
		BookCover:   p.BookCover,
		PublishedAt: published,
	}

	result, err := s.repo.CreateBook(ctx, book)
	if err != nil {
		s.logs.Printf("ERROR: repo create book: %s", err)
		// todo: return Internal server error return
		return &librarian.Createbookresponse{}, nil
	}

	return &librarian.Createbookresponse{
		ID:          result.Id,
		Title:       result.Title,
		Author:      result.Author,
		BookCover:   result.BookCover,
		PublishedAt: result.PublishedAt.Format("2006-01-02"),
	}, nil
}

// Updates a book by the given id.
func (s *librariansrvc) UpdateBook(ctx context.Context, p *librarian.UpdateBookPayload) (res *librarian.Updatebookresponse, err error) {
	res = &librarian.Updatebookresponse{}
	s.logs.Print("librarian.update-book")
	return
}

// Delete a single book.
func (s *librariansrvc) DeleteBook(ctx context.Context, p *librarian.DeleteBookPayload) (err error) {
	s.logs.Print("librarian.delete-book")
	return
}
