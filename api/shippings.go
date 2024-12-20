package api

import (
	"net/http"
	"strconv"

	db "github.com/adityaputra42/e-commerce_backend/db/sqlc"
	"github.com/adityaputra42/e-commerce_backend/dto"
	"github.com/adityaputra42/e-commerce_backend/dto/request"
	"github.com/adityaputra42/e-commerce_backend/dto/response"
	"github.com/adityaputra42/e-commerce_backend/helper"
	"github.com/gofiber/fiber/v2"
)

type ShippingController interface {
	Create(c *fiber.Ctx) error
	Update(c *fiber.Ctx) error
	Delete(c *fiber.Ctx) error
	GetAll(c *fiber.Ctx) error
	GetById(c *fiber.Ctx) error
}

type ShippingControllerImpl struct {
	Server Server
}

// Create implements ShippingController.
func (s *ShippingControllerImpl) Create(c *fiber.Ctx) error {
	req := new(request.CreateShipping)

	err := c.BodyParser(req)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(dto.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: "Invalid Message Body",
		})
	}

	shippingParam := db.CreateShippingParams{
		Name:  req.Name,
		Price: float64(req.Price),
		State: req.State,
	}
	shipping, err := s.Server.Store.CreateShipping(c.Context(), shippingParam)
	if err != nil {
		return c.Status(403).JSON(dto.ErrorResponse{
			Status:  403,
			Message: "Failed Create Shipping",
		})
	}

	return c.Status(http.StatusCreated).JSON(dto.SuccessResponse{
		Status:  http.StatusCreated,
		Message: "Success Create Shipping",
		Data:    helper.ToShippingRespone(shipping),
	})

}

// Delete implements ShippingController.
func (s *ShippingControllerImpl) Delete(c *fiber.Ctx) error {
	shippingId := c.Params("id")

	id, err := strconv.Atoi(shippingId)
	if err != nil {
		return c.Status(400).JSON(dto.ErrorResponse{
			Status:  400,
			Message: err.Error(),
		})
	}
	err = s.Server.Store.DeleteShipping(c.Context(), int64(id))
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

// GetAll implements ShippingController.
func (s *ShippingControllerImpl) GetAll(c *fiber.Ctx) error {
	var shippings []response.ShippingResponse
	listShipping, err := s.Server.Store.ListShipping(c.Context(), db.ListShippingParams{})

	if err != nil {
		return c.Status(400).JSON(dto.ErrorResponse{
			Status:  400,
			Message: err.Error(),
		})
	}

	for _, v := range listShipping {
		shippings = append(shippings, helper.ToShippingRespone(v))
	}

	return c.Status(http.StatusCreated).JSON(dto.SuccessResponse{
		Status:  http.StatusCreated,
		Message: "Success",
		Data:    shippings,
	})
}

// GetById implements ShippingController.
func (s *ShippingControllerImpl) GetById(c *fiber.Ctx) error {
	shippingId := c.Params("id")

	id, err := strconv.Atoi(shippingId)
	if err != nil {
		return c.Status(400).JSON(dto.ErrorResponse{
			Status:  400,
			Message: err.Error(),
		})
	}
	shipping, err := s.Server.Store.GetShipping(c.Context(), int64(id))

	if err != nil {
		return c.Status(400).JSON(dto.ErrorResponse{
			Status:  400,
			Message: err.Error(),
		})
	}

	return c.Status(http.StatusCreated).JSON(dto.SuccessResponse{
		Status:  http.StatusCreated,
		Message: "Success",
		Data:    helper.ToShippingRespone(shipping),
	})
}

// Update implements ShippingController.
func (s *ShippingControllerImpl) Update(c *fiber.Ctx) error {
	req := new(request.UpadateShipping)

	err := c.BodyParser(req)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(dto.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: "Invalid Message Body",
		})
	}
	shipping, err := s.Server.Store.GetShippingForUpdate(c.Context(), req.Id)
	if err != nil {
		return c.Status(400).JSON(dto.ErrorResponse{
			Status:  400,
			Message: err.Error(),
		})
	}

	updateShip := db.UpdateShippingParams{
		ID:    shipping.ID,
		Name:  req.Name,
		Price: float64(req.Price),
		State: req.State,
	}

	shippingUpdate, err := s.Server.Store.UpdateShipping(c.Context(), updateShip)
	if err != nil {
		return c.Status(400).JSON(dto.ErrorResponse{
			Status:  400,
			Message: err.Error(),
		})
	}

	return c.Status(200).JSON(dto.SuccessResponse{
		Status:  200,
		Message: "Success",
		Data:    helper.ToShippingRespone(shippingUpdate),
	})

}

func NewShippingController(server Server) ShippingController {
	return &ShippingControllerImpl{Server: server}
}
