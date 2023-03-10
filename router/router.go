package router

import (
	"fiberever/controller"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api", logger.New())

	// Auth
	auth := api.Group("/auth")
	auth.Post("/login", controller.Login)

	//user
	user := api.Group("/user")
	user.Post("/", controller.CreateUser)
	user.Get("/:id", controller.GetUser)
	user.Put("/:id", controller.UpdateUser)
	user.Delete("/:id", controller.DeleteUser)

	//dokter
	dokter := api.Group("/dokter")
	dokter.Get("/", controller.GetDokterAll)
	dokter.Get("/:id", controller.GetDokter)
	dokter.Post("/", controller.AddDokter)
	dokter.Delete("/:id", controller.DeleteDokter)
	dokter.Put("/:id", controller.UpdateDokter)
}
