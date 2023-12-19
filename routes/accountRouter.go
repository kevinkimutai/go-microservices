package router

import (
	handler "simple_bank/handlers"

	"github.com/gofiber/fiber/v2"
)

func AccountRouter(api fiber.Router) {
	api.Post("/", handler.CreateAccount)
}
