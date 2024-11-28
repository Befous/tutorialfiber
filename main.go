package main

import (
	"github.com/Befous/tutorialfiber/middleware"
	"github.com/Befous/tutorialfiber/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()
	app.Use(cors.New(middleware.CorsAllowAll))
	routes.ContohRoute(app)
	app.Listen(":3000")
}
