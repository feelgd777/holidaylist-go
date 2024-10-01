package myapi

import (
    "encoding/json"
    "fmt"
    "net/http"
    "net/url"
    "io/ioutil"
)

type API struct {
    apiKey string
}

func NewAPI(apiKey string) *API {
    return &API{apiKey: apiKey}
}

func (a *API) GetHolidays(params map[string]interface{}) (Response, error) {
    baseURL := "https://back.holidaylist.io/api/v1/holidays"

    // Prepare query parameters
    reqURL, _ := url.Parse(baseURL)
    query := reqURL.Query()
    query.Add("key", a.apiKey)

    for k, v := range params {
        query.Add(k, fmt.Sprintf("%v", v))
    }
    reqURL.RawQuery = query.Encode()

    // Make the GET request
    resp, err := http.Get(reqURL.String())
    if err != nil {
        return Response{}, err
    }
    defer resp.Body.Close()

    // Read and parse response body
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return Response{}, err
    }

    var response Response
    json.Unmarshal(body, &response)

    // Include HTTP status in the response object
    response.Status = resp.StatusCode

    return response, nil
}
