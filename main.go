package main

import (
	"kingscross/pkg/configs"
	"kingscross/pkg/middlewares"
	"kingscross/pkg/routes"
	"kingscross/pkg/utils"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	config := configs.FiberConfig()

	app := fiber.New(config)

	middlewares.FiberMiddleware(app)

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	routes.SignInRoute(app)
	routes.NotFoundRoute(app)

	if os.Getenv("STAGE_STATUS") == "dev" {
		utils.StartServer(app)
	}
}
