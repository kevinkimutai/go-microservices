package main

import (
	"log"
	"os"
	"simple_bank/db"
	router "simple_bank/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"

	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	//Init Migration
	db.InitMigration()

	app := fiber.New()

	//Env variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	//LOGGER MIDDLEWARE
	app.Use(logger.New())

	//API ROUTES
	app.Route("/api/v1/account", router.AccountRouter)

	PORT := os.Getenv("PORT")

	app.Listen(":" + PORT)

}
