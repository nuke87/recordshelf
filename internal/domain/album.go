package domain

import "time"

// Album represents a vinyl record or CD in the collection.
type Album struct {
	ID            int64
	Title         string
	Artist        string
	Year          int
	Label         string
	Format        Format
	CatalogNumber string
	Pressing      string
	PressingCount int
	CoverURL      string
	DiscogsID     string
	MusicBrainzID string
	Notes         string
	CreatedAt     time.Time
	UpdatedAt     time.Time
	Tracks        []Track
}
