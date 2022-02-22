package userRoutes

import (
	"github.com/gofiber/fiber/v2"

	"app/models/user"
)

func SetUserRoutes(api fiber.Router) {
	userGroup := api.Group("/user")
	userGroup.Post("/create", user.CreateUser)
	userGroup.Post("/login", user.Login)
}
