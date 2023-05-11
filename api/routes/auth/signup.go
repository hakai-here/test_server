package auth

import (
	"demoproject/api/db"
	"demoproject/api/models"
	"demoproject/api/utils"
	"encoding/json"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

// the function below will deals with the signup functionality
// future function :
// 	- adding validating functions to validate the userinput
//	- cache existing usernames in database to redis for increase performance

func signup(c *fiber.Ctx) error {
	var signup models.User

	if err := json.Unmarshal(c.Body(), &signup); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(models.Error{Error: "unable to parse the data"})
	}

	exists, err := db.CheckUsername(signup.Username)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(models.Error{Error: err.Error()})
	} else if exists {
		return c.Status(http.StatusBadRequest).JSON(models.Error{Error: "Username already found in the database"})
	}

	signup.ID = utils.GenerateUUID()                    // generating random username
	signup.Password = utils.HashArgon2(signup.Password) // generating hash password
	if err := db.InsertUser(signup); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(models.Error{Error: "unable to insert the data to the database"})
	}
	return c.JSON(models.Success{Success: "Created the user Successfully"})
}
