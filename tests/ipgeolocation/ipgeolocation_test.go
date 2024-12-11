package ipgeolocation

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/drakoRRR/user-auth-go/internal/ipgeolocation"
	"github.com/stretchr/testify/assert"
)

func mockIPAPIResponse(w http.ResponseWriter, r *http.Request) {
	ipResponse := ipgeolocation.IPResponse{
		Country: "US",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(ipResponse)
}

func TestGetCountry_Success(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(mockIPAPIResponse))
	defer server.Close()

	ipgeolocation.IpApiUrl = server.URL + "/json/"

	ipAddress := "8.8.8.8"
	country, err := ipgeolocation.GetCountry(ipAddress)

	assert.NoError(t, err)
	assert.Equal(t, "US", country)
}

func TestGetCountry_ErrorOnRequest(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}))
	defer server.Close()

	ipgeolocation.IpApiUrl = server.URL + "/json/"

	ipAddress := "8.8.8.8"
	country, err := ipgeolocation.GetCountry(ipAddress)

	assert.Error(t, err)
	assert.Empty(t, country)
}

func TestGetCountry_ErrorOnDecode(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "invalid json")
	}))
	defer server.Close()

	ipgeolocation.IpApiUrl = server.URL + "/json/"

	ipAddress := "8.8.8.8"
	country, err := ipgeolocation.GetCountry(ipAddress)

	assert.Error(t, err)
	assert.Empty(t, country)
	assert.Contains(t, err.Error(), "error decoding response")
}
