package app

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/rohanmathur91/golang/internal/controllers"
	"github.com/rohanmathur91/golang/internal/store"
)

type Application struct {
	Logger          *log.Logger
	UsersController *controllers.UsersController
	DB              *sql.DB
}

func RawApplication() (*Application, error) {
	db, err := store.Open()

	if err != nil {
		return nil, err
	}

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	// store - db

	// handlers
	usersController := controllers.NewUsersController()

	app := &Application{
		Logger:          logger,
		UsersController: usersController,
		DB:              db,
	}

	return app, nil
}

func (a *Application) HealthCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "health check\n")
}
