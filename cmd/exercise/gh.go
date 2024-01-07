package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

type Issue struct {
	URL       string    `json:"url"`
	Title     string    `json:"title"`
	State     string    `json:"state"`
	Body      string    `json:"body"`
	UpdatedAt time.Time `json:"updated_at"`
}

func main() {
	req, _ := http.NewRequest(http.MethodGet, "https://api.github.com/repos/octocat/Spoon-Knife/issues", nil)
	req.Header.Add("Authorization", "Bearer "+os.Getenv("GH_TOKEN"))
	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("oops!!")
		res.Body.Close()
		os.Exit(1)
	}
	if res.StatusCode != http.StatusOK {
		fmt.Println("oops!!")
		fmt.Printf("%s\n", res.Status)
		os.Exit(1)
	}
	defer res.Body.Close()
	var issues []Issue
	data, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("oops!!")
		os.Exit(1)
	}
	json.Unmarshal(data, &issues)
	if err != nil {
		fmt.Println("oops!!")
		os.Exit(1)
	}
	fmt.Println(issues[0:2])

}

