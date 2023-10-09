// Code generated by goa v3.13.2, DO NOT EDIT.
//
// librarian HTTP server types
//
// Command:
// $ goa gen github.com/dgdraganov/super-librarian/design

package server

import (
	"unicode/utf8"

	librarian "github.com/dgdraganov/super-librarian/gen/librarian"
	librarianviews "github.com/dgdraganov/super-librarian/gen/librarian/views"
	goa "goa.design/goa/v3/pkg"
)

// CreateBookRequestBody is the type of the "librarian" service "create-book"
// endpoint HTTP request body.
type CreateBookRequestBody struct {
	// The title of the book.
	Title *string `form:"title,omitempty" json:"title,omitempty" xml:"title,omitempty"`
	// The author of the book.
	Author *string `form:"author,omitempty" json:"author,omitempty" xml:"author,omitempty"`
	// The cover image of the book.
	BookCover *string `form:"book_cover,omitempty" json:"book_cover,omitempty" xml:"book_cover,omitempty"`
	// The date at which the book was published.
	PublishedAt *string `form:"published_at,omitempty" json:"published_at,omitempty" xml:"published_at,omitempty"`
}

// UpdateBookRequestBody is the type of the "librarian" service "update-book"
// endpoint HTTP request body.
type UpdateBookRequestBody struct {
	// The unique id of the book.
	ID *int `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// The title of the book.
	Title *string `form:"title,omitempty" json:"title,omitempty" xml:"title,omitempty"`
	// The author of the book.
	Author *string `form:"author,omitempty" json:"author,omitempty" xml:"author,omitempty"`
	// The cover image of the book.
	BookCover *string `form:"book_cover,omitempty" json:"book_cover,omitempty" xml:"book_cover,omitempty"`
	// The date at which the book was published.
	PublishedAt *string `form:"published_at,omitempty" json:"published_at,omitempty" xml:"published_at,omitempty"`
}

// GetBookResponseBody is the type of the "librarian" service "get-book"
// endpoint HTTP response body.
type GetBookResponseBody struct {
	// The unique id of the book.
	ID int `form:"id" json:"id" xml:"id"`
	// The title of the book.
	Title string `form:"title" json:"title" xml:"title"`
	// The author of the book.
	Author string `form:"author" json:"author" xml:"author"`
	// The cover image of the book.
	BookCover string `form:"book_cover" json:"book_cover" xml:"book_cover"`
	// The date at which the book was published.
	PublishedAt string `form:"published_at" json:"published_at" xml:"published_at"`
}

// GetBooksResponseBody is the type of the "librarian" service "get-books"
// endpoint HTTP response body.
type GetBooksResponseBody struct {
	// List of paginated books.
	Books ResultbookResponseBodyCollection `form:"books,omitempty" json:"books,omitempty" xml:"books,omitempty"`
}

// CreateBookResponseBody is the type of the "librarian" service "create-book"
// endpoint HTTP response body.
type CreateBookResponseBody struct {
	// The unique id of the book.
	ID int `form:"id" json:"id" xml:"id"`
	// The title of the book.
	Title string `form:"title" json:"title" xml:"title"`
	// The author of the book.
	Author string `form:"author" json:"author" xml:"author"`
	// The cover image of the book.
	BookCover string `form:"book_cover" json:"book_cover" xml:"book_cover"`
	// The date at which the book was published.
	PublishedAt string `form:"published_at" json:"published_at" xml:"published_at"`
}

// UpdateBookResponseBody is the type of the "librarian" service "update-book"
// endpoint HTTP response body.
type UpdateBookResponseBody struct {
	// The unique id of the book.
	ID int `form:"id" json:"id" xml:"id"`
	// The title of the book.
	Title string `form:"title" json:"title" xml:"title"`
	// The author of the book.
	Author string `form:"author" json:"author" xml:"author"`
	// The cover image of the book.
	BookCover string `form:"book_cover" json:"book_cover" xml:"book_cover"`
	// The date at which the book was published.
	PublishedAt string `form:"published_at" json:"published_at" xml:"published_at"`
}

// ResultbookResponseBodyCollection is used to define fields on response body
// types.
type ResultbookResponseBodyCollection []*ResultbookResponseBody

// ResultbookResponseBody is used to define fields on response body types.
type ResultbookResponseBody struct {
	// The unique id of the book.
	ID int `form:"id" json:"id" xml:"id"`
	// The title of the book.
	Title string `form:"title" json:"title" xml:"title"`
	// The author of the book.
	Author string `form:"author" json:"author" xml:"author"`
	// The cover image of the book.
	BookCover string `form:"book_cover" json:"book_cover" xml:"book_cover"`
	// The date at which the book was published.
	PublishedAt string `form:"published_at" json:"published_at" xml:"published_at"`
}

// NewGetBookResponseBody builds the HTTP response body from the result of the
// "get-book" endpoint of the "librarian" service.
func NewGetBookResponseBody(res *librarianviews.GetbookresponseView) *GetBookResponseBody {
	body := &GetBookResponseBody{
		ID:          *res.ID,
		Title:       *res.Title,
		Author:      *res.Author,
		BookCover:   *res.BookCover,
		PublishedAt: *res.PublishedAt,
	}
	return body
}

// NewGetBooksResponseBody builds the HTTP response body from the result of the
// "get-books" endpoint of the "librarian" service.
func NewGetBooksResponseBody(res *librarianviews.GetbooksresponseView) *GetBooksResponseBody {
	body := &GetBooksResponseBody{}
	if res.Books != nil {
		body.Books = make([]*ResultbookResponseBody, len(res.Books))
		for i, val := range res.Books {
			body.Books[i] = marshalLibrarianviewsResultbookViewToResultbookResponseBody(val)
		}
	}
	return body
}

// NewCreateBookResponseBody builds the HTTP response body from the result of
// the "create-book" endpoint of the "librarian" service.
func NewCreateBookResponseBody(res *librarianviews.CreatebookresponseView) *CreateBookResponseBody {
	body := &CreateBookResponseBody{
		ID:          *res.ID,
		Title:       *res.Title,
		Author:      *res.Author,
		BookCover:   *res.BookCover,
		PublishedAt: *res.PublishedAt,
	}
	return body
}

// NewUpdateBookResponseBody builds the HTTP response body from the result of
// the "update-book" endpoint of the "librarian" service.
func NewUpdateBookResponseBody(res *librarianviews.UpdatebookresponseView) *UpdateBookResponseBody {
	body := &UpdateBookResponseBody{
		ID:          *res.ID,
		Title:       *res.Title,
		Author:      *res.Author,
		BookCover:   *res.BookCover,
		PublishedAt: *res.PublishedAt,
	}
	return body
}

// NewGetBookPayload builds a librarian service get-book endpoint payload.
func NewGetBookPayload(id int) *librarian.GetBookPayload {
	v := &librarian.GetBookPayload{}
	v.ID = id

	return v
}

// NewGetBooksPayload builds a librarian service get-books endpoint payload.
func NewGetBooksPayload(skip int, take int) *librarian.GetBooksPayload {
	v := &librarian.GetBooksPayload{}
	v.Skip = skip
	v.Take = take

	return v
}

// NewCreateBookPayload builds a librarian service create-book endpoint payload.
func NewCreateBookPayload(body *CreateBookRequestBody) *librarian.CreateBookPayload {
	v := &librarian.CreateBookPayload{
		Title:       *body.Title,
		Author:      *body.Author,
		BookCover:   *body.BookCover,
		PublishedAt: *body.PublishedAt,
	}

	return v
}

// NewUpdateBookPayload builds a librarian service update-book endpoint payload.
func NewUpdateBookPayload(body *UpdateBookRequestBody) *librarian.UpdateBookPayload {
	v := &librarian.UpdateBookPayload{
		ID:          body.ID,
		Title:       *body.Title,
		Author:      *body.Author,
		BookCover:   *body.BookCover,
		PublishedAt: *body.PublishedAt,
	}

	return v
}

// NewDeleteBookPayload builds a librarian service delete-book endpoint payload.
func NewDeleteBookPayload(id int) *librarian.DeleteBookPayload {
	v := &librarian.DeleteBookPayload{}
	v.ID = id

	return v
}

// ValidateCreateBookRequestBody runs the validations defined on
// Create-BookRequestBody
func ValidateCreateBookRequestBody(body *CreateBookRequestBody) (err error) {
	if body.Title == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("title", "body"))
	}
	if body.Author == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("author", "body"))
	}
	if body.BookCover == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("book_cover", "body"))
	}
	if body.PublishedAt == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("published_at", "body"))
	}
	if body.Title != nil {
		if utf8.RuneCountInString(*body.Title) < 1 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.title", *body.Title, utf8.RuneCountInString(*body.Title), 1, true))
		}
	}
	if body.Title != nil {
		if utf8.RuneCountInString(*body.Title) > 255 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.title", *body.Title, utf8.RuneCountInString(*body.Title), 255, false))
		}
	}
	if body.Author != nil {
		if utf8.RuneCountInString(*body.Author) < 3 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.author", *body.Author, utf8.RuneCountInString(*body.Author), 3, true))
		}
	}
	if body.Author != nil {
		if utf8.RuneCountInString(*body.Author) > 100 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.author", *body.Author, utf8.RuneCountInString(*body.Author), 100, false))
		}
	}
	if body.BookCover != nil {
		if utf8.RuneCountInString(*body.BookCover) < 15 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.book_cover", *body.BookCover, utf8.RuneCountInString(*body.BookCover), 15, true))
		}
	}
	if body.BookCover != nil {
		if utf8.RuneCountInString(*body.BookCover) > 1024 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.book_cover", *body.BookCover, utf8.RuneCountInString(*body.BookCover), 1024, false))
		}
	}
	if body.PublishedAt != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("body.published_at", *body.PublishedAt, goa.FormatDate))
	}
	return
}

// ValidateUpdateBookRequestBody runs the validations defined on
// Update-BookRequestBody
func ValidateUpdateBookRequestBody(body *UpdateBookRequestBody) (err error) {
	if body.Title == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("title", "body"))
	}
	if body.Author == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("author", "body"))
	}
	if body.BookCover == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("book_cover", "body"))
	}
	if body.PublishedAt == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("published_at", "body"))
	}
	if body.ID != nil {
		if *body.ID < 1 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("body.id", *body.ID, 1, true))
		}
	}
	if body.Title != nil {
		if utf8.RuneCountInString(*body.Title) < 1 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.title", *body.Title, utf8.RuneCountInString(*body.Title), 1, true))
		}
	}
	if body.Title != nil {
		if utf8.RuneCountInString(*body.Title) > 255 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.title", *body.Title, utf8.RuneCountInString(*body.Title), 255, false))
		}
	}
	if body.Author != nil {
		if utf8.RuneCountInString(*body.Author) < 3 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.author", *body.Author, utf8.RuneCountInString(*body.Author), 3, true))
		}
	}
	if body.Author != nil {
		if utf8.RuneCountInString(*body.Author) > 100 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.author", *body.Author, utf8.RuneCountInString(*body.Author), 100, false))
		}
	}
	if body.BookCover != nil {
		if utf8.RuneCountInString(*body.BookCover) < 15 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.book_cover", *body.BookCover, utf8.RuneCountInString(*body.BookCover), 15, true))
		}
	}
	if body.BookCover != nil {
		if utf8.RuneCountInString(*body.BookCover) > 1024 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.book_cover", *body.BookCover, utf8.RuneCountInString(*body.BookCover), 1024, false))
		}
	}
	if body.PublishedAt != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("body.published_at", *body.PublishedAt, goa.FormatDate))
	}
	return
}
