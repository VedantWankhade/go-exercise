package main

import "net/http"

func (app *application) routes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/list", app.list)
	mux.HandleFunc("/submit", app.submit)
	mux.HandleFunc("/update", app.update)
	mux.HandleFunc("/submissionState", app.submissionState)
	return mux
}
