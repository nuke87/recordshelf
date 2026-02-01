CREATE TABLE IF NOT EXISTS albums (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    title TEXT NOT NULL,
    artist TEXT NOT NULL,
    year INTEGER,
    label TEXT,
    format TEXT NOT NULL CHECK (format IN ('vinyl', 'cd')),
    catalog_number TEXT,
    pressing TEXT,
    pressing_count INTEGER,
    cover_url TEXT,
    discogs_id TEXT,
    musicbrainz_id TEXT,
    notes TEXT,
    created_at TEXT NOT NULL,
    updated_at TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS tracks (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    album_id INTEGER NOT NULL,
    position TEXT,
    title TEXT NOT NULL,
    duration_seconds INTEGER,
    created_at TEXT NOT NULL,
    updated_at TEXT NOT NULL,
    FOREIGN KEY (album_id) REFERENCES albums(id) ON DELETE CASCADE
);

CREATE INDEX IF NOT EXISTS idx_albums_artist_title ON albums(artist, title);
CREATE INDEX IF NOT EXISTS idx_tracks_album_id ON tracks(album_id);
