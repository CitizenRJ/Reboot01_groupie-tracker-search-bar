package groupie

import (
	"encoding/json"
	gpd "groupie/internal/models"
	gps "groupie/internal/search"
	gpf "groupie/internal/utils"
	"io"
	"net/http"
	"strings"
)

// global variables that holds data
var ApiData gpd.API
var GroupData gpd.Data

// GetData fetches data from the Groupie Trackers API and returns the parsed Date, Artists, Location, and Relations data.
// The function makes HTTP GET requests to the API endpoints specified in the ApiData struct, reads the response bodies,
// and unmarshals the JSON data into the corresponding data structures.
// The function returns the fetched data, which can then be used to set the application's data state.
func GetData() (gpd.Date, []gpd.Artists, gpd.GetLocation, gpd.Relations) {
	response, _ := http.Get("https://groupietrackers.herokuapp.com/api")

	responseData, _ := io.ReadAll(response.Body)
	json.Unmarshal(responseData, &ApiData)

	responseDates, _ := http.Get(ApiData.Dates)
	responseDataDates, _ := io.ReadAll(responseDates.Body)
	DateData := gpd.Date{}
	json.Unmarshal(responseDataDates, &DateData)

	responseArtists, _ := http.Get(ApiData.Artists)
	responseDataArtists, _ := io.ReadAll(responseArtists.Body)
	ArtistsData := []gpd.Artists{}
	json.Unmarshal(responseDataArtists, &ArtistsData)

	responseLocation, _ := http.Get(ApiData.Locations)
	responseDataLocation, _ := io.ReadAll(responseLocation.Body)
	LocationData := gpd.GetLocation{}
	json.Unmarshal(responseDataLocation, &LocationData)

	responseRelation, _ := http.Get(ApiData.Relations)
	responseDataRelation, _ := io.ReadAll(responseRelation.Body)
	RelationData := gpd.Relations{}
	json.Unmarshal(responseDataRelation, &RelationData)

	return DateData, ArtistsData, LocationData, RelationData
}

// SetData sets the global GroupData variable with the provided date, artists, location, and relation data.
// It also generates a SearchData slice that contains the location and date information for each relation,
// and a list of all unique locations and countries represented in the data.
// The function returns the updated GroupData struct.
func SetData(date gpd.Date, artists []gpd.Artists, location gpd.GetLocation, relation gpd.Relations) gpd.Data {

	GroupData.Date = date.Index
	for i := 0; i < (len(artists)); i++ {
		GroupData.Artist = append(GroupData.Artist, artists[i])
	}
	GroupData.Location = location.Index
	GroupData.Relation = relation.Index

	SearchData := make([][][]string, 52)
	for i := 0; i < len(GroupData.Location); i++ {
		SearchData[i] = make([][]string, len(GroupData.Relation[i].DatesLocations))
		counter := 0
		for location, dates := range GroupData.Relation[i].DatesLocations {
			SearchData[i][counter] = append(SearchData[i][counter], location+" : ")
			for j := 0; j < len(dates); j++ {
				if j == 0 {
					SearchData[i][counter] = append(SearchData[i][counter], dates[j])
				} else if j >= 1 {
					SearchData[i][counter] = append(SearchData[i][counter], ", "+dates[j])
				} else {
					SearchData[i][counter] = append(SearchData[i][counter], dates[j])
				}
			}
			counter++
		}
	}

	GroupData.Locs = SearchData

	for i := 0; i < (len(GroupData.Artist)); i++ {
		GroupData.NumMembers = append(GroupData.NumMembers, len(GroupData.Artist[i].Members))
	}

	GroupData.All = gps.GetAll(GroupData)

	for i := 0; i < len(GroupData.Locs); i++ {
		for j := 0; j < len(GroupData.Locs[i]); j++ {
			for k := 0; k < len(GroupData.Locs[i][j]); k++ {
				if strings.Contains(GroupData.Locs[i][j][k], " : ") {
					location := strings.Split(GroupData.Locs[i][j][k], " : ")[0]
					if !gpf.Isin(location, GroupData.All) {
						GroupData.All = append(GroupData.All, location)
					}
				}
			}
		}
	}

	var countriesList []string
	for i := 0; i < (len(GroupData.Location)); i++ {
		for j := 0; j < (len(GroupData.Location[i].Locations)); j++ {
			country := strings.Split(GroupData.Location[i].Locations[j], "-")[1]
			if !gpf.Isin(country, countriesList) {
				countriesList = append(countriesList, country)
			}
		}
	}
	GroupData.Country = countriesList

	return GroupData
}
