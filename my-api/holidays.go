package myapi

// GetHolidays fetches holiday data
func (a *API) GetHolidays(params map[string]interface{}) Response {
    endpoint := "https://back.holidaylist.io/api/v1/holidays"
    response, err := a.getRequest(endpoint, params)
    if err != nil {
        return Response{}
    }
    return response
}
