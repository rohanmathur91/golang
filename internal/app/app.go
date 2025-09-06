package app

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/rohanmathur91/golang/internal/controllers"
	"github.com/rohanmathur91/golang/internal/db"
	"github.com/rohanmathur91/golang/internal/db/migrations"
)

type Application struct {
	Logger          *log.Logger
	UsersController *controllers.UsersController
	DB              *sql.DB
}

func New() (*Application, error) {
	pgdb, err := db.Open()

	if err != nil {
		return nil, err
	}

	err = db.MigrateFS(pgdb, migrations.FS, ".")

	if err != nil {
		panic(err)
	}

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	// models - db

	// controllers will go here
	usersController := controllers.NewUsersController()

	app := &Application{
		Logger:          logger,
		UsersController: usersController,
		DB:              pgdb,
	}

	return app, nil
}

func (a *Application) HealthCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "health check\n")
}
