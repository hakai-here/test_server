package todo

import (
	"demoproject/api/db"
	"demoproject/api/models"
	"encoding/json"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func updatetodo(c *fiber.Ctx) error {
	id := c.Query("id")
	uid := c.Locals("userid").(string)
	if id == "" {
		return c.Status(http.StatusBadRequest).JSON(models.Error{Error: "no id was provided"})
	}
	var data models.TodoActivity
	if err := json.Unmarshal(c.Body(), &data); err != nil {
		return c.Status(http.StatusBadRequest).JSON(models.Error{Error: "unable to parse the data"})
	}

	data.ID = id
	data.UserID = uid
	if err := db.Update(data); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(models.Error{Error: "Unable to update the data"})
	}

	return c.JSON(models.Success{Success: "Successfully updated the data"})

}
