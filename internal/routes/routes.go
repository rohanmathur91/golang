package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/rohanmathur91/golang/internal/app"
)

func SetupRoutes(app *app.Application) *chi.Mux {
	r := chi.NewRouter()

	r.Get("/health", app.HealthCheck)

	// users
	r.Get("/users/{username}", app.UsersController.GetUserByUsername)

	return r
}
