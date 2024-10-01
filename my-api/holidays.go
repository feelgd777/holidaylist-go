package holidaylist

// GetHolidays is the function for fetching holiday data
func (a *API) GetHolidays(params map[string]interface{}) Response {
    // Define the holidays endpoint
    endpoint := "https://back.holidaylist.io/api/v1/holidays"

    // Fetch the holidays using the core request logic
    response, err := a.getRequest(endpoint, params)
    if err != nil {
        return Response{} // If there's an error, return an empty response
    }

    return response
}