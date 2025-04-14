package middleware

import (
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
)

func LoggerMiddleware(c *fiber.Ctx) error {
	start := time.Now()
	err := c.Next()
	duration := time.Since(start)

	log.Printf("[%s] %s - %s (%d)  %s - %s", c.Method(), c.OriginalURL(), c.IP(), c.Response().StatusCode(),c.Response().Body(), duration)
	return err
}