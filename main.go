package main

import (
	"flag"
	"fmt"
	"net/http"
	"time"

	"github.com/rohanmathur91/golang/internal/app"
	"github.com/rohanmathur91/golang/internal/routes"
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

	r := routes.SetupRoutes(app)

	server := http.Server{
		Addr:        fmt.Sprintf(":%d", port),
		IdleTimeout: time.Minute,
		Handler:     r,
	}

	app.Logger.Printf("server running on port %d...", port)

	error = server.ListenAndServe()

	if error != nil {
		panic("Something went wrong with server!")
	}
}
