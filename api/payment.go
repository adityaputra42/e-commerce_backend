package api

import (
	"net/http"
	"strconv"

	db "github.com/adityaputra42/e-commerce_backend/db/sqlc"
	"github.com/adityaputra42/e-commerce_backend/dto"
	"github.com/adityaputra42/e-commerce_backend/dto/request"
	"github.com/adityaputra42/e-commerce_backend/helper"
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
	Server Server
}

// Create implements PaymentController.
func (p *PaymentControllerImpl) Create(c *fiber.Ctx) error {
	return p.Server.Store.ExecTx(c.Context(), func(q *db.Queries) error {

		req := new(request.CreatePayment)

		err := c.BodyParser(req)
		if err != nil {
			return c.Status(http.StatusBadRequest).JSON(dto.ErrorResponse{
				Status:  http.StatusBadRequest,
				Message: "Invalid Message Body",
			})
		}

		transaction, err := q.GetTransactionForUpdate(c.Context(), req.TransactionID)
		if err != nil {
			return c.Status(http.StatusNotFound).JSON(dto.ErrorResponse{
				Status:  http.StatusNotFound,
				Message: "Transaction not found",
			})
		}

		if transaction.TotalPrice != req.TotalPayment {
			return c.Status(http.StatusBadRequest).JSON(dto.ErrorResponse{
				Status:  http.StatusBadRequest,
				Message: "Total payment didn't match",
			})
		}
		arg := db.CreatePaymentParams{
			TransactionID: transaction.TxID,
			TotalPayment:  req.TotalPayment,
			Status:        helper.Pending,
		}

		payment, err := q.CreatePayment(c.Context(), arg)
		if err != nil {
			return c.Status(http.StatusNotFound).JSON(dto.ErrorResponse{
				Status:  http.StatusNotFound,
				Message: "Failed create payment",
			})
		}
		updateTx := db.UpdateTransactionParams{
			TxID:   transaction.TxID,
			Status: helper.WaitingConfirPayment,
		}
		newTx, err := q.UpdateTransaction(c.Context(), updateTx)

		if err != nil {
			return c.Status(http.StatusNotFound).JSON(dto.ErrorResponse{
				Status:  http.StatusNotFound,
				Message: "Failed to update transaction status",
			})
		}

		return c.Status(http.StatusCreated).JSON(dto.SuccessResponse{
			Status:  http.StatusCreated,
			Message: "Ok",
			Data:    helper.ToPaymentResponse(payment, newTx),
		})

	})

}

// Delete implements PaymentController.
func (p *PaymentControllerImpl) Delete(c *fiber.Ctx) error {
	paymentId := c.Params("id")

	id, err := strconv.Atoi(paymentId)
	if err != nil {
		return c.Status(500).JSON(dto.ErrorResponse{
			Status:  500,
			Message: err.Error(),
		})
	}
	err = p.Server.Store.DeletePayment(c.Context(), int64(id))

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
}

// GetAll implements PaymentController.
func (p *PaymentControllerImpl) GetAll(c *fiber.Ctx) error {

	page := c.QueryInt("page")
	limit := c.QueryInt("limit", 10)

	arg := db.ListPaymentParams{Limit: int32(limit), Offset: int32(page)}

	payment, err := p.Server.Store.ListPayment(c.Context(), arg)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(dto.ErrorResponse{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		})
	}
	return c.Status(200).JSON(dto.SuccessResponse{
		Status:  200,
		Message: "Ok",
		Data:    payment,
	})
}

// GetById implements PaymentController.
func (p *PaymentControllerImpl) GetById(c *fiber.Ctx) error {
	productId := c.Params("id")

	id, err := strconv.Atoi(productId)
	if err != nil {
		return c.Status(400).JSON(dto.ErrorResponse{
			Status:  400,
			Message: err.Error(),
		})
	}
	payment, err := p.Server.Store.GetPayment(c.Context(), int64(id))

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(dto.ErrorResponse{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	transaction, err := p.Server.Store.GetTransactionForUpdate(c.Context(), payment.TransactionID)

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(dto.ErrorResponse{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	return c.Status(200).JSON(dto.SuccessResponse{
		Status:  200,
		Message: "Ok",
		Data:    helper.ToPaymentResponse(payment, transaction),
	})

}

// Update implements PaymentController.
func (p *PaymentControllerImpl) Update(c *fiber.Ctx) error {

	req := new(db.UpdatePaymentParams)

	err := c.BodyParser(req)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(dto.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: "Invalid Message Body",
		})
	}

	payment, err := p.Server.Store.GetPaymentForUpdate(c.Context(), req.ID)

	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(dto.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: "Failed to get data payment",
		})
	}

	if payment.Status == req.Status {
		return c.Status(http.StatusBadRequest).JSON(dto.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: "Failed to update payment",
		})
	}
	Payment, err := p.Server.Store.UpdatePayment(c.Context(), *req)

	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(dto.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: "Failed to update payment",
		})
	}

	return c.Status(http.StatusOK).JSON(dto.SuccessResponse{
		Status:  http.StatusOK,
		Message: "Ok",
		Data:    Payment,
	})
}

func NewPaymentController(server Server) PaymentController {
	return &PaymentControllerImpl{Server: server}
}
