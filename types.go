package main

// Response struct for parsing the holiday data and status
type Response struct {
    Status     int       json:"status"
    TotalCount int       json:"totalCount"
    Country    string    json:"country"
    Year       int       json:"year"
    Data       []Holiday json:"data"
}

// Holiday represents individual holiday data
type Holiday struct {
    Name     string json:"name"
    Date     string json:"date"
    Observed string json:"observed"
    IsPublic bool   json:"isPublic"
    Country  struct {
        Code string json:"code"
        Name string json:"name"
        Flag string json:"flag"
    } json:"country"
}