package api

import (
	"github.com/adityaputra42/e-commerce_backend/routes"
	"github.com/gofiber/fiber/v2"
)

type PaymentMethodController interface {
	Create(c *fiber.Ctx) error
	Update(c *fiber.Ctx) error
	Delete(c *fiber.Ctx) error
	GetAll(c *fiber.Ctx) error
	GetById(c *fiber.Ctx) error
}

type PaymentMethodControllerImpl struct {
	Server routes.Server
}

// Create implements PaymentMethodController.
func (p *PaymentMethodControllerImpl) Create(c *fiber.Ctx) error {
	panic("unimplemented")
}

// Delete implements PaymentMethodController.
func (p *PaymentMethodControllerImpl) Delete(c *fiber.Ctx) error {
	panic("unimplemented")
}

// GetAll implements PaymentMethodController.
func (p *PaymentMethodControllerImpl) GetAll(c *fiber.Ctx) error {
	panic("unimplemented")
}

// GetById implements PaymentMethodController.
func (p *PaymentMethodControllerImpl) GetById(c *fiber.Ctx) error {
	panic("unimplemented")
}

// Update implements PaymentMethodController.
func (p *PaymentMethodControllerImpl) Update(c *fiber.Ctx) error {
	panic("unimplemented")
}

func NewPaymentMethodController(server routes.Server) PaymentMethodController {
	return &PaymentMethodControllerImpl{Server: server}
}
