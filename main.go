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
	var port int
	flag.IntVar(&port, "port", 8080, "backend server port")
	flag.Parse()

	fmt.Println("starting new app")
	app, err := app.New()

	if err != nil {
		panic("something went wrong!")
	}

	defer app.DB.Close() // when everything is executed in main then at last run this

	app.Logger.Println("app started")
	app.Logger.Println("starting server...")

	r := routes.SetupRoutes(app)

	server := http.Server{
		Addr:        fmt.Sprintf(":%d", port),
		IdleTimeout: time.Minute,
		Handler:     r,
	}

	app.Logger.Printf("server running on port %d...", port)

	err = server.ListenAndServe()

	if err != nil {
		panic("Something went wrong with server!")
	}
}
