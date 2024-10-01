# HolidayList Go API

## Usage
```go
import "github.com/yourusername/holidaylist-go"

api := myapi.NewAPI("your_api_key") // Initialize API with your key

holidays := api.GetHolidays(map[string]interface{}{
    "country": "US",
    "year": 2023,
})
