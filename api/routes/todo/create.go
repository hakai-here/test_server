package todo

import (
	"demoproject/api/cache"
	"demoproject/api/db"
	"demoproject/api/models"
	"demoproject/api/utils"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
)

// future work
// - Put validations to the input
// todo create activity

func createTodo(c *fiber.Ctx) error {
	var data models.TodoActivity
	if err := json.Unmarshal(c.Body(), &data); err != nil {
		return c.Status(http.StatusBadRequest).JSON(models.Error{Error: "Unable to parse the data"})
	}
	data.ID = utils.GenerateUUID()
	data.UserID = c.Locals("userid").(string) // converting to string
	if err := db.InsertTodo(data); err != nil {
		log.Println(err.Error())
		return c.Status(http.StatusInternalServerError).JSON(models.Error{Error: "Unable to Create the todo activity"})
	}
	datam := []models.TodoActivity{data}
	if err := cache.DeleteKey(data.UserID); err != nil {
		log.Printf("No key found : %s", err.Error())
	}
	if err := cache.TodoSetKey(data.ID, datam, 2*time.Hour); err != nil {
		log.Printf("Unable to cache : %s", err.Error())
	}

	return c.JSON(models.Success{Success: "created the entry"})
}
