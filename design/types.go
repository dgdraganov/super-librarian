package design

import (
	. "goa.design/goa/v3/dsl"
)

var UpdateBookPayload = Type("UpdateBookPayload", func() {
	Description("A single book payload.")
	Reference(CreateBookPayload)

	Attribute("id", Int, "The unique id of the book.", func() {
		Minimum(1)
	})

	Attribute("title")
	Attribute("author")
	Attribute("book_cover")
	Attribute("published_at")
})

var GetBookPayload = Type("GetBookPayload", func() {
	Description("ID of a book.")
	Attribute("id", Int, "The ID of the book.", func() {
		Minimum(1)
	})

	Required("id")
})

var GetBooksPayload = Type("GetBooksPayload", func() {
	Description("Number for 'skip' and 'take'")
	Attribute("skip", Int, "The numbers of records to skip before return.", func() {
		Minimum(0)
	})
	Attribute("take", Int, "The numbers of records to return.", func() {
		Minimum(1)
	})
	Required("skip", "take")
})

var CreateBookPayload = Type("CreateBookPayload", func() {
	Description("A single book payload.")
	Attribute("title", String, "The title of the book.", func() {
		MinLength(1)
		MaxLength(255)
	})
	Attribute("author", String, "The author of the book.", func() {
		MinLength(3)
		MaxLength(100)
	})
	Attribute("book_cover", String, "The URL to the cover image.", func() {
		MinLength(1)
		MaxLength(2048)
	})
	Attribute("published_at", String, "The date at which the book was published.", func() {
		Format(FormatDate)
	})
	Required("title", "author", "book_cover", "published_at")
})

var ResultBook = ResultType("ResultBook", func() {
	Attribute("id", Int, "The unique id of the book.")
	Attribute("title", String, "The title of the book.")
	Attribute("author", String, "The author of the book.")
	Attribute("book_cover", String, "The cover image of the book.")
	Attribute("published_at", String, "The date at which the book was published.")

	Required("id", "title", "book_cover", "author", "published_at")
})

var CreateBookResponse = ResultType("CreateBookResponse", func() {
	Description("A created book response.")
	Reference(ResultBook)

	Attributes(func() {
		Attribute("id")
		Attribute("title")
		Attribute("author")
		Attribute("book_cover")
		Attribute("published_at")
	})
})

var GetBookResponse = ResultType("GetBookResponse", func() {
	Description("A single book response.")
	Reference(ResultBook)
	Attributes(func() {
		Attribute("id")
		Attribute("title")
		Attribute("author")
		Attribute("book_cover")
		Attribute("published_at")
	})
})

var GetBooksResponse = ResultType("GetBooksResponse", func() {
	Attribute("books", CollectionOf(ResultBook), "List of paginated books.")
})

var UpdateBookResponse = ResultType("UpdateBookResponse", func() {
	Description("An updated book response.")
	Reference(ResultBook)
	Attributes(func() {
		Attribute("id")
		Attribute("title")
		Attribute("author")
		Attribute("book_cover")
		Attribute("published_at")
	})
})

var DeleteBookResponse = ResultType("DeleteBookResponse", func() {
	Description("A delete book response.")
	Attribute("id", Int, "The unique id of the book.", func() {
		Minimum(1)
	})
})
