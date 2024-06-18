package main

import (
	"groupie/pkg/funcs"
	"log"
	"net/http"
)

var data []funcs.Band

func main() {
	apiURL := "https://groupietrackers.herokuapp.com/api/artists"
	data = funcs.Gather(apiURL)
	funcs.SetData(data)
	http.HandleFunc("/suggestions", funcs.Suggestions)
	http.HandleFunc("/info", funcs.InfoPage)
	http.HandleFunc("/band/", funcs.BandDetails)
	http.HandleFunc("/", funcs.HomePage)
	http.HandleFunc("/style.css", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "../../web/html/style.css")
	})
	log.Println("Server listening on port http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
