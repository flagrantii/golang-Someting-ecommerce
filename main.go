package main

import (
	"ecommerce-api/routes"
	"ecommerce-api/utils"

	"github.com/gofiber/fiber/v2"
)

func main() {

	db, err := utils.ConnectDB()
	if err != nil {
		panic(err)
	}

	app := fiber.New()

	routes.Setup(app, db)

	app.Listen(":8000")
}
