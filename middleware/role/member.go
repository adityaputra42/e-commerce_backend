package role

import (
	"fmt"
	"net/http"

	db "github.com/adityaputra42/e-commerce_backend/db/sqlc"
	"github.com/adityaputra42/e-commerce_backend/dto"
	"github.com/gofiber/fiber/v2"
)

func MemberAuth(c *fiber.Ctx) error {
	user := c.Locals("CurrentUser").(db.User)
	fmt.Printf("user role => %s", user.Role)
	if user.Role != "member" {
		return c.Status(fiber.StatusForbidden).JSON(&dto.ErrorResponse{Status: http.StatusUnauthorized, Message: "You don't have permission to access this resource"})
	}
	return c.Next()
}
