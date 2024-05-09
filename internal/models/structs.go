package groupie

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

type ArtistInfo struct {
	Artist   interface{}
	Location interface{}
	Date     interface{}
	Relation interface{}
	All      [][]string
	Maps     string
}

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

type Date struct {
	Index []struct {
		Id    int      `json:"id"`
		Dates []string `json:"dates"`
	} `json:"index"`
}

type API struct {
	Artists   string `json:"artists"`
	Locations string `json:"locations"`
	Dates     string `json:"dates"`
	Relations string `json:"relation"`
}

type FeatureCollection struct {
	Features []struct {
		Center []float64 `json:"center"`
	} `json:"features"`
}

type GetLocation struct {
	Index []struct {
		Id        int      `json:"id"`
		Locations []string `json:"locations"`
	} `json:"index"`
}

type Relations struct {
	Index []struct {
		Id             int                 `json:"id"`
		DatesLocations map[string][]string `json:"datesLocations"`
	} `json:"index"`
}
