package main

import (
	"github.com/mixedmachine/GoalsBackend/api/cmd/v1/api"
)

func main() {
	api.Init()
	api.RunUserAuthApiServer()
}
