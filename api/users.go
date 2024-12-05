package api

import (
	db "github.com/adityaputra42/e-commerce_backend/db/sqlc"
	"github.com/adityaputra42/e-commerce_backend/helper"
	"github.com/adityaputra42/e-commerce_backend/model"
	"github.com/adityaputra42/e-commerce_backend/model/request"
	"github.com/adityaputra42/e-commerce_backend/routes"
	"github.com/gofiber/fiber"
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
	server routes.Server
}

// CreateAdmin implements UserController.
func (u *UserControllerImpl) CreateAdmin(c *fiber.Ctx) error {
	req := new(request.CreateUser)

	err := c.BodyParser(req)
	if err != nil {
		return c.Status(403).JSON(model.ErrorResponse{
			Status:  403,
			Message: "Invalid Message Body",
		})
	}

	userParam := db.CreateUserParams{
		Uid:      helper.Generate("UID"),
		FullName: req.FullName,
		Email:    req.Email,
		Username: req.Username,
		Password: req.Password,
		Role:     "admin",
	}
	user, err := u.server.Store.CreateUser(c.Context(), userParam)
	if err != nil {
		return c.Status(403).JSON(model.ErrorResponse{
			Status:  403,
			Message: "Failed Create User",
		})
	}

	return c.Status(201).JSON(model.SuccessResponse{
		Status:  201,
		Message: "Success Create User",
		Data:    user,
	})

}

// CreateUser implements UserController.
func (u *UserControllerImpl) CreateUser(c *fiber.Ctx) error {
	req := new(request.CreateUser)

	err := c.BodyParser(req)
	if err != nil {
		return c.Status(403).JSON(model.ErrorResponse{
			Status:  403,
			Message: "Invalid Message Body",
		})
	}

	userParam := db.CreateUserParams{
		Uid:      helper.Generate("UID"),
		FullName: req.FullName,
		Email:    req.Email,
		Username: req.Username,
		Password: req.Password,
		Role:     "user",
	}
	user, err := u.server.Store.CreateUser(c.Context(), userParam)
	if err != nil {
		return c.Status(403).JSON(model.ErrorResponse{
			Status:  403,
			Message: "Failed Create User",
		})
	}

	return c.Status(201).JSON(model.SuccessResponse{
		Status:  201,
		Message: "Success Create User",
		Data:    user,
	})
}

// Delete implements UserController.
func (u *UserControllerImpl) Delete(c *fiber.Ctx) error {
	panic("unimplemented")
}

// FetchAllUSer implements UserController.
func (u *UserControllerImpl) FetchAllUSer(c *fiber.Ctx) error {
	panic("unimplemented")
}

// FetchUser implements UserController.
func (u *UserControllerImpl) FetchUser(c *fiber.Ctx) error {
	panic("unimplemented")
}

// Login implements UserController.
func (u *UserControllerImpl) Login(c *fiber.Ctx) error {
	panic("unimplemented")
}

// UpdatePassword implements UserController.
func (u *UserControllerImpl) UpdatePassword(c *fiber.Ctx) error {
	panic("unimplemented")
}

func NewUserController(server routes.Server) UserController {
	return &UserControllerImpl{server: server}
}
