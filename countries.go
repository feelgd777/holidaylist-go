package holidaylist

import (
    "encoding/json"
    "log"
    "net/http"
    "net/url"
)

// GetCountries fetches the list of countries, allowing optional parameters to be passed as a JSON object
func (a *API) GetCountries(params map[string]interface{}) (CountryResponse, error) {
    endpoint := "https://back.holidaylist.io/api/v1/countries"
    
    // Build the query parameters from the passed JSON object
    reqURL, _ := url.Parse(endpoint)
    query := reqURL.Query()
    query.Add("key", a.apiKey)
    
    // Loop through the passed parameters and add them to the query string
    for key, value := range params {
        query.Add(key, value.(string))
    }
    
    reqURL.RawQuery = query.Encode()
    
    // Make the request
    resp, err := http.Get(reqURL.String())
    if err != nil {
        log.Println("Error making request:", err)
        return CountryResponse{}, err
    }
    defer resp.Body.Close()

    // Parse the response
    var response CountryResponse
    err = json.NewDecoder(resp.Body).Decode(&response)
    if err != nil {
        log.Println("Error decoding response:", err)
        return CountryResponse{}, err
    }

    return response, nil
}
