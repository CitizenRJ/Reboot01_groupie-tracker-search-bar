package groupie

import gpd "groupie/internal/models"

// DisplayData takes an index, the current data, and all data, and returns a slice of Artists.
// It creates a new Artist struct with the name, image, and ID from the all data at the given index,
// appends it to the current data's Artist slice, and returns the updated current data's Artist slice.
func DisplayData(index int, currentData gpd.Data, allData gpd.Data) []gpd.Artists {
	var artist gpd.Artists
	artist.Name = allData.Artist[index].Name
	artist.Image = allData.Artist[index].Image
	artist.Id = allData.Artist[index].Id

	currentData.Artist = append(currentData.Artist, artist)

	return currentData.Artist
}
