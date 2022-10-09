package routes

import (
	"github.com/mixedmachine/GoalsBackend/api/pkg/controllers"

	"fmt"

	"github.com/gofiber/fiber/v2"
)

const apiVersion = "v1"

type Routes interface {
	Install(app *fiber.App)
}

type authRoutes struct {
	goalController controllers.GoalController
}

func NewAuthRoutes(goalController controllers.GoalController) Routes {
	return &authRoutes{
		goalController: goalController,
	}
}

func (r *authRoutes) Install(app *fiber.App) {
	app.Get("/ping", r.goalController.Ping)
	api := app.Group(fmt.Sprintf("/api/%s", apiVersion))

	// Health check
	api.Get("/ping", r.goalController.Ping)

	// Goals management
	goalsGroup := api.Group("/goals")
	goalsGroup.Post("/", r.goalController.PostGoal)
	goalsGroup.Get("/", r.goalController.GetGoals)
	goalsGroup.Get("/:id", r.goalController.GetGoal)
	goalsGroup.Put("/:id", r.goalController.PutGoal)
	goalsGroup.Delete("/:id", r.goalController.DeleteGoal)
}
