// Code generated by goa v3.13.2, DO NOT EDIT.
//
// books endpoints
//
// Command:
// $ goa gen github.com/dgdraganov/super-librarian/design

package books

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Endpoints wraps the "books" service endpoints.
type Endpoints struct {
	GetBook    goa.Endpoint
	GetBooks   goa.Endpoint
	CreateBook goa.Endpoint
	UpdateBook goa.Endpoint
	DeleteBook goa.Endpoint
}

// NewEndpoints wraps the methods of the "books" service with endpoints.
func NewEndpoints(s Service) *Endpoints {
	return &Endpoints{
		GetBook:    NewGetBookEndpoint(s),
		GetBooks:   NewGetBooksEndpoint(s),
		CreateBook: NewCreateBookEndpoint(s),
		UpdateBook: NewUpdateBookEndpoint(s),
		DeleteBook: NewDeleteBookEndpoint(s),
	}
}

// Use applies the given middleware to all the "books" service endpoints.
func (e *Endpoints) Use(m func(goa.Endpoint) goa.Endpoint) {
	e.GetBook = m(e.GetBook)
	e.GetBooks = m(e.GetBooks)
	e.CreateBook = m(e.CreateBook)
	e.UpdateBook = m(e.UpdateBook)
	e.DeleteBook = m(e.DeleteBook)
}

// NewGetBookEndpoint returns an endpoint function that calls the method
// "get-book" of service "books".
func NewGetBookEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		p := req.(*GetBookPayload)
		res, err := s.GetBook(ctx, p)
		if err != nil {
			return nil, err
		}
		vres := NewViewedGetbookresponse(res, "default")
		return vres, nil
	}
}

// NewGetBooksEndpoint returns an endpoint function that calls the method
// "get-books" of service "books".
func NewGetBooksEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		p := req.(*GetBooksPayload)
		res, err := s.GetBooks(ctx, p)
		if err != nil {
			return nil, err
		}
		vres := NewViewedGetbooksresponse(res, "default")
		return vres, nil
	}
}

// NewCreateBookEndpoint returns an endpoint function that calls the method
// "create-book" of service "books".
func NewCreateBookEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		p := req.(*CreateBookPayload)
		res, err := s.CreateBook(ctx, p)
		if err != nil {
			return nil, err
		}
		vres := NewViewedCreatebookresponse(res, "default")
		return vres, nil
	}
}

// NewUpdateBookEndpoint returns an endpoint function that calls the method
// "update-book" of service "books".
func NewUpdateBookEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		p := req.(*UpdateBookPayload)
		res, err := s.UpdateBook(ctx, p)
		if err != nil {
			return nil, err
		}
		vres := NewViewedUpdatebookresponse(res, "default")
		return vres, nil
	}
}

// NewDeleteBookEndpoint returns an endpoint function that calls the method
// "delete-book" of service "books".
func NewDeleteBookEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		p := req.(*DeleteBookPayload)
		return nil, s.DeleteBook(ctx, p)
	}
}
