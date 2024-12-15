package api

import (
	"github.com/adityaputra42/e-commerce_backend/routes"
	"github.com/gofiber/fiber/v2"
)

type SizeVarianController interface {
	Create(c *fiber.Ctx) error
	Update(c *fiber.Ctx) error
	Delete(c *fiber.Ctx) error
	GetAll(c *fiber.Ctx) error
	GetById(c *fiber.Ctx) error
}

type SizeVarianControllerImpl struct {
	Server routes.Server
}

// Create implements SizeVarianController.
func (s *SizeVarianControllerImpl) Create(c *fiber.Ctx) error {
	panic("unimplemented")
}

// Delete implements SizeVarianController.
func (s *SizeVarianControllerImpl) Delete(c *fiber.Ctx) error {
	panic("unimplemented")
}

// GetAll implements SizeVarianController.
func (s *SizeVarianControllerImpl) GetAll(c *fiber.Ctx) error {
	panic("unimplemented")
}

// GetById implements SizeVarianController.
func (s *SizeVarianControllerImpl) GetById(c *fiber.Ctx) error {
	panic("unimplemented")
}

// Update implements SizeVarianController.
func (s *SizeVarianControllerImpl) Update(c *fiber.Ctx) error {
	panic("unimplemented")
}

func NewSizeVarianController(server routes.Server) SizeVarianController {
	return &SizeVarianControllerImpl{Server: server}
}
