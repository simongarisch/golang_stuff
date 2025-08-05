package main

import (
	"log"
	"net/http"
)

// go get github.com/go-chi/chi/v5
// go get github.com/go-chi/chi/v5/middleware
// go run ./cmd/web

type application struct{}

func main() {
	// set up an app config
	app := application{}

	// get application routes
	mux := app.routes()

	// print out a message
	log.Println("Starting server on port 8080")

	// start the server
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal(err)
	}
}
