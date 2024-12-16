package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	db "github.com/adityaputra42/e-commerce_backend/db/sqlc"
	"github.com/adityaputra42/e-commerce_backend/dto"
	"github.com/adityaputra42/e-commerce_backend/dto/request"
	"github.com/adityaputra42/e-commerce_backend/dto/response"
	"github.com/adityaputra42/e-commerce_backend/helper"
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

		names := strings.Fields(req.Name)
		var filePath string

		folder := "./assets/product/" + helper.Generate(names[0])

		if err := os.MkdirAll(folder, os.ModePerm); err != nil {
			return c.Status(500).JSON(dto.ErrorResponse{
				Status:  500,
				Message: "Failed create directory",
			})
		}

		if req.Images != nil {
			file := req.Images
			filePath = filepath.Join(folder, file.Filename)
			if err = c.SaveFile(file, filePath); err != nil {
				return c.Status(500).JSON(dto.ErrorResponse{
					Status:  500,
					Message: "Failed to save product image",
				})
			}
		}

		Category, err := q.GetCategories(c.Context(), req.CategoryID)
		if err != nil {
			return c.Status(500).JSON(dto.ErrorResponse{
				Status:  500,
				Message: "Failed get Category",
			})
		}

		productParam := db.CreateProductParams{
			CategoryID:  Category.ID,
			Name:        req.Name,
			Description: req.Description,
			Images:      filePath,
			Rating:      float64(req.Rating),
			Price:       req.Price,
		}

		productResult, err := q.CreateProduct(c.Context(), productParam)
		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(dto.ErrorResponse{
				Status:  http.StatusInternalServerError,
				Message: "Failed to create product",
			})
		}

		var resultColorVarians []response.ColorVarianResponse
		var colorVarians []request.CreateColorVarianProduct
		if err := json.Unmarshal([]byte(req.ColorVarian), &colorVarians); err != nil {
			return c.Status(400).JSON(dto.ErrorResponse{
				Status:  400,
				Message: "Invalid color_varians format",
			})
		}

		for i := range colorVarians {
			formFileKey := fmt.Sprintf("color_varians[%d].images", i)
			file, err := c.FormFile(formFileKey)
			var colorVarianPath string
			if err == nil {
				colorVarianPath = folder + file.Filename
				if saveErr := c.SaveFile(file, colorVarianPath); saveErr != nil {
					return c.Status(500).JSON(dto.ErrorResponse{
						Status:  500,
						Message: fmt.Sprintf("Failed to save image for color variant %d", i),
					})
				}
			}

			colorVarianParam := db.CreateColorVarianProductParams{
				ProductID: productResult.ID,
				Name:      colorVarians[i].Name,
				Color:     colorVarians[i].Color,
				Images:    colorVarianPath,
			}

			colorVarianResult, err := q.CreateColorVarianProduct(c.Context(), colorVarianParam)
			if err != nil {
				return c.Status(http.StatusInternalServerError).JSON(dto.ErrorResponse{
					Status:  http.StatusInternalServerError,
					Message: fmt.Sprintf("Failed to create color varian ke %d", i),
				})
			}

			var resultSizeVarians []response.SizeVarianResponse
			var sizeVarians []request.CreateSizeVarianProduct
			if err := json.Unmarshal([]byte(colorVarians[i].Sizes), &sizeVarians); err != nil {
				return c.Status(400).JSON(dto.ErrorResponse{
					Status:  400,
					Message: "Invalid size_varians format",
				})
			}

			for i := range colorVarians {
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

			resultColorVarians = append(resultColorVarians, helper.ToColorVarianResponse(colorVarianResult, resultSizeVarians))

		}

		return c.Status(201).JSON(dto.SuccessResponse{
			Status:  201,
			Message: "Ok",
			Data:    helper.ToProductDetailResponse(productResult, helper.ToCategoryRespone(Category), resultColorVarians),
		})
	})
}

// Delete implements ProductController.
func (p *ProductControllerImpl) Delete(c *fiber.Ctx) error {
	return p.Server.Store.ExecTx(c.Context(), func(q *db.Queries) error {
		productId := c.Params("id")

		id, err := strconv.Atoi(productId)
		if err != nil {
			return c.Status(400).JSON(dto.ErrorResponse{
				Status:  400,
				Message: err.Error(),
			})
		}
		product, err := q.GetProduct(c.Context(), int64(id))
		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(dto.ErrorResponse{
				Status:  http.StatusInternalServerError,
				Message: err.Error(),
			})
		}

		listColorVarian, err := q.ListColorVarianProduct(c.Context(), db.ListColorVarianProductParams{ProductID: product.ID})
		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(dto.ErrorResponse{
				Status:  http.StatusInternalServerError,
				Message: err.Error(),
			})
		}
		for _, v := range listColorVarian {

			sizes, err := q.ListSizeVarianProduct(c.Context(), db.ListSizeVarianProductParams{ColorVarianID: v.ID})
			if err != nil {
				return c.Status(http.StatusInternalServerError).JSON(dto.ErrorResponse{
					Status:  http.StatusInternalServerError,
					Message: err.Error(),
				})
			}
			for _, i := range sizes {
				err = q.DeleteSizeVarianProduct(c.Context(), i.ID)
				if err != nil {
					return c.Status(http.StatusInternalServerError).JSON(dto.ErrorResponse{
						Status:  http.StatusInternalServerError,
						Message: err.Error(),
					})
				}
			}
			err = q.DeleteColorVarianProduct(c.Context(), v.ID)
			if err != nil {
				return c.Status(http.StatusInternalServerError).JSON(dto.ErrorResponse{
					Status:  http.StatusInternalServerError,
					Message: err.Error(),
				})
			}
		}

		err = q.DeleteProduct(c.Context(), product.ID)
		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(dto.ErrorResponse{
				Status:  http.StatusInternalServerError,
				Message: err.Error(),
			})
		}

		return c.Status(200).JSON(dto.ErrorResponse{
			Status:  200,
			Message: "Ok",
		})
	})
}

// FetchListProduct implements ProductController.
func (p *ProductControllerImpl) FetchListProduct(c *fiber.Ctx) error {

	page := c.QueryInt("page")
	limit := c.QueryInt("limit", 10)
	products, err := p.Server.Store.ListProduct(c.Context(), db.ListProductParams{
		Limit:  int32(limit),
		Offset: int32(page),
	})
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(dto.ErrorResponse{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		})
	}
	return c.Status(200).JSON(dto.SuccessResponse{
		Status:  200,
		Message: "Ok",
		Data:    products,
	})
}

// FetchProduct implements ProductController.
func (p *ProductControllerImpl) FetchProduct(c *fiber.Ctx) error {
	productId := c.Params("id")

	id, err := strconv.Atoi(productId)
	if err != nil {
		return c.Status(400).JSON(dto.ErrorResponse{
			Status:  400,
			Message: err.Error(),
		})
	}
	product, err := p.Server.Store.GetProductWithDetail(c.Context(), int64(id))
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(dto.ErrorResponse{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		})
	}
	return c.Status(200).JSON(dto.SuccessResponse{
		Status:  200,
		Message: "Ok",
		Data:    product,
	})
}

// UpdateProduct implements ProductController.
func (p *ProductControllerImpl) UpdateProduct(c *fiber.Ctx) error {
	return p.Server.Store.ExecTx(c.Context(), func(q *db.Queries) error {

		return nil
	})
}

func NewProductController(server routes.Server) ProductController {
	return &ProductControllerImpl{Server: server}
}
