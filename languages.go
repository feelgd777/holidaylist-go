package holidaylist

import (
    "encoding/json"
    "log"
    "net/http"
)

// Language represents a single language returned by the API
type Language struct {
    Code string `json:"code"`
    Name string `json:"name"`
}

// LanguageResponse represents the structure of the response from the API for the languages endpoint
type LanguageResponse struct {
    Status int        `json:"status"`
    Data   []Language `json:"data"`
}

// GetLanguages fetches the list of languages from the API
func (a *API) GetLanguages() (LanguageResponse, error) {
    endpoint := "https://back.holidaylist.io/api/v1/languages"
    reqURL := endpoint + "?key=" + a.apiKey

    resp, err := http.Get(reqURL)
    if err != nil {
        log.Println("Error making request:", err)
        return LanguageResponse{}, err
    }
    defer resp.Body.Close()

    var response LanguageResponse
    err = json.NewDecoder(resp.Body).Decode(&response)
    if err != nil {
        log.Println("Error decoding response:", err)
        return LanguageResponse{}, err
    }

    response.Status = resp.StatusCode
    return response, nil
}
