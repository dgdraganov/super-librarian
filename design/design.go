package design

import (
	. "goa.design/goa/v3/dsl"
)

var _ = API("librarian", func() {
	Title("Books Service")
	Description("Service for managing Books with CRUD operations.")
	Server("librarian", func() {
		Host("localhost", func() {
			URI("http://localhost:8000")
		})
	})
})

var _ = Service("librarian", func() {
	Description("The books service performs CRUD operations on books.")

	Error("internal_server_error", func() {
		Description("Something went wrong on our end.")
		Fault()
	})
	Error("bad_request", func() {
		Description("Bad request.")
		Fault()
	})
	Error("not_found", func() {
		Description("Book not found.")
		Fault()
	})

	// get a single book
	Method("get-book", func() {
		Description("Retrieve a book by id.")
		Payload(GetBookPayload)

		Result(GetBookResponse)

		HTTP(func() {
			GET("/book/{id}")
			Response("not_found", StatusNotFound)
			Response("internal_server_error", StatusInternalServerError)
		})
	})

	// get books by pages
	Method("get-books", func() {
		Description("Get paginated books by specifying the number of books to skip and take.")
		Payload(GetBooksPayload)

		Result(GetBooksResponse)

		HTTP(func() {
			GET("/books/{skip}/{take}")
		})
	})

	// create a new book
	Method("create-book", func() {
		Description("Create a single book.")
		Payload(CreateBookPayload)

		Result(CreateBookResponse)
		HTTP(func() {
			POST("/book")
			Response("not_found", StatusNotFound)
			Response("bad_request", StatusBadRequest)
			Response("internal_server_error", StatusInternalServerError)
		})
	})

	// update an existing book
	Method("update-book", func() {
		Description("Updates a book by the given id.")
		Payload(UpdateBookPayload)

		Result(UpdateBookResponse)

		HTTP(func() {
			PATCH("/book")
		})
	})

	// delete a single book by id
	Method("delete-book", func() {
		Description("Delete a single book.")
		Payload(func() {
			Field(1, "id", Int, "Book id")
			Required("id")
		})

		HTTP(func() {
			DELETE("/book/{id}")
		})
	})

	Files("/openapi.json", "./gen/http/openapi.json")
})
