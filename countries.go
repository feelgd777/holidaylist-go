package holidaylist

import (
    "encoding/json"
    "fmt"
    "net/http"
    "net/url"
)

// GetCountries fetches the list of countries, allowing optional parameters to be passed
func (a *API) GetCountries(params map[string]interface{}) (CountryResponse, error) {
    endpoint := "https://back.holidaylist.io/api/v1/countries"
    
    // Build the query parameters from the passed JSON object
    reqURL, _ := url.Parse(endpoint)
    query := reqURL.Query()
    query.Add("key", a.apiKey)
    
    // Handle nil or optional parameters
    if params != nil {
        for key, value := range params {
            query.Add(key, fmt.Sprintf("%v", value)) // Convert any type to string
        }
    }
    
    reqURL.RawQuery = query.Encode()
    
    // Make the request
    resp, err := http.Get(reqURL.String())
    if err != nil {
        return CountryResponse{}, fmt.Errorf("error making request: %v", err)
    }
    defer resp.Body.Close()

    // Parse the response
    var response CountryResponse
    err = json.NewDecoder(resp.Body).Decode(&response)
    if err != nil {
        return CountryResponse{}, fmt.Errorf("error decoding response: %v", err)
    }

    return response, nil
}
