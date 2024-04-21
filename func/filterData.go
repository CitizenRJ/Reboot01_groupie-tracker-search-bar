package groupie

import (
	gpd "groupie/data"
	"strconv"
	"strings"
)

// FilterData filters the provided data based on the criteria specified in the parameters.
// It returns a new Data struct containing only the filtered data.
func FilterData(allData gpd.Data, buttonAll string, tabButtons []int, selectedCity string, minCreationYear int, minAlbumYear int) gpd.Data {
	var filteredData gpd.Data
	var albumYears []int
	var filteredNames []string

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
