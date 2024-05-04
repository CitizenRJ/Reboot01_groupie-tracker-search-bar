package groupie

import (
	gpd "groupie/data"
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

// SearchData filters the dataset based on the search term and search type.
// It returns a subset of the data where the search term matches any artist's name,
// member name, or location, depending on the search type specified.
// If searchType is 0, it performs a general search across multiple fields.
// If searchType is non-zero, it assumes the value represents a creation date and filters accordingly.
// The function ensures no duplicate artists are included in the results.
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
