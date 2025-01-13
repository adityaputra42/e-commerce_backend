package role

import (
	"fmt"
	"net/http"

	"github.com/adityaputra42/e-commerce_backend/dto"
	"github.com/adityaputra42/e-commerce_backend/helper"
	"github.com/adityaputra42/e-commerce_backend/token"
	"github.com/gofiber/fiber/v2"
)

func MemberAuth(c *fiber.Ctx) error {
	authPayload := c.Locals(helper.GetPayloadKey()).(*token.Payload)

	fmt.Printf("user role => %s", authPayload.Role)
	if authPayload.Role != "member" {
		return c.Status(fiber.StatusForbidden).JSON(&dto.ErrorResponse{Status: http.StatusUnauthorized, Message: "You don't have permission to access this resource"})
	}
	return c.Next()
}
