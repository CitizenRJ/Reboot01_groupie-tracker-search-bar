package groupie

// Artists represents information about a musical artist, including their ID, image, name, members,
// creation date, first album, locations, concert dates, and related artists.
type Artists struct {
	Id            int      `json:"id"`
	Image         string   `json:"image"`
	Name          string   `json:"name"`
	Members       []string `json:"members"`
	Creation_date int      `json:"creationDate"`
	First_album   string   `json:"firstAlbum"`
	Locations     []string `json:"locations"`
	ConcertDates  []string `json:"concertDates"`
	Relations     []int    `json:"relations"`
}

// ArtistInfo is a struct that holds various information about an artist, including their name,
// location, dates, and relationships. The All field contains a 2D slice of strings, and the
// Maps field is a string that represents mapping and location data.
type ArtistInfo struct {
	Artist   interface{}
	Location interface{}
	Date     interface{}
	Relation interface{}
	All      [][]string
	Maps     string
}

// Data represents a collection of data related to dates, artists, locations, and relations.
// The struct contains the following fields:
//
// - Date: a slice of structs, each containing an ID and a slice of date strings
// - Artist: a slice of Artist structs (not defined in the provided code)
// - Location: a slice of structs, each containing an ID and a slice of location strings
// - Relation: a slice of structs, each containing an ID and a map of date-location pairs
// - Locs: a 3D slice of strings representing locations
// - NumMembers: a slice of integers representing the number of members
// - All: a slice of strings
// - Country: a slice of strings representing countries
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

// Date represents a collection of dates, where each date has an ID and a list of dates.
type Date struct {
	Index []struct {
		Id    int      `json:"id"`
		Dates []string `json:"dates"`
	} `json:"index"`
}

// API represents a set of fields that can be used to query the application's data.
// The fields correspond to different types of data that can be retrieved, such as
// artists, locations, dates, and relations.
type API struct {
	Artists   string `json:"artists"`
	Locations string `json:"locations"`
	Dates     string `json:"dates"`
	Relations string `json:"relation"`
}

// FeatureCollection represents a collection of geographic features.
// Each feature has a center coordinate.
type FeatureCollection struct {
	Features []struct {
		Center []float64 `json:"center"`
	} `json:"features"`
}

// GetLocation represents a collection of locations associated with an identifier.
// The Index field contains a slice of structs, each with an Id and a slice of Locations.
// This type is likely used to represent a set of locations for some entity in the application.
type GetLocation struct {
	Index []struct {
		Id        int      `json:"id"`
		Locations []string `json:"locations"`
	} `json:"index"`
}

// Relations represents a collection of relationships, where each relationship has an ID and a map of dates to locations.
// The Index field contains a slice of structs, each with an Id and a DatesLocations map that associates dates with locations.
// This type is likely used to represent a set of relationships between entities in the application.
type Relations struct {
	Index []struct {
		Id             int                 `json:"id"`
		DatesLocations map[string][]string `json:"datesLocations"`
	} `json:"index"`
}
