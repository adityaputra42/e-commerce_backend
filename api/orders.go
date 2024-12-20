package api

import (
	"net/http"

	db "github.com/adityaputra42/e-commerce_backend/db/sqlc"
	"github.com/adityaputra42/e-commerce_backend/dto"
	"github.com/adityaputra42/e-commerce_backend/dto/request"
	"github.com/adityaputra42/e-commerce_backend/helper"
	"github.com/gofiber/fiber/v2"
)

type OrderController interface {
	Cancel(c *fiber.Ctx) error
	Update(c *fiber.Ctx) error
	Delete(c *fiber.Ctx) error
	GetAll(c *fiber.Ctx) error
	GetById(c *fiber.Ctx) error
}

type OrderControllerImpl struct {
	Server Server
}

// Cancel implements OrderController.
func (o *OrderControllerImpl) Cancel(c *fiber.Ctx) error {
	orderId := c.Params("order_id")

	order, err := o.Server.Store.GetOrderForUpdate(c.Context(), orderId)

	if err != nil {
		return c.Status(http.StatusNotFound).JSON(dto.ErrorResponse{
			Status:  http.StatusNotFound,
			Message: "Order not found",
		},
		)
	}

	if !helper.ValidateStatusOrder(order.Status) {
		return c.Status(http.StatusForbidden).JSON(dto.ErrorResponse{
			Status:  http.StatusForbidden,
			Message: "Can not cancel order",
		},
		)
	}

	arg := db.UpdateOrderParams{
		ID:     order.ID,
		Status: helper.Cancelled}
	_, err = o.Server.Store.UpdateOrder(c.Context(), arg)

	if err != nil {
		return c.Status(http.StatusNotFound).JSON(dto.ErrorResponse{
			Status:  http.StatusNotFound,
			Message: "Order not found",
		},
		)
	}
	return c.Status(http.StatusOK).JSON(dto.ErrorResponse{
		Status:  http.StatusOK,
		Message: "Ok",
	})
}

// Delete implements OrderController.
func (o *OrderControllerImpl) Delete(c *fiber.Ctx) error {
	orderId := c.Params("order_id")

	order, err := o.Server.Store.GetOrderForUpdate(c.Context(), orderId)

	if err != nil {
		return c.Status(http.StatusNotFound).JSON(dto.ErrorResponse{
			Status:  http.StatusNotFound,
			Message: "Order not found",
		},
		)
	}
	err = o.Server.Store.DeleteOrder(c.Context(), order.ID)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(dto.ErrorResponse{
			Status:  http.StatusInternalServerError,
			Message: "Failed delete order",
		},
		)
	}
	return c.Status(http.StatusOK).JSON(dto.ErrorResponse{
		Status:  http.StatusOK,
		Message: "Ok",
	})
}

// GetAll implements OrderController.
func (o *OrderControllerImpl) GetAll(c *fiber.Ctx) error {

	status := c.Query("status")
	page := c.QueryInt("page")
	limit := c.QueryInt("limit", 10)

	arg := db.ListOrderParams{
		Status: status,
		Offset: int32(page),
		Limit:  int32(limit),
	}
	listOrders, err := o.Server.Store.ListOrder(c.Context(), arg)

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(dto.ErrorResponse{
			Status:  http.StatusInternalServerError,
			Message: "Failed get order list",
		},
		)
	}
	return c.Status(http.StatusOK).JSON(dto.SuccessResponse{
		Status:  http.StatusOK,
		Message: "Ok",
		Data:    listOrders,
	})
}

// GetById implements OrderController.
func (o *OrderControllerImpl) GetById(c *fiber.Ctx) error {
	orderId := c.Params("order_id")

	order, err := o.Server.Store.GetOrder(c.Context(), orderId)

	if err != nil {
		return c.Status(http.StatusNotFound).JSON(dto.ErrorResponse{
			Status:  http.StatusNotFound,
			Message: "Order not found",
		},
		)
	}
	return c.Status(http.StatusOK).JSON(dto.SuccessResponse{
		Status:  http.StatusOK,
		Message: "Ok",
		Data:    order,
	})
}

// Update implements OrderController.
func (o *OrderControllerImpl) Update(c *fiber.Ctx) error {

	req := new(request.UpdateOrder)

	err := c.BodyParser(req)
	if err != nil {
		return c.Status(400).JSON(dto.ErrorResponse{
			Status:  400,
			Message: "Invalid Message Body",
		})
	}

	order, err := o.Server.Store.GetOrderForUpdate(c.Context(), req.ID)

	if err != nil {
		return c.Status(http.StatusNotFound).JSON(dto.ErrorResponse{
			Status:  http.StatusNotFound,
			Message: "Order not found",
		},
		)
	}

	arg := db.UpdateOrderParams{
		ID:     order.ID,
		Status: req.Status}
	_, err = o.Server.Store.UpdateOrder(c.Context(), arg)

	if err != nil {
		return c.Status(http.StatusNotFound).JSON(dto.ErrorResponse{
			Status:  http.StatusNotFound,
			Message: "Failed to update order",
		},
		)
	}
	return c.Status(http.StatusOK).JSON(dto.ErrorResponse{
		Status:  http.StatusOK,
		Message: "Ok",
	})
}

func NewOrderController(server Server) OrderController {
	return &OrderControllerImpl{Server: server}
}
