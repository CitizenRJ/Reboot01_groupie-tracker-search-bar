package funcs

import (
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func HomePage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		NotFound(w, r)
		return
	}
	if r.Method != "GET" {
		MethodNotAllowed(w, r)
		return
	}

	// Pass the data to the template
	tmpl, err := template.ParseFiles("../../web/html/home.html")
	if err != nil {
		InternalServer(w, r)
		return
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		InternalServer(w, r)
		return
	}
}

func BandDetails(w http.ResponseWriter, r *http.Request) {
	// Extract the band ID from the URL
	bandID := r.URL.Path[len("/band/"):]

	// Find the band with the matching ID
	var band Band
	for _, b := range data {
		if strconv.Itoa(b.Id) == bandID {
			band = b
			break
		}
	}

	// Render the band details template
	tmpl, err := template.ParseFiles("../../web/html/band.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, band)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func InfoPage(w http.ResponseWriter, r *http.Request) {
	query := r.FormValue("q")
	log.Println("Search query:", query)

	if query == "" {
		log.Println("Empty search query, redirecting to homepage")
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	// Split the search query into value and search type
	parts := strings.Split(query, " - ")
	value := parts[0]
	searchType := "location"
	if len(parts) > 1 {
		searchType = strings.TrimSpace(parts[1])
	}

	var bands []Band
	for _, b := range data {
		switch searchType {
		case "location":
			if containsLocation(b.Concerts, value) {
				bands = append(bands, b)
			}
		case "creation date":
			if strconv.FormatUint(uint64(b.CreationDate), 10) == value {
				bands = append(bands, b)
			}
		case "first album":
			if strings.EqualFold(b.FirstAlbum, value) {
				bands = append(bands, b)
			}
		case "artist/band":
			if strings.EqualFold(b.Name, value) {
				bands = append(bands, b)
			}
		case "member":
			for _, member := range b.Members {
				if strings.EqualFold(member, value) {
					bands = append(bands, b)
					break
				}
			}
		}
	}

	log.Println("Bands matching the search query:", bands)
	log.Println("Search type:", searchType)

	if len(bands) == 0 {
		log.Println("No bands found for the search query")
		http.NotFound(w, r)
		return
	}

	var resultType string
	if searchType == "artist/band" {
		if len(bands[0].Members) == 1 {
			resultType = "artist"
		} else {
			resultType = "band"
		}
	} else {
		resultType = searchType
	}

	tmpl, err := template.ParseFiles("../../web/html/info.html")
	if err != nil {
		log.Println("Error parsing template:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, map[string]interface{}{
		"Query":      value,
		"Bands":      bands,
		"SearchType": resultType,
	})
	if err != nil {
		log.Println("Error executing template:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Println("Info page rendered successfully")
}

func containsLocation(concerts map[string][]string, location string) bool {
	for l := range concerts {
		if strings.EqualFold(l, location) {
			return true
		}
	}
	return false
}
