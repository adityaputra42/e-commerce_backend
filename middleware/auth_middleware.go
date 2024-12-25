package middleware

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/adityaputra42/e-commerce_backend/dto"
	"github.com/adityaputra42/e-commerce_backend/helper"
	"github.com/adityaputra42/e-commerce_backend/token"
	"github.com/gofiber/fiber/v2"
)

func  AuthMiddleware(c *fiber.Ctx) error {
	tokenMaker := token.GetTokenMaker()
	authorizationHeader := c.Get(helper.GetHeaderKey())
	if len(authorizationHeader) == 0 {
		err := errors.New("authorization header is not provided")
		return c.Status(http.StatusUnauthorized).JSON(dto.ErrorResponse{
			Status:  http.StatusUnauthorized,
			Message: err.Error(),
		})
	}

	fields := strings.Fields(authorizationHeader)
	if len(fields) < 2 {
		err := errors.New("invalid authorization header format")
		return c.Status(http.StatusUnauthorized).JSON(dto.ErrorResponse{
			Status:  http.StatusUnauthorized,
			Message: err.Error(),
		})
	}

	authorizationType := strings.ToLower(fields[0])
	if authorizationType != helper.GetTypeBearer() {
		err := fmt.Errorf("unsupported authorization type %s", authorizationType)
		return c.Status(http.StatusUnauthorized).JSON(dto.ErrorResponse{
			Status:  http.StatusUnauthorized,
			Message: err.Error(),
		})
	}

	accessToken := fields[1]
	payload, err := tokenMaker.VerifyToken(accessToken)
	if err != nil {
		return c.Status(http.StatusUnauthorized).JSON(dto.ErrorResponse{
			Status:  http.StatusUnauthorized,
			Message: err.Error(),
		})
	}
	c.Locals(helper.GetPayloadKey(), payload)
	return c.Next()

}

// func AddAuthorization(
// 	t *testing.T,
// 	request *http.Request,
// 	tokenMaker token.Maker,
// 	authorizationType string,
// 	username string,
// 	uid string,
// 	duration time.Duration,
// ) {
// 	token, err := tokenMaker.CreateToken(username, uid, duration)
// 	require.NoError(t, err)
// 	authorizationHeader := fmt.Sprintf("%s %s", authorizationType, token)
// 	request.Header.Set(helper.GetHeaderKey(), authorizationHeader)
// }
