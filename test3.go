package main

import (
    "github.com/feelgd777/holidaylist-go/my-api"
)

func main() {
    // Initialize API with your key
    api := myapi.NewAPI("21deeb8a-183d-444c-bba3-9f853b81ad86")

    // Fetch holidays without handling errors directly
    holidays := api.GetHolidays(map[string]interface{}{
        "country": "US",
        "year":    2023,
        // Optional parameters can be added here
    })

    // The holidays variable should hold the result, and the package will handle
    // the display or other necessary steps internally.
}
