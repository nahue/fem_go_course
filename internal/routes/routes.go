package routes

import (
	"github.com/go-chi/chi"
	"github.com/nahue/go_http_server/internal/app"
)

func SetupRoutes(app *app.Application) *chi.Mux {
	r := chi.NewRouter()
	r.Get("/health", app.HealthCheck)
	r.Get("/workouts/{id}", app.WorkoutHandler.HandleGetWorkoutById)
	r.Post("/workouts", app.WorkoutHandler.HandleCreateWorkout)
	return r
}
