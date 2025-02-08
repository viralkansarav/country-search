package tests

import (
	"testing"

	"github.com/viralkansarav/country-search/services"
)

func TestGetCountryInfo(t *testing.T) {
	_, err := services.GetCountryInfo("India")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}

func TestGetCountryInfo_Cache(t *testing.T) {
	services.GetCountryInfo("India")           // First call to populate cache
	_, err := services.GetCountryInfo("India") // Should hit cache
	if err != nil {
		t.Errorf("Expected cached response, got error: %v", err)
	}
}
