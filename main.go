package main

import (
	"simple_bank/db"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	//Init Migration
	db.InitMigration()

	app := fiber.New()

	//LOGGER MIDDLEWARE
	app.Use(logger.New())

	//API ROUTES

}
