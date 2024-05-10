package groupie

import (
	"strconv"
	"strings"

	gpd "groupie/internal/models"
)

// FilterData filters the provided data based on the given criteria and returns the filtered data.
func FilterData(allData gpd.Data, buttonAll string, tabButtons []int, selectedCity string, minCreationYear int, minAlbumYear int) gpd.Data {
	var filteredData gpd.Data
	var albumYears []int
	filteredNames := []string{}

	albumYears = getAlbumYears(allData)
	filteredData = filterDataByConditions(allData, buttonAll, tabButtons, selectedCity, minCreationYear, minAlbumYear, albumYears, filteredNames)
	filteredData.All = allData.All
	filteredData.Country = allData.Country

	return filteredData
}

func getAlbumYears(allData gpd.Data) []int {
	var albumYears []int
	for i := 0; i < len(allData.Artist); i++ {
		albumYearStr := strings.Split(allData.Artist[i].First_album, "-")[2]
		albumYear, _ := strconv.Atoi(albumYearStr)
		albumYears = append(albumYears, albumYear)
	}
	return albumYears
}

func filterDataByConditions(allData gpd.Data, buttonAll string, tabButtons []int, selectedCity string, minCreationYear int, minAlbumYear int, albumYears []int, filteredNames []string) gpd.Data {
	var filteredData gpd.Data
	for i := 0; i < len(allData.Artist); i++ {
		for _, location := range allData.Location[i].Locations {
			country := strings.Split(location, "-")[1]
			for j := 0; j < 8; j++ {
				if meetsCriteria(allData, buttonAll, tabButtons, selectedCity, minCreationYear, minAlbumYear, albumYears, i, j, country, filteredNames) {
					filteredData.Artist = DisplayData(i, filteredData, allData)
					filteredNames = append(filteredNames, allData.Artist[i].Name)
				}
			}
		}
	}
	return filteredData
}

func meetsCriteria(allData gpd.Data, buttonAll string, tabButtons []int, selectedCity string, minCreationYear int, minAlbumYear int, albumYears []int, i, j int, country string, filteredNames []string) bool {
	if string(buttonAll) != "All" && selectedCity != "All" {
		if j+1 == tabButtons[j] && len(allData.Artist[i].Members) == tabButtons[j] && allData.Artist[i].Creation_date >= minCreationYear && int(albumYears[i]) >= minAlbumYear && country == selectedCity && !Isin(allData.Artist[i].Name, filteredNames) {
			return true
		}
	} else if buttonAll == "All" && selectedCity != "All" {
		if allData.Artist[i].Creation_date >= minCreationYear && int(albumYears[i]) >= minAlbumYear && country == selectedCity && !Isin(allData.Artist[i].Name, filteredNames) {
			return true
		}
	} else if buttonAll == "All" && selectedCity == "All" {
		if allData.Artist[i].Creation_date >= minCreationYear && int(albumYears[i]) >= minAlbumYear && !Isin(allData.Artist[i].Name, filteredNames) {
			return true
		}
	} else if string(buttonAll) != "All" && selectedCity == "All" {
		if len(allData.Artist[i].Members) == tabButtons[j] && allData.Artist[i].Creation_date >= minCreationYear && int(albumYears[i]) >= minAlbumYear && !Isin(allData.Artist[i].Name, filteredNames) {
			return true
		}
	}
	return false
}
