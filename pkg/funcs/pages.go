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

	parts := strings.Split(query, " - ")
	value := strings.ToLower(parts[0])
	searchType := "all" // Default to searching all fields
	if len(parts) > 1 {
		searchType = strings.TrimSpace(parts[1])
	}

	var results struct {
		Query      string
		SearchType string
		Bands      []Band
		Locations  []string
		Albums     []string
		Members    []string
		Dates      []uint
	}

	results.Query = value
	results.SearchType = searchType

	for _, b := range data {
		switch searchType {
		case "location":
			for loc := range b.Concerts {
				if strings.Contains(strings.ToLower(loc), value) {
					results.Locations = append(results.Locations, loc)
					results.Bands = append(results.Bands, b)
				}
			}
		case "creation date":
			if strconv.FormatUint(uint64(b.CreationDate), 10) == value {
				results.Dates = append(results.Dates, b.CreationDate)
				results.Bands = append(results.Bands, b)
			}
		case "first album":
			if strings.Contains(strings.ToLower(b.FirstAlbum), value) {
				results.Albums = append(results.Albums, b.FirstAlbum)
				results.Bands = append(results.Bands, b)
			}
		case "artist/band":
			if strings.Contains(strings.ToLower(b.Name), value) {
				results.Bands = append(results.Bands, b)
			}
		case "member":
			for _, member := range b.Members {
				if strings.Contains(strings.ToLower(member), value) {
					results.Members = append(results.Members, member)
					results.Bands = append(results.Bands, b)
				}
			}
		case "all":
			if strings.Contains(strings.ToLower(b.Name), value) {
				results.Bands = append(results.Bands, b)
			}
			for loc := range b.Concerts {
				if strings.Contains(strings.ToLower(loc), value) {
					results.Locations = append(results.Locations, loc)
					results.Bands = append(results.Bands, b)
				}
			}
			if strconv.FormatUint(uint64(b.CreationDate), 10) == value {
				results.Dates = append(results.Dates, b.CreationDate)
				results.Bands = append(results.Bands, b)
			}
			if strings.Contains(strings.ToLower(b.FirstAlbum), value) {
				results.Albums = append(results.Albums, b.FirstAlbum)
				results.Bands = append(results.Bands, b)
			}
			for _, member := range b.Members {
				if strings.Contains(strings.ToLower(member), value) {
					results.Members = append(results.Members, member)
					results.Bands = append(results.Bands, b)
				}
			}
		
		}
	}

	log.Printf("Search results: %+v", results)

	funcMap := template.FuncMap{
		"join": strings.Join,
	}

	tmpl, err := template.New("info.html").Funcs(funcMap).ParseFiles("../../web/html/info.html")
	if err != nil {
		log.Println("Error parsing template:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, results)
	if err != nil {
		log.Println("Error executing template:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Println("Info page rendered successfully")
}
