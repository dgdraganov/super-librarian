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

	// get a single book
	Method("get-book", func() {
		Description("Retrieve a book by id.")
		Payload(GetBookPayload)

		Result(GetBookResponse)

		HTTP(func() {
			GET("/book/{id}")
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
		})
	})

	// update an existing book
	Method("update-book", func() {
		Description("Updates a book by the given id.")
		Payload(UpdateBookPayload)

		Result(UpdateBookResponse)

		HTTP(func() {
			PATCH("/book")
			// Body(Book)
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
