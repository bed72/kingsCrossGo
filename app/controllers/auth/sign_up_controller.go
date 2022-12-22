package auth

import (
	"fmt"
	"kingscross/app/models/request"
	"kingscross/app/models/responses"
	"kingscross/pkg/utils"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/imroc/req/v3"
)

func SignUp(context *fiber.Ctx) error {
	request := &request.SignUpRequestModel{}

	var response responses.SignUpResponseModel
	var failure responses.FailureMessageResponseModel

	if err := context.BodyParser(request); err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(
			responses.NewFailureResponse(
				fiber.StatusBadRequest,
				err.Error(),
			),
		)
	}

	validate := utils.NewValidator()

	if err := validate.Struct(request); err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(
			responses.NewFailureResponse(
				fiber.StatusBadRequest,
				err.Error(),
			),
		)
	}

	url := fmt.Sprintf("%s/%s", os.Getenv("BASE_URL"), "auth/v1/signup")

	client := req.C().DevMode()
	resp, err := client.R().
		SetHeader("apiKey", os.Getenv("API_KEY")).
		SetBody(&request).
		SetResult(&response).
		SetError(&failure).
		Post(url)

	if err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(
			responses.NewFailureResponse(
				fiber.StatusBadRequest,
				err.Error(),
			),
		)
	}

	if resp.IsError() {
		response := &responses.ResponseModel[responses.FailureMessageResponseModel]{
			Status: resp.StatusCode,
			Data:   failure,
		}

		return context.Status(resp.StatusCode).JSON(
			response,
		)
	}

	return context.Status(resp.StatusCode).JSON(
		responses.NewSignUpResponse(
			resp.StatusCode, response,
		),
	)
}
