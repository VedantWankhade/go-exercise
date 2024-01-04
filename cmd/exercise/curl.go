package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func main() {
    curl3()
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
            os.Exit(1)
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
