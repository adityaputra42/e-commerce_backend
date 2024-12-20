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
	var param db.CreatePaymentMethodParams

	param.AccountName = c.FormValue("account_name")
	param.AccountNumber = c.FormValue("account_number")
	param.BankName = c.FormValue("bank_name")

	file, err := c.FormFile("bank_images")
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
	param.BankImages = filePath
	paymentMethod, err := p.Server.Store.CreatePaymentMethod(c.Context(), param)
	if err != nil {
		return c.Status(403).JSON(dto.ErrorResponse{
			Status:  403,
			Message: "Failed Create PaymentMethod",
		})
	}

	return c.Status(http.StatusCreated).JSON(dto.SuccessResponse{
		Status:  http.StatusCreated,
		Message: "Success Create PaymentMethod",
		Data:    helper.ToPaymentMethodRespone(paymentMethod),
	})

}

// Delete implements PaymentMethodController.
func (p *PaymentMethodControllerImpl) Delete(c *fiber.Ctx) error {
	pmId := c.Params("id")

	id, err := strconv.Atoi(pmId)
	if err != nil {
		return c.Status(400).JSON(dto.ErrorResponse{
			Status:  400,
			Message: err.Error(),
		})
	}
	err = p.Server.Store.DeletePaymentMethod(c.Context(), int64(id))
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

// GetAll implements PaymentMethodController.
func (p *PaymentMethodControllerImpl) GetAll(c *fiber.Ctx) error {
	var ListPaymentMethod []response.PaymentMethodResponse

	page := c.QueryInt("page")
	limit := c.QueryInt("limit", 10)

	param := db.ListPaymentMethodParams{

		Offset: int32(page),
		Limit:  int32(limit),
	}
	paymentMethods, err := p.Server.Store.ListPaymentMethod(c.Context(), param)

	if err != nil {
		return c.Status(400).JSON(dto.ErrorResponse{
			Status:  400,
			Message: err.Error(),
		})
	}

	for _, v := range paymentMethods {
		ListPaymentMethod = append(ListPaymentMethod, helper.ToPaymentMethodRespone(v))
	}

	return c.Status(http.StatusCreated).JSON(dto.SuccessResponse{
		Status:  http.StatusCreated,
		Message: "Success",
		Data:    ListPaymentMethod,
	})
}

// GetById implements PaymentMethodController.
func (p *PaymentMethodControllerImpl) GetById(c *fiber.Ctx) error {
	pmId := c.Params("id")
	id, err := strconv.Atoi(pmId)
	if err != nil {
		return c.Status(400).JSON(dto.ErrorResponse{
			Status:  400,
			Message: err.Error(),
		})
	}

	pm, err := p.Server.Store.GetPaymentMethod(c.Context(), int64(id))

	if err != nil {
		return c.Status(400).JSON(dto.ErrorResponse{
			Status:  400,
			Message: err.Error(),
		})
	}

	return c.Status(http.StatusCreated).JSON(dto.SuccessResponse{
		Status:  http.StatusCreated,
		Message: "Success",
		Data:    helper.ToPaymentMethodRespone(pm),
	})
}

// Update implements PaymentMethodController.
func (p *PaymentMethodControllerImpl) Update(c *fiber.Ctx) error {
	var param db.UpdatePaymentMethodParams
	pmId := c.FormValue("id")
	id, err := strconv.Atoi(pmId)
	if err != nil {
		return c.Status(400).JSON(dto.ErrorResponse{
			Status:  400,
			Message: err.Error(),
		})
	}
	param.ID = int64(id)
	param.AccountName = c.FormValue("account_name")
	param.AccountNumber = c.FormValue("account_number")
	param.BankName = c.FormValue("bank_name")

	// Ambil file dari form file
	file, err := c.FormFile("bank_images")
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
	param.BankImages = filePath
	paymentMethod, err := p.Server.Store.UpdatePaymentMethod(c.Context(), param)
	if err != nil {
		return c.Status(403).JSON(dto.ErrorResponse{
			Status:  403,
			Message: "Failed Create PaymentMethod",
		})
	}

	return c.Status(http.StatusCreated).JSON(dto.SuccessResponse{
		Status:  http.StatusCreated,
		Message: "Success Create PaymentMethod",
		Data:    helper.ToPaymentMethodRespone(paymentMethod),
	})
}

func NewPaymentMethodController(server routes.Server) PaymentMethodController {
	return &PaymentMethodControllerImpl{Server: server}
}
