package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/mixedmachine/GoalsBackend/api/pkg/models"
	"github.com/mixedmachine/GoalsBackend/api/pkg/repository"
	"github.com/mixedmachine/GoalsBackend/api/pkg/util"
	"github.com/nats-io/nats.go"
)

// GoalController defines the interface for user controller
type GoalController interface {
	Ping(ctx *fiber.Ctx) error
	PostGoal(ctx *fiber.Ctx) error
	GetGoal(ctx *fiber.Ctx) error
	GetGoals(ctx *fiber.Ctx) error
	PutGoal(ctx *fiber.Ctx) error
	DeleteGoal(ctx *fiber.Ctx) error
}

type goalController struct {
	goalsRepo repository.GoalsRepository
	client    *http.Client
}

func NewGoalController(repos map[string]interface{}) GoalController {
	return &goalController{
		goalsRepo: repos["goals"].(repository.GoalsRepository),
		client:    &http.Client{}, // TODO: use timeout
	}
}

func (c *goalController) Ping(ctx *fiber.Ctx) error {
	return ctx.
		Status(http.StatusOK).
		JSON(map[string]string{"message": "pong"})
}

/********************************************************
 *				Handler Functions for Goals				*
 ********************************************************/

func (c *goalController) PostGoal(ctx *fiber.Ctx) error {
	user, err := AuthRequest(ctx, c.client)
	if err != nil {
		return ctx.
			Status(http.StatusInternalServerError).
			JSON(util.NewJError(err))
	}
	goal := models.NewUserGoal(user)
	err = ctx.BodyParser(&goal)
	if err != nil {
		return ctx.
			Status(http.StatusBadRequest).
			JSON(util.NewJError(err))
	}

	err = c.goalsRepo.Create(goal)
	if err != nil {
		return ctx.
			Status(http.StatusInternalServerError).
			JSON(util.NewJError(err))
	}

	extendGoal(goal)

	return ctx.
		Status(http.StatusOK).
		JSON(goal)
}

func (c *goalController) GetGoal(ctx *fiber.Ctx) error {
	user, err := AuthRequest(ctx, c.client)
	if err != nil {
		return ctx.
			Status(http.StatusInternalServerError).
			JSON(util.NewJError(err))
	}
	goal, err := c.goalsRepo.RetrieveById(user, ctx.Params("id"))
	if err != nil {
		return ctx.
			Status(http.StatusInternalServerError).
			JSON(util.NewJError(err))
	}
	return ctx.
		Status(http.StatusOK).
		JSON(goal)
}

func (c *goalController) GetGoals(ctx *fiber.Ctx) error {
	user, err := AuthRequest(ctx, c.client)
	if err != nil {
		return ctx.
			Status(http.StatusInternalServerError).
			JSON(util.NewJError(err))
	}
	goals, err := c.goalsRepo.RetrieveAll(user)
	if err != nil {
		return ctx.
			Status(http.StatusInternalServerError).
			JSON(util.NewJError(err))
	}
	return ctx.
		Status(http.StatusOK).
		JSON(goals)
}

func (c *goalController) PutGoal(ctx *fiber.Ctx) error {
	user, err := AuthRequest(ctx, c.client)
	if err != nil {
		return ctx.
			Status(http.StatusInternalServerError).
			JSON(util.NewJError(err))
	}
	if user == "internal" {
		user = ctx.Params("user")
	}

	goal := models.NewUserGoal(user)
	err = ctx.BodyParser(&goal)
	if err != nil {
		return ctx.
			Status(http.StatusBadRequest).
			JSON(util.NewJError(err))
	}
	goal.Id = ctx.Params("id")
	goal.UpdatedAt = time.Now()
	err = c.goalsRepo.Update(goal)
	if err != nil {
		return ctx.
			Status(http.StatusInternalServerError).
			JSON(util.NewJError(err))
	}

	return ctx.
		Status(http.StatusOK).
		JSON(goal)
}

func (c *goalController) DeleteGoal(ctx *fiber.Ctx) error {
	user, err := AuthRequest(ctx, c.client)
	if err != nil {
		return ctx.
			Status(http.StatusInternalServerError).
			JSON(util.NewJError(err))
	}
	goal := models.NewUserGoal(user)
	goal.Id = ctx.Params("id")
	err = c.goalsRepo.Delete(goal.UserId, goal.Id)
	if err != nil {
		return ctx.
			Status(http.StatusInternalServerError).
			JSON(util.NewJError(err))
	}

	return ctx.
		Status(http.StatusNoContent).
		JSON(goal)
}

type ExtendGoalRequest struct {
	ReturnEndpoint string      `json:"return_endpoint"`
	Goal           models.Goal `json:"goal"`
}

func extendGoal(goal *models.Goal) {
	log.Printf("Extending goal %s...\n", goal.Id)
	msg := &ExtendGoalRequest{
		ReturnEndpoint: os.Getenv("API_HOST") + os.Getenv("API_PORT") + "/api/v1/goals/" + goal.Id,
		Goal:           *goal,
	}
	msgBytes, err := json.Marshal(msg)
	if err != nil {
		log.Println("error marshalling message to extend goal: ", err)
		return
	}
	nc, err := nats.Connect(os.Getenv("NATS_URL"))
	if err != nil {
		log.Println("error connecting to nats: ", err)
		return
	}
	err = nc.Publish(os.Getenv("NATS_GOAL_REC_TOPIC"), msgBytes)
	if err != nil {
		log.Println("error publishing message to extend goal: ", err)
		return
	}
	log.Println("message successfully published to extend goal")
}
