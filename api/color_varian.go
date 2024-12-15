package api

import (
	"github.com/adityaputra42/e-commerce_backend/routes"
	"github.com/gofiber/fiber/v2"
)

type ColorVarianController interface {
	Create(c *fiber.Ctx) error
	Update(c *fiber.Ctx) error
	Delete(c *fiber.Ctx) error
	GetALl(c *fiber.Ctx) error
	GetById(c *fiber.Ctx) error
}

type ColorVarianControllerImpl struct {
	Server routes.Server
}

// Create implements ColorVarianController.
func (*ColorVarianControllerImpl) Create(c *fiber.Ctx) error {
	panic("unimplemented")
}

// Delete implements ColorVarianController.
func (*ColorVarianControllerImpl) Delete(c *fiber.Ctx) error {
	panic("unimplemented")
}

// GetALl implements ColorVarianController.
func (*ColorVarianControllerImpl) GetALl(c *fiber.Ctx) error {
	panic("unimplemented")
}

// GetById implements ColorVarianController.
func (*ColorVarianControllerImpl) GetById(c *fiber.Ctx) error {
	panic("unimplemented")
}

// Update implements ColorVarianController.
func (*ColorVarianControllerImpl) Update(c *fiber.Ctx) error {
	panic("unimplemented")
}

func NewColorVarianController(server routes.Server) ColorVarianController {
	return &ColorVarianControllerImpl{Server: server}
}
