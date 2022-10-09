package main

import (
	"github.com/mixedmachine/GoalsBackend/recommender/cmd/v1/api"
)

func main() {
	api.Init()
	api.RunApiServer()
}
