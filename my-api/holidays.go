package myapi

import (
    "encoding/json"
    "net/http"
    "log"
)

// GetHolidays fetches holiday data based on the provided parameters
func (a *API) GetHolidays(params map[string]interface{}) Response {
    // Construct the endpoint
    endpoint := "https://back.holidaylist.io/api/v1/holidays"

    // Prepare the URL parameters for GET request
    reqURL := endpoint + "?country=" + params["country"].(string) + "&year=" + params["year"].(string) + "&key=" + a.apiKey

    // Make the API request
    resp, err := http.Get(reqURL)
    if err != nil {
        log.Println("Error:", err)
        return Response{}
    }
    defer resp.Body.Close()

    // Parse the response
    var response Response
    err = json.NewDecoder(resp.Body).Decode(&response)
    if err != nil {
        log.Println("Error decoding response:", err)
        return Response{}
    }

    // Include HTTP status in the response object
    response.Status = resp.StatusCode

    return response
}
