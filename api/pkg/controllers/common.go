package controllers

import (
	"encoding/json"
	"errors"
	"github.com/gofiber/fiber/v2"
	"io"
	"log"
	"net/http"
	"os"
)

const (
	AuthUrlEndpoint = "/api/v1/auth"
)

type Response struct {
	UserId string `json:"user_id"`
}

func AuthRequest(ctx *fiber.Ctx, c *http.Client) (string, error) {
	var res Response
	log.Println("Authorizing user...")
	token := string(ctx.Request().Header.Peek("Authorization"))
	if token == "" {
		return "", errors.New("Authorization header is missing")
	}
	if token == os.Getenv("SERVICE_AUTH_TOKEN") {
		return "internal", nil
	}

	req, err := http.NewRequest(
		"GET",
		os.Getenv("SERVICE_USER_AUTH_HOST") + AuthUrlEndpoint,
		nil,
	)
	if err != nil {
		return "", err
	}
	req.Header.Add("Authorization", token)

	response, err := c.Do(req)
	if err != nil {
		return "", err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Printf("Error closing body: %s\n", err)
		}
	}(response.Body)

	if response.StatusCode != http.StatusOK {
		return "", errors.New("unauthorized")
	}

	body, _ := io.ReadAll(response.Body)
	err = json.Unmarshal(body, &res)
	if err != nil {
		return "", err
	}

	return res.UserId, nil
}
