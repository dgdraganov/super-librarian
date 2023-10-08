// Code generated by goa v3.13.2, DO NOT EDIT.
//
// books HTTP client CLI support package
//
// Command:
// $ goa gen github.com/dgdraganov/super-librarian/design

package client

import (
	"encoding/json"
	"fmt"
	"strconv"
	"unicode/utf8"

	books "github.com/dgdraganov/super-librarian/gen/books"
	goa "goa.design/goa/v3/pkg"
)

// BuildGetBookPayload builds the payload for the books get-book endpoint from
// CLI flags.
func BuildGetBookPayload(booksGetBookID string) (*books.GetBookPayload, error) {
	var err error
	var id int
	{
		var v int64
		v, err = strconv.ParseInt(booksGetBookID, 10, strconv.IntSize)
		id = int(v)
		if err != nil {
			return nil, fmt.Errorf("invalid value for id, must be INT")
		}
	}
	v := &books.GetBookPayload{}
	v.ID = id

	return v, nil
}

// BuildGetBooksPayload builds the payload for the books get-books endpoint
// from CLI flags.
func BuildGetBooksPayload(booksGetBooksSkip string, booksGetBooksTake string) (*books.GetBooksPayload, error) {
	var err error
	var skip int
	{
		var v int64
		v, err = strconv.ParseInt(booksGetBooksSkip, 10, strconv.IntSize)
		skip = int(v)
		if err != nil {
			return nil, fmt.Errorf("invalid value for skip, must be INT")
		}
	}
	var take int
	{
		var v int64
		v, err = strconv.ParseInt(booksGetBooksTake, 10, strconv.IntSize)
		take = int(v)
		if err != nil {
			return nil, fmt.Errorf("invalid value for take, must be INT")
		}
	}
	v := &books.GetBooksPayload{}
	v.Skip = skip
	v.Take = take

	return v, nil
}

// BuildCreateBookPayload builds the payload for the books create-book endpoint
// from CLI flags.
func BuildCreateBookPayload(booksCreateBookBody string) (*books.CreateBookPayload, error) {
	var err error
	var body CreateBookRequestBody
	{
		err = json.Unmarshal([]byte(booksCreateBookBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"author\": \"gtu\",\n      \"book_cover\": \"49g\",\n      \"published_at\": \"2009-04-11\",\n      \"title\": \"yb\"\n   }'")
		}
		if utf8.RuneCountInString(body.Title) < 1 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.title", body.Title, utf8.RuneCountInString(body.Title), 1, true))
		}
		if utf8.RuneCountInString(body.Title) > 100 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.title", body.Title, utf8.RuneCountInString(body.Title), 100, false))
		}
		if utf8.RuneCountInString(body.Author) < 3 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.author", body.Author, utf8.RuneCountInString(body.Author), 3, true))
		}
		if utf8.RuneCountInString(body.Author) > 50 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.author", body.Author, utf8.RuneCountInString(body.Author), 50, false))
		}
		if utf8.RuneCountInString(body.BookCover) < 15 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.book_cover", body.BookCover, utf8.RuneCountInString(body.BookCover), 15, true))
		}
		err = goa.MergeErrors(err, goa.ValidateFormat("body.published_at", body.PublishedAt, goa.FormatDate))
		if err != nil {
			return nil, err
		}
	}
	v := &books.CreateBookPayload{
		Title:       body.Title,
		Author:      body.Author,
		BookCover:   body.BookCover,
		PublishedAt: body.PublishedAt,
	}

	return v, nil
}

// BuildUpdateBookPayload builds the payload for the books update-book endpoint
// from CLI flags.
func BuildUpdateBookPayload(booksUpdateBookBody string) (*books.UpdateBookPayload, error) {
	var err error
	var body UpdateBookRequestBody
	{
		err = json.Unmarshal([]byte(booksUpdateBookBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"author\": \"67b\",\n      \"book_cover\": \"995\",\n      \"id\": 4031548831220571457,\n      \"published_at\": \"1993-10-28\",\n      \"title\": \"0j\"\n   }'")
		}
	}
	v := &books.UpdateBookPayload{
		ID: body.ID,
	}
	if body.Title != nil {
		v.Title = *body.Title
	}
	if body.Author != nil {
		v.Author = *body.Author
	}
	if body.BookCover != nil {
		v.BookCover = *body.BookCover
	}
	if body.PublishedAt != nil {
		v.PublishedAt = *body.PublishedAt
	}

	return v, nil
}

// BuildDeleteBookPayload builds the payload for the books delete-book endpoint
// from CLI flags.
func BuildDeleteBookPayload(booksDeleteBookID string) (*books.DeleteBookPayload, error) {
	var err error
	var id int
	{
		var v int64
		v, err = strconv.ParseInt(booksDeleteBookID, 10, strconv.IntSize)
		id = int(v)
		if err != nil {
			return nil, fmt.Errorf("invalid value for id, must be INT")
		}
	}
	v := &books.DeleteBookPayload{}
	v.ID = id

	return v, nil
}
