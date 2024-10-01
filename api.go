package main

import (
    "bytes"
    "encoding/json"
    "net/http"
)

type API struct {
    apiKey string
}

type Response struct {
    Status     int       `json:"status"`
    TotalCount int       `json:"totalCount"`
    Country    string    `json:"country"`
    Year       int       `json:"year"`
    Data       []Holiday `json:"data"`
}

type Holiday struct {
    Name     string `json:"name"`
    Date     string `json:"date"`
    Observed string `json:"observed"`
    IsPublic bool   `json:"isPublic"`
    Country  struct {
        Code string `json:"code"`
        Name string `json:"name"`
        Flag string `json:"flag"`
    } `json:"country"`
}

func NewAPI(apiKey string) *API {
    return &API{apiKey: apiKey}
}

func (a *API) GetHolidays(params map[string]interface{}) (Response, error) {
    // Convert parameters to JSON
    jsonData, _ := json.Marshal(params)

    // Make the API request
    resp, err := http.Post("https://back.holidaylist.io/api/v1/holidays", "application/json", bytes.NewBuffer(jsonData))
    if err != nil {
        return Response{}, err
    }
    defer resp.Body.Close()

    // Parse the response
    var response Response
    json.NewDecoder(resp.Body).Decode(&response)

    // Get status code from HTTP response
    response.Status = resp.StatusCode

    return response, nil
}
