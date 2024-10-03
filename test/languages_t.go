package main

import (
    "fmt"
    "log"
    "github.com/feelgd777/holidaylist-go"
)

func main() {
    api := holidaylist.NewAPI("21deeb8a-183d-444c-bba3-9f853b81ad86") // Replace with your actual API key

    // Fetch the languages
    languagesResponse, err := api.GetLanguages()

    // Check if there was an error
    if err != nil {
        log.Fatalf("Expected no error, got %v", err)
    }

    // Access the Data field which holds the slice of languages
    languages := languagesResponse.Data

    // Check if we got any languages back
    if len(languages) == 0 {
        log.Fatalf("Expected some languages, got 0")
    }

    // Print out the first language
    fmt.Printf("First language: %v\n", languages[0].Name)
}
