// Code generated by goa v3.13.2, DO NOT EDIT.
//
// books views
//
// Command:
// $ goa gen github.com/dgdraganov/super-librarian/design

package views

import (
	goa "goa.design/goa/v3/pkg"
)

// Getbookresponse is the viewed result type that is projected based on a view.
type Getbookresponse struct {
	// Type to project
	Projected *GetbookresponseView
	// View to render
	View string
}

// Getbooksresponse is the viewed result type that is projected based on a view.
type Getbooksresponse struct {
	// Type to project
	Projected *GetbooksresponseView
	// View to render
	View string
}

// Createbookresponse is the viewed result type that is projected based on a
// view.
type Createbookresponse struct {
	// Type to project
	Projected *CreatebookresponseView
	// View to render
	View string
}

// Updatebookresponse is the viewed result type that is projected based on a
// view.
type Updatebookresponse struct {
	// Type to project
	Projected *UpdatebookresponseView
	// View to render
	View string
}

// GetbookresponseView is a type that runs validations on a projected type.
type GetbookresponseView struct {
	// The unique id of the book.
	ID *int
	// The title of the book.
	Title *string
	// The author of the book.
	Author *string
	// The cover image of the book.
	BookCover *string
	// The date at which the book was published.
	PublishedAt *string
}

// GetbooksresponseView is a type that runs validations on a projected type.
type GetbooksresponseView struct {
	// List of paginated books.
	Books ResultbookCollectionView
}

// ResultbookCollectionView is a type that runs validations on a projected type.
type ResultbookCollectionView []*ResultbookView

// ResultbookView is a type that runs validations on a projected type.
type ResultbookView struct {
	// The unique id of the book.
	ID *int
	// The title of the book.
	Title *string
	// The author of the book.
	Author *string
	// The cover image of the book.
	BookCover *string
	// The date at which the book was published.
	PublishedAt *string
}

// CreatebookresponseView is a type that runs validations on a projected type.
type CreatebookresponseView struct {
	// The unique id of the book.
	ID *int
	// The title of the book.
	Title *string
	// The author of the book.
	Author *string
	// The cover image of the book.
	BookCover *string
	// The date at which the book was published.
	PublishedAt *string
}

// UpdatebookresponseView is a type that runs validations on a projected type.
type UpdatebookresponseView struct {
	// The unique id of the book.
	ID *int
	// The title of the book.
	Title *string
	// The author of the book.
	Author *string
	// The cover image of the book.
	BookCover *string
	// The date at which the book was published.
	PublishedAt *string
}

var (
	// GetbookresponseMap is a map indexing the attribute names of Getbookresponse
	// by view name.
	GetbookresponseMap = map[string][]string{
		"default": {
			"id",
			"title",
			"author",
			"book_cover",
			"published_at",
		},
	}
	// GetbooksresponseMap is a map indexing the attribute names of
	// Getbooksresponse by view name.
	GetbooksresponseMap = map[string][]string{
		"default": {
			"books",
		},
	}
	// CreatebookresponseMap is a map indexing the attribute names of
	// Createbookresponse by view name.
	CreatebookresponseMap = map[string][]string{
		"default": {
			"id",
			"title",
			"author",
			"book_cover",
			"published_at",
		},
	}
	// UpdatebookresponseMap is a map indexing the attribute names of
	// Updatebookresponse by view name.
	UpdatebookresponseMap = map[string][]string{
		"default": {
			"id",
			"title",
			"author",
			"book_cover",
			"published_at",
		},
	}
	// ResultbookCollectionMap is a map indexing the attribute names of
	// ResultbookCollection by view name.
	ResultbookCollectionMap = map[string][]string{
		"default": {
			"id",
			"title",
			"author",
			"book_cover",
			"published_at",
		},
	}
	// ResultbookMap is a map indexing the attribute names of Resultbook by view
	// name.
	ResultbookMap = map[string][]string{
		"default": {
			"id",
			"title",
			"author",
			"book_cover",
			"published_at",
		},
	}
)

// ValidateGetbookresponse runs the validations defined on the viewed result
// type Getbookresponse.
func ValidateGetbookresponse(result *Getbookresponse) (err error) {
	switch result.View {
	case "default", "":
		err = ValidateGetbookresponseView(result.Projected)
	default:
		err = goa.InvalidEnumValueError("view", result.View, []any{"default"})
	}
	return
}

// ValidateGetbooksresponse runs the validations defined on the viewed result
// type Getbooksresponse.
func ValidateGetbooksresponse(result *Getbooksresponse) (err error) {
	switch result.View {
	case "default", "":
		err = ValidateGetbooksresponseView(result.Projected)
	default:
		err = goa.InvalidEnumValueError("view", result.View, []any{"default"})
	}
	return
}

// ValidateCreatebookresponse runs the validations defined on the viewed result
// type Createbookresponse.
func ValidateCreatebookresponse(result *Createbookresponse) (err error) {
	switch result.View {
	case "default", "":
		err = ValidateCreatebookresponseView(result.Projected)
	default:
		err = goa.InvalidEnumValueError("view", result.View, []any{"default"})
	}
	return
}

// ValidateUpdatebookresponse runs the validations defined on the viewed result
// type Updatebookresponse.
func ValidateUpdatebookresponse(result *Updatebookresponse) (err error) {
	switch result.View {
	case "default", "":
		err = ValidateUpdatebookresponseView(result.Projected)
	default:
		err = goa.InvalidEnumValueError("view", result.View, []any{"default"})
	}
	return
}

// ValidateGetbookresponseView runs the validations defined on
// GetbookresponseView using the "default" view.
func ValidateGetbookresponseView(result *GetbookresponseView) (err error) {
	if result.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "result"))
	}
	if result.Title == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("title", "result"))
	}
	if result.BookCover == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("book_cover", "result"))
	}
	if result.Author == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("author", "result"))
	}
	if result.PublishedAt == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("published_at", "result"))
	}
	return
}

// ValidateGetbooksresponseView runs the validations defined on
// GetbooksresponseView using the "default" view.
func ValidateGetbooksresponseView(result *GetbooksresponseView) (err error) {

	if result.Books != nil {
		if err2 := ValidateResultbookCollectionView(result.Books); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// ValidateResultbookCollectionView runs the validations defined on
// ResultbookCollectionView using the "default" view.
func ValidateResultbookCollectionView(result ResultbookCollectionView) (err error) {
	for _, item := range result {
		if err2 := ValidateResultbookView(item); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// ValidateResultbookView runs the validations defined on ResultbookView using
// the "default" view.
func ValidateResultbookView(result *ResultbookView) (err error) {
	if result.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "result"))
	}
	if result.Title == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("title", "result"))
	}
	if result.BookCover == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("book_cover", "result"))
	}
	if result.Author == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("author", "result"))
	}
	if result.PublishedAt == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("published_at", "result"))
	}
	return
}

// ValidateCreatebookresponseView runs the validations defined on
// CreatebookresponseView using the "default" view.
func ValidateCreatebookresponseView(result *CreatebookresponseView) (err error) {
	if result.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "result"))
	}
	if result.Title == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("title", "result"))
	}
	if result.BookCover == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("book_cover", "result"))
	}
	if result.Author == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("author", "result"))
	}
	if result.PublishedAt == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("published_at", "result"))
	}
	return
}

// ValidateUpdatebookresponseView runs the validations defined on
// UpdatebookresponseView using the "default" view.
func ValidateUpdatebookresponseView(result *UpdatebookresponseView) (err error) {
	if result.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "result"))
	}
	if result.Title == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("title", "result"))
	}
	if result.BookCover == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("book_cover", "result"))
	}
	if result.Author == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("author", "result"))
	}
	if result.PublishedAt == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("published_at", "result"))
	}
	return
}
