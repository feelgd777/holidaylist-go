package main

import (
    "fmt"
    "github.com/feelgd777/holidaylist-go"
)

func main() {
    // Initialize API with your key
    api := holidaylist.NewAPI("21deeb8a-183d-444c-bba3-9f853b81ad86")

    // Fetch holidays
    holidays, err := api.GetHolidays(map[string]interface{}{
        "country": "US",
        "year":    2023,
        // Optional parameters can be added here
    })

    // Handle error
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    // **USE the holidays variable by printing it**
    fmt.Println("Holidays Response:", holidays)
}
