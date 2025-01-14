package routes

import (
	"rbac-go/controllers"
	"rbac-go/middleware"

	"github.com/gofiber/fiber/v2"
)

func ProductRoutes(app *fiber.App){
	app.Get("/products", middleware.AuthMiddleware, controllers.GetAllProducts)
	app.Get("/products/:id", middleware.AuthMiddleware, controllers.GetSingleProduct)
	app.Post("/products", middleware.AuthMiddleware, middleware.IsAdmin, controllers.CreateProduct)
	app.Patch("/products/:id", middleware.AuthMiddleware, middleware.IsAdmin, controllers.UpdateProduct)
	app.Delete("/products/:id", middleware.AuthMiddleware, middleware.IsAdmin , controllers.DeleteProduct)
}
