package repo

import "errors"

var ErrBookNotFound error = errors.New("no book exist with the given id")
