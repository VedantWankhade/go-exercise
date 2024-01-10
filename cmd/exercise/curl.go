package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	// parallelCurl()

	curl1()
}

func curl1() {
	for _, url := range os.Args[1:] {
		res, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "curl: %v\n", err)
			os.Exit(1)
		}

		byteData, err := ioutil.ReadAll(res.Body)
		res.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "curl: reading %s: %v\n", url, err)
			os.Exit(1)
		}
		fmt.Printf("%s", byteData)
	}
}

// use io.Copy instead of ioutil.ReadAll
func curl2() {
	for _, url := range os.Args[1:] {
		res, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "curl: %v\n", err)
			os.Exit(1)
		}
		_, err = io.Copy(os.Stdout, res.Body)
		res.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "curl: writing os.Stdout from res.Body: %v", err)
		}
	}
}

// add https:// prefix it it doesnt exist
func curl3() {
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, "https://") {
			url = "https://" + url
		}
		res, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "curl: %v\n", err)
			// os.Exit(1)
			continue
		}
		fmt.Println("HTTP STATUS")
		fmt.Println(res.Status)
		_, err = io.Copy(os.Stdout, res.Body)
		res.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "curl: writing os.Stdout from res.Body: %v", err)
		}
	}
}

// parallelCurl fetches all urls concurrently
func parallelCurl() {
	start := time.Now()
	ch := make(chan string)
	for i, url := range os.Args[1:] {
		go fetch(url, ch, i+1)
	}
	for range os.Args[1:] {
		fmt.Printf(<-ch)
	}
	fmt.Printf("Total time elapsed: %.2fs", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string, n int) {
	start := time.Now()
	fmt.Printf("thread %d started\n", n)
	res, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprintf("%v\n", err)
	} else {
		nbytes, err := io.Copy(ioutil.Discard, res.Body)
		res.Body.Close()
		if err != nil {
			ch <- fmt.Sprintf("error reading response: %v\n", err)
		} else {
			ch <- fmt.Sprintf("%.2fs\t%7d\t%s\n", time.Since(start).Seconds(), nbytes, url)
		}
	}
}
