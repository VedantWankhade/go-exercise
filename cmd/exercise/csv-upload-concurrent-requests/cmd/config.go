package main

import (
	"flag"
	"html/template"
	"log"
	"net/http"
	"os"
)

type application struct {
	errLogger  *log.Logger
	infoLogger *log.Logger
	server     *http.Server
	templates  struct{ home, list, submissions *template.Template }
	workers    *int
}

func initializeApp() *application {
	var app application
	app.infoLogger = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	app.errLogger = log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	app.workers = flag.Int("workers", 5, "number of workers")
	addr := flag.String("port", ":8080", "port number to run server on")
	flag.Parse()
	app.server = &http.Server{
		Addr:     *addr,
		ErrorLog: app.errLogger,
		Handler:  app.routes(),
	}
	app.templates = app.templs()
	app.infoLogger.Println("Customer loggers initialized")
	app.infoLogger.Println("Initializing application")
	app.infoLogger.Println("Port set to", app.server.Addr)
	app.infoLogger.Println("Templates set to", app.templates)
	return &app
}
