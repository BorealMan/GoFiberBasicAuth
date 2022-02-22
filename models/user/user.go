package user

import (
	"fmt"

	"github.com/gofiber/fiber/v2"

	"app/api/auth"
)

func Login(c *fiber.Ctx) error {

	userId := "1"
	userRole := "free"

	t, err := auth.IssueJWT(userId, userRole)

	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(fiber.Map{"token": t})
}

func CreateUser(c *fiber.Ctx) error {
	email := c.FormValue("email")

	fmt.Println(email)

	return c.JSON(fiber.Map{
		"sucess": "yay",
	})
}
