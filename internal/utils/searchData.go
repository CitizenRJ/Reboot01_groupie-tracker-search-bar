package groupie

import (
	gpd "groupie/internal/models"
	"strings"
)

// isArtistUnique checks if an artist is unique based on both the artist's name and the member's name.
func isArtistUnique(artist gpd.Artists, bandMembers map[string]bool) bool {
	for _, member := range artist.Members {
		key := artist.Name + member
		if bandMembers[key] {
			return false
		}
	}
	return true
}

func SearchData(searchTerm string, searchType int, allData gpd.Data) gpd.Data {
	var dataSearch gpd.Data
	bandMembers := make(map[string]bool)
	searchTerm = strings.ToLower(searchTerm)

	if searchType == 0 {
		for i := 0; i < len(allData.Artist); i++ {
			artist := allData.Artist[i]
			artistName := strings.ToLower(artist.Name)
			firstAlbum := strings.ToLower(artist.First_album)

			for _, location := range allData.Location[i].Locations {
				location := strings.ToLower(location)
				if strings.Contains(location, searchTerm) || strings.Contains(artistName, searchTerm) || strings.Contains(firstAlbum, searchTerm) {
					if isArtistUnique(artist, bandMembers) {
						dataSearch.Artist = append(dataSearch.Artist, artist)
						for _, member := range artist.Members {
							bandMembers[artist.Name+member] = true
						}
					}
				}
			}

			for _, member := range artist.Members {
				memberName := strings.ToLower(member)
				if strings.Contains(memberName, searchTerm) {
					if isArtistUnique(artist, bandMembers) {
						dataSearch.Artist = append(dataSearch.Artist, artist)
						for _, member := range artist.Members {
							bandMembers[artist.Name+member] = true
						}
					}
				}
			}
		}
	}

	if searchType != 0 {
		for i := 0; i < len(allData.Artist); i++ {
			artist := allData.Artist[i]
			if artist.Creation_date == searchType && isArtistUnique(artist, bandMembers) {
				dataSearch.Artist = append(dataSearch.Artist, artist)
				for _, member := range artist.Members {
					bandMembers[artist.Name+member] = true
				}
			}
		}
	}

	dataSearch.All = allData.All
	dataSearch.Country = allData.Country
	return dataSearch
}
