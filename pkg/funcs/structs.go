package funcs

type Band struct {
	Id            int                 `json:"id"`
	Image         string              `json:"image"`
	Name          string              `json:"name"`
	Members       []string            `json:"members"`
	CreationDate  uint                `json:"creationDate"`
	FirstAlbum    string              `json:"firstAlbum"`
	Relations     string              `json:"relations"`
	Concerts      map[string][]string `json:"datesLocations"`
	LocationsLink string              `json:"locations"`
	DatesLink     string              `json:"concertDates"`
	Type          string              `json:"type"`
	Location      []string
	Dates         []string
}

type Relation struct {
	Id             int                 `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}

type Location struct {
	Id        int      `json:"id"`
	Locations []string `json:"locations"`
}

type Date struct {
	Id   int      `json:"id"`
	Date []string `json:"dates"`
}

var data []Band

func SetData(d []Band) {
	data = d
}
