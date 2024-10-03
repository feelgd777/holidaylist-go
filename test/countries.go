package main

import (
    "fmt"
    "log"
    "github.com/feelgd777/holidaylist-go"
)

func main() {
    api := holidaylist.NewAPI("21deeb8a-183d-444c-bba3-9f853b81ad86")  // Replace with your actual API key

    // Fetch countries with the optional language parameter set to "es" (Spanish)
    countriesResponse, err := api.GetCountries(map[string]interface{}{
        "language": "es",  // Optional parameter: fetch countries in Spanish
    })

    // Check for any errors
    if err != nil {
        log.Fatalf("Error fetching countries: %v", err)
    }

    // Check if we got any countries back
    countries := countriesResponse.Data
    if len(countries) == 0 {
        log.Fatalf("Expected some countries, got none")
    }

    // Print out the first country to verify the result
    fmt.Printf("First country (in Spanish): %v\n", countries[0].Name)
}
