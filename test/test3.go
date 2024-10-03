package main

import (
    "github.com/feelgd777/holidaylist-go"
    "fmt"
)

func main() {
    api := holidaylist.NewAPI("21deeb8a-183d-444c-bba3-9f853b81ad86")
    holidays, _ := api.GetHolidays(map[string]interface{}{
        "country": "US",
        "year": 2023,
    })
    fmt.Println(holidays)
}