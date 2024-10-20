package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/xyztavo/go-gym/internal/handlers"
	"github.com/xyztavo/go-gym/internal/middlewares"
)

func SetupRoutes() *chi.Mux {
	r := chi.NewRouter()
	r.Get("/", handlers.HelloWorld)
	r.Post("/users", handlers.CreateUser)
	r.Post("/auth", handlers.GetAuth)
	r.Group(AdminRouter)
	r.Group(AdminOrGymAdminRouter)
	r.Group(AuthRouter)
	r.Group(GymAdminRouter)
	return r
}

func AuthRouter(r chi.Router) {
	r.Use(middlewares.AuthMiddleware)
	r.Get("/testauth", handlers.TestAuth)
	r.Get("/user/gym", handlers.GetUserGym)
	r.Post("/user/gym/check-in", handlers.CheckIn)
	r.Get("/user/gym/plans", handlers.GetUserGymPlans)
}

func AdminRouter(r chi.Router) {
	r.Use(middlewares.AdminAuthMiddleware)
	r.Get("/admin/testauth", handlers.TestAuth)
	r.Post("/users/gym-admin", handlers.SetUserGymAdmin)
}

func AdminOrGymAdminRouter(r chi.Router) {
	r.Use(middlewares.AdminOrGymAdminAuthMiddleware)
	r.Post("/exercises", handlers.CreateExercise)
	r.Post("/routines", handlers.CreateRoutine)
	r.Post("/routines/exercises", handlers.CreateExerciseRoutine)
}

func GymAdminRouter(r chi.Router) {
	r.Use(middlewares.GymAdminAuthMiddleware)
	r.Post("/gym", handlers.CreateGym)
	r.Post("/gym/plans", handlers.CreatePlan)
	r.Post("/gym/user", handlers.SetGymUser)
	r.Get("/gym/users", handlers.GetGymUsers)
	r.Patch("/gym/user/plan", handlers.SetUserPlan)
}
