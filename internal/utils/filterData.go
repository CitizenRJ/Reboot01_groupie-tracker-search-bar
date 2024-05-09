package groupie

import (
	gpd "groupie/internal/models"
	"strconv"
	"strings"
)

// FilterData filters the provided data based on the given criteria and returns the filtered data.
//
// allData is the full set of data to be filtered.
// buttonAll is a string indicating whether the "All" button is selected.
// tabButtons is a slice of integers representing the number of members in the artist groups to filter by.
// selectedCity is a string representing the selected city to filter by.
// minCreationYear is an integer representing the minimum creation year for the artist to be included.
// minAlbumYear is an integer representing the minimum album year for the artist to be included.
//
// The function returns the filtered data as a gpd.Data struct.
func FilterData(allData gpd.Data, buttonAll string, tabButtons []int, selectedCity string, minCreationYear int, minAlbumYear int) gpd.Data {
	var filteredData gpd.Data
	var albumYears []int
	filteredNames := []string{}

	for i := 0; i < (len(allData.Artist)); i++ {
		albumYearStr := strings.Split(allData.Artist[i].First_album, "-")[2]
		albumYear, _ := strconv.Atoi(albumYearStr)
		albumYears = append(albumYears, albumYear)
	}
	for i := 0; i < len(allData.Artist); i++ {
		for _, location := range allData.Location[i].Locations {
			country := strings.Split(location, "-")[1]
			for j := 0; j < 8; j++ {
				if string(buttonAll) != "All" && selectedCity != "All" {
					if j+1 == tabButtons[j] {
						if len(allData.Artist[i].Members) == tabButtons[j] && allData.Artist[i].Creation_date >= minCreationYear && int(albumYears[i]) >= minAlbumYear && country == selectedCity && !Isin(allData.Artist[i].Name, filteredNames) {
							filteredData.Artist = DisplayData(i, filteredData, allData)
							filteredNames = append(filteredNames, allData.Artist[i].Name)
						}
					}
				} else if buttonAll == "All" && selectedCity != "All" {
					if allData.Artist[i].Creation_date >= minCreationYear && int(albumYears[i]) >= minAlbumYear && country == selectedCity && !Isin(allData.Artist[i].Name, filteredNames) {
						filteredData.Artist = DisplayData(i, filteredData, allData)
						filteredNames = append(filteredNames, allData.Artist[i].Name)
					}
				} else if buttonAll == "All" && selectedCity == "All" {
					if allData.Artist[i].Creation_date >= minCreationYear && int(albumYears[i]) >= minAlbumYear && !Isin(allData.Artist[i].Name, filteredNames) {
						filteredData.Artist = DisplayData(i, filteredData, allData)
						filteredNames = append(filteredNames, allData.Artist[i].Name)
					}
				} else if string(buttonAll) != "All" && selectedCity == "All" {
					if len(allData.Artist[i].Members) == tabButtons[j] && allData.Artist[i].Creation_date >= minCreationYear && int(albumYears[i]) >= minAlbumYear && !Isin(allData.Artist[i].Name, filteredNames) {
						filteredData.Artist = DisplayData(i, filteredData, allData)
						filteredNames = append(filteredNames, allData.Artist[i].Name)
					}
				}
			}
		}
	}

	filteredData.All = allData.All
	filteredData.Country = allData.Country
	return filteredData
}
