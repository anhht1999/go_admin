package controller

import (
	"github.com/anhht1999/go_admin/models"
	"github.com/gofiber/fiber/v2"
)

func Register(c *fiber.Ctx) error {

	user := models.User{
		FirstName: "John",
	}
	user.LastName="Doe"
	return c.JSON(user)
  }