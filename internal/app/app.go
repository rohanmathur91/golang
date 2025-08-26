package app

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/rohanmathur91/golang/internal/controllers"
)

type Application struct {
	Logger          *log.Logger
	UsersController *controllers.UsersController
}

func RawApplication() (*Application, error) {
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	usersController := controllers.NewUsersController()

	app := &Application{
		Logger:          logger,
		UsersController: usersController,
	}

	return app, nil
}

func (a *Application) HealthCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "health check\n")
}
