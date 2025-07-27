package sources

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"delphinium/internal"
)

const baseURL = "https://www.whoisxmlapi.com/whoisserver/WhoisService"

type WhoisResponse struct {
	// desired data from WHOIS
	WhoisRecord struct {
		CreatedDate   string `json:"createdDate"`
		UpdatedDate   string `json:"updatedDate"`
		ExpiresDate   string `json:"expiresDate"`
		RegistrarName string `json:"registrarName"`
	} `json:"WhoisRecord"`
}

func FetchWHOIS(domain string) []internal.Event {
	apiKey := os.Getenv("WHOIS_API_KEY") // can be hardcoded variable as well

	url := fmt.Sprintf("%s?apiKey=%s&domainName=%s&outputFormat=JSON", baseURL, apiKey, domain)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error fetching WHOIS data: %v\n", err)
		return nil
	}
	defer resp.Body.Close()

	var data WhoisResponse
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		fmt.Printf("Error parsing WHOIS JSON: %v\n", err)
		return nil
	}

	events := []internal.Event{}

	parseAndAppend := func(dateStr, title string) {
		if dateStr == "" {
			return
		}

		layout := "2006-01-02T15:04:05Z0700" // specified layout to ensure parse
		if t, err := time.Parse(layout, dateStr); err == nil {
			events = append(events, internal.Event{
				Source:    "WHOIS",
				Title:     title,
				Detail:    fmt.Sprintf("WHOIS %s for %s", title, domain),
				Timestamp: t,
			})
		}
	}

	parseAndAppend(data.WhoisRecord.CreatedDate, "Domain Created")
	parseAndAppend(data.WhoisRecord.UpdatedDate, "WHOIS Updated")
	parseAndAppend(data.WhoisRecord.ExpiresDate, "Domain Expiry")

	return events
}
