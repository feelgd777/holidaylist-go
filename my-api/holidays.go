package holidaylist

import (
    "bytes"
    "encoding/json"
    "net/http"
)


// GetHolidays is the function for fetching holiday data
func (a *API) GetHolidays(params map[string]interface{}) Response {
    // Make API request here
    jsonData, _ := json.Marshal(params)

    resp, err := http.Get(endpoint) // Use GET
    if err != nil {
        log.Println("Error:", err) // Log the error instead of returning it
        return Response{} // Return an empty response on error
    }
    defer resp.Body.Close()

    // Parse the response
    var response Response
    json.NewDecoder(resp.Body).Decode(&response)

    response.Status = resp.StatusCode
    return response
}
