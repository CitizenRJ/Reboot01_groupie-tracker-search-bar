package funcs

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func FetchAPI(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func Gather(url string) []Band {
	data, err := FetchAPI(url)
	if err != nil {
		log.Println("Error fetching data:", err)
		return nil
	}

	var art []Band
	err = json.Unmarshal(data, &art)
	if err != nil {
		log.Println("Error unmarshaling data:", err)
		return nil
	}

	log.Println("Fetched bands:", art)

	for i := 0; i < len(art); i++ {
		rel := Relation{}
		date := Date{}
		Location := Location{}
		art[i].Type = "artist/band"

		relData, err := FetchAPI(art[i].Relations)
		if err != nil {
			log.Println("Error fetching relation data:", err)
			continue
		}

		dateData, err := FetchAPI(art[i].DatesLink)
		if err != nil {
			log.Println("Error fetching date data:", err)
			continue
		}

		LocationData, err := FetchAPI(art[i].LocationsLink)
		if err != nil {
			log.Println("Error fetching location data:", err)
			continue
		}

		err = json.Unmarshal(relData, &rel)
		if err != nil {
			log.Println("Error unmarshaling relation data:", err)
			continue
		}

		err = json.Unmarshal(dateData, &date)
		if err != nil {
			log.Println("Error unmarshaling date data:", err)
			continue
		}

		err = json.Unmarshal(LocationData, &Location)
		if err != nil {
			log.Println("Error unmarshaling location data:", err)
			continue
		}

		log.Println("Relation data:", rel)
		log.Println("Date data:", date)
		log.Println("Location data:", Location)

		art[i].Concerts = make(map[string][]string)
		for j := 0; j < len(Location.Locations); j++ {
			if date.Date[j] != "" {
				art[i].Concerts[Location.Locations[j]] = append(art[i].Concerts[Location.Locations[j]], date.Date[j])
			}
		}

		log.Println("Processed band:", art[i])
	}

	return art
}
