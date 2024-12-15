package api

import (
	"github.com/adityaputra42/e-commerce_backend/routes"
	"github.com/gofiber/fiber/v2"
)

type OrderController interface {
	Create(c *fiber.Ctx) error
	Update(c *fiber.Ctx) error
	Delete(c *fiber.Ctx) error
	GetAll(c *fiber.Ctx) error
	GetById(c *fiber.Ctx) error
}

type OrderControllerImpl struct {
	Server routes.Server
}

// Create implements OrderController.
func (o *OrderControllerImpl) Create(c *fiber.Ctx) error {
	panic("unimplemented")
}

// Delete implements OrderController.
func (o *OrderControllerImpl) Delete(c *fiber.Ctx) error {
	panic("unimplemented")
}

// GetAll implements OrderController.
func (o *OrderControllerImpl) GetAll(c *fiber.Ctx) error {
	panic("unimplemented")
}

// GetById implements OrderController.
func (o *OrderControllerImpl) GetById(c *fiber.Ctx) error {
	panic("unimplemented")
}

// Update implements OrderController.
func (o *OrderControllerImpl) Update(c *fiber.Ctx) error {
	panic("unimplemented")
}

func NewOrderController(server routes.Server) OrderController {
	return &OrderControllerImpl{Server: server}
}
