package groupie

// Data is a struct that contains various fields related to artist data:
//
// - Date contains date IDs and date strings
// - Artist contains artist objects
// - Location contains location IDs and location strings
// - Relation contains relations between dates and locations
// - Locs contains nested location strings
// - NumMembers contains number of members
// - All contains generic strings
// - Country contains country strings
type Data struct {
	Date []struct {
		Id    int      `json:"id"`
		Dates []string `json:"dates"`
	}
	Artist   []Artists
	Location []struct {
		Id        int      `json:"id"`
		Locations []string `json:"locations"`
	}
	Relation []struct {
		Id             int                 `json:"id"`
		DatesLocations map[string][]string `json:"datesLocations"`
	}

	Locs       [][][]string
	NumMembers []int
	All        []string
	Country    []string
}
