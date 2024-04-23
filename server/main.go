package main

import (
	"fmt"
	gp "groupie/api"
	gpd "groupie/data"
	gpf "groupie/func"
	"html/template"
	"net/http"
	"strconv"
)

// Initializes the data set and starts the HTTP server.
// date, artists, getLocation, and relations are initialized by calling gp.GetData().
// allData is initialized by calling gp.SetData() with the data returned from gp.GetData().
// Routes and handlers are set up for /, /search, /filter, /infos.
// The file server is set up to serve static files from /static.
// The HTTP server is started on port 8080.
var allData gpd.Data
var date gpd.Date
var artists []gpd.Artists
var getLocation gpd.GetLocation
var relations gpd.Relations

func main() {
	date, artists, getLocation, relations = gp.GetData()
	allData = gp.SetData(date, artists, getLocation, relations)
	fmt.Println("Starting server on port 8080")
	http.HandleFunc("/", handleIndex)
	http.HandleFunc("/search", handleSearch)
	http.HandleFunc("/filter", handleFilter)
	http.HandleFunc("/info", handleInfos)
	http.HandleFunc("/style.css", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "../pages/style.css")
	})
	fs := http.FileServer(http.Dir("./"))
	http.Handle("/static/", http.StripPrefix("/", fs))
	http.ListenAndServe(":8080", nil)
}

// handleIndex handles requests to the root path "/".
// It executes the html template, passing the allData struct as data.
func handleIndex(responseWriter http.ResponseWriter, request *http.Request) {
	var templateInstance *template.Template
	templateInstance = template.Must(template.ParseFiles("../pages/home.html")) //home html page
	templateInstance.Execute(responseWriter, allData)
}

// handleSearch handles requests to /search
// It gets the search query from the request form value "input"
// It converts the query to an int if possible using Atoi
// It calls SearchData to get search results for the query
// It chooses a template based on if results were found
// It executes the template passing the results or a "no results" page
func handleSearch(w http.ResponseWriter, r *http.Request) {
	query := r.FormValue("input")
	parsedQuery, _ := strconv.Atoi(query)
	resultSet := gpf.SearchData(query, parsedQuery, allData)
	var pageTemplate *template.Template
	if len(resultSet.Artist) == 0 {
		pageTemplate = template.Must(template.ParseFiles("../pages/home.html")) // no results page
	} else {
		pageTemplate = template.Must(template.ParseFiles("../pages/home.html")) // artists page
	}
	pageTemplate.Execute(w, resultSet)
}

// handleFilter handles requests to /filter
// It gets filter values from the request form
// It calls FilterData to filter the data set
// It chooses a template based on if results were found
// It executes the template passing the filtered data or no results page
func handleFilter(w http.ResponseWriter, r *http.Request) {
	allMembersButton := r.FormValue("MemberAll")
	cityFilter := r.FormValue("city")
	memberButton1, _ := strconv.Atoi(r.FormValue("Member1"))
	memberButton2, _ := strconv.Atoi(r.FormValue("Member2"))
	memberButton3, _ := strconv.Atoi(r.FormValue("Member3"))
	memberButton4, _ := strconv.Atoi(r.FormValue("Member4"))
	memberButton5, _ := strconv.Atoi(r.FormValue("Member5"))
	memberButton6, _ := strconv.Atoi(r.FormValue("Member6"))
	memberButton7, _ := strconv.Atoi(r.FormValue("Member7"))
	memberButton8, _ := strconv.Atoi(r.FormValue("Member8"))
	creationDateFilter, _ := strconv.Atoi(r.FormValue("creationdate"))
	albumDateFilter, _ := strconv.Atoi(r.FormValue("albumdate"))
	var tabButton []int
	tabButton = append(tabButton, memberButton1, memberButton2, memberButton3, memberButton4, memberButton5, memberButton6, memberButton7, memberButton8)
	filteredData := gpf.FilterData(allData, allMembersButton, tabButton, cityFilter, creationDateFilter, albumDateFilter)
	var templateInstance *template.Template
	templateInstance = template.Must(template.ParseFiles("../pages/home.html")) // no results page
	if len(filteredData.Artist) == 0 {
		templateInstance = template.Must(template.ParseFiles("../pages/home.html")) // artists page
	}
	templateInstance.Execute(w, filteredData)
}

// handleInfos handles requests to /info
// It gets the artist identifier from the request form value "id"
// It converts the identifier to an int
// It calls InfoArtist to get artist information for the given identifier
// It executes the info template passing the artist information
func handleInfos(responseWriter http.ResponseWriter, request *http.Request) {
	identifier := request.FormValue("id")
	artistIdentifier, _ := strconv.Atoi(identifier)
	artistIdentifier = artistIdentifier - 1
	artistInformation := gpf.InfoArtist(allData, artistIdentifier)
	var templatePointer *template.Template
	templatePointer = template.Must(template.ParseFiles("../pages/info.html")) // info html page
	templatePointer.Execute(responseWriter, artistInformation)
}