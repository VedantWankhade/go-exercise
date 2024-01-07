package main

import (
	"encoding/json"
	"fmt"
)

type Movie struct {
    Title string
    Year int `json:"released"`
    Color bool `json"color,omitempty"`
}

func main() {
    movies := []Movie{
        {
            Title: "A", Year: 1942, Color: false,
        },
        {
            Title: "B", Year: 2000, Color: true,
        },
        {
            Title: "C", Year: 2011,
        },
    }
    j, _ := json.Marshal(movies)
    fmt.Printf("%s", j)
}
