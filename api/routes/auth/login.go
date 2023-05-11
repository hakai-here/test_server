package auth

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
	"gorm.io/gorm"
)

// the function below deals with login functions

func login(c *fiber.Ctx) error {
	var login models.Login
	cookie := c.Cookies("session_id") // checking the cookies to resist duplicate authentication
	if cookie != "" {
		data, err := cache.AuthGetKey(cookie)
		if err != nil {
			log.Println(err.Error())
		}
		if data.Authenticated {
			return c.JSON(fiber.Map{"error": "Session is already authenticated"})
		}
	}
	if err := json.Unmarshal(c.Body(), &login); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(models.Error{Error: "Unable to unmarshal the data"})
	}
	data, err := db.GetUserDetials(login.Username) // getting the user data from postgres database for further authentication
	if err == gorm.ErrRecordNotFound {
		return c.Status(http.StatusBadRequest).JSON(models.Error{Error: "Username doesnot exists in the database"})
	} else if err != nil {
		return c.Status(http.StatusBadRequest).JSON(models.Error{Error: err.Error()})
	}
	login.Password = utils.HashArgon2(login.Password) // hashing password
	if login.Password == data.Password {
		rkey := utils.RandStr() // creating a random key
		rvalue := models.RedisType{
			UserID:        data.ID,
			Authenticated: true,
		}
		rtime := time.Hour
		if err := cache.AuthSetKey(rkey, rvalue, rtime); err != nil { // registering the session to the redis database
			return c.Status(http.StatusInternalServerError).JSON(models.Error{Error: err.Error()})
		}
		c.Cookie(&fiber.Cookie{Name: "session_id", Value: rkey, Expires: time.Now().Add(rtime)}) // setting up the cookie
		sessiondata := models.SessionAuthdata{
			SessionId: rkey,
			Validity:  int(rtime / time.Millisecond),
		}
		return c.Status(http.StatusOK).JSON(sessiondata) // returning
	}
	return c.Status(http.StatusUnauthorized).JSON(models.Error{Error: "Unable to authenticate the user"})
}
