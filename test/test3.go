package main

import (
    "github.com/feelgd777/holidaylist-go"
    "fmt"
)

func main() {
    api := holidaylist.NewAPI("YOUR_API_KEY")
    holidays, _ := api.GetHolidays(map[string]interface{}{
        "country": "US",
        "year":    2023,
    })
    fmt.Println(holidays)
}