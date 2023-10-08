package model

import "time"

type Book struct {
	Id          int
	Title       string
	Author      string
	BookCover   string
	PublishedAt time.Time
}
