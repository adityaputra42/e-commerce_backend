package api

import (
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
	panic("unimplemented")
}

// CreateUser implements UserController.
func (u *UserControllerImpl) CreateUser(c *fiber.Ctx) error {
	panic("unimplemented")
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
	return &UserControllerImpl{
		server: server,
	}
}
