package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"text/template"
	"time"
)

type Issue struct {
	Number int
	URL    string `json:"url"`
	User   struct {
		Login   string `json:"login"`
		HTMLURL string `json:"html_url"`
	} `json:"user"`
	Title     string    `json:"title"`
	State     string    `json:"state"`
	Body      string    `json:"body"`
	UpdatedAt time.Time `json:"updated_at"`
	HTMLURL   string    `json:"html_url"`
}

type Report int

// Set implements flag.Value.
func (r *Report) Set(s string) error {
	if s == "stdout" {
		*r = stdout
	} else if s == "html" {
		*r = html
	} else {
		return errors.New("Invalid Report type")
	}
	return nil
}

// String implements flag.Value.
func (r Report) String() string {
	if r == stdout {
		return "stdout"
	} else {
		return "html"
	}
}

const (
	stdout Report = iota
	html
)

var daysOld *int = flag.Int("days", 15, "Get at most 'days' old issues.")
var repo *string = flag.String("repo", "", "Repository name (owner:reponame)")
var serve *bool = flag.Bool("serve", false, "Start a webserver to navigate issues")
var port *int = flag.Int("port", 8080, "Port to start server on")
var issues []Issue

func main() {
	var reportType Report
	flag.Var(&reportType, "report", "Type of report to produce (stdout | html)")
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
		// fmt.Printf("[updatedAt:%v, duration:%f\n", issue.UpdatedAt, time.Since(issue.UpdatedAt).Hours())
		if int(time.Since(issue.UpdatedAt).Hours()) < (*daysOld)*24 {
			active = append(active, issue)
		} else {
			old = append(old, issue)
		}
	}
	if !*serve {
		report(os.Stdout, issues, reportType)
		os.Exit(0)
	}
	// server
	http.HandleFunc("/issue/", showIssue)
	fmt.Printf("Starting server on : http://localhost:%d", *port)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", *port), nil); err != nil {
		log.Fatal("Error starting the server:", err)
	}
}

func showIssue(w http.ResponseWriter, r *http.Request) {
	urlParts := strings.Split(fmt.Sprintf("%s", r.URL), "/")
	if len(urlParts) != 3 {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	issueNumber, err := strconv.Atoi(urlParts[2])
	if err != nil || issueNumber <= 0 || issueNumber >= len(issues) {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	_showIssue(w, issueNumber)
}

func _showIssue(w http.ResponseWriter, issueNumber int) {
	report(w, []Issue{issues[issueNumber-1]}, html)
}

func daysAgo(t time.Time) int {
	return int(time.Since(t).Hours() / 24)
}

func report(w io.Writer, issues []Issue, reportType Report) {
	var reportStdout string = `{{.TotalCount}} issues:
	{{range .Items}}--------------------------------------
	Title: {{.Title}}
	User: {{.User.Login}}
	Updated: {{.UpdatedAt | daysAgo}} days ago
	{{end}}`

	var reportHtml string = `
		<h1>{{.TotalCount}} issues</h1>
		<table>
		<tr style='text-align: left'>
			<th>#</th>
			<th>State</th>
			<th>User</th>
			<th>Title</th>
			<th>Last Updated</th>
		</tr>
		{{range .Items}}
		<tr>
			<td><a href='{{.HTMLURL}}'>{{.Number}}</a></td>
			<td>{{.State}}</td>
			<td><a href='{{.User.HTMLURL}}'>{{.User.Login}}</a></td>
			<td><a href='{{.HTMLURL}}'>{{.Title}}</a></td>
			<td>{{.UpdatedAt | daysAgo}} days ago</td>
		</tr>
		{{end}}
		</table>
	`
	var reportTemplate = template.New("report").Funcs(template.FuncMap{"daysAgo": daysAgo})
	var reportTempl string
	fmt.Println("Report will be", reportType)
	if reportType == stdout {
		reportTempl = reportStdout
	} else {
		reportTempl = reportHtml
	}
	report, err := reportTemplate.Parse(reportTempl)
	if err != nil {
		log.Fatal("Error parsing template: ", err)
	}
	if err := report.Execute(w, struct {
		TotalCount int
		Items      []Issue
	}{len(issues), issues}); err != nil {
		log.Fatal(err)
	}
}
