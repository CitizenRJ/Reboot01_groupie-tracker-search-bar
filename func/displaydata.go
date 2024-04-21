package groupie

import gpd "groupie/data"

// DisplayData adds the artist data at the given index from allData to currentData.
// It copies the name, image, and id fields from the artist at index in allData
// to a new artist struct, appends that struct to currentData.Artist,
// and returns the updated currentData.Artist slice.
func DisplayData(index int, currentData gpd.Data, allData gpd.Data) []gpd.Artists {
	var Artist gpd.Artists
	Artist.Name = allData.Artist[index].Name
	Artist.Image = allData.Artist[index].Image
	Artist.Id = allData.Artist[index].Id
	currentData.Artist = append(currentData.Artist, Artist)
	return currentData.Artist
}
