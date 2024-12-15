package api

import (
	"github.com/adityaputra42/e-commerce_backend/routes"
	"github.com/gofiber/fiber/v2"
)

type TransactionsController interface {
	Create(c *fiber.Ctx) error
	Update(c *fiber.Ctx) error
	Delete(c *fiber.Ctx) error
	GetAll(c *fiber.Ctx) error
	GetById(c *fiber.Ctx) error
}

type TransactionsControllerImpl struct {
	Server routes.Server
}

// Create implements TransactionsController.
func (t *TransactionsControllerImpl) Create(c *fiber.Ctx) error {
	panic("unimplemented")
}

// Delete implements TransactionsController.
func (t *TransactionsControllerImpl) Delete(c *fiber.Ctx) error {
	panic("unimplemented")
}

// GetAll implements TransactionsController.
func (t *TransactionsControllerImpl) GetAll(c *fiber.Ctx) error {
	panic("unimplemented")
}

// GetById implements TransactionsController.
func (t *TransactionsControllerImpl) GetById(c *fiber.Ctx) error {
	panic("unimplemented")
}

// Update implements TransactionsController.
func (t *TransactionsControllerImpl) Update(c *fiber.Ctx) error {
	panic("unimplemented")
}

func NewTransactionsController(server routes.Server) TransactionsController {
	return &TransactionsControllerImpl{Server: server}
}
