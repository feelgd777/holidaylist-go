package holidaylist

import (
    "bytes"
    "encoding/json"
    "net/http"
)

type API struct {
    apiKey string
}

func NewAPI(apiKey string) *API {
    return &API{apiKey: apiKey}
}

// getRequest sends a request to the API with parameters
func (a *API) getRequest(endpoint string, params map[string]interface{}) (Response, error) {
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