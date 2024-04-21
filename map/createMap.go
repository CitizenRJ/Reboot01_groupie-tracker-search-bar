package groupie

import (
	"encoding/json"
	gpd "groupie/data"
	"io"
	"net/http"
	"strconv"
)

// CreateMap generates a map URL with pins for the given locations.
// It takes the full data set and an index to select the desired location group.
// It geocodes each location name into latitude/longitude coordinates.
// It then constructs a map URL using the map API that pins each location.
// The returned URL can be used to display an interactive map of the locations.
func CreateMap(allData gpd.Data, index int) string {
	var featureCollections []gpd.FeatureCollection
	mapURL := ""

	for _, location := range allData.Location[index].Locations {

		var featureCollection gpd.FeatureCollection
		locationName := ""
		apiURL := ""
		hasDash := false
		for _, letter := range location {
			if string(letter) != "-" && !hasDash {
				locationName += string(letter)
			} else if string(letter) == "-" {
				hasDash = true
			}
		}
		apiURL = "MAP API LINK" + locationName + "MAP API TOKEN" // insert map api plz

		response, _ := http.Get(apiURL)
		responseData, _ := io.ReadAll(response.Body)
		json.Unmarshal(responseData, &featureCollection)

		featureCollections = append(featureCollections, featureCollection)
	}

	mapURL = "MAP API LINK" + "MAP API TOKEN" // insert map api plz

	for i, feature := range featureCollections {
		longitude := strconv.FormatFloat(feature.Features[0].Center[0], 'g', 9, 32)
		latitude := strconv.FormatFloat(feature.Features[0].Center[1], 'g', 9, 32)
		if i == len(featureCollections)-1 {
			mapURL += "Insert map pin" + longitude + "," + latitude + ")" + "map functions token zoom etc" // insert map api plz
		} else {
			mapURL += "Insert map pin" + longitude + "," + latitude + ")," // insert map api plz
		}
	}

	return mapURL
}
