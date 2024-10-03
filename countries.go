package holidaylist

import (
    "encoding/json"
    "log"
    "net/http"
)

// Country represents a single country returned by the API
type Country struct {
    Code string `json:"code"`
    Name string `json:"name"`
    Flag string `json:"flag"`
}

// CountryResponse represents the structure of the response from the API for the countries endpoint
type CountryResponse struct {
    Status int       `json:"status"`
    Data   []Country `json:"data"`
}

// GetCountries fetches the list of countries from the API
func (a *API) GetCountries() (CountryResponse, error) {
    endpoint := "https://back.holidaylist.io/api/v1/countries"
    reqURL := endpoint + "?key=" + a.apiKey

    resp, err := http.Get(reqURL)
    if err != nil {
        log.Println("Error making request:", err)
        return CountryResponse{}, err
    }
    defer resp.Body.Close()

    var response CountryResponse
    err = json.NewDecoder(resp.Body).Decode(&response)
    if err != nil {
        log.Println("Error decoding response:", err)
        return CountryResponse{}, err
    }

    response.Status = resp.StatusCode
    return response, nil
}
