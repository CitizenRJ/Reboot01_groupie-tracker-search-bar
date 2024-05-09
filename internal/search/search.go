package groupie

import (
	gpd "groupie/internal/models"
	gpf "groupie/internal/utils"
	"strconv"
)

// GetAll processes the artist and location data to generate a list
// of unique strings containing all artist names, member names,
// locations, first albums, and creation dates. It replaces spaces
// with slashes in the names, converts the creation dates to strings,
// and de-duplicates the list before returning it.
func GetAll(data gpd.Data) []string {
	artistNames := []string{}
	memberNames := []string{}
	positions := []string{}
	firstAlbums := []string{}
	creationDates := []string{}
	combinedList := []string{}
	uniqueList := []string{}

	for _, artist := range data.Artist {
		processedName := ""
		for _, letter := range artist.Name {
			if string(letter) == " " {
				processedName += "/"
			} else {
				processedName += string(letter)
			}
		}
		artistNames = append(artistNames, processedName)
		firstAlbums = append(firstAlbums, artist.First_album)
		creationDates = append(creationDates, strconv.Itoa(artist.Creation_date))
		for _, member := range artist.Members {
			processedMember := ""
			for _, letter := range member {
				if string(letter) == " " {
					processedMember += "/"
				} else {
					processedMember += string(letter)
				}
			}
			memberNames = append(memberNames, processedMember)
		}
	}

	for _, location := range data.Location {
		for _, position := range location.Locations {
			positions = append(positions, position)
		}

	}

	combinedList = append(combinedList, artistNames...)
	combinedList = append(combinedList, memberNames...)
	combinedList = append(combinedList, positions...)
	combinedList = append(combinedList, firstAlbums...)
	combinedList = append(combinedList, creationDates...)
	for _, element := range combinedList {
		if !gpf.Isin(element, uniqueList) {
			uniqueList = append(uniqueList, element)
		}
	}

	return uniqueList
}
