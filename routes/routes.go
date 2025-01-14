package routes

import (
	"rbac-go/middleware"

	"github.com/gofiber/fiber/v2"
)

func Routes(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	authRoutes(app)
	ProductRoutes(app)

	app.Get("/protected", middleware.AuthMiddleware, func(c *fiber.Ctx) error {
		userId := c.Locals("user_id").(uint)
		role := c.Locals("role").(string)
		return c.JSON(fiber.Map{
			"message": "Access granted",
			"user_id": userId,
			"role": role,
		})
	})
}