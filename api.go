package holidaylist

// API holds the user's API key
type API struct {
    apiKey string
}

// NewAPI initializes the API with the provided key
func NewAPI(apiKey string) *API {
    return &API{apiKey: apiKey}
}
