package main

import (
    "fmt"
    "log"
    "github.com/feelgd777/holidaylist-go"
)

func main() {
    api := holidaylist.NewAPI("21deeb8a-183d-444c-bba3-9f853b81ad86") // Replace with your actual API key

    // Fetch the countries
    countriesResponse, err := api.GetCountries(map[string]interface{}{
        "language":"es",
    })

    // Check if there was an error
    if err != nil {
        log.Fatalf("Expected no error, got %v", err)
    }

    // Access the Data field which holds the slice of countries
    countries := countriesResponse.Data

    // Check if we got any countries back
    if len(countries) == 0 {
        log.Fatalf("Expected some countries, got 0")
    }

    // Print out the first country
    fmt.Printf("First country: %v\n", countries[0].Name)
}
