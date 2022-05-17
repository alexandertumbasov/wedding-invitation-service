package controllers

import (
	"context"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
	"time"
	"wedding-invitation-service/configs"
	"wedding-invitation-service/models"
	"wedding-invitation-service/responses"
)

var guestsCollection *mongo.Collection = configs.GetCollection(configs.DB, "guests")
var validate = validator.New()

func CreateGuest(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var guest models.Guest
	defer cancel()

	//validate the request body
	if err := c.BodyParser(&guest); err != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	if validationErr := validate.Struct(&guest); validationErr != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": validationErr.Error()}})
	}

	newGuest := bson.M{
		"name":     guest.Name,
		"surname":  guest.Surname,
		"message":  guest.Message,
		"isActive": guest.IsActive,
	}

	result, err := guestsCollection.InsertOne(ctx, newGuest)

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.UserResponse{Status: http.StatusInternalServerError, Message: "error when insert to DB", Data: &fiber.Map{"data": err.Error()}})
	}

	return c.Status(http.StatusCreated).JSON(responses.UserResponse{Status: http.StatusCreated, Message: "Guest created!", Data: &fiber.Map{"data": result}})

}
