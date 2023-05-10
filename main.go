package main

import (
	"demoproject/api/cache"
	"demoproject/api/db"
	"demoproject/api/router"
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
	// utils.InsertData()                           //  Parsing and inserting data to postgresql
	if _, err := cache.InitRedis(); err != nil { // initilizing redis
		log.Fatalf("Error Occured in redis : %s", err.Error())
	}

	// creating app
	app := fiber.New(fiber.Config{})

	app.Get("/ping", func(c *fiber.Ctx) error {
		return c.SendString("ok")
	})

	app.Mount("/api", router.MountRoute())

	// starting the app
	log.Fatal(app.Listen(fmt.Sprintf(":%s", viper.GetString("PORT"))))
}
