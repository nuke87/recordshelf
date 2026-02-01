package domain

import "time"

// Track represents a single track on an album.
type Track struct {
	ID              int64
	AlbumID         int64
	Position        string
	Title           string
	DurationSeconds int
	CreatedAt       time.Time
	UpdatedAt       time.Time
}
