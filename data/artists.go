package groupie

// Artists represents the artist data including id, image, name, members,
// creation date, first album, locations, concert dates, and relations.
// data/artists.go
type Artists struct {
	Id            int      `json:"id"`
	Image         string   `json:"image"`
	Name          string   `json:"name"`
	Members       []string `json:"members"`
	Creation_date int      `json:"creationDate"`
	First_album   string   `json:"firstAlbum"`
	Locations     []string `json:"locations"`    // Changed from string to []string
	ConcertDates  []string `json:"concertDates"` // Changed from string to []string
	Relations     []int    `json:"relations"`
}
