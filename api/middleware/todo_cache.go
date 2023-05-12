package middleware

import (
	"demoproject/api/cache"
	"log"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func CacheMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		err := c.Next()

		if strings.HasPrefix(c.Path(), "/todo/delete") || strings.HasPrefix(c.Path(), "/todo/update") {
			id := c.Query("id")
			uid := c.Locals("userid").(string)

			if id != "" {
				if err := cache.DeleteKey(id); err != nil {
					log.Println(err)
				}

			}
			if err := cache.DeleteKey(uid); err != nil {
				log.Println(err)
			}
		}

		return err
	}
}
