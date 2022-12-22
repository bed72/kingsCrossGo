package routes

import (
	"kingscross/app/models/responses"

	"github.com/gofiber/fiber/v2"
)

func NotFoundRoute(app *fiber.App) {
	app.Use(
		func(context *fiber.Ctx) error {
			return context.Status(fiber.StatusNotFound).JSON(
				responses.NewFailureResponse(
					fiber.StatusNotFound,
					"sorry, endpoint is not found",
				),
			)
		},
	)
}
