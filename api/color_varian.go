package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"

	db "github.com/adityaputra42/e-commerce_backend/db/sqlc"
	"github.com/adityaputra42/e-commerce_backend/dto"
	"github.com/adityaputra42/e-commerce_backend/dto/request"
	"github.com/adityaputra42/e-commerce_backend/dto/response"
	"github.com/adityaputra42/e-commerce_backend/helper"
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
	Server Server
}

// Create implements ColorVarianController.
func (cv *ColorVarianControllerImpl) Create(c *fiber.Ctx) error {
	return cv.Server.Store.ExecTx(c.Context(), func(q *db.Queries) error {
		req := new(request.CreateColorVarianProduct)

		err := c.BodyParser(req)
		if err != nil {
			return c.Status(http.StatusBadRequest).JSON(dto.ErrorResponse{
				Status:  http.StatusBadRequest,
				Message: "Invalid Message Body",
			})
		}
		product, err := q.GetProduct(c.Context(), req.ProductId)

		if err != nil {
			return c.Status(400).JSON(dto.ErrorResponse{
				Status:  400,
				Message: err.Error(),
			})
		}

		names := strings.Fields(product.Name)
		folder := "./assets/product/" + helper.Generate(names[0])

		if err := os.MkdirAll(folder, os.ModePerm); err != nil {
			return c.Status(500).JSON(dto.ErrorResponse{
				Status:  500,
				Message: "Failed create directory",
			})
		}

		file, err := c.FormFile("images")
		var colorVarianPath string
		if err == nil {
			colorVarianPath = folder + file.Filename
			if saveErr := c.SaveFile(file, colorVarianPath); saveErr != nil {
				return c.Status(500).JSON(dto.ErrorResponse{
					Status:  500,
					Message: fmt.Sprintf("Failed to save image for color variant"),
				})
			}
		}

		colorVarianParam := db.CreateColorVarianProductParams{
			ProductID: product.ID,
			Name:      req.Name,
			Color:     req.Color,
			Images:    colorVarianPath,
		}

		colorVarianResult, err := q.CreateColorVarianProduct(c.Context(), colorVarianParam)
		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(dto.ErrorResponse{
				Status:  http.StatusInternalServerError,
				Message: fmt.Sprintf("Failed to create color varian"),
			})
		}
		var resultSizeVarians []response.SizeVarianResponse
		var sizeVarians []request.CreateSizeVarianProduct
		if err := json.Unmarshal([]byte(req.Sizes), &sizeVarians); err != nil {
			return c.Status(400).JSON(dto.ErrorResponse{
				Status:  400,
				Message: "Invalid size_varians format",
			})
		}

		for i := range sizeVarians {
			sizeParam := db.CreateSizeVarianProductParams{
				ColorVarianID: colorVarianResult.ID,
				Size:          sizeVarians[i].Size,
				Stock:         sizeVarians[i].Stock,
			}
			size, err := q.CreateSizeVarianProduct(c.Context(), sizeParam)
			if err != nil {
				return c.Status(http.StatusInternalServerError).JSON(dto.ErrorResponse{
					Status:  http.StatusInternalServerError,
					Message: "Failed to create size varian",
				})
			}
			resultSizeVarians = append(resultSizeVarians, helper.ToSizeVarianResponse(size))
		}

		return c.Status(201).JSON(dto.SuccessResponse{
			Status:  201,
			Message: "Ok",
			Data:    helper.ToColorVarianResponse(colorVarianResult, resultSizeVarians),
		})
	})
}

// Delete implements ColorVarianController.
func (cv *ColorVarianControllerImpl) Delete(c *fiber.Ctx) error {
	return cv.Server.Store.ExecTx(c.Context(), func(q *db.Queries) error {
		productId := c.Params("id")

		id, err := strconv.Atoi(productId)
		if err != nil {
			return c.Status(400).JSON(dto.ErrorResponse{
				Status:  400,
				Message: err.Error(),
			})
		}

		colorVarian, err := q.GetColorVarianProduct(c.Context(), int64(id))
		if err != nil {
			return c.Status(500).JSON(dto.ErrorResponse{
				Status:  500,
				Message: err.Error(),
			})
		}

		var sizeVarians []db.SizeVarian
		if err := json.Unmarshal([]byte(colorVarian.SizeVarians), &sizeVarians); err != nil {
			return c.Status(400).JSON(dto.ErrorResponse{
				Status:  400,
				Message: "Invalid size_varians format",
			})
		}
		for _, size := range sizeVarians {
			if err := q.DeleteSizeVarianProduct(c.Context(), size.ID); err != nil {
				return c.Status(500).JSON(dto.ErrorResponse{
					Status:  500,
					Message: err.Error(),
				})
			}
		}
		err = q.DeleteColorVarianProduct(c.Context(), colorVarian.ID)
		if err != nil {
			return c.Status(500).JSON(dto.ErrorResponse{
				Status:  500,
				Message: err.Error(),
			})
		}

		return c.Status(200).JSON(dto.ErrorResponse{
			Status:  200,
			Message: "Ok",
		})
	})
}

// GetALl implements ColorVarianController.
func (cv *ColorVarianControllerImpl) GetALl(c *fiber.Ctx) error {
	productId := c.Params("id")

	id, err := strconv.Atoi(productId)
	if err != nil {
		return c.Status(400).JSON(dto.ErrorResponse{
			Status:  400,
			Message: err.Error(),
		})
	}
	args := db.ListColorVarianProductParams{ProductID: int64(id)}
	colorVarians, err := cv.Server.Store.ListColorVarianProduct(c.Context(), args)
	if err != nil {
		return c.Status(500).JSON(dto.ErrorResponse{
			Status:  500,
			Message: err.Error(),
		})
	}

	return c.Status(200).JSON(dto.SuccessResponse{
		Status:  200,
		Message: "Ok",
		Data:    colorVarians,
	})

}

// GetById implements ColorVarianController.
func (cv *ColorVarianControllerImpl) GetById(c *fiber.Ctx) error {
	cvId := c.Params("id")

	id, err := strconv.Atoi(cvId)
	if err != nil {
		return c.Status(400).JSON(dto.ErrorResponse{
			Status:  400,
			Message: err.Error(),
		})
	}

	colorVarian, err := cv.Server.Store.GetColorVarianProduct(c.Context(), int64(id))
	if err != nil {
		return c.Status(500).JSON(dto.ErrorResponse{
			Status:  500,
			Message: err.Error(),
		})
	}

	return c.Status(200).JSON(dto.SuccessResponse{
		Status:  200,
		Message: "Ok",
		Data:    colorVarian,
	})
}

// Update implements ColorVarianController.
func (cv *ColorVarianControllerImpl) Update(c *fiber.Ctx) error {
	return cv.Server.Store.ExecTx(c.Context(), func(q *db.Queries) error {
		productId := c.Params("product_id")

		id, err := strconv.Atoi(productId)
		if err != nil {
			return c.Status(400).JSON(dto.ErrorResponse{
				Status:  400,
				Message: err.Error(),
			})
		}

		req := new(request.UpdateColorVarianProduct)

		err = c.BodyParser(req)
		if err != nil {
			return c.Status(http.StatusBadRequest).JSON(dto.ErrorResponse{
				Status:  http.StatusBadRequest,
				Message: "Invalid Message Body",
			})
		}

		product, err := q.GetProduct(c.Context(), int64(id))

		if err != nil {
			return c.Status(400).JSON(dto.ErrorResponse{
				Status:  400,
				Message: err.Error(),
			})
		}

		names := strings.Fields(product.Name)
		folder := "./assets/product/" + helper.Generate(names[0])

		if err := os.MkdirAll(folder, os.ModePerm); err != nil {
			return c.Status(500).JSON(dto.ErrorResponse{
				Status:  500,
				Message: "Failed create directory",
			})
		}

		file, err := c.FormFile("images")
		var colorVarianPath string
		if err == nil {
			colorVarianPath = folder + file.Filename
			if saveErr := c.SaveFile(file, colorVarianPath); saveErr != nil {
				return c.Status(500).JSON(dto.ErrorResponse{
					Status:  500,
					Message: fmt.Sprintf("Failed to save image for color variant"),
				})
			}
		}

		colorVarianParam := db.UpdateColorVarianProductParams{
			ID:     req.Id,
			Name:   req.Name,
			Color:  req.Color,
			Images: colorVarianPath,
		}

		colorVarianResult, err := q.UpdateColorVarianProduct(c.Context(), colorVarianParam)
		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(dto.ErrorResponse{
				Status:  http.StatusInternalServerError,
				Message: fmt.Sprintf("Failed to create color varian"),
			})
		}
		var resultSizeVarians []response.SizeVarianResponse
		var sizeVarians []request.UpdateSizeVarianProduct
		if err := json.Unmarshal([]byte(req.Sizes), &sizeVarians); err != nil {
			return c.Status(400).JSON(dto.ErrorResponse{
				Status:  400,
				Message: "Invalid size_varians format",
			})
		}

		for i := range sizeVarians {
			sizeParam := db.UpdateSizeVarianProductParams{
				ID:    sizeVarians[i].ID,
				Size:  sizeVarians[i].Size,
				Stock: sizeVarians[i].Stock,
			}
			size, err := q.UpdateSizeVarianProduct(c.Context(), sizeParam)
			if err != nil {
				return c.Status(http.StatusInternalServerError).JSON(dto.ErrorResponse{
					Status:  http.StatusInternalServerError,
					Message: "Failed to create size varian",
				})
			}
			resultSizeVarians = append(resultSizeVarians, helper.ToSizeVarianResponse(size))
		}

		return c.Status(201).JSON(dto.SuccessResponse{
			Status:  201,
			Message: "Ok",
			Data:    helper.ToColorVarianResponse(colorVarianResult, resultSizeVarians),
		})
	})
}

func NewColorVarianController(server Server) ColorVarianController {
	return &ColorVarianControllerImpl{Server: server}
}
