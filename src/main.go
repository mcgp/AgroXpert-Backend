package main

import (
	"github.com/gofiber/fiber/v2"

	"AgroXpert-Backend/src/configs"
	"AgroXpert-Backend/src/database"
	"AgroXpert-Backend/src/routes"
)

func main() {
	configs.InitEnv()
	database.InitMongoConnection()

	app := fiber.New()
	routes.SetupTestRoutes(app)
	app.Listen(":5000")

	defer database.Db.CloseConnection()
}