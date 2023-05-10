package router

import (
	"demoproject/api/db"
	"demoproject/api/structs"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func GetAllData(c *fiber.Ctx) error {
	query := c.Query("id")
	var data []structs.Proceedingentry
	var err error
	if query != "" {
		data, err = db.GetQueriedEntry(query)
	} else {
		data, err = db.GetAllEntrys()
	}
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.Status(http.StatusOK).JSON(data)
}
