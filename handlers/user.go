package handlers

import (
	"github.com/darkykek/db"
	"github.com/darkykek/models"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

func GetUsers(c *fiber.Ctx) error {

	query := bson.D{{}}
	cursor, err := db.Mg.Database.Collection("users").Find(c.Context(), query)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	var users []models.User = make([]models.User, 0)
	if err := cursor.All(c.Context(), &users); err != nil {
		return c.Status(500).SendString(err.Error())
	}

	if len(users) <= 0 {
		return c.Status(200).JSON(fiber.Map{"msg": "No user has been created yet"})
	}

	return c.Status(200).JSON(users)
}

func GetUser(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"Update": "User",
	})
}

func UpdateUser(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"Update": "User",
	})
}

func DeleteUser(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"Delete": "User",
	})
}
