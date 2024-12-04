package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/xyztavo/go-gym/internal/handlers"
	"github.com/xyztavo/go-gym/internal/middlewares"
)

func SetupRoutes() *chi.Mux {
	r := chi.NewRouter()
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"https://*", "http://*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))
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
	r.Get("/user/gym/details", handlers.GetUserGymDetails)
	r.Get("/user/plan/details", handlers.GetUserPlanDetails)
	r.Post("/user/gym/check-in", handlers.CheckIn)
	r.Get("/user/gym/plans", handlers.GetUserGymPlans)
	r.Get("/user/gym/routines", handlers.GetUserGymRoutines)
	r.Get("/exercises", handlers.GetExercises)
	r.Get("/collections", handlers.GetCollections)
	r.Get("/collections/{id}/exercises-reps", handlers.GetExercisesRepsCollectionsByCollectionId)
	r.Get("/routines/{id}/collections", handlers.GetCollectionsByRoutineId)
	r.Get("/exercises-reps/collections", handlers.GetExercisesRepsCollections)
	r.Get("/routines", handlers.GetRoutines)
}

func AdminRouter(r chi.Router) {
	r.Use(middlewares.AdminAuthMiddleware)
	r.Get("/admin/test-auth", handlers.TestAuth)
	r.Post("/users/gym-admin", handlers.SetUserGymAdmin)
	r.Post("/users/gym-admin/email", handlers.SetUserGymAdminByEmail)
}

func AdminOrGymAdminRouter(r chi.Router) {
	r.Use(middlewares.AdminOrGymAdminAuthMiddleware)
	r.Post("/exercises", handlers.CreateExercise)
	r.Post("/collections", handlers.CreateCollection)
	r.Get("/user/collections", handlers.GetAdminCollections)
	r.Get("/user/routines", handlers.GetUserRoutines)
	r.Post("/exercises-reps/collections", handlers.CreateExercisesRepsCollection)
	r.Post("/exercises-reps/collections/multiple", handlers.CreateMultipleExercisesRepCollection)
	r.Post("/routines", handlers.CreateRoutine)
	r.Post("/routines/{routineId}/collections/{collectionId}", handlers.CreateRoutineCollection)
}

func GymAdminRouter(r chi.Router) {
	r.Use(middlewares.GymAdminAuthMiddleware)
	r.Get("/gym/admin/test-auth", handlers.TestAuth)
	r.Post("/gym", handlers.CreateGym)
	r.Post("/gym/plans", handlers.CreatePlan)
	r.Post("/gym/user", handlers.SetGymUser)
	r.Post("/gym/user/email", handlers.SetGymUserByEmail)
	r.Get("/gym/users", handlers.GetGymUsers)
	r.Patch("/gym/user/plan", handlers.SetUserPlan)
	r.Post("/gym/routines", handlers.CreateGymRoutine)
}
