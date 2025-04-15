package api

import (
	"net/http"
	"time"

	"github.com/adityaputra42/e-commerce_backend/dto"
	"github.com/adityaputra42/e-commerce_backend/dto/request"
	"github.com/adityaputra42/e-commerce_backend/dto/response"
	"github.com/gofiber/fiber/v2"
)

type SessionController interface {
	RenewSession(c *fiber.Ctx) error
}

type SessionControllerImpl struct {
	Server Server
}

// RenewSession implements SessionController.
func (s *SessionControllerImpl) RenewSession(c *fiber.Ctx) error {

	req := new(request.RenewAccessTokenRequest)

	err := c.BodyParser(req)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(dto.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: "Invalid Message Body",
		})
	}

	refreshPayload, err := s.Server.TokenMaker.VerifyToken(req.RefreshToken)

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(dto.ErrorResponse{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	session, err := s.Server.Store.GetSessionById(c.Context(), refreshPayload.ID)

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(dto.ErrorResponse{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	if session.IsBlocked {
		return c.Status(http.StatusUnauthorized).JSON(dto.ErrorResponse{
			Status:  http.StatusUnauthorized,
			Message: "Session is blocked",
		})

	}

	if session.UserUid != refreshPayload.Uid {

		return c.Status(http.StatusUnauthorized).JSON(dto.ErrorResponse{
			Status:  http.StatusUnauthorized,
			Message: "Incorrect Session user",
		})
	}

	if session.RefreshToken != req.RefreshToken {
		return c.Status(http.StatusUnauthorized).JSON(dto.ErrorResponse{
			Status:  http.StatusUnauthorized,
			Message: "miss match Session token",
		})
	}

	if time.Now().After(session.ExpiredAt) {

		return c.Status(http.StatusUnauthorized).JSON(dto.ErrorResponse{
			Status:  http.StatusUnauthorized,
			Message: "Expired Session",
		})
	}

	accessToken, accessPayload, err := s.Server.TokenMaker.CreateToken(refreshPayload.Username, session.UserUid, refreshPayload.Role, s.Server.Config.AccessTokenDuration)

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(dto.ErrorResponse{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(dto.SuccessResponse{
		Status:  http.StatusOK,
		Message: "Ok",
		Data: response.SessionResponse{
			AccessToken:          accessToken,
			AccessTokenExpiredAt: accessPayload.ExpiredAt,
		},
	})

}

func NewSessionController(server Server) SessionController {
	return &SessionControllerImpl{Server: server}
}
