package controllers

import (
	"context"
	"ecommerce-api/models"
	"ecommerce-api/utils"
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetProducts(c *fiber.Ctx) error {
	collection := utils.DB.Collection("products")
	cursor, err := collection.Find(context.Background(), bson.D{})
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": fmt.Sprintf("Error retrieving products: %v", err),
		})
	}
	defer cursor.Close(context.Background())

	var products []models.Product
	if err = cursor.All(context.Background(), &products); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": fmt.Sprintf("Error decoding products: %v", err),
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"data": products,
	})
}

func CreateProduct(c *fiber.Ctx) error {
	var product models.Product
	if err := c.BodyParser(&product); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": fmt.Sprintf("Error parsing request body: %v", err),
		})
	}

	collection := utils.DB.Collection("products")
	result, err := collection.InsertOne(context.Background(), product)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": fmt.Sprintf("Error creating product: %v", err),
		})
	}

	switch id := result.InsertedID.(type) {
	case primitive.ObjectID:
		product.ID = id
	default:
		product.ID = primitive.NewObjectID()
	}

	return c.Status(http.StatusCreated).JSON(fiber.Map{
		"data": product,
	})
}
