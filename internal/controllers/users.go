package controllers

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type UsersController struct{}

// constructor
func NewUsersController() *UsersController {
	return &UsersController{}
}

func (u *UsersController) GetUserByUsername(w http.ResponseWriter, r *http.Request) {
	username := chi.URLParam(r, "username")

	if username == "" { // 404
		http.NotFound(w, r)
		fmt.Fprintf(w, "not found: %s\n", username)
		return
	}

	// TODO: find user in db

	w.WriteHeader(http.StatusOK) // 200
	fmt.Fprintf(w, "found: %s\n", username)
}
