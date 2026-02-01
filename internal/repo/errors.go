package repo

import "errors"

// ErrNotFound signals that a requested entity does not exist.
var ErrNotFound = errors.New("not found")
