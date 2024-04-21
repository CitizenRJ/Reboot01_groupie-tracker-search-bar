package groupie

import (
	gpd "groupie/data"
	gpm "groupie/map"
)

// InfoArtist returns an ArtistInfo struct containing detailed information
// about the artist at the given index in the provided Data struct. It
// populates the struct with data from the Data as well as calling CreateMap
// to generate related map data.
func InfoArtist(allData gpd.Data, index int) gpd.ArtistInfo {
	// Check if the index is within the valid range for all slices
	if index < 0 || index >= len(allData.Date) || index >= len(allData.Artist) || index >= len(allData.Location) || index >= len(allData.Locs) {
		// Handle the out-of-range error, e.g., return an empty ArtistInfo struct or log an error message
		return gpd.ArtistInfo{}
	}

	dateInfo := allData.Date[index]
	artistDetails := gpd.ArtistInfo{}
	artistDetails.Artist = allData.Artist[index]
	artistDetails.Location = allData.Location[index]
	artistDetails.All = allData.Locs[index]

	for i := 0; i < len(dateInfo.Dates); i++ {
		if string(dateInfo.Dates[i][0]) == "*" {
			dateInfo.Dates[i] = dateInfo.Dates[i][1:]
		}
	}

	artistDetails.Date = dateInfo
	artistDetails.Maps = gpm.CreateMap(allData, index)

	return artistDetails
}
