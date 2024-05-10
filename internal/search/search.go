package groupie

import (
    gpd "groupie/internal/models"
    "strconv"
    "strings"
)

func GetAll(data gpd.Data) []string {
    artistNames := []string{}
    memberNames := []string{}
    positions := []string{}
    firstAlbums := []string{}
    creationDates := []string{}
    combinedList := []string{}
    uniqueMap := make(map[string]bool)

    for _, artist := range data.Artist {
        processedName := strings.Replace(artist.Name, " ", "/", -1)
        artistNames = append(artistNames, processedName)
        firstAlbums = append(firstAlbums, artist.First_album)
        creationDates = append(creationDates, strconv.Itoa(artist.Creation_date))
        for _, member := range artist.Members {
            processedMember := strings.Replace(member, " ", "/", -1)
            memberNames = append(memberNames, processedMember)
        }
    }

    for _, location := range data.Location {
        positions = append(positions, location.Locations...)
    }

    combinedList = append(combinedList, strings.Join(artistNames, ""))
    combinedList = append(combinedList, strings.Join(memberNames, ""))
    combinedList = append(combinedList, strings.Join(positions, ""))
    combinedList = append(combinedList, firstAlbums...)
    combinedList = append(combinedList, creationDates...)

    for _, element := range combinedList {
        uniqueMap[element] = true
    }

    uniqueList := make([]string, 0, len(uniqueMap))
    for element := range uniqueMap {
        uniqueList = append(uniqueList, element)
    }

    return uniqueList
}
