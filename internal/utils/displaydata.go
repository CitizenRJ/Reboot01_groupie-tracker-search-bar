package groupie

import gpd "groupie/internal/models"

func DisplayData(index int, currentData gpd.Data, allData gpd.Data) []gpd.Artists {
	var Artist gpd.Artists
	Artist.Name = allData.Artist[index].Name
	Artist.Image = allData.Artist[index].Image
	Artist.Id = allData.Artist[index].Id
	currentData.Artist = append(currentData.Artist, Artist)
	return currentData.Artist
}
