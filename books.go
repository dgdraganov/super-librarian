package booksapi

import (
	"context"
	"log"

	books "github.com/dgdraganov/super-librarian/gen/books"
)

// books service example implementation.
// The example methods log the requests and return zero values.
type bookssrvc struct {
	logger *log.Logger
}

// NewBooks returns the books service implementation.
func NewBooks(logger *log.Logger) books.Service {
	return &bookssrvc{logger}
}

// Retrieve a book by id.
func (s *bookssrvc) GetBook(ctx context.Context, book *books.GetBookPayload) (*books.Getbookresponse, error) {
	res := &books.Getbookresponse{
		ID:          7,
		Title:       "It",
		Author:      "Penko Penkov",
		BookCover:   "http://alabala.com/my-image",
		PublishedAt: "2016-05-06",
	}
	s.logger.Print("books.get-book")
	return res, nil
}

// Get paginated books by specifying the number of books to skip and take.
func (s *bookssrvc) GetBooks(ctx context.Context, book *books.GetBooksPayload) (*books.Getbooksresponse, error) {
	res := books.Getbooksresponse{
		Books: []*books.Resultbook{
			{ID: 1, Title: "Love", Author: "Ivan Vazov", PublishedAt: "1905-05-04"},
			{ID: 2, Title: "Truth", Author: "haled Hussein", PublishedAt: "2018-05-04"},
		},
	}
	s.logger.Print("books.get-books")
	return &res, nil
}

// Create a single book.
func (s *bookssrvc) CreateBook(ctx context.Context, book *books.CreateBookPayload) (*books.Createbookresponse, error) {
	res := books.Createbookresponse{
		Title: "test_title",
	}
	s.logger.Print("books.create-book")
	s.logger.Printf("%+v", *book)
	return &res, nil
}

// Create a new book.
func (s *bookssrvc) UpdateBook(ctx context.Context, p *books.UpdateBookPayload) (res *books.Updatebookresponse, err error) {
	res = &books.Updatebookresponse{
		Title: "test_title",
	}
	s.logger.Print("books.update-book")
	return
}

// Delete a single book.
func (s *bookssrvc) DeleteBook(ctx context.Context, p *books.DeleteBookPayload) (err error) {
	s.logger.Print("books.delete-book")
	return
}
