package routes

import (
	"ecommerce-api/controllers"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

func Setup(app *fiber.App, db *mongo.Database) {
	app.Get("/api/products", controllers.GetProducts)
	app.Post("/api/products", controllers.CreateProduct)
}
