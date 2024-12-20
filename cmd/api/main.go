package main

import (
	"fmt"
	"net/http"
	"simple-go-server/internals/routes"
	"simple-go-server/pkg/constant"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

func main() {
	constant.LoadEnv()

	r := chi.NewRouter()
	// A good base middleware stack
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
		r.Mount("/", routes.Router())

	fmt.Println("Started the Go server on port " + constant.PORT)
	http.ListenAndServe(":"+constant.PORT, r)
}