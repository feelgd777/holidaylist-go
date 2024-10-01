package holidaylist

import (
    "bytes"
    "encoding/json"
    "net/http"
)

// API holds the user's API key
type API struct {
    apiKey string
}

// NewAPI initializes the API with the provided key
func NewAPI(apiKey string) *API {
    return &API{apiKey: apiKey}
}

func (a *API) getRequest(endpoint string, params map[string]interface{}) (Response, error) {
    // Convert parameters to JSON
    jsonData, _ := json.Marshal(params)

    // Make the API request
    resp, err := http.Post(endpoint, "application/json", bytes.NewBuffer(jsonData))
    if err != nil {
        return Response{}, err
    }
    defer resp.Body.Close()

    var response Response
    json.NewDecoder(resp.Body).Decode(&response)

    // Include HTTP status in the response object
    response.Status = resp.StatusCode

    return response, nil
}

// GetHolidays fetches holiday data based on the provided parameters

    jsonData, _ := json.Marshal(params)

    // Make the API request as a GET request
    req, err := http.NewRequest("GET", endpoint, bytes.NewBuffer(jsonData))
    if err != nil {
        return Response{}, err
    }
    req.Header.Set("Content-Type", "application/json")

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        return Response{}, err
    }
    defer resp.Body.Close()

    // Parse the response
    var response Response
    if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
        return Response{}, err
    }

    // Include HTTP status in the response object
    response.Status = resp.StatusCode

    return response, nil
}
