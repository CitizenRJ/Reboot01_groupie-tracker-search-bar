package groupie

// Relations holds the index of relations between entities.
// The Index field contains the relation ID and associated date/location pairs.
type Relations struct {
	Index []struct {
		Id             int                 `json:"id"`
		DatesLocations map[string][]string `json:"datesLocations"`
	} `json:"index"`
}
