package repo

import (
	"context"

	"recordshelf/internal/domain"
)

// AlbumFilter defines optional filters for listing albums.
type AlbumFilter struct {
	Query  string
	Artist string
	Title  string
	Format *domain.Format
	Limit  int
	Offset int
}

// AlbumRepository defines persistence behavior for albums and tracks.
type AlbumRepository interface {
	List(ctx context.Context, filter AlbumFilter) ([]domain.Album, error)
	Get(ctx context.Context, id int64) (domain.Album, error)
	Save(ctx context.Context, album domain.Album) (int64, error)
	Delete(ctx context.Context, id int64) error
}
