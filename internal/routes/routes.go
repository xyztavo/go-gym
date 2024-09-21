package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/xyztavo/go-gym/internal/handlers"
)

func SetupRoutes() *chi.Mux {
	r := chi.NewRouter()
	r.Get("/", handlers.HelloWorld)
	r.Post("/users", handlers.CreateUser)
	r.Post("/auth", handlers.GetAuth)
	return r
}
