package groupie

// API holds information about artists playing shows
// in different locations and dates.
type API struct {
	Artists   string `json:"artists"`
	Locations string `json:"locations"`
	Dates     string `json:"dates"`
	Relations string `json:"relation"`
}
