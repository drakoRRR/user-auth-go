package ipgeolocation

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

const IpApiUrl = "http://ip-api.com/json/"

type IPResponse struct {
	Country string `json:"country"`
}

func GetCountry(ipAddress string) (string, error) {
	requestURL := IpApiUrl + ipAddress
	client := &http.Client{Timeout: 10 * time.Second}
	res, err := client.Get(requestURL)
	if err != nil {
		return "", fmt.Errorf("error making request: %w", err)
	}
	defer res.Body.Close()

	var ipResponse IPResponse
	if err := json.NewDecoder(res.Body).Decode(&ipResponse); err != nil {
		return "", fmt.Errorf("error decoding response: %w", err)
	}

	return ipResponse.Country, nil
}
