package main

import (
	"fmt"
	gp "groupie/internal/api"
	gpd "groupie/internal/models"
	gpf "groupie/internal/utils"
	"html/template"
	"net/http"
	"path/filepath"
	"runtime"
	"strconv"
)

var allData gpd.Data
var date gpd.Date
var artists []gpd.Artists
var getLocation gpd.GetLocation
var relations gpd.Relations

func main() {
	date, artists, getLocation, relations = gp.GetData()
	allData = gp.SetData(date, artists, getLocation, relations)
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
