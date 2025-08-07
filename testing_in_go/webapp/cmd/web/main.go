package main

import (
	"log"
	"log/slog"
	"net/http"
	"os"
)

// go get github.com/go-chi/chi/v5
// go get github.com/go-chi/chi/v5/middleware
// go run ./cmd/web
// go test ./...

type application struct{}

func initLog() {
	handlerOptions := &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}

	logger := slog.New(slog.NewTextHandler(os.Stdout, handlerOptions))
	slog.SetDefault(logger)
}

func main() {
	initLog()

	// set up an app config
	app := application{}

	// get application routes
	mux := app.routes()

	// print out a message
	slog.Info("Starting server on port 8080")

	// start the server
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal(err)
	}
}
