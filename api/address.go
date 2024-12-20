package api

import (
	"database/sql"
	"net/http"
	"strconv"

	db "github.com/adityaputra42/e-commerce_backend/db/sqlc"
	"github.com/adityaputra42/e-commerce_backend/dto"
	"github.com/adityaputra42/e-commerce_backend/dto/request"
	"github.com/adityaputra42/e-commerce_backend/dto/response"
	"github.com/adityaputra42/e-commerce_backend/helper"
	"github.com/adityaputra42/e-commerce_backend/token"
	"github.com/gofiber/fiber/v2"
)

type AddressController interface {
	CreateAddress(c *fiber.Ctx) error
	Update(c *fiber.Ctx) error
	Delete(c *fiber.Ctx) error
	FetchAddress(c *fiber.Ctx) error
	FetchAllAddressByUser(c *fiber.Ctx) error
	FetchAllAddressFromAdmin(c *fiber.Ctx) error
}

type AddressControllerImpl struct {
	server Server
}

// CreateAddress implements AddressController.
func (a *AddressControllerImpl) CreateAddress(c *fiber.Ctx) error {
	req := new(request.CreateAddress)

	err := c.BodyParser(req)
	if err != nil {
		return c.Status(400).JSON(dto.ErrorResponse{
			Status:  400,
			Message: "Invalid Message Body",
		})
	}

	authPayload := c.Locals(helper.GetPayloadKey()).(*token.Payload)
	addressParam := db.CreateAddressParams{
		Uid:                  authPayload.Uid,
		RecipientName:        req.RecipientName,
		RecipientPhoneNumber: req.RecipientPhoneNumber,
		Province:             req.Province,
		City:                 req.City,
		District:             req.District,
		Village:              req.Village,
		PostalCode:           req.PostalCode,
		FullAddress:          req.FullAddress,
	}

	address, err := a.server.Store.CreateAddress(c.Context(), addressParam)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(dto.ErrorResponse{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		})
	}
	return c.Status(http.StatusCreated).JSON(dto.SuccessResponse{
		Status:  http.StatusCreated,
		Message: "Ok",
		Data:    helper.ToAddressResponse(address),
	})
}

// Delete implements AddressController.
func (a *AddressControllerImpl) Delete(c *fiber.Ctx) error {
	addressId := c.Params("id")

	id, err := strconv.Atoi(addressId)
	if err != nil {
		return c.Status(500).JSON(dto.ErrorResponse{
			Status:  500,
			Message: err.Error(),
		})
	}
	err = a.server.Store.DeleteAddress(c.Context(), int64(id))

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

// FetchAddress implements AddressController.
func (a *AddressControllerImpl) FetchAddress(c *fiber.Ctx) error {
	addressId := c.Params("id")

	id, err := strconv.Atoi(addressId)
	if err != nil {
		return c.Status(500).JSON(dto.ErrorResponse{
			Status:  500,
			Message: err.Error(),
		})
	}
	address, err := a.server.Store.GetAddress(c.Context(), int64(id))
	if err != nil {
		return c.Status(500).JSON(dto.ErrorResponse{
			Status:  500,
			Message: err.Error(),
		})
	}
	return c.Status(200).JSON(dto.SuccessResponse{
		Status:  200,
		Message: "Ok",
		Data:    helper.ToAddressResponse(address),
	})

}

// FetchAllAddressByUser implements AddressController.
func (a *AddressControllerImpl) FetchAllAddressByUser(c *fiber.Ctx) error {
	addresses := []response.AddressResponse{}

	authPayload := c.Locals(helper.GetPayloadKey()).(*token.Payload)

	addressListParam := db.ListAddressParams{Uid: authPayload.Uid}
	addresslist, err := a.server.Store.ListAddress(c.Context(), addressListParam)
	if err != nil {
		return c.Status(500).JSON(dto.ErrorResponse{
			Status:  500,
			Message: err.Error(),
		})
	}
	for _, address := range addresslist {
		addresses = append(addresses, helper.ToAddressResponse(address))
	}

	return c.Status(http.StatusOK).JSON(dto.SuccessResponse{
		Status:  http.StatusOK,
		Message: "Ok",
		Data:    addresses,
	})
}

// FetchAllAddressFromAdmin implements AddressController.
func (a *AddressControllerImpl) FetchAllAddressFromAdmin(c *fiber.Ctx) error {
	addresses := []response.AddressResponse{}
	uid := c.Params("uid")
	addressListParam := db.ListAddressParams{Uid: uid}
	addresslist, err := a.server.Store.ListAddress(c.Context(), addressListParam)
	if err != nil {
		return c.Status(500).JSON(dto.ErrorResponse{
			Status:  500,
			Message: err.Error(),
		})
	}
	for _, address := range addresslist {
		addresses = append(addresses, helper.ToAddressResponse(address))
	}

	return c.Status(http.StatusOK).JSON(dto.SuccessResponse{
		Status:  http.StatusOK,
		Message: "Ok",
		Data:    addresses,
	})
}

// Update implements AddressController.
func (a *AddressControllerImpl) Update(c *fiber.Ctx) error {
	req := new(request.CreateAddress)

	err := c.BodyParser(req)
	if err != nil {
		return c.Status(400).JSON(dto.ErrorResponse{
			Status:  400,
			Message: "Invalid Message Body",
		})
	}

	addressId := c.Params("id")

	id, err := strconv.Atoi(addressId)
	if err != nil {
		return c.Status(500).JSON(dto.ErrorResponse{
			Status:  500,
			Message: err.Error(),
		})
	}
	address, err := a.server.Store.GetAddressForUpdate(c.Context(), int64(id))
	if err != nil {
		if err == sql.ErrNoRows {
			return c.Status(http.StatusNotFound).JSON(dto.ErrorResponse{
				Status:  http.StatusNotFound,
				Message: "Address not found",
			})
		}
		return c.Status(http.StatusInternalServerError).JSON(dto.ErrorResponse{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		})
	}
	addressParam := &db.UpdateAddressParams{
		ID:                   address.ID,
		RecipientName:        req.RecipientName,
		RecipientPhoneNumber: req.RecipientPhoneNumber,
		Province:             req.Province,
		City:                 req.City,
		District:             req.District,
		Village:              req.Village,
		PostalCode:           req.PostalCode,
		FullAddress:          req.FullAddress,
	}
	address, err = a.server.Store.UpdateAddress(c.Context(), *addressParam)
	if err != nil {
		return c.Status(500).JSON(dto.ErrorResponse{
			Status:  500,
			Message: err.Error(),
		})
	}
	return c.Status(http.StatusOK).JSON(dto.SuccessResponse{
		Status:  http.StatusOK,
		Message: "Ok",
		Data:    helper.ToAddressResponse(address),
	})
}

func NewAddressController(server Server) AddressController {
	return &AddressControllerImpl{server: server}
}
