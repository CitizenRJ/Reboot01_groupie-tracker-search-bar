package groupie

import (
	gpm "groupie/internal/maps"
	gpd "groupie/internal/models"
)

// InfoArtist returns an ArtistInfo struct containing information about an artist at the given index in the provided Data.
// It extracts the artist name, location, and all locations from the Data, and creates a map of the artist's information.
// It also processes the date information, removing any leading "*" characters.
func InfoArtist(allData gpd.Data, index int) gpd.ArtistInfo {
	dateInfo := allData.Date[index]
	artistDetails := gpd.ArtistInfo{}

	artistDetails.Artist = allData.Artist[index]
	artistDetails.Location = allData.Location[index]
	artistDetails.All = allData.Locs[index]

	artistDetails.Date = processDateInfo(dateInfo)
	artistDetails.Maps = gpm.CreateMap(allData, index)

	return artistDetails
}

func processDateInfo(dateInfo struct {
	Id    int      `json:"id"`
	Dates []string `json:"dates"`
}) struct {
	Id    int      `json:"id"`
	Dates []string `json:"dates"`
} {
	for i := 0; i < len(dateInfo.Dates); i++ {
		if string(dateInfo.Dates[i][0]) == "*" {
			dateInfo.Dates[i] = dateInfo.Dates[i][1:]
		}
	}

	return dateInfo
}
