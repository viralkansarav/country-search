package clients

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"time"
)

// CountryResponse represents the structure of the response received from the REST Countries API
// It contains nested fields to correctly parse API response data
type CountryResponse struct {
	Name struct {
		Common string `json:"common"`
	} `json:"name"`
	Capital    []string `json:"capital"`
	Currencies map[string]struct {
		Name   string `json:"name"`
		Symbol string `json:"symbol"`
	} `json:"currencies"`
	Population int `json:"population"`
}

// Country represents the formatted response to be sent to API consumers
type Country struct {
	Name       string `json:"name"`
	Capital    string `json:"capital"`
	Currency   string `json:"currency"`
	Population int    `json:"population"`
}

// HTTP client with timeout to prevent slow network requests
var httpClient = &http.Client{
	Timeout: 5 * time.Second,
}

// FetchCountryData retrieves country details from the REST Countries API
// If the API response is valid, it extracts relevant fields and returns a formatted Country object
func FetchCountryData(name string) (*Country, error) {
	url := "https://restcountries.com/v3.1/name/" + name
	resp, err := httpClient.Get(url)
	if err != nil {
		log.Printf("HTTP request failed: %v", err)
		return nil, errors.New("failed to fetch country data")
	}
	defer resp.Body.Close()

	// Validate HTTP response status
	if resp.StatusCode != http.StatusOK {
		log.Printf("Invalid API response status: %d", resp.StatusCode)
		return nil, errors.New("invalid response from API")
	}

	var countries []CountryResponse
	if err := json.NewDecoder(resp.Body).Decode(&countries); err != nil || len(countries) == 0 {
		log.Printf("JSON decoding error or empty response: %v", err)
		return nil, errors.New("invalid response format")
	}

	// Extract data safely, handling missing fields gracefully
	country := countries[0]
	capital := "Unknown"
	if len(country.Capital) > 0 {
		capital = country.Capital[0]
	}

	currency := "Unknown"
	for _, c := range country.Currencies {
		currency = c.Symbol + " (" + c.Name + ")"
		break // Pick the first available currency
	}

	return &Country{
		Name:       country.Name.Common,
		Capital:    capital,
		Currency:   currency,
		Population: country.Population,
	}, nil
}
