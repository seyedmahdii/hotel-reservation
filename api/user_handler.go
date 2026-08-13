package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/seyedmahdii/hotel-reservation/types"
)

func HandleGetUsers(c *fiber.Ctx) error {
	user := types.User{
		FirstName: "Mahdi",
		LastName:  "Jalali",
	}

	return c.JSON(user)
}

func HandleGetUser(c *fiber.Ctx) error {
	return c.JSON(map[string]string{"user": "Mahdi Khan"})
}
