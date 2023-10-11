// Code generated by goa v3.13.2, DO NOT EDIT.
//
// librarian HTTP server
//
// Command:
// $ goa gen github.com/dgdraganov/super-librarian/design

package server

import (
	"context"
	"net/http"

	librarian "github.com/dgdraganov/super-librarian/gen/librarian"
	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// Server lists the librarian service endpoint HTTP handlers.
type Server struct {
	Mounts             []*MountPoint
	GetBook            http.Handler
	GetBooks           http.Handler
	CreateBook         http.Handler
	UpdateBook         http.Handler
	DeleteBook         http.Handler
	GenHTTPOpenapiJSON http.Handler
}

// MountPoint holds information about the mounted endpoints.
type MountPoint struct {
	// Method is the name of the service method served by the mounted HTTP handler.
	Method string
	// Verb is the HTTP method used to match requests to the mounted handler.
	Verb string
	// Pattern is the HTTP request path pattern used to match requests to the
	// mounted handler.
	Pattern string
}

// New instantiates HTTP handlers for all the librarian service endpoints using
// the provided encoder and decoder. The handlers are mounted on the given mux
// using the HTTP verb and path defined in the design. errhandler is called
// whenever a response fails to be encoded. formatter is used to format errors
// returned by the service methods prior to encoding. Both errhandler and
// formatter are optional and can be nil.
func New(
	e *librarian.Endpoints,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(ctx context.Context, err error) goahttp.Statuser,
	fileSystemGenHTTPOpenapiJSON http.FileSystem,
) *Server {
	if fileSystemGenHTTPOpenapiJSON == nil {
		fileSystemGenHTTPOpenapiJSON = http.Dir(".")
	}
	return &Server{
		Mounts: []*MountPoint{
			{"GetBook", "GET", "/book/{id}"},
			{"GetBooks", "GET", "/books/{skip}/{take}"},
			{"CreateBook", "POST", "/book"},
			{"UpdateBook", "PATCH", "/book"},
			{"DeleteBook", "DELETE", "/book/{id}"},
			{"./gen/http/openapi.json", "GET", "/openapi.json"},
		},
		GetBook:            NewGetBookHandler(e.GetBook, mux, decoder, encoder, errhandler, formatter),
		GetBooks:           NewGetBooksHandler(e.GetBooks, mux, decoder, encoder, errhandler, formatter),
		CreateBook:         NewCreateBookHandler(e.CreateBook, mux, decoder, encoder, errhandler, formatter),
		UpdateBook:         NewUpdateBookHandler(e.UpdateBook, mux, decoder, encoder, errhandler, formatter),
		DeleteBook:         NewDeleteBookHandler(e.DeleteBook, mux, decoder, encoder, errhandler, formatter),
		GenHTTPOpenapiJSON: http.FileServer(fileSystemGenHTTPOpenapiJSON),
	}
}

// Service returns the name of the service served.
func (s *Server) Service() string { return "librarian" }

// Use wraps the server handlers with the given middleware.
func (s *Server) Use(m func(http.Handler) http.Handler) {
	s.GetBook = m(s.GetBook)
	s.GetBooks = m(s.GetBooks)
	s.CreateBook = m(s.CreateBook)
	s.UpdateBook = m(s.UpdateBook)
	s.DeleteBook = m(s.DeleteBook)
}

// MethodNames returns the methods served.
func (s *Server) MethodNames() []string { return librarian.MethodNames[:] }

// Mount configures the mux to serve the librarian endpoints.
func Mount(mux goahttp.Muxer, h *Server) {
	MountGetBookHandler(mux, h.GetBook)
	MountGetBooksHandler(mux, h.GetBooks)
	MountCreateBookHandler(mux, h.CreateBook)
	MountUpdateBookHandler(mux, h.UpdateBook)
	MountDeleteBookHandler(mux, h.DeleteBook)
	MountGenHTTPOpenapiJSON(mux, goahttp.Replace("", "/./gen/http/openapi.json", h.GenHTTPOpenapiJSON))
}

// Mount configures the mux to serve the librarian endpoints.
func (s *Server) Mount(mux goahttp.Muxer) {
	Mount(mux, s)
}

// MountGetBookHandler configures the mux to serve the "librarian" service
// "get-book" endpoint.
func MountGetBookHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("GET", "/book/{id}", f)
}

// NewGetBookHandler creates a HTTP handler which loads the HTTP request and
// calls the "librarian" service "get-book" endpoint.
func NewGetBookHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(ctx context.Context, err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodeGetBookRequest(mux, decoder)
		encodeResponse = EncodeGetBookResponse(encoder)
		encodeError    = EncodeGetBookError(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "get-book")
		ctx = context.WithValue(ctx, goa.ServiceKey, "librarian")
		payload, err := decodeRequest(r)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		res, err := endpoint(ctx, payload)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			errhandler(ctx, w, err)
		}
	})
}

// MountGetBooksHandler configures the mux to serve the "librarian" service
// "get-books" endpoint.
func MountGetBooksHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("GET", "/books/{skip}/{take}", f)
}

// NewGetBooksHandler creates a HTTP handler which loads the HTTP request and
// calls the "librarian" service "get-books" endpoint.
func NewGetBooksHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(ctx context.Context, err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodeGetBooksRequest(mux, decoder)
		encodeResponse = EncodeGetBooksResponse(encoder)
		encodeError    = goahttp.ErrorEncoder(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "get-books")
		ctx = context.WithValue(ctx, goa.ServiceKey, "librarian")
		payload, err := decodeRequest(r)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		res, err := endpoint(ctx, payload)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			errhandler(ctx, w, err)
		}
	})
}

// MountCreateBookHandler configures the mux to serve the "librarian" service
// "create-book" endpoint.
func MountCreateBookHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("POST", "/book", f)
}

// NewCreateBookHandler creates a HTTP handler which loads the HTTP request and
// calls the "librarian" service "create-book" endpoint.
func NewCreateBookHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(ctx context.Context, err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodeCreateBookRequest(mux, decoder)
		encodeResponse = EncodeCreateBookResponse(encoder)
		encodeError    = EncodeCreateBookError(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "create-book")
		ctx = context.WithValue(ctx, goa.ServiceKey, "librarian")
		payload, err := decodeRequest(r)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		res, err := endpoint(ctx, payload)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			errhandler(ctx, w, err)
		}
	})
}

// MountUpdateBookHandler configures the mux to serve the "librarian" service
// "update-book" endpoint.
func MountUpdateBookHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("PATCH", "/book", f)
}

// NewUpdateBookHandler creates a HTTP handler which loads the HTTP request and
// calls the "librarian" service "update-book" endpoint.
func NewUpdateBookHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(ctx context.Context, err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodeUpdateBookRequest(mux, decoder)
		encodeResponse = EncodeUpdateBookResponse(encoder)
		encodeError    = goahttp.ErrorEncoder(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "update-book")
		ctx = context.WithValue(ctx, goa.ServiceKey, "librarian")
		payload, err := decodeRequest(r)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		res, err := endpoint(ctx, payload)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			errhandler(ctx, w, err)
		}
	})
}

// MountDeleteBookHandler configures the mux to serve the "librarian" service
// "delete-book" endpoint.
func MountDeleteBookHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("DELETE", "/book/{id}", f)
}

// NewDeleteBookHandler creates a HTTP handler which loads the HTTP request and
// calls the "librarian" service "delete-book" endpoint.
func NewDeleteBookHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(ctx context.Context, err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodeDeleteBookRequest(mux, decoder)
		encodeResponse = EncodeDeleteBookResponse(encoder)
		encodeError    = goahttp.ErrorEncoder(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "delete-book")
		ctx = context.WithValue(ctx, goa.ServiceKey, "librarian")
		payload, err := decodeRequest(r)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		res, err := endpoint(ctx, payload)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			errhandler(ctx, w, err)
		}
	})
}

// MountGenHTTPOpenapiJSON configures the mux to serve GET request made to
// "/openapi.json".
func MountGenHTTPOpenapiJSON(mux goahttp.Muxer, h http.Handler) {
	mux.Handle("GET", "/openapi.json", h.ServeHTTP)
}
