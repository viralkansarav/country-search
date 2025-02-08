# Country Search API (Golang)

## Setup
1. Install Go
2. Clone repository
3. Run `go mod tidy`
4. Start server: `go run main.go`

## API Usage
`GET /api/countries/search?name=India`

Response:
```json
{
  "name": "India",
  "capital": "New Delhi",
  "currency": "₹",
  "population": 1380004385
}
