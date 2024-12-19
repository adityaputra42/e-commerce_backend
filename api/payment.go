package api

import (
	"github.com/adityaputra42/e-commerce_backend/routes"
	"github.com/gofiber/fiber/v2"
)

type PaymentController interface {
	Create(c *fiber.Ctx) error
	Update(c *fiber.Ctx) error
	Delete(c *fiber.Ctx) error
	GetAll(c *fiber.Ctx) error
	GetById(c *fiber.Ctx) error
}

type PaymentControllerImpl struct {
	Server routes.Server
}

// Create implements PaymentController.
func (p *PaymentControllerImpl) Create(c *fiber.Ctx) error {
	return nil

}

// Delete implements PaymentController.
func (p *PaymentControllerImpl) Delete(c *fiber.Ctx) error {
	panic("unimplemented")
}

// GetAll implements PaymentController.
func (p *PaymentControllerImpl) GetAll(c *fiber.Ctx) error {
	panic("unimplemented")
}

// GetById implements PaymentController.
func (p *PaymentControllerImpl) GetById(c *fiber.Ctx) error {
	panic("unimplemented")
}

// Update implements PaymentController.
func (p *PaymentControllerImpl) Update(c *fiber.Ctx) error {
	panic("unimplemented")
}

func NewPaymentController(server routes.Server) PaymentController {
	return &PaymentControllerImpl{Server: server}
}
