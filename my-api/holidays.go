package myapi

// GetHolidays fetches holiday data using the holidays endpoint
func (a *API) GetHolidays(params map[string]interface{}) (Response, error) {
    endpoint := "https://back.holidaylist.io/api/v1/holidays"
    return a.getRequest(endpoint, params)
}
