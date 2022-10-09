package api

import (
	"log"
	"os"

	"github.com/mixedmachine/GoalsBackend/recommender/pkg/controllers"
	"github.com/mixedmachine/GoalsBackend/recommender/pkg/mb"

	"github.com/joho/godotenv"
)

const Subscription = "goals"

func Init() {
	err := godotenv.Load()
	if err != nil {
		log.Panicln(err)
	}
}

func RunApiServer() {

	/*
		create a new message connection
	*/
	Conn := mb.NewMessageConnection()
	defer Conn.Close()

	/*
		tie the controller to the message connection
	*/
	// controller that uses the repo
	goalController := controllers.NewGoalController()

	// subscriber that uses controller
	var sub string
	if sub = os.Getenv("NATS_GOALS_REC_TOPIC"); sub == "" {
		sub = Subscription
	}
	Conn.Subscribe(sub, goalController.GoalExtendHandler)

	/*
		start the message connection listener
	*/
	Conn.Listen()

}
