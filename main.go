package main

import (
	"demoproject/api/cache"
	"demoproject/api/db"
	"demoproject/api/routes/auth"
	"demoproject/api/routes/todo"
	"demoproject/api/utils"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
)

func main() {
	utils.ImportEnv() // loading envionment variables

	if err := db.ConnectDB(); err != nil { // connecting the postgresql database
		log.Fatalf("Error : %s", err.Error())
	}
	if err := cache.InitConnection(); err != nil {
		log.Fatalf("Error : %s", err.Error())
	}

	app := fiber.New(fiber.Config{})

	app.Get("/ping", func(c *fiber.Ctx) error {
		return c.SendString("ok")
	})

	app.Mount("/auth", auth.MountApp())
	app.Mount("/todo", todo.MountApp())

	// starting the app
	log.Fatal(app.Listen(fmt.Sprintf(":%s", viper.GetString("PORT"))))
}
