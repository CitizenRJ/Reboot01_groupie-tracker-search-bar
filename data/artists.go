package groupie

// Artists represents the artist data including id, image, name, members,
// creation date, first album, locations, concert dates, and relations.
type Artists struct {
	Id            int      `json:"id"`
	Image         string   `json:"image"`
	Name          string   `json:"name"`
	Members       []string `json:"members"`
	Creation_date int      `json:"creationDate"`
	First_album   string   `json:"firstAlbum"`
	Locations     string   `json:"locations"`
	Concert_dates string   `json:"concertDates"`
	Relations     string   `json:"relations"`
}
