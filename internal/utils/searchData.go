package groupie

import (
	gpd "groupie/internal/models"
	"strings"
)

// isArtistUnique checks if the given artist is unique within the provided bandMembers map.
// It iterates through the artist's members and checks if any of them have already been seen.
// If a member is found that has already been seen, the function returns false, indicating the artist is not unique.
// Otherwise, it returns true, indicating the artist is unique.
func isArtistUnique(artist gpd.Artists, bandMembers map[string]bool) bool {
	for _, member := range artist.Members {
		key := artist.Name + member
		if bandMembers[key] {
			return false
		}
	}
	return true
}

// SearchData searches the provided data for matches based on the given search term and search type.
// If the search type is 0, it searches for the search term in the artist name, first album, and location.
// If the search type is not 0, it searches for artists with a creation date matching the search type.
// The function returns a new Data struct containing the matching artists.
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
