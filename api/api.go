package groupie

import (
	"encoding/json"
	gpd "groupie/data"
	gpf "groupie/func"
	gps "groupie/search-bar"
	"io"
	"net/http"
	"strings"
)

var ApiData gpd.API
var GroupData gpd.Data

// GetData fetches and unmarshals data from the groupietrackers API.
// GetData fetches data from the groupietrackers API, unmarshals the
// JSON responses into structs, and returns the data. It gets and
// unmarshals the dates, artists, locations, and relations data.
func GetData() (gpd.Date, []gpd.Artists, gpd.GetLocation, gpd.Relations) {
	// Fetch the API data
	response, _ := http.Get("https://groupietrackers.herokuapp.com/api")

	// Read the response body
	responseData, _ := io.ReadAll(response.Body)
	// Unmarshal the JSON data into ApiData
	json.Unmarshal(responseData, &ApiData)

	// Fetch and unmarshal Dates data
	responseDates, _ := http.Get(ApiData.Dates)
	responseDataDates, _ := io.ReadAll(responseDates.Body)
	DateData := gpd.Date{}
	json.Unmarshal(responseDataDates, &DateData)

	// Fetch and unmarshal Artists data
	responseArtists, _ := http.Get(ApiData.Artists)
	responseDataArtists, _ := io.ReadAll(responseArtists.Body)
	ArtistsData := []gpd.Artists{}
	json.Unmarshal(responseDataArtists, &ArtistsData)

	// Populate additional fields for each artist
	for i := range ArtistsData {
		// Populate Locations field
		if ArtistsData[i].Id < len(GroupData.Location) {
			ArtistsData[i].Locations = GroupData.Location[ArtistsData[i].Id].Locations
		} else {
			// Handle the case where the artist ID is out of range
			continue
		}
		// Populate ConcertDates field
		concertDates := []string{}
		for _, dateId := range ArtistsData[i].Relations {
			for _, date := range GroupData.Date {
				if date.Id == dateId {
					concertDates = append(concertDates, date.Dates...)
				}
			}
		}
		ArtistsData[i].ConcertDates = concertDates
	}

	// Fetch and unmarshal Locations data
	responseLocation, _ := http.Get(ApiData.Locations)
	responseDataLocation, _ := io.ReadAll(responseLocation.Body)
	LocationData := gpd.GetLocation{}
	json.Unmarshal(responseDataLocation, &LocationData)

	// Fetch and unmarshal Relations data
	responseRelation, _ := http.Get(ApiData.Relations)
	responseDataRelation, _ := io.ReadAll(responseRelation.Body)
	RelationData := gpd.Relations{}
	json.Unmarshal(responseDataRelation, &RelationData)

	// Return all fetched and unmarshaled data
	return DateData, ArtistsData, LocationData, RelationData
}

// SetData processes the fetched API data into a usable structure for the application.
// It takes the date, artists, location and relation data and processes it into a single
// Data struct that can be used for display and search functionality. Key steps:
// - Set the main indexes from the fetched data
// - Initialize and populate search data
// - Calculate number of members
// - Generate 'all' data for search
// - Build country list
// Returns the fully populated Data struct.
func SetData(date gpd.Date, artists []gpd.Artists, location gpd.GetLocation, relation gpd.Relations) gpd.Data {
	// Set the Date index
	GroupData.Date = date.Index
	// Append each artist to GroupData.Artist
	for i := 0; i < (len(artists)); i++ {
		GroupData.Artist = append(GroupData.Artist, artists[i])
	}
	// Set the Location index
	GroupData.Location = location.Index
	// Set the Relation index
	GroupData.Relation = relation.Index

	// Initialize SearchData for search functionality
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

	// Assign processed SearchData to GroupData.Locs
	GroupData.Locs = SearchData

	// Calculate and set the number of members for each artist
	for i := 0; i < (len(GroupData.Artist)); i++ {
		GroupData.NumMembers = append(GroupData.NumMembers, len(GroupData.Artist[i].Members))
	}

	// Generate and set the All data for search functionality
	GroupData.All = gps.GetAll(GroupData)

	// Initialize and set the list of countries
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

	// Return the processed GroupData
	return GroupData
}
