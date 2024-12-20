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

type SizeVarianController interface {
	Create(c *fiber.Ctx) error
	Update(c *fiber.Ctx) error
	Delete(c *fiber.Ctx) error
	GetAll(c *fiber.Ctx) error
	GetById(c *fiber.Ctx) error
}

type SizeVarianControllerImpl struct {
	Server Server
}

// Create implements SizeVarianController.
func (s *SizeVarianControllerImpl) Create(c *fiber.Ctx) error {

	req := new(request.CreateSizeVarianProduct)

	err := c.BodyParser(req)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(dto.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: "Invalid Message Body",
		})
	}

	colorVarian, err := s.Server.Store.GetColorVarianProduct(c.Context(), req.ColorVarianId)

	if err != nil {
		return c.Status(404).JSON(dto.ErrorResponse{
			Status:  404,
			Message: "Color varian doesn't exist",
		})
	}
	param := db.CreateSizeVarianProductParams{
		ColorVarianID: colorVarian.ID,
		Size:          req.Size,
		Stock:         req.Stock,
	}
	sizeVarian, err := s.Server.Store.CreateSizeVarianProduct(c.Context(), param)
	if err != nil {
		return c.Status(403).JSON(dto.ErrorResponse{
			Status:  403,
			Message: "Failed Create Size Varian",
		})
	}

	return c.Status(http.StatusCreated).JSON(dto.SuccessResponse{
		Status:  http.StatusCreated,
		Message: "Success Create Size varian",
		Data:    helper.ToSizeVarianResponse(sizeVarian),
	})
}

// Delete implements SizeVarianController.
func (s *SizeVarianControllerImpl) Delete(c *fiber.Ctx) error {
	sizeId := c.Params("id")
	id, err := strconv.Atoi(sizeId)
	if err != nil {
		return c.Status(400).JSON(dto.ErrorResponse{
			Status:  400,
			Message: err.Error(),
		})
	}
	err = s.Server.Store.DeleteSizeVarianProduct(c.Context(), int64(id))
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

// GetAll implements SizeVarianController.
func (s *SizeVarianControllerImpl) GetAll(c *fiber.Ctx) error {
	var listSize []response.SizeVarianResponse
	colorId := c.Params("color_id")
	id, err := strconv.Atoi(colorId)
	if err != nil {
		return c.Status(400).JSON(dto.ErrorResponse{
			Status:  400,
			Message: err.Error(),
		})
	}
	page := c.QueryInt("page")
	limit := c.QueryInt("limit", 10)

	param := db.ListSizeVarianProductParams{
		ColorVarianID: int64(id),
		Offset:        int32(page),
		Limit:         int32(limit),
	}
	sizes, err := s.Server.Store.ListSizeVarianProduct(c.Context(), param)

	if err != nil {
		return c.Status(400).JSON(dto.ErrorResponse{
			Status:  400,
			Message: err.Error(),
		})
	}

	for _, v := range sizes {
		listSize = append(listSize, helper.ToSizeVarianResponse(v))
	}

	return c.Status(http.StatusCreated).JSON(dto.SuccessResponse{
		Status:  http.StatusCreated,
		Message: "Success",
		Data:    listSize,
	})
}

// GetById implements SizeVarianController.
func (s *SizeVarianControllerImpl) GetById(c *fiber.Ctx) error {
	sizeId := c.Params("id")
	id, err := strconv.Atoi(sizeId)
	if err != nil {
		return c.Status(400).JSON(dto.ErrorResponse{
			Status:  400,
			Message: err.Error(),
		})
	}

	size, err := s.Server.Store.GetSizeVarianProduct(c.Context(), int64(id))

	if err != nil {
		return c.Status(400).JSON(dto.ErrorResponse{
			Status:  400,
			Message: err.Error(),
		})
	}

	return c.Status(http.StatusCreated).JSON(dto.SuccessResponse{
		Status:  http.StatusCreated,
		Message: "Success",
		Data:    helper.ToSizeVarianResponse(size),
	})
}

// Update implements SizeVarianController.
func (s *SizeVarianControllerImpl) Update(c *fiber.Ctx) error {
	req := new(request.UpdateSizeVarianProduct)

	err := c.BodyParser(req)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(dto.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: "Invalid Message Body",
		})
	}
	data, err := s.Server.Store.GetSizeVarianProductForUpdate(c.Context(), req.ID)

	if err != nil {
		return c.Status(400).JSON(dto.ErrorResponse{
			Status:  400,
			Message: err.Error(),
		})
	}

	param := db.UpdateSizeVarianProductParams{
		ID:    data.ID,
		Size:  req.Size,
		Stock: req.Stock,
	}

	sizeUpdate, err := s.Server.Store.UpdateSizeVarianProduct(c.Context(), param)

	if err != nil {
		return c.Status(400).JSON(dto.ErrorResponse{
			Status:  400,
			Message: err.Error(),
		})
	}

	return c.Status(http.StatusCreated).JSON(dto.SuccessResponse{
		Status:  http.StatusCreated,
		Message: "Success",
		Data:    helper.ToSizeVarianResponse(sizeUpdate),
	})
}

func NewSizeVarianController(server Server) SizeVarianController {
	return &SizeVarianControllerImpl{Server: server}
}
