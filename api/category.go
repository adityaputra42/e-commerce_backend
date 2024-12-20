package api

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	db "github.com/adityaputra42/e-commerce_backend/db/sqlc"
	"github.com/adityaputra42/e-commerce_backend/dto"
	"github.com/adityaputra42/e-commerce_backend/dto/response"
	"github.com/adityaputra42/e-commerce_backend/helper"
	"github.com/gofiber/fiber/v2"
)

type CategoryController interface {
	Create(c *fiber.Ctx) error
	Update(c *fiber.Ctx) error
	Delete(c *fiber.Ctx) error
	GetAll(c *fiber.Ctx) error
	GetById(c *fiber.Ctx) error
}

type CategoryControllerImpl struct {
	Server Server
}

// Create implements CategoryController.
func (s *CategoryControllerImpl) Create(c *fiber.Ctx) error {
	var param db.CreateCategoriesParams

	param.Name = c.FormValue("name")

	file, err := c.FormFile("icon")
	if err != nil {
		if err == http.ErrMissingFile {
			return fmt.Errorf("bank_images is required")
		}
		return err
	}

	folder := "./assets/bank/"

	if err := os.MkdirAll(folder, os.ModePerm); err != nil {
		return c.Status(500).JSON(dto.ErrorResponse{
			Status:  500,
			Message: "Failed create directory",
		})
	}

	filePath := filepath.Join(folder, file.Filename)
	if err = c.SaveFile(file, filePath); err != nil {
		return c.Status(500).JSON(dto.ErrorResponse{
			Status:  500,
			Message: "Failed to save product image",
		})

	}
	param.Icon = filePath

	Category, err := s.Server.Store.CreateCategories(c.Context(), param)
	if err != nil {
		return c.Status(403).JSON(dto.ErrorResponse{
			Status:  403,
			Message: "Failed Create Category",
		})
	}

	return c.Status(http.StatusCreated).JSON(dto.SuccessResponse{
		Status:  http.StatusCreated,
		Message: "Success Create Category",
		Data:    helper.ToCategoryRespone(Category),
	})

}

// Delete implements CategoryController.
func (s *CategoryControllerImpl) Delete(c *fiber.Ctx) error {
	CategoryId := c.Params("id")

	id, err := strconv.Atoi(CategoryId)
	if err != nil {
		return c.Status(400).JSON(dto.ErrorResponse{
			Status:  400,
			Message: err.Error(),
		})
	}
	err = s.Server.Store.DeleteCategories(c.Context(), int64(id))
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

// GetAll implements CategoryController.
func (s *CategoryControllerImpl) GetAll(c *fiber.Ctx) error {
	var Categorys []response.Category

	page := c.QueryInt("page")
	limit := c.QueryInt("limit", 10)

	arg := db.ListCategoriesParams{Limit: int32(limit), Offset: int32(page)}
	listCategory, err := s.Server.Store.ListCategories(c.Context(), arg)

	if err != nil {
		return c.Status(400).JSON(dto.ErrorResponse{
			Status:  400,
			Message: err.Error(),
		})
	}

	for _, v := range listCategory {
		Categorys = append(Categorys, helper.ToCategoryRespone(v))
	}

	return c.Status(http.StatusCreated).JSON(dto.SuccessResponse{
		Status:  http.StatusCreated,
		Message: "Success",
		Data:    Categorys,
	})
}

// GetById implements CategoryController.
func (s *CategoryControllerImpl) GetById(c *fiber.Ctx) error {
	CategoryId := c.Params("id")

	id, err := strconv.Atoi(CategoryId)
	if err != nil {
		return c.Status(400).JSON(dto.ErrorResponse{
			Status:  400,
			Message: err.Error(),
		})
	}
	Category, err := s.Server.Store.GetCategories(c.Context(), int64(id))

	if err != nil {
		return c.Status(400).JSON(dto.ErrorResponse{
			Status:  400,
			Message: err.Error(),
		})
	}

	return c.Status(http.StatusCreated).JSON(dto.SuccessResponse{
		Status:  http.StatusCreated,
		Message: "Success",
		Data:    helper.ToCategoryRespone(Category),
	})
}

// Update implements CategoryController.
func (s *CategoryControllerImpl) Update(c *fiber.Ctx) error {
	var param db.UpdateCategoriesParams

	id := c.FormValue("id")
	paramId, err := strconv.Atoi(id)
	if err != nil {
		return c.Status(400).JSON(dto.ErrorResponse{
			Status:  400,
			Message: err.Error(),
		})
	}

	Category, err := s.Server.Store.GetCategoriesForUpdate(c.Context(), int64(paramId))
	if err != nil {
		return c.Status(400).JSON(dto.ErrorResponse{
			Status:  400,
			Message: err.Error(),
		})
	}

	param.ID = Category.ID
	param.Name = c.FormValue("name")

	file, err := c.FormFile("icon")
	if err != nil {
		if err == http.ErrMissingFile {
			return fmt.Errorf("bank_images is required")
		}
		return err
	}

	folder := "./assets/bank/"

	if err := os.MkdirAll(folder, os.ModePerm); err != nil {
		return c.Status(500).JSON(dto.ErrorResponse{
			Status:  500,
			Message: "Failed create directory",
		})
	}

	filePath := filepath.Join(folder, file.Filename)
	if err = c.SaveFile(file, filePath); err != nil {
		return c.Status(500).JSON(dto.ErrorResponse{
			Status:  500,
			Message: "Failed to save product image",
		})

	}
	param.Icon = filePath

	CategoryUpdate, err := s.Server.Store.UpdateCategories(c.Context(), param)
	if err != nil {
		return c.Status(400).JSON(dto.ErrorResponse{
			Status:  400,
			Message: err.Error(),
		})
	}

	return c.Status(200).JSON(dto.SuccessResponse{
		Status:  200,
		Message: "Success",
		Data:    helper.ToCategoryRespone(CategoryUpdate),
	})

}

func NewCategoryController(server Server) CategoryController {
	return &CategoryControllerImpl{Server: server}
}
