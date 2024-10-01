package holidaylist

import (
    "encoding/json"
    "net/http"
    "log"
    "strconv"
)

func (a *API) GetHolidays(params map[string]interface{}) (Response, error) {
    // Construct the endpoint
    endpoint := "https://back.holidaylist.io/api/v1/holidays"

    // Prepare URL parameters for GET request with optional parameters
    reqURL := endpoint + "?country=" + params["country"].(string) + "&year=" + strconv.Itoa(params["year"].(int))

    // Check for optional parameters and append them to the request URL if they exist
    if val, ok := params["month"]; ok {
        reqURL += "&month=" + val.(string)
    }
    if val, ok := params["day"]; ok {
        reqURL += "&day=" + val.(string)
    }
    if val, ok := params["public"]; ok {
        reqURL += "&public=" + val.(string)
    }

    // Append the API key
    reqURL += "&key=" + a.apiKey

    // Make the API request
    resp, err := http.Get(reqURL)
    if err != nil {
        log.Println("Error making request:", err)
        return Response{}, err
    }
    defer resp.Body.Close()

    // Parse the response
    var response Response
    err = json.NewDecoder(resp.Body).Decode(&response)
    if err != nil {
        log.Println("Error decoding response:", err)
        return Response{}, err
    }

    // Include HTTP status in the response object
    response.Status = resp.StatusCode

    return response, nil
}