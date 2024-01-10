package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

type Issue struct {
	URL       string    `json:"url"`
	Title     string    `json:"title"`
	State     string    `json:"state"`
	Body      string    `json:"body"`
	UpdatedAt time.Time `json:"updated_at"`
}

var daysOld *int  = flag.Int("days", 15, "Get at most 'days' old issues.")
var repo *string = flag.String("repo", "", "Repository name (owner:reponame)")

func main() {
    flag.Parse()
    if *repo == "" {
        flag.Usage()
        os.Exit(2)
    }
    owner, reponame, found := strings.Cut(*repo, ":")
    if !found {
        flag.Usage()
        log.Fatal("Please provide a proper repository in the form (owner:reponame)")
    }
    url := fmt.Sprintf("https://api.github.com/repos/%s/%s/issues", owner, reponame)
	req, _ := http.NewRequest(http.MethodGet, url, nil)
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
    var active, old []Issue
    for _, issue := range issues {
        fmt.Printf("[updatedAt:%v, duration:%f\n", issue.UpdatedAt, time.Since(issue.UpdatedAt).Hours())
        if int(time.Since(issue.UpdatedAt).Hours()) < (*daysOld) * 24 {
            active = append(active, issue)
        } else {
            old = append(old, issue)
        }
    }
    
    fmt.Printf("%d, %d, %d", len(issues), len(active), len(old))
}

