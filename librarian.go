package librarianapi

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	librarian "github.com/dgdraganov/super-librarian/gen/librarian"
	"github.com/dgdraganov/super-librarian/internal/model"
	"github.com/dgdraganov/super-librarian/internal/storage/repo"
)

var ErrInternalServerError = errors.New("internal server error")

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

	if errors.Is(err, repo.ErrNotFound) {
		s.logs.Printf("ERROR: not found: %s", err)
		return nil, librarian.MakeNotFound(errors.New("book not found"))
	}
	if err != nil {
		s.logs.Printf("ERROR: repo get book: %s", err)
		return nil, librarian.MakeInternalServerError(ErrInternalServerError)
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
	books, err := s.repo.GetBooks(ctx, p.Skip, p.Take)
	if err != nil {
		s.logs.Printf("repo get books: %s", err)
		return nil, librarian.MakeInternalServerError(ErrInternalServerError)
	}

	res.Books = make([]*librarian.Resultbook, 0, len(books))
	for i := 0; i < len(books); i++ {
		current := books[i]
		res.Books = append(res.Books, &librarian.Resultbook{
			ID:          current.Id,
			Author:      current.Author,
			Title:       current.Title,
			PublishedAt: current.PublishedAt.Format("2006-01-02"),
			BookCover:   current.BookCover,
		})
	}

	return res, nil
}

// Create a single book.
func (s *librariansrvc) CreateBook(ctx context.Context, p *librarian.CreateBookPayload) (res *librarian.Createbookresponse, err error) {
	published, err := time.Parse("2006-01-02", p.PublishedAt)
	if err != nil {
		s.logs.Printf("ERROR: time parse: %s", err)
		return nil, librarian.MakeBadRequest(fmt.Errorf("invalid request data: published_at: %s", p.PublishedAt))
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
		return nil, librarian.MakeInternalServerError(ErrInternalServerError)
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

	published, err := time.Parse("2006-01-02", p.PublishedAt)
	if err != nil {
		s.logs.Printf("ERROR: time parse: %s", err)
		return nil, librarian.MakeBadRequest(fmt.Errorf("invalid request data: published_at: %s", p.PublishedAt))
	}

	book := model.Book{
		Id:          *p.ID,
		Title:       p.Title,
		Author:      p.Author,
		PublishedAt: published,
		BookCover:   p.BookCover,
	}
	book, err = s.repo.UpdateBook(ctx, book)
	if err != nil {
		s.logs.Printf("ERROR: repo create book: %s", err)
		return nil, librarian.MakeInternalServerError(ErrInternalServerError)
	}

	return &librarian.Updatebookresponse{
		ID:          book.Id,
		Title:       book.Title,
		Author:      book.Author,
		PublishedAt: book.PublishedAt.Format("2006-01-02"),
		BookCover:   book.BookCover,
	}, nil
}

// Delete a single book.
func (s *librariansrvc) DeleteBook(ctx context.Context, p *librarian.DeleteBookPayload) error {
	err := s.repo.DeleteBook(ctx, p.ID)
	if err != nil {
		s.logs.Printf("ERROR: repo delete book: %s", err)
		return librarian.MakeInternalServerError(ErrInternalServerError)
	}
	return nil
}
