package funcs

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

func Suggestions(w http.ResponseWriter, r *http.Request) {
	query := r.FormValue("q")
	if query == "" {
		json.NewEncoder(w).Encode([]map[string]string{})
		return
	}

	query = strings.ToLower(query)

	var suggestions []map[string]string
	for _, band := range data {
		if strings.Contains(strings.ToLower(band.Name), query) {
			suggestions = append(suggestions, map[string]string{
				"name": band.Name,
				"type": "artist/band",
			})
		}
		for _, member := range band.Members {
			if strings.Contains(strings.ToLower(member), query) {
				suggestions = append(suggestions, map[string]string{
					"name": member,
					"type": "member",
				})
			}
		}
		for location := range band.Concerts {
			if strings.Contains(strings.ToLower(location), query) {
				suggestions = append(suggestions, map[string]string{
					"name": location,
					"type": "location",
				})
			}
		}
		if strings.Contains(strings.ToLower(band.FirstAlbum), query) {
			suggestions = append(suggestions, map[string]string{
				"name": band.FirstAlbum,
				"type": "first album",
			})
		}
		if strings.Contains(strings.ToLower(strconv.FormatUint(uint64(band.CreationDate), 10)), query) {
			suggestions = append(suggestions, map[string]string{
				"name": strconv.FormatUint(uint64(band.CreationDate), 10),
				"type": "creation date",
			})
		}
	}

	// Remove duplicate suggestions
	uniqueSuggestions := make([]map[string]string, 0, len(suggestions))
	seen := make(map[string]bool)
	for _, suggestion := range suggestions {
		key := suggestion["name"] + "-" + suggestion["type"]
		if !seen[key] {
			seen[key] = true
			uniqueSuggestions = append(uniqueSuggestions, suggestion)
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(uniqueSuggestions)
}
