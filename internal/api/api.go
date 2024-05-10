package groupie

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"

	gpd "groupie/internal/models"
	gps "groupie/internal/search"
	gpf "groupie/internal/utils"
)

// Global variables that hold data
var ApiData gpd.API
var GroupData gpd.Data

// GetData fetches data from the Groupie Trackers API and returns the parsed Date, Artists, Location, and Relations data.
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
func SetData(date gpd.Date, artists []gpd.Artists, location gpd.GetLocation, relation gpd.Relations) gpd.Data {
	GroupData.Date = date.Index
	for i := 0; i < len(artists); i++ {
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

	for i := 0; i < len(GroupData.Artist); i++ {
		GroupData.NumMembers = append(GroupData.NumMembers, len(GroupData.Artist[i].Members))
	}

	GroupData.All = gps.GetAll(GroupData)

	// Deduplicate locations
	uniqueLocations := make(map[string]bool)
	for i := 0; i < len(GroupData.Location); i++ {
		for _, location := range GroupData.Location[i].Locations {
			uniqueLocations[location] = true
		}
	}

	var deduplicatedLocations []string
	for location := range uniqueLocations {
		deduplicatedLocations = append(deduplicatedLocations, location)
	}

	GroupData.Location = make([]struct {
		Id        int      `json:"id"`
		Locations []string `json:"locations"`
	}, len(deduplicatedLocations))

	for i, location := range deduplicatedLocations {
		GroupData.Location[i] = struct {
			Id        int      `json:"id"`
			Locations []string `json:"locations"`
		}{
			Id:        i + 1,
			Locations: []string{location},
		}
	}

	var countriesList []string
	for i := 0; i < len(GroupData.Location); i++ {
		for j := 0; j < len(GroupData.Location[i].Locations); j++ {
			country := strings.Split(GroupData.Location[i].Locations[j], "-")[1]
			if !gpf.Isin(country, countriesList) {
				countriesList = append(countriesList, country)
			}
		}
	}
	GroupData.Country = countriesList

	return GroupData
}
