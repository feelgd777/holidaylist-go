package main

import (
    "fmt"
    "github.com/feelgd777/holidaylist-go" // Import your API package
)

func main() {
    api := holidaylist.NewAPI("21deeb8a-183d-444c-bba3-9f853b81ad86") // Initialize API with your key

    holidays := api.GetHolidays(map[string]interface{}{
        "country": "US",
        "year":    2023,
    })

    fmt.Println(holidays) // Print the holidays response
}
