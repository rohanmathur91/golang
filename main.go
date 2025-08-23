package main

import (
	"flag"
	"fmt"
	"net/http"
	"time"

	"github.com/rohanmathur91/golang/internal/app"
)

func main() {
	// 1. parse command line port flag
	var port int
	flag.IntVar(&port, "port", 8080, "backend server port")
	flag.Parse()

	// 2. then start the server
	fmt.Println("starting new app")
	app, error := app.RawApplication()

	if error != nil {
		panic("something went wrong!")
	}

	app.Logger.Println("app started")
	app.Logger.Println("starting server")

	http.HandleFunc("/health", HealthCheck)

	server := http.Server{
		Addr:        fmt.Sprintf(":%d", port),
		IdleTimeout: time.Minute,
	}

	app.Logger.Printf("server running on port %d...", port)

	error = server.ListenAndServe()

	if error != nil {
		panic("Something went wrong with server!")
	}
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	// fmt.Println(w, "health check")
	fmt.Fprintf(w, "health check\n")
}
