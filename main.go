package main

import (
	"fmt"
	"golang/internal/app"
)

func main() {
	fmt.Println("Starting new app")
	app, error := app.RawApplication()

	if error != nil {
		panic("Something went wrong!")
	}

	app.Logger.Println("App started")
}
