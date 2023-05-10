package router

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func MountRoute() *fiber.App {
	route := fiber.New(fiber.Config{}) // creating a new app for handeling routes
	route.Hooks().OnMount(func(a *fiber.App) error {
		log.Println("[SERVER] Route '/' is connected to the server")
		return nil
	})

	route.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Connected app")
	})
	route.Get("/proceedingentries", GetAllData)

	return route
}
