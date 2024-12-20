package routes

import (
	"net/http"
	"simple-go-server/internals/controllers"

	"github.com/go-chi/chi"
)

func UserRoutes() http.Handler {
	r := chi.NewRouter()
	r.Post("/", controllers.CreateUser)
	r.Patch("/{id}", controllers.EditUser)
	r.Get("/", controllers.GetUsers)
	return r
}