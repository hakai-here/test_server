package middleware

import (
	"demoproject/api/cache"
	"demoproject/api/models"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func SessionMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		session := c.Cookies("session_id")
		if session != "" {
			value, err := cache.AuthGetKey(session)
			if err != nil {
				log.Println(err.Error())
			} else if value.Authenticated {
				c.Locals("userid", value.UserID)
				return c.Next()
			}
		}

		return c.Status(http.StatusUnauthorized).JSON(models.Error{Error: "No session defined, please login"})
	}
}
