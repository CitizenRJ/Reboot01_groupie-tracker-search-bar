package main

import (
	"fmt"
	gpi "groupie/internal/api"
	gpd "groupie/internal/models"
	gpf "groupie/internal/utils"
	"html/template"
	"net/http"
	"path/filepath"
	"runtime"
	"strconv"
)

// global variables that holds data from the GetData function.
var allData gpd.Data
var date gpd.Date
var artists []gpd.Artists
var getLocation gpd.GetLocation
var relations gpd.Relations

// main is the entry point for the server application. It sets up the necessary data, configures the HTTP server,
// and starts listening for incoming requests on port 8080.
// The main function performs the following tasks:
// 1. Retrieves data from the gpi package and stores it in the allData variable.
// 2. Prints a message indicating that the server is starting on localhost:8080.
// 3. Registers HTTP handlers for the following routes:
//   - "/" and "/home": Handles the index page.
//   - "/search": Handles the search functionality.
//   - "/filter": Handles the filtering functionality.
//   - "/info": Handles the information page.
//   - "/style.css": Serves the CSS file for the web application.
//
// 4. Starts the HTTP server and listens for incoming requests.
// 5. If an error occurs while starting the server, it prints the error message and returns.
func main() {
	date, artists, getLocation, relations = gpi.GetData()
	allData = gpi.SetData(date, artists, getLocation, relations)
	const port = ":8080"
	fmt.Println("Starting server at http://localhost" + port)
	http.HandleFunc("/", handleIndex)
	http.HandleFunc("/home", handleIndex)
	http.HandleFunc("/search", handleSearch)
	http.HandleFunc("/filter", handleFilter)
	http.HandleFunc("/info", handleInfos)
	http.HandleFunc("/style.css", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "../../web/templates/style.css")
	})
	err := http.ListenAndServe(port, nil)
	if err != nil {
		fmt.Printf("Error starting server: %v", err)
		return
	}
}

// handleIndex is an HTTP handler function that renders the home.html template with the allData
// data structure. If there is an error parsing the template or executing it, the function
// will return a 500 Internal Server Error response.
func handleIndex(w http.ResponseWriter, r *http.Request) {
	var templateInstance *template.Template
	templatePath, err := getTemplatePath("home.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	templateInstance = template.Must(template.ParseFiles(templatePath))
	err = templateInstance.Execute(w, allData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// handleSearch handles the search functionality for the application.
// It takes an HTTP request with a "input" form value, parses the query,
// and uses the gpf.SearchData function to retrieve a resultSet of matching
// data. It then renders the "home.html" template with the resultSet.
// If the resultSet is empty, it still renders the "home.html" template.
// Any errors that occur during template parsing or execution are
// returned as an HTTP 500 Internal Server Error.
func handleSearch(w http.ResponseWriter, r *http.Request) {
	query := r.FormValue("input")
	parsedQuery, _ := strconv.Atoi(query)
	resultSet := gpf.SearchData(query, parsedQuery, allData)

	var pageTemplate *template.Template
	if len(resultSet.Artist) == 0 {
		templatePath, err := getTemplatePath("home.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		pageTemplate = template.Must(template.ParseFiles(templatePath))
	} else {
		templatePath, err := getTemplatePath("home.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		pageTemplate = template.Must(template.ParseFiles(templatePath))
	}
	err := pageTemplate.Execute(w, resultSet)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// handleFilter is an HTTP handler function that processes filter parameters from a request and
// generates a filtered data set, which is then rendered using a template.
// The function extracts various filter parameters from the request, such as member buttons,
// city, creation date, and album date. It then calls the gpf.FilterData function to generate
// a filtered data set based on these parameters.
// If the filtered data set is empty, the function renders the "home.html" template. Otherwise,
// it renders the "home.html" template with the filtered data set.
// The function returns an HTTP error if there is an issue parsing the request parameters or
// rendering the template.
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
	if len(filteredData.Artist) == 0 {
		templatePath, err := getTemplatePath("home.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		templateInstance = template.Must(template.ParseFiles(templatePath))
	} else {
		templatePath, err := getTemplatePath("home.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		templateInstance = template.Must(template.ParseFiles(templatePath))
	}
	err := templateInstance.Execute(w, filteredData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// handleInfos is an HTTP handler function that retrieves artist information and renders an HTML template with the information.
// The function expects a request parameter "id" which is used to look up the artist information.
// The artist information is then passed to an HTML template for rendering.
// If there is an error retrieving the template path or rendering the template, the function will return a 500 Internal Server Error response.
func handleInfos(w http.ResponseWriter, r *http.Request) {
	identifier := r.FormValue("id")
	artistIdentifier, _ := strconv.Atoi(identifier)
	artistIdentifier = artistIdentifier - 1

	artistInformation := gpf.InfoArtist(allData, artistIdentifier)

	var templatePointer *template.Template
	templatePath, err := getTemplatePath("info.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	templatePointer = template.Must(template.ParseFiles(templatePath))
	err = templatePointer.Execute(w, artistInformation)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// getTemplatePath returns the absolute path to the specified template file.
// It determines the current directory of the calling code, and then constructs the
// path to the template file relative to the current directory.
// If there is an error resolving the absolute path, an error is returned.
func getTemplatePath(templateFile string) (string, error) {
	_, currentFile, _, ok := runtime.Caller(0)
	if !ok {
		return "", fmt.Errorf("failed to get current file path")
	}
	currentDir := filepath.Dir(currentFile)
	templatePath := filepath.Join(currentDir, "../../web/templates", templateFile)
	absolutePath, err := filepath.Abs(templatePath)
	if err != nil {
		return "", fmt.Errorf("failed to resolve absolute path: %v", err)
	}
	return absolutePath, nil
}
