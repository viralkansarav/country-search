package services

import (
	"time"

	"github.com/viralkansarav/country-search/cache"
	"github.com/viralkansarav/country-search/clients"
)

var countryCache = cache.NewCache()

// Gets the country info from cache if availaible, if not then stores the data into cache for next 10 minutes to ensure the data in cache should be updated. calls the main api to get resp.
func GetCountryInfo(name string) (*clients.Country, error) {
	if data, found := countryCache.Get(name); found {
		return data.(*clients.Country), nil
	}
	country, err := clients.FetchCountryData(name)
	if err != nil {
		return nil, err
	}
	countryCache.Set(name, country, 10*time.Minute)
	return country, nil
}
