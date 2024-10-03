package main

import (
    "fmt"
    "github.com/feelgd777/holidaylist-go"
)

func main() {
    api := holidaylist.NewAPI("YOUR_API_KEY")

    // Fetch countries with the optional parameter
    countries, err := api.GetCountries(map[string]interface{}{
        "language": "es",  // Optional parameter: fetch countries in Spanish
    })

    if err != nil {
        fmt.Println("Error fetching countries:", err)
        return
    }

    // Print the first country name
    fmt.Println("First country:", countries.Data[0].Name)
}
