package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"

	"github.com/vedantwankhade/go-exercise/gopl/lissajous"
)

var count int
var mu sync.Mutex

func main() {
	server1()
}

func server1() {
	a := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "text/html")
		_, err := fmt.Fprintln(w, "<h2>URL PATH:", r.URL.Path, "</h2>")
		if err != nil {
			fmt.Println("err writing response", err)
		}
	}
	b := func(w http.ResponseWriter, r *http.Request) {
		mu.Lock()
		count++
		fmt.Fprintln(w, "count", count)
		mu.Unlock()
	}
	c := func(w http.ResponseWriter, r *http.Request) {
		_ = r.ParseForm()
		cycles, _ := strconv.Atoi(r.FormValue("cycles"))
		lissajous.Lissajous(w, cycles)
	}
	http.HandleFunc("/", a)
	http.HandleFunc("/count", b)
	http.HandleFunc("/gif", c)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
