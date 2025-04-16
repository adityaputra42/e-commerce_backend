package api

import (
	"database/sql"
	"net/http"

	db "github.com/adityaputra42/e-commerce_backend/db/sqlc"
	"github.com/adityaputra42/e-commerce_backend/dto"
	"github.com/adityaputra42/e-commerce_backend/dto/request"
	"github.com/adityaputra42/e-commerce_backend/dto/response"
	"github.com/adityaputra42/e-commerce_backend/helper"
	"github.com/adityaputra42/e-commerce_backend/token"
	"github.com/gofiber/fiber/v2"
)

type UserController interface {
	CreateUser(c *fiber.Ctx) error
	CreateAdmin(c *fiber.Ctx) error
	UpdatePassword(c *fiber.Ctx) error
	Delete(c *fiber.Ctx) error
	Login(c *fiber.Ctx) error
	FetchUser(c *fiber.Ctx) error
	FetchAllUSer(c *fiber.Ctx) error
}

type UserControllerImpl struct {
	server Server
}

// CreateAdmin implements UserController.
func (u *UserControllerImpl) CreateAdmin(c *fiber.Ctx) error {
	req := new(request.CreateUser)

	err := c.BodyParser(req)
	if err != nil {
		return c.Status(400).JSON(dto.ErrorResponse{
			Status:  400,
			Message: "Invalid Message Body",
		})
	}

	hasPassword, err := helper.HashPassword(req.Password)

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(dto.ErrorResponse{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	userParam := db.CreateUserParams{
		Uid:      helper.Generate("UID"),
		FullName: req.FullName,
		Email:    req.Email,
		Username: req.Username,
		Password: hasPassword,
		Role:     "admin",
	}
	user, err := u.server.Store.CreateUser(c.Context(), userParam)
	if err != nil {
		return c.Status(403).JSON(dto.ErrorResponse{
			Status:  403,
			Message: "Failed Create User",
		})
	}

	return c.Status(201).JSON(dto.SuccessResponse{
		Status:  201,
		Message: "Success Create User",
		Data:    helper.ToUserResponse(user),
	})

}

// CreateUser implements UserController.
func (u *UserControllerImpl) CreateUser(c *fiber.Ctx) error {
	req := new(request.CreateUser)

	err := c.BodyParser(req)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(dto.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: "Invalid Message Body",
		})
	}

	hasPassword, err := helper.HashPassword(req.Password)

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(dto.ErrorResponse{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	userParam := db.CreateUserParams{
		Uid:      helper.Generate("UID"),
		FullName: req.FullName,
		Email:    req.Email,
		Username: req.Username,
		Password: hasPassword,
		Role:     "member",
	}
	user, err := u.server.Store.CreateUser(c.Context(), userParam)
	if err != nil {
		return c.Status(403).JSON(dto.ErrorResponse{
			Status:  403,
			Message: "Failed Create User",
		})
	}

	return c.Status(http.StatusCreated).JSON(dto.SuccessResponse{
		Status:  http.StatusCreated,
		Message: "Success Create User",
		Data:    helper.ToUserResponse(user),
	})
}

// Delete implements UserController.
func (u *UserControllerImpl) Delete(c *fiber.Ctx) error {
	uid := c.Params("uid")

	err := u.server.Store.DeleteUser(c.Context(), uid)
	if err != nil {
		return c.Status(400).JSON(dto.ErrorResponse{
			Status:  400,
			Message: err.Error(),
		})
	}
	return c.Status(200).JSON(dto.ErrorResponse{
		Status:  200,
		Message: "Ok",
	})

}

// FetchAllUSer implements UserController.
func (u *UserControllerImpl) FetchAllUSer(c *fiber.Ctx) error {
	userList := []response.UserResponse{}
	users, err := u.server.Store.ListUser(c.Context(), db.ListUserParams{Role: "user"})
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(dto.ErrorResponse{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	for _, value := range users {
		userList = append(userList, helper.ToUserResponse(value))

	}
	return c.Status(http.StatusOK).JSON(dto.SuccessResponse{
		Status:  http.StatusOK,
		Message: "Ok",
		Data:    userList,
	})

}

// FetchUser implements UserController.
func (u *UserControllerImpl) FetchUser(c *fiber.Ctx) error {
	authPayload := c.Locals(helper.GetPayloadKey()).(*token.Payload)

	response, err := u.server.Store.GetUser(c.Context(), authPayload.Username)
	if err != nil {
		return c.Status(400).JSON(dto.ErrorResponse{
			Status:  400,
			Message: err.Error(),
		})
	}
	return c.Status(200).JSON(dto.SuccessResponse{
		Status:  200,
		Message: "Success",
		Data:    helper.ToUserResponse(response),
	})
}

// Login implements UserController.
func (u *UserControllerImpl) Login(c *fiber.Ctx) error {
	req := new(request.LoginUser)
	err := c.BodyParser(req)
	if err != nil {
		return c.Status(400).JSON(dto.ErrorResponse{
			Status:  400,
			Message: "Invalid Message Body",
		})
	}
	user, err := u.server.Store.GetUserLogin(c.Context(), req.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			return c.Status(http.StatusNotFound).JSON(dto.ErrorResponse{
				Status:  http.StatusNotFound,
				Message: "User not found",
			})
		}
		return c.Status(http.StatusInternalServerError).JSON(dto.ErrorResponse{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	err = helper.CheckPassword(req.Password, user.Password)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(dto.ErrorResponse{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	accessToken, _, err := u.server.TokenMaker.CreateToken(user.Username, user.Uid, user.Role, u.server.Config.AccessTokenDuration)

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(dto.ErrorResponse{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	refreshToken, refreshPayload, err := u.server.TokenMaker.CreateToken(user.Username, user.Uid, user.Role, u.server.Config.RefreshTokenDuration)

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(dto.ErrorResponse{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	session := db.CreateSessionParams{
		ID:           refreshPayload.ID,
		RefreshToken: refreshToken,
		UserUid:      user.Uid,
		UserAgent:    "",
		ClientIp:     "",
		IsBlocked:    false,
		ExpiredAt:    refreshPayload.ExpiredAt,
	}
	_, err = u.server.Store.CreateSession(c.Context(), session)

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(dto.ErrorResponse{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	return c.Status(200).JSON(dto.SuccessResponse{
		Status:  200,
		Message: "Success",
		Data: response.LoginResponse{
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
			User:         helper.ToUserResponse(user),
		},
	})
}

// UpdatePassword implements UserController.
func (u *UserControllerImpl) UpdatePassword(c *fiber.Ctx) error {
	req := new(request.UpdateUser)
	authPayload := c.Locals(helper.GetPayloadKey()).(*token.Payload)
	err := c.BodyParser(req)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(dto.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: "Invalid Message Body",
		})
	}

	user, err := u.server.Store.GetUserForUpdate(c.Context(), authPayload.Uid)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(dto.ErrorResponse{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		})
	}
	err = helper.CheckPassword(req.OldPassword, user.Password)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(dto.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: "hash and password doesn't match",
		})
	}

	user, err = u.server.Store.UpdateUser(c.Context(), db.UpdateUserParams{Uid: user.Uid, Password: req.Password})
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(dto.ErrorResponse{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		})
	}
	return c.Status(http.StatusOK).JSON(dto.SuccessResponse{
		Status:  http.StatusOK,
		Message: "Ok",
		Data:    helper.ToUserResponse(user),
	})
}

func NewUserController(server Server) UserController {
	return &UserControllerImpl{server: server}
}
