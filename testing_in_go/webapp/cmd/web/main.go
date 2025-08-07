package main

import (
	"log"
	"log/slog"
	"net/http"
	"os"

	"github.com/alexedwards/scs/v2"
)

// go get github.com/go-chi/chi/v5
// go get github.com/go-chi/chi/v5/middleware
// go run ./cmd/web
// go test ./...

type application struct {
	Session *scs.SessionManager
}

func initLog() {
	handlerOptions := &slog.HandlerOptions{
		Level: slog.LevelInfo,
	}

	logger := slog.New(slog.NewTextHandler(os.Stdout, handlerOptions))
	slog.SetDefault(logger)
}

func main() {
	initLog()

	// set up an app config
	app := application{}

	// get a session manager
	app.Session = getSession()

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
