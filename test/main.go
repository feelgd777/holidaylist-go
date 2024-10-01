package main

import (
    "fmt"
    "github.com/feelgd777/holidaylist-go/my-api"
)

func main() {
    api := myapi.NewAPI("21deeb8a-183d-444c-bba3-9f853b81ad86")

    holidays, err := api.GetHolidays(map[string]interface{}{
        "country": "US",
        "year":    2023,
    })

    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    fmt.Println(holidays)
}

