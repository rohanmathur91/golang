package app

import (
	"log"
	"os"
)

type Application struct {
	Logger *log.Logger
}

func RawApplication() (*Application, error) {
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	app := &Application{
		Logger: logger,
	}

	return app, nil
}
