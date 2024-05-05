package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strings"
	"time"
)

func daysAgo(t time.Time) int {
	return int(time.Since(t).Hours() / 24)
}

var IssueListTemplate = template.Must(template.New("issueList").Funcs(template.FuncMap{"dasyAgo": daysAgo}).Parse(`
	<h1>Issues</h1>
	<table>
		<tr style='text-align:left'>
			<th>#</th>
			<th>State</th>
			<th>User</th>
			<th>Title</th>
		</tr>
		{{range .}}
		<tr>
		<td><a href='{{.HTMLURL}}'>{{.Number}}</a></td>
		<td>{{.State}}</td>
		<td><a href='{{.User.HTMLURL}}'>{{.User.Login}}</a></td>
		<td><a href='{{.HTMLURL}}'>{{.Title}}</a></td>
		</tr>
		{{end}}
	</table>
`))

var GITHUB_TOKEN = "Bearer " + os.Getenv("GITHUB_TOKEN")

func getIssuesURL(owner, repo string) string {
	return fmt.Sprintf("https://api.github.com/repos/%s/%s/issues", owner, repo)
}

type Issue struct {
	Number    int
	HTMLURL   string `json:"html_url"`
	Title     string
	State     string
	User      *User
	CreatedAt time.Time `json:"created_at"`
	Body      string
}

type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}

func SearchIssues(owner, repo string, params RequestParams) ([]Issue, error) {
	url := getIssuesURL(owner, repo)
	if len(params) != 0 {
		url = fmt.Sprintf("%s?%s", url, &params)
	}
	res, err := http.Get(url)
	if err != nil || res.StatusCode != 200 {
		return nil, fmt.Errorf("error with get request: status: %d: %v", res.StatusCode, err)
	}
	// result, err := io.ReadAll(res.Body)
	// if err != nil {
	// 	return nil, fmt.Errorf("error reading from body: %v", err)
	// }
	var issues []Issue
	// err = json.Unmarshal(result, &issues)
	err = json.NewDecoder(res.Body).Decode(&issues)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling json: %v", err)
	}
	return issues, nil
}

func main() {
	action := flag.String("type", "issues", "issues|users")
	owner := flag.String("owner", "octocat", "github id of owner")
	repo := flag.String("repo", "Spoon-Knife", "repo name")
	title := flag.String("title", "Test issue created with API", "Issue title")
	since := flag.Bool("since", false, "--since")
	months := flag.Int("m", 0, "-m10")
	days := flag.Int("d", 0, "-d10")
	flag.Parse()
	switch *action {
	case "issues":
		issues(*owner, *repo, *since, *months, *days)
	case "createIssue":
		createIssue(*owner, *repo, *title, []string{"test", "api"})
	default:
		fmt.Println("quiting")
	}
}

func createIssue(owner, repo, title string, labels []string) {
	fmt.Printf("init create issue at %s/%s\n", owner, repo)
	editor := os.Getenv("EDITOR")
	tempFile := "/tmp/tempFile.txt"
	cmd := exec.Command(editor, tempFile)
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	// err := cmd.Run()
	// if err != nil {
	// 	fmt.Println(err)
	// }
	var body = &struct {
		Title  string   `json:"title"`
		Body   string   `json:"body"`
		Labels []string `json:"labels"`
	}{
		Title:  title,
		Labels: labels,
		Body:   "Issue body TBI",
	}
	b, err := json.Marshal(body)
	fmt.Println(string(b))
	if err != nil {
		fmt.Println(err)
	}
	req, err := http.NewRequest(http.MethodPost, getIssuesURL(owner, repo), strings.NewReader(string(b)))
	if err != nil {
		fmt.Println(err)
	}
	req.Header.Add("Accept", "application/vnd.github+json")
	req.Header.Add("Authorization", GITHUB_TOKEN)
	c := http.Client{}
	res, err := c.Do(req)
	fmt.Println(err, res.Status)
	result, err := io.ReadAll(req.Body)
	if err != nil {
		fmt.Println(err)
	}
	defer res.Body.Close()
	fmt.Println(result)
}

func issues(owner, repo string, since bool, months, days int) {
	params := make(map[string]string)
	params["owner"] = owner
	params["repo"] = repo
	if since {
		params["since"] = time.Now().AddDate(0, -months, -days).Format("2006-01-02T15:04:05Z")
	}
	res, err := SearchIssues(owner, repo, params)
	if err != nil {
		log.Fatalln("Error searching issues", err)
	}
	for _, issue := range res {
		fmt.Printf("#%-5d %9.9s %.55s\n", issue.Number, issue.User.Login, issue.Title)
	}
	tempFile := "/tmp/temp.html"
	f, err := os.Create(tempFile)
	defer f.Close()
	if err := IssueListTemplate.Execute(f, res); err != nil {
		log.Fatal(err)
	}
	c := exec.Command("xdg-open", tempFile)
	c.Run()
}

type RequestParams map[string]string

func (rp *RequestParams) String() string {
	var params bytes.Buffer
	count := 0
	last := len(*rp)
	for k, v := range *rp {
		count++
		params.WriteString(k + "=" + v)
		if count != last {
			params.WriteString("&")
		}
	}
	fmt.Println(params.String())
	return params.String()
}
