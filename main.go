package main

import (
	"log"
	"net/http"

	"github.com/viralkansarav/country-search/config"
	"github.com/viralkansarav/country-search/handlers"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/api/countries/search", handlers.SearchCountryHandler).Methods("GET")

	log.Printf("Server running on port %s", config.Port)
	http.ListenAndServe(":"+config.Port, r)
}
