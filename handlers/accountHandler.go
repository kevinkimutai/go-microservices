package handler

import (
	"simple_bank/db"
	"simple_bank/model"

	"github.com/gofiber/fiber/v2"
)

func CreateAccount(c *fiber.Ctx) error {
	account := new(model.Account)

	if err := c.BodyParser(&account); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	err := db.DB.Create(&account).Error

	if err != nil {
		c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal Server Error",
			"error":   err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(&account)

}
