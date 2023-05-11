package todo

import (
	"demoproject/api/middleware"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func MountApp() *fiber.App {
	app := fiber.New()
	app.Use(middleware.SessionMiddleware()) // using session middleware

	app.Get("/", func(c *fiber.Ctx) error {
		fmt.Println(c.Locals("userid"))
		return c.SendString("hello")
	})
	app.Post("/create", createTodo)
	app.Get("/read", readTodo)
	app.Delete("/delete", deletetodo)
	app.Put("/update", updatetodo)
	return app
}
