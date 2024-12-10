package api

import (
	"fmt"

	db "github.com/adityaputra42/e-commerce_backend/db/sqlc"
	"github.com/adityaputra42/e-commerce_backend/dto"
	"github.com/adityaputra42/e-commerce_backend/dto/request"
	"github.com/adityaputra42/e-commerce_backend/routes"
	"github.com/gofiber/fiber/v2"
)

type ProductController interface {
	CreateProduct(c *fiber.Ctx) error
	UpdateProduct(c *fiber.Ctx) error
	Delete(c *fiber.Ctx) error
	FetchProduct(c *fiber.Ctx) error
	FetchListProduct(c *fiber.Ctx) error
}

type ProductControllerImpl struct {
	Server routes.Server
}

// CreateProduct implements ProductController.
func (p *ProductControllerImpl) CreateProduct(c *fiber.Ctx) error {
	return p.Server.Store.ExecTx(c.Context(), func(q *db.Queries) error {
		req := new(request.CreateProduct)
		err := c.BodyParser(req)
		if err != nil {
			return c.Status(400).JSON(dto.ErrorResponse{
				Status:  400,
				Message: "Invalid form data",
			})
		}
		var uploadPath string
		if req.Images != nil {
			file := req.Images
			uploadPath = fmt.Sprintf("./uploads/products/%s", file.Filename)
			if err := c.SaveFile(file, uploadPath); err != nil {
				return fiber.NewError(fiber.StatusInternalServerError, "failed to save product image")
			}
			req.Images = nil
		}

		return nil
	})
}

// Delete implements ProductController.
func (p *ProductControllerImpl) Delete(c *fiber.Ctx) error {
	panic("unimplemented")
}

// FetchListProduct implements ProductController.
func (p *ProductControllerImpl) FetchListProduct(c *fiber.Ctx) error {
	panic("unimplemented")
}

// FetchProduct implements ProductController.
func (p *ProductControllerImpl) FetchProduct(c *fiber.Ctx) error {
	panic("unimplemented")
}

// UpdateProduct implements ProductController.
func (p *ProductControllerImpl) UpdateProduct(c *fiber.Ctx) error {
	panic("unimplemented")
}

func NewProductController(server routes.Server) ProductController {
	return &ProductControllerImpl{Server: server}
}
