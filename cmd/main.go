package main

import (
	"UnknownProject/internal/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/sirupsen/logrus"
	"net/http"
)

var r = chi.NewRouter()

func main() {
	startServer()
	r.Get("/schedules/", handlers.Example)
}

func startServer() {
	err := http.ListenAndServe("localhost:8080", r)
	if err != nil {
		logrus.Fatalf("Failed with start server. Error: %s", err)
	}
	logrus.Info("Server started")
}
