package controllers

import (
	"github.com/mixedmachine/GoalsBackend/recommender/pkg/models"

	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	gpt3 "github.com/sashabaranov/go-gpt3"
	"github.com/nats-io/nats.go"
)

// GoalController defines the interface for user controller
type GoalController interface {
	GoalExtendHandler(m *nats.Msg)
}

type goalController struct {
	client *http.Client
}

func NewGoalController() GoalController {
	return &goalController{
		client: &http.Client{},
	}
}

/********************************************************
 *				Handler Functions for Goals				*
 ********************************************************/

func (c *goalController) GoalExtendHandler(m *nats.Msg) {
	var message models.Message
	log.Println("----------------------------------------------------")
	println("*** GoalExtendHandler ***")

	// decode the message
	err := json.Unmarshal(m.Data, &message)
	if err != nil {
		log.Println(err.Error())
	}
	fmt.Printf("Message:\n")
	fmt.Printf("%s\n", message.ReturnEndpoint)
	goal, err := json.MarshalIndent(message.Goal, "", "    ")
	fmt.Printf("%s\n", goal)

	// extend the goal
	extendedGoal := c.getGoalExtension(message.Goal.Content)
	fmt.Printf("Result: %s\n", extendedGoal)

	// build body for update request
	reqGoal := models.Goal{
		Id:        message.Goal.Id,
		UserId:    message.Goal.UserId,
		Completed: message.Goal.Completed,
		Content:   message.Goal.Content,
		Extended:  extendedGoal,
	}
	marshaledGoal, err := json.Marshal(reqGoal)
	body := bytes.NewBuffer(marshaledGoal)

	// send the update request
	req, err := http.NewRequest("PUT", message.ReturnEndpoint, body)
	if err != nil {
		log.Println(err.Error())
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", os.Getenv("SERVICE_AUTH_TOKEN"))
	resp, err := c.client.Do(req)
	if err != nil {
		log.Println(err.Error())
	}
	log.Printf("%s\n", resp.Status)
}

// GPT3RequestBodySchema defines the request body for the GPT-3 API
type GPT3RequestBodySchema struct {
	Model       string  `json:"model"`
	Instruction string  `json:"instruction"`
	Input       string  `json:"input"`
	Temperature float32 `json:"temperature"`
	MaxTokens   int     `json:"max_tokens"`
}

func (c *goalController) getGoalExtension(initialGoal string) (extendedGoal string) {
	// create a new get request
	prompt := "Expand on the goal"
	gpt3Params := GPT3RequestBodySchema{
		Model:       os.Getenv("GPT3_MODEL"),
		Instruction: prompt,
		Input:       initialGoal,
		Temperature: 0.8,
		MaxTokens:   100,
	}
	gpt3Client := gpt3.NewClient(os.Getenv("GPT3_TOKEN"))
	ctx := context.Background()

	// create the request
	req := gpt3.CompletionRequest{
		Model:       gpt3Params.Model,
		Prompt:      gpt3Params.Instruction + ": " + gpt3Params.Input,
		Temperature: gpt3Params.Temperature,
		MaxTokens:   gpt3Params.MaxTokens,
	}

	// send the request
	resp, err := gpt3Client.CreateCompletion(ctx, req)
	if err != nil {
		log.Println(err.Error())
	}

	// return the response
	return strings.TrimSpace(resp.Choices[0].Text)
}
