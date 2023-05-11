package todo

import (
	"demoproject/api/cache"
	"demoproject/api/db"
	"demoproject/api/models"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func deletetodo(c *fiber.Ctx) error {
	id := c.Query("id")
	uid := c.Locals("userid").(string)
	if id == "" {
		return c.Status(http.StatusBadRequest).JSON(models.Error{Error: "Provide an id to delete the requested data"})
	}

	if err := db.Deletetodo(id); err != nil {
		return c.Status(http.StatusBadRequest).JSON(models.Error{Error: err.Error()})
	}
	if err := cache.DeleteKey(id); err != nil {
		log.Printf("Unable to delete : %s", err.Error())
	}
	if err := cache.DeleteKey(uid); err != nil {
		log.Printf("Unable to delete : %s", err.Error())
	}
	return c.JSON(models.Success{Success: "successfully deleted the id"})
}
