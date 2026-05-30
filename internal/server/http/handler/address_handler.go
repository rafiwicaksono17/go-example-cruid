package handler

import (
	"net/http"
	"strconv"
	"tugas_akhir_example/internal/helper"
	"tugas_akhir_example/internal/pkg/model"
	"tugas_akhir_example/internal/pkg/usecase"
	"tugas_akhir_example/internal/utils"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type addressHandler struct {
	addressUseCase usecase.AddressUseCase
	validate       *validator.Validate
}

func NewAddressHandler(addressUseCase usecase.AddressUseCase, validate *validator.Validate) *addressHandler {
	return &addressHandler{
		addressUseCase: addressUseCase,
		validate:       validate,
	}
}

func (h *addressHandler) CreateAddress(c *fiber.Ctx) error {
	userID, err := utils.ExtractToken(c)
	if err != nil {
		return helper.Response(c, http.StatusUnauthorized, "unauthorized", nil, err.Error())
	}

	req := new(model.CreateAddressRequest)
	if err := c.BodyParser(req); err != nil {
		return helper.Response(c, http.StatusBadRequest, "invalid request", nil, err.Error())
	}

	if err := h.validate.Struct(req); err != nil {
		return helper.Response(c, http.StatusBadRequest, "validation error", nil, err.Error())
	}

	result, err := h.addressUseCase.CreateAddress(userID, *req)
	if err != nil {
		return helper.Response(c, http.StatusBadRequest, "create address failed", nil, err.Error())
	}

	return helper.Response(c, http.StatusCreated, "create address success", result, "")
}

func (h *addressHandler) GetMyAddresses(c *fiber.Ctx) error {
	userID, err := utils.ExtractToken(c)
	if err != nil {
		return helper.Response(c, http.StatusUnauthorized, "unauthorized", nil, err.Error())
	}

	limit := helper.QueryInt(c, "limit", 10)
	offset := helper.QueryInt(c, "offset", 0)

	result, err := h.addressUseCase.GetAllAddressesByUser(userID, limit, offset)
	if err != nil {
		return helper.Response(c, http.StatusInternalServerError, "get addresses failed", nil, err.Error())
	}

	return helper.Response(c, http.StatusOK, "get addresses success", result, "")
}

func (h *addressHandler) GetMyAddressByID(c *fiber.Ctx) error {
	userID, err := utils.ExtractToken(c)
	if err != nil {
		return helper.Response(c, http.StatusUnauthorized, "unauthorized", nil, err.Error())
	}

	addressID, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return helper.Response(c, http.StatusBadRequest, "invalid address id", nil, err.Error())
	}

	result, err := h.addressUseCase.GetAddressByID(userID, uint(addressID))
	if err != nil {
		return helper.Response(c, http.StatusNotFound, "get address failed", nil, err.Error())
	}

	return helper.Response(c, http.StatusOK, "get address success", result, "")
}

func (h *addressHandler) UpdateMyAddress(c *fiber.Ctx) error {
	userID, err := utils.ExtractToken(c)
	if err != nil {
		return helper.Response(c, http.StatusUnauthorized, "unauthorized", nil, err.Error())
	}

	addressID, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return helper.Response(c, http.StatusBadRequest, "invalid address id", nil, err.Error())
	}

	req := new(model.UpdateAddressRequest)
	if err := c.BodyParser(req); err != nil {
		return helper.Response(c, http.StatusBadRequest, "invalid request", nil, err.Error())
	}

	result, err := h.addressUseCase.UpdateAddress(userID, uint(addressID), *req)
	if err != nil {
		return helper.Response(c, http.StatusBadRequest, "update address failed", nil, err.Error())
	}

	return helper.Response(c, http.StatusOK, "update address success", result, "")
}

func (h *addressHandler) DeleteMyAddress(c *fiber.Ctx) error {
	userID, err := utils.ExtractToken(c)
	if err != nil {
		return helper.Response(c, http.StatusUnauthorized, "unauthorized", nil, err.Error())
	}

	addressID, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return helper.Response(c, http.StatusBadRequest, "invalid address id", nil, err.Error())
	}

	if err := h.addressUseCase.DeleteAddress(userID, uint(addressID)); err != nil {
		return helper.Response(c, http.StatusBadRequest, "delete address failed", nil, err.Error())
	}

	return helper.Response(c, http.StatusOK, "delete address success", nil, "")
}
