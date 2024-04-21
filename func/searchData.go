package groupie

import (
	gpd "groupie/data"
)

// SearchData filters the dataset based on the search term and search type.
// It returns a subset of the data where the search term matches any artist's name,
// member name, or location, depending on the search type specified.
// If searchType is 0, it performs a general search across multiple fields.
// If searchType is non-zero, it assumes the value represents a creation date and filters accordingly.
// The function ensures no duplicate artists are included in the results.
func SearchData(searchTerm string, searchType int, allData gpd.Data) gpd.Data {
	var dataSearch gpd.Data
	bandMembers := []string{}
	if searchType == 0 {
		for i := 0; i < (len(allData.Artist)); i++ {
			for _, location := range allData.Location[i].Locations {
				if (location == searchTerm || allData.Artist[i].Name == searchTerm || allData.Artist[i].First_album == searchTerm) && !Isin(allData.Artist[i].Name, bandMembers) {
					dataSearch.Artist = DisplayData(i, dataSearch, allData)
					bandMembers = append(bandMembers, allData.Artist[i].Name)
				}
			}
			for _, member := range allData.Artist[i].Members {
				if member == searchTerm && !Isin(member, bandMembers) {
					dataSearch.Artist = DisplayData(i, dataSearch, allData)
					bandMembers = append(bandMembers, allData.Artist[i].Name)
				}
			}
		}
	}

	if searchType != 0 {
		for i := 0; i < len(allData.Artist); i++ {
			if allData.Artist[i].Creation_date == searchType && !Isin(allData.Artist[i].Name, bandMembers) {
				dataSearch.Artist = DisplayData(i, dataSearch, allData)
				bandMembers = append(bandMembers, allData.Artist[i].Name)
			}
		}
	}

	dataSearch.All = allData.All
	dataSearch.Country = allData.Country
	return dataSearch
}
