package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/rohanmathur91/golang/internal/app"
)

func main() {
	fmt.Println("starting new app")
	app, error := app.RawApplication()

	if error != nil {
		panic("something went wrong!")
	}

	app.Logger.Println("app started")
	app.Logger.Println("starting server")

	http.HandleFunc("/health", HealthCheck)

	server := http.Server{
		Addr:        ":8080",
		IdleTimeout: time.Minute,
	}

	error = server.ListenAndServe()

	if error != nil {
		panic("Something went wrong with server!")
	}

	app.Logger.Println("server running on port 8080...")
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	// fmt.Println(w, "health check")
	fmt.Fprintf(w, "health check\n")
}
