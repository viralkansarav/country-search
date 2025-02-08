package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/viralkansarav/country-search/services"
)

// handler to handle get request and perform validation over incoming request
func SearchCountryHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		http.Error(w, "Missing country name", http.StatusBadRequest)
		return
	}
	country, err := services.GetCountryInfo(name)
	if err != nil {
		http.Error(w, "Country not found", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(country)
}
