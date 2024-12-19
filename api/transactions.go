package api

import (
	"database/sql"
	"net/http"
	"sync"

	db "github.com/adityaputra42/e-commerce_backend/db/sqlc"
	"github.com/adityaputra42/e-commerce_backend/dto"
	"github.com/adityaputra42/e-commerce_backend/dto/request"
	"github.com/adityaputra42/e-commerce_backend/dto/response"
	"github.com/adityaputra42/e-commerce_backend/helper"
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
	return t.Server.Store.ExecTx(c.Context(), func(q *db.Queries) error {
		req := new(request.CreateTransaction)
		if err := c.BodyParser(req); err != nil {
			return c.Status(http.StatusBadRequest).JSON(dto.ErrorResponse{
				Status:  http.StatusBadRequest,
				Message: "Invalid Message Body",
			})
		}

		var (
			address       db.Address
			shipping      db.Shipping
			paymentMethod db.PaymentMethod
			wg            sync.WaitGroup
			mu            sync.Mutex
			errChan       = make(chan error, 3)
		)

		wg.Add(1)
		go func() {
			defer wg.Done()
			addr, err := q.GetAddress(c.Context(), req.AddressID)
			mu.Lock()
			address = addr
			mu.Unlock()
			if err != nil {
				if err == sql.ErrNoRows {
					errChan <- c.Status(http.StatusNotFound).JSON(dto.ErrorResponse{
						Status:  http.StatusNotFound,
						Message: "Address not found",
					})
				} else {
					errChan <- c.Status(http.StatusInternalServerError).JSON(dto.ErrorResponse{
						Status:  http.StatusInternalServerError,
						Message: err.Error(),
					})
				}
			}
		}()

		wg.Add(1)
		go func() {
			defer wg.Done()
			ship, err := q.GetShipping(c.Context(), req.ShippingID)
			mu.Lock()
			shipping = ship
			mu.Unlock()
			if err != nil {
				if err == sql.ErrNoRows {
					errChan <- c.Status(http.StatusNotFound).JSON(dto.ErrorResponse{
						Status:  http.StatusNotFound,
						Message: "Shipping not found",
					})
				} else {
					errChan <- c.Status(http.StatusInternalServerError).JSON(dto.ErrorResponse{
						Status:  http.StatusInternalServerError,
						Message: err.Error(),
					})
				}
			}
		}()

		wg.Add(1)
		go func() {
			defer wg.Done()
			pm, err := q.GetPaymentMethod(c.Context(), req.PaymentMethodID)
			mu.Lock()
			paymentMethod = pm
			mu.Unlock()
			if err != nil {
				if err == sql.ErrNoRows {
					errChan <- c.Status(http.StatusNotFound).JSON(dto.ErrorResponse{
						Status:  http.StatusNotFound,
						Message: "Payment Method not found",
					})
				} else {
					errChan <- c.Status(http.StatusInternalServerError).JSON(dto.ErrorResponse{
						Status:  http.StatusInternalServerError,
						Message: err.Error(),
					})
				}
			}
		}()

		wg.Wait()

		close(errChan)
		for err := range errChan {
			if err != nil {
				return err
			}
		}

		transactionParam := db.CreateTransactionParams{
			TxID:            helper.Generate("TRX-"),
			AddressID:       address.ID,
			ShippingID:      shipping.ID,
			PaymentMethodID: paymentMethod.ID,
			ShippingPrice:   req.ShippingPrice,
			TotalPrice:      req.TotalPrice,
			Status:          helper.WaitingPayment,
		}
		transaction, err := q.CreateTransaction(c.Context(), transactionParam)
		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(dto.ErrorResponse{
				Status:  http.StatusInternalServerError,
				Message: err.Error(),
			})
		}

		var orderProduct []response.OrderResponse
		var wgOrders sync.WaitGroup
		var muOrders sync.Mutex
		errorsChan := make(chan error, len(req.ProductOrders))

		for _, v := range req.ProductOrders {
			wgOrders.Add(1)
			go func(productOrder request.CreateOrder) {
				defer wgOrders.Done()
				Product, err := q.GetProduct(c.Context(), v.ProductID)
				if err != nil {
					errorsChan <- c.Status(http.StatusNotFound).JSON(dto.ErrorResponse{
						Status:  http.StatusNotFound,
						Message: "Product not found",
					})
					return
				}

				Category, err := q.GetCategories(c.Context(), Product.CategoryID)
				if err != nil {
					errorsChan <- c.Status(http.StatusNotFound).JSON(dto.ErrorResponse{
						Status:  http.StatusNotFound,
						Message: "Category not found",
					})
					return
				}

				ColorVarian, err := q.GetColorVarianProductForUpdate(c.Context(), v.ColorVarianID)
				if err != nil {
					errorsChan <- c.Status(http.StatusNotFound).JSON(dto.ErrorResponse{
						Status:  http.StatusNotFound,
						Message: "Color varian not found",
					})
					return
				}

				Size, err := q.GetSizeVarianProduct(c.Context(), v.SizeVarianID)
				if err != nil {
					errorsChan <- c.Status(http.StatusNotFound).JSON(dto.ErrorResponse{
						Status:  http.StatusNotFound,
						Message: "Size varian not found",
					})
					return
				}

				if Size.Stock < v.Quantity {
					errorsChan <- c.Status(http.StatusBadRequest).JSON(dto.ErrorResponse{
						Status:  http.StatusBadRequest,
						Message: "Insufficient stock",
					})
					return
				}

				orderParam := db.CreateOrderParams{
					ID:            helper.Generate("TXO"),
					TransactionID: transaction.TxID,
					ProductID:     Product.ID,
					ColorVarianID: ColorVarian.ID,
					SizeVarianID:  Size.ID,
					UnitPrice:     v.UnitPrice,
					Quantity:      v.Quantity,
					Subtotal:      v.Subtotal,
					Status:        helper.Pending,
				}

				updateStockParam := db.UpdateSizeVarianProductParams{
					ID:    Size.ID,
					Stock: Size.Stock - v.Quantity,
				}

				muOrders.Lock()
				orders, err := q.CreateOrder(c.Context(), orderParam)
				if err == nil {
					_, err = q.UpdateSizeVarianProduct(c.Context(), updateStockParam)
				}
				orderProduct = append(orderProduct, helper.ToOrderResponse(orders, Size.Size, helper.ToProductOrderResponse(Product, helper.ToCategoryRespone(Category), helper.ToColorVarianOrderResponse(ColorVarian))))
				muOrders.Unlock()

				if err != nil {
					errorsChan <- c.Status(http.StatusInternalServerError).JSON(dto.ErrorResponse{
						Status:  http.StatusInternalServerError,
						Message: "Failed to create order or update stock: " + err.Error(),
					})
					return
				}
			}(v)
		}

		wgOrders.Wait()
		close(errorsChan)

		for err := range errorsChan {
			if err != nil {
				return err
			}
		}

		return c.Status(200).JSON(dto.SuccessResponse{
			Status:  200,
			Message: "Ok",
			Data: helper.ToTransactionResponse(
				transaction, helper.ToAddressResponse(address),
				helper.ToShippingRespone(shipping),
				helper.ToPaymentMethodRespone(paymentMethod),
				orderProduct,
			),
		})
	})
}

// Delete implements TransactionsController.
func (t *TransactionsControllerImpl) Delete(c *fiber.Ctx) error {
	return t.Server.Store.ExecTx(c.Context(), func(q *db.Queries) error {

		return c.Status(200).JSON(dto.ErrorResponse{
			Status:  200,
			Message: "Ok",
		})
	})
}

// GetAll implements TransactionsController.
func (t *TransactionsControllerImpl) GetAll(c *fiber.Ctx) error {

	page := c.QueryInt("page")
	limit := c.QueryInt("limit", 10)

	arg := db.ListTransactionParams{Limit: int32(limit), Offset: int32(page)}

	TransactionList, err := t.Server.Store.ListTransaction(c.Context(), arg)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(dto.ErrorResponse{
			Status:  http.StatusInternalServerError,
			Message: "Failed to get list transaction " + err.Error(),
		})
	}

	return c.Status(200).JSON(dto.SuccessResponse{
		Status:  200,
		Message: "Ok",
		Data:    TransactionList,
	})
}

// GetById implements TransactionsController.
func (t *TransactionsControllerImpl) GetById(c *fiber.Ctx) error {
	tx_id := c.Params("tx_id")
	transaction, err := t.Server.Store.GetTransaction(c.Context(), tx_id)

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(dto.ErrorResponse{
			Status:  http.StatusInternalServerError,
			Message: "Failed to get transaction " + err.Error(),
		})
	}
	return c.Status(200).JSON(dto.SuccessResponse{
		Status:  200,
		Message: "Ok",
		Data:    transaction,
	})

}

// Update implements TransactionsController.
func (t *TransactionsControllerImpl) Update(c *fiber.Ctx) error {
	req := new(db.UpdateTransactionParams)

	if err := c.BodyParser(req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(dto.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: "Invalid Message Body",
		})
	}

	transacction, err := t.Server.Store.UpdateTransaction(c.Context(), *req)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(dto.ErrorResponse{
			Status:  http.StatusInternalServerError,
			Message: "Failed to update transaction",
		})
	}
	return c.Status(200).JSON(dto.SuccessResponse{
		Status:  200,
		Message: "Ok",
		Data:    transacction,
	})

}

func NewTransactionsController(server routes.Server) TransactionsController {
	return &TransactionsControllerImpl{Server: server}
}
