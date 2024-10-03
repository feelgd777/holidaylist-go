package holidaylist

import (
    "encoding/json"
    "net/http"
    "log"
)

// GetLanguages fetches the list of languages from the API
func (a *API) GetLanguages() ([]Language, error) {
    endpoint := "https://back.holidaylist.io/api/v1/languages"
    reqURL := endpoint + "?key=" + a.apiKey

    resp, err := http.Get(reqURL)
    if err != nil {
        log.Println("Error making request:", err)
        return nil, err
    }
    defer resp.Body.Close()

    var response LanguageResponse
    err = json.NewDecoder(resp.Body).Decode(&response)
    if err != nil {
        log.Println("Error decoding response:", err)
        return nil, err
    }

    return response.Data, nil
}
