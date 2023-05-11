package todo

import (
	"demoproject/api/cache"
	"demoproject/api/db"
	"demoproject/api/models"
	"log"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
)

func readTodo(c *fiber.Ctx) error {
	uid := c.Locals("userid").(string)
	id := c.Query("id")
	var data []models.TodoActivity
	var err error

	if id != "" {
		data, err = cache.TodoGetKey(id)
		if err == nil {
			return c.Status(http.StatusOK).JSON(data)
		}
	} else {
		data, err = cache.TodoGetKey(uid)
		if err == nil && len(data) > 0 {
			return c.Status(http.StatusOK).JSON(data)
		}
	}

	data, err = db.GetTodo(uid, id)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(models.Error{Error: "unable to fetch any data"})
	}
	if len(data) == 0 {
		return c.Status(http.StatusNotFound).JSON(models.Error{Error: "Requested resource not found"})
	} else {
		if err := cache.TodoSetKey(uid, data, 2*time.Hour); err != nil {
			log.Println("Error : unable to cache the data ") // unable to cache data of user --> continue
		}
		return c.JSON(data)
	}

}
