package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/xyztavo/go-gym/internal/handlers"
	"github.com/xyztavo/go-gym/internal/middlewares"
)

func SetupRoutes() *chi.Mux {
	r := chi.NewRouter()
	r.Get("/", handlers.HelloWorld)
	r.Post("/users", handlers.CreateUser)
	r.Post("/auth", handlers.GetAuth)

	r.Mount("/", AuthRouter())
	r.Mount("/admin", AdminRouter())
	return r
}

func AuthRouter() http.Handler {
	r := chi.NewRouter()
	r.Use(middlewares.AuthMiddleware)
	r.Get("/testauth", handlers.TestAuth)
	return r
}

func AdminRouter() http.Handler {
	r := chi.NewRouter()
	r.Use(middlewares.AdminAuthMiddleware)
	r.Get("/testauth", handlers.TestAuth)
	r.Post("/users/gym-admin", handlers.SetUserGymAdmin)
	return r
}
