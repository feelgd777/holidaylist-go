package myapi

import (
    "bytes"
    "encoding/json"
    "fmt"
    "io/ioutil"
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

// getRequest sends a request to the API with parameters
func (a *API) getRequest(endpoint string, params map[string]interface{}) (Response, error) {
    jsonData, _ := json.Marshal(params)

    // Debugging: Print the JSON data
    fmt.Println("Sending JSON Data: ", string(jsonData))

    // Make the API request
    resp, err := http.Post(endpoint, "application/json", bytes.NewBuffer(jsonData))
    if err != nil {
        return Response{}, err
    }
    defer resp.Body.Close()

    // Debugging: Print response status and body
    fmt.Println("Response Status:", resp.StatusCode)
    body, _ := ioutil.ReadAll(resp.Body)
    fmt.Println("Response Body:", string(body))

    var response Response
    json.NewDecoder(bytes.NewBuffer(body)).Decode(&response)

    // Include HTTP status in the response object
    response.Status = resp.StatusCode

    return response, nil
}
