package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

func exec(fn func([]string), urls []string) {
	start := time.Now().UnixMilli()
	fn(urls)
	fmt.Printf("[Took %v]\n", time.Now().UnixMilli()-start)
}

func main() {
	urls := os.Args[1:]
	exec(fetch1, urls)
	// fetch2(urls)
	// fetch3(urls)
	exec(concurrentFetch, urls)
}

func fetch1(urls []string) {
	type Res struct {
		resp   []byte
		err    error
		status int
	}
	res := make(map[string]*Res)
	for _, url := range urls {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Println("skipping "+url, err)
			continue
		}
		res[url] = &Res{
			err:    nil,
			status: resp.StatusCode,
		}
		res[url].resp, err = io.ReadAll(resp.Body)
		defer resp.Body.Close()
	}
	for k, v := range res {
		fmt.Println(k, string(*&v.resp))
	}
}

/*
The function call io.Copy(dst, src) reads from src and writes to dst. Use it
instead of ioutil.ReadAll to copy the response body to os.Stdout without requir ing a
buffer large enough to hold the entire stream. Be sure to che ck the error result of io.Copy.
*/
func fetch2(urls []string) {
	for _, url := range urls {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Println("skipping "+url, err)
			continue
		}
		fmt.Println("[" + url + "]")
		_, err = io.Copy(os.Stdin, resp.Body)
		if err != nil {
			fmt.Println("error writing to stdin", err)
		}
		defer resp.Body.Close()
	}
}

/*
Modify fetch to add the prefix http:// to each argument URL if it is missing.
You might want to use strings.HasPrefix.
*/
func fetch3(urls []string) {
	type Res struct {
		resp   []byte
		err    error
		status int
	}
	res := make(map[string]*Res)
	for _, url := range urls {
		if !strings.HasPrefix(url, "http") {
			url = "http://" + url
		}
		resp, err := http.Get(url)
		if err != nil {
			fmt.Println("skipping "+url, err)
			continue
		}
		res[url] = &Res{
			err:    nil,
			status: resp.StatusCode,
		}
		res[url].resp, err = io.ReadAll(resp.Body)
		defer resp.Body.Close()
	}
	for k, v := range res {
		fmt.Println(k, string(*&v.resp))
	}
}

type Res struct {
	res        []byte
	err        error
	statuscode int
}

/*
concurrent fetch
*/
func concurrentFetch(urls []string) {

	_ = make(map[string]*Res)
	ch := make(chan *Res)
	for _, url := range urls {
		go fetch(url, ch)
	}
	for range urls {
		fmt.Println(string((<-ch).res))
	}
}

func fetch(url string, ch chan<- *Res) {
	res, err := http.Get(url)
	if err != nil {
		fmt.Println("Skipping ", err)
	}
	ret := &Res{
		statuscode: res.StatusCode,
		err:        nil,
	}
	ret.res, err = io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("skipping ", err)
	}
	ch <- ret
}
