package domain

// Format indicates the physical media type for an album.
type Format string

const (
	FormatVinyl Format = "vinyl"
	FormatCD    Format = "cd"
)

// Valid reports whether the format is one of the supported values.
func (f Format) Valid() bool {
	switch f {
	case FormatVinyl, FormatCD:
		return true
	default:
		return false
	}
}
