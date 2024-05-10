package groupie

import (
	"strconv"
	"strings"

	gpd "groupie/internal/models"
)

// GetAll retrieves all relevant data from the provided Data struct and returns a slice of unique strings.
// It processes the artist names, member names, locations, first albums, and creation dates,
// and combines them into a single list of unique strings.
func GetAll(data gpd.Data) []string {
	artistNames := []string{}
	memberNames := []string{}
	positions := []string{}
	firstAlbums := []string{}
	creationDates := []string{}

	processArtistData(data, &artistNames, &memberNames, &firstAlbums, &creationDates)
	processLocationData(data, &positions)
	combinedData := combineData(artistNames, memberNames, positions, firstAlbums, creationDates)
	uniqueList := makeUniqueList(combinedData)

	return uniqueList
}

func processArtistData(data gpd.Data, artistNames, memberNames, firstAlbums, creationDates *[]string) {
	for _, artist := range data.Artist {
		processedName := strings.Replace(artist.Name, " ", "/", -1)
		*artistNames = append(*artistNames, processedName)
		*firstAlbums = append(*firstAlbums, artist.First_album)
		*creationDates = append(*creationDates, strconv.Itoa(artist.Creation_date))
		for _, member := range artist.Members {
			processedMember := strings.Replace(member, " ", "/", -1)
			*memberNames = append(*memberNames, processedMember)
		}
	}
}

func processLocationData(data gpd.Data, positions *[]string) {
	locationMap := make(map[string]bool)
	for _, location := range data.Location {
		for _, loc := range location.Locations {
			if !locationMap[loc] {
				locationMap[loc] = true
				*positions = append(*positions, loc)
			}
		}
	}
}

func combineData(artistNames, memberNames, positions, firstAlbums, creationDates []string) []string {
	combinedData := []string{}
	combinedData = append(combinedData, strings.Join(artistNames, ""))
	combinedData = append(combinedData, strings.Join(memberNames, ""))
	combinedData = append(combinedData, strings.Join(positions, ""))
	combinedData = append(combinedData, firstAlbums...)
	combinedData = append(combinedData, creationDates...)
	return combinedData
}

func makeUniqueList(combinedData []string) []string {
	uniqueMap := make(map[string]bool)
	for _, element := range combinedData {
		uniqueMap[element] = true
	}

	uniqueList := make([]string, 0, len(uniqueMap))
	for element := range uniqueMap {
		uniqueList = append(uniqueList, element)
	}

	return uniqueList
}
