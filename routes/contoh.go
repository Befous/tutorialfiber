package routes

import (
	"github.com/Befous/tutorialfiber/handler"
	"github.com/gofiber/fiber/v2"
)

func ContohRoute(page *fiber.App) {
	page.Get("/routeget/:z", handler.ContohHandler)
	page.Post("/routepost/:z", handler.ContohHandler)
	page.Put("/routeput/:z", handler.ContohHandler)
	page.Delete("/routedelete/:z", handler.ContohHandler)
}
