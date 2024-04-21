package groupie

// ArtistInfo stores information about an artist.
//
// Artist is the name of the artist.
// Location is where the artist is based.
// Date is when the artist started performing.
// Relation lists other artists related to this artist.
// All lists all data associated with the artist.
// Maps maps artist data to a string representation.
type ArtistInfo struct {
	Artist   interface{}
	Location interface{}
	Date     interface{}
	Relation interface{}
	All      [][]string
	Maps     string
}
