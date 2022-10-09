package api

import (
	"github.com/mixedmachine/GoalsBackend/api/pkg/controllers"
	"github.com/mixedmachine/GoalsBackend/api/pkg/db"
	"github.com/mixedmachine/GoalsBackend/api/pkg/repository"
	"github.com/mixedmachine/GoalsBackend/api/pkg/routes"

	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
)

func Init() {
	err := godotenv.Load()
	if err != nil {
		log.Panicln(err)
	}
}

func RunUserAuthApiServer() {

	sConn := db.NewSqlConnection()
	defer sConn.Close()

	app := fiber.New()
	app.Use(cors.New())
	app.Use(logger.New())

	goalRepo := repository.NewUserRepository(sConn)
	repos := map[string]interface{}{
		"goals": goalRepo,
	}
	goalController := controllers.NewGoalController(repos)

	authRoutes := routes.NewAuthRoutes(goalController)
	authRoutes.Install(app)

	log.Fatal(app.Listen(":" + os.Getenv("API_PORT")))
}
