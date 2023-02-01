package main

import (
	_ "fiberever/docs"
	"fiberever/middleware"
	"fiberever/model"
	"fiberever/router"
	"log"

	"github.com/alisholihindev/go-lib"
	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

// @title Swagger for Fiber-Ever
// @version 1.0
// @description Swagger for backend API service
// @description Get the Bearer token on the Authentication Service
// @description JSON Link: <a href=/swagger/doc.json>docs.json</a>

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization

// @docExpansion none
// @BasePath /api

func main() {
	// Load .env file
	if err := godotenv.Load(".env"); err != nil {
		panic("Error loading .env file")
	}

	app := fiber.New()

	// Middlewares.
	middleware.FiberMiddleware(app)
	//establish pooling connection
	lib.DBConn = lib.DBEstablish()

	//set swagger
	app.Get("/swagger/*", swagger.HandlerDefault)

	router.SetupRoutes(app)
	router.NotFoundRoute(app)

	if err := lib.DBConn.AutoMigrate(&model.User{}); err != nil {
		log.Fatalln(err)
	}

	log.Fatal(app.Listen(":3000"))
}
