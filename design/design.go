package design

import (
	. "goa.design/goa/dsl"
)

var _ = API("books", func() {
	Title("Books Service")
	Description("Service for managing Books with CRUD operations.")
	Server("books", func() {
		Host("localhost", func() {
			URI("http://localhost:8000")
		})
	})
})

var Book = Type("Book", func() {
	Attribute("id", Int, "The unique id of the book.", func() {
		Minimum(0)
	})
	Attribute("title", String, "The title of the book.", func() {
		MinLength(1)
		MaxLength(100)
	})
	Attribute("author", String, "The author of the book.", func() {
		MinLength(3)
		MaxLength(50)
	})
	Attribute("book_cover", String, "The cover image of the book.", func() {
		MinLength(15)
	})
	// // FormatDate: RFC3339 date
	Attribute("published_at", String, "The date at which the book was published.", func() {
		Format("RFC3339")
	})
})

var CreateBookPayload = Type("BookPayload", func() {
	Attribute("title", String, "The title of the book.", func() {
		MinLength(1)
		MaxLength(100)
	})
	Attribute("author", String, "The author of the book.", func() {
		MinLength(3)
		MaxLength(50)
	})
	Attribute("book_cover", String, "The cover image of the book.", func() {
		MinLength(15)
	})
	// FormatDate: RFC3339 date
	Attribute("published_at", String, "The date at which the book was published.", func() {
		Format("RFC3339")
	})
	Required("title", "author", "book_cover", "published_at")
})

var UpdateBookPayload = Type("BookPayload", func() {
	// Field(1, "id", Int, "The unique id of the book.")
	// Field(2, "title", String, "The title of the book.")
	// Field(3, "author", String, "The author of the book.")
	// Field(4, "book_cover", String, "The cover image of the book.")
	// Field(5, "published_at", String, "The date at which the book was published.")
	Attribute("id", Int, "The unique id of the book.", func() {
		Minimum(1)
	})
	Attribute("title", String, "The title of the book.", func() {
		MinLength(1)
		MaxLength(100)
	})
	Attribute("author", String, "The author of the book.", func() {
		MinLength(3)
		MaxLength(50)
	})
	Attribute("book_cover", String, "The cover image of the book.", func() {
		MinLength(15)
	})
	// FormatDate: RFC3339 date
	Attribute("published_at", String, "The date at which the book was published.", func() {
		Format("RFC3339")
	})
	Required("id", "title", "book_cover", "author", "published_at")
})

var _ = Service("books", func() {
	Description("The books service performs CRUD operations on books.")

	// get a single book
	Method("book", func() {
		Description("Retrieve a book by id.")
		Payload(func() {
			Field(1, "id", Int, "Book id")
			Required("id")
		})

		HTTP(func() {
			GET("/book/{id}")
		})
	})

	// get books by pages
	Method("books", func() {
		Description("Get paginated books by specifying the number of books to skip and take.")
		Payload(func() {
			Field(1, "skip", Int, "Number of books to skip")
			Field(2, "take", Int, "Number of books to take after skip")
			Required("skip")
			Required("take")
		})
		Result(CollectionOf(Book))
		HTTP(func() {
			GET("/books/{skip}/{take}")
		})
	})

	// create a new book
	Method("book", func() {
		Description("Create a single book.")
		Payload(CreateBookPayload)

		HTTP(func() {
			POST("/book")
			Body(Book)
		})
	})

	// update an existing book
	Method("book", func() {
		Description("Create a new book.")
		Payload(UpdateBookPayload)

		Result(Book)

		HTTP(func() {
			PATCH("/book")
			Body(Book)
		})
	})

	Method("delete", func() {
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
