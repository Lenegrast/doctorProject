package main

import (
	config2 "UnknownProject/config"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
)

func main() {
	r := chi.NewRouter()
	config := config2.ReadConfig()

	r.Get("/schedules", func(rw http.ResponseWriter, r *http.Request) {
		id := r.URL.Query().Get("user_id")
		io.WriteString(rw, fmt.Sprintf("user_id = %s", id))
	})

	// Запуск сервера
	logrus.Infof("Server started. Port %s", config.Port)
	err := http.ListenAndServe("localhost:"+config.Port, r)
	if err != nil {
		logrus.Fatalf("Failed with start server. Error: %s", err)
	}
}
