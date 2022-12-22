package routes

import (
	"kingscross/app/controllers/auth"

	"github.com/gofiber/fiber/v2"
)

func SignInRoute(app *fiber.App) {
	route := app.Group("/v1")

	route.Post("/auth/sign/up", auth.SignUp)
}
