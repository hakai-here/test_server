package todo

import (
	"demoproject/api/db"
	"demoproject/api/models"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func deletetodo(c *fiber.Ctx) error {
	id := c.Query("id")

	if id == "" {
		return c.Status(http.StatusBadRequest).JSON(models.Error{Error: "Provide an id to delete the requested data"})
	}

	if err := db.Deletetodo(id); err != nil {
		return c.Status(http.StatusBadRequest).JSON(models.Error{Error: err.Error()})
	}

	return c.JSON(models.Success{Success: "successfully deleted the id"})
}
