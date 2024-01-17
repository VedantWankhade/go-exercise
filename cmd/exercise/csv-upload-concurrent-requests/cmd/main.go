package main

func main() {
	app := initializeApp()
	app.infoLogger.Println("Starting server on http://localhost" + app.server.Addr)
	err := app.server.ListenAndServe()
	if err != nil {
		app.errLogger.Fatal("Failed to start server", err)
	}
}
