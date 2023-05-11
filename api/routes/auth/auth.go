package auth

import "github.com/gofiber/fiber/v2"

func MountApp() *fiber.App {
	auth := fiber.New() // creating auth app
	auth.Post("/signup", signup)
	auth.Post("/login", login)
	return auth

}
