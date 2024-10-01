package myapi

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

// getRequest sends a GET request to the API with parameters
func (a *API) getRequest(endpoint string, params map[string]interface{}) (Response, error) {
    jsonData, _ := json.Marshal(params)

    req, err := http.NewRequest("GET", endpoint, bytes.NewBuffer(jsonData))
    if err != nil {
        return Response{}, err
    }
    req.Header.Set("Content-Type", "application/json")
    req.Header.Set("Authorization", "Bearer "+a.apiKey)

    client := &http.Client{}
    resp, err := client.Do(req)
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

