package holidaylist

import (
    "encoding/json"
    "net/http"
    "log"
)

// GetCountries fetches the list of countries from the API
func (a *API) GetCountries() ([]Country, error) {
    endpoint := "https://back.holidaylist.io/api/v1/countries"
    reqURL := endpoint + "?key=" + a.apiKey

    resp, err := http.Get(reqURL)
    if err != nil {
        log.Println("Error making request:", err)
        return nil, err
    }
    defer resp.Body.Close()

    var response CountryResponse
    err = json.NewDecoder(resp.Body).Decode(&response)
    if err != nil {
        log.Println("Error decoding response:", err)
        return nil, err
    }

    return response.Data, nil
}
