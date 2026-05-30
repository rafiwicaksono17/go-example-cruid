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

type storeHandler struct {
	storeUseCase usecase.StoreUseCase
	validate     *validator.Validate
}

func NewStoreHandler(storeUseCase usecase.StoreUseCase, validate *validator.Validate) *storeHandler {
	return &storeHandler{
		storeUseCase: storeUseCase,
		validate:     validate,
	}
}

func (h *storeHandler) CreateStore(c *fiber.Ctx) error {
	userID, err := utils.ExtractToken(c)
	if err != nil {
		return helper.Response(c, http.StatusUnauthorized, "unauthorized", nil, err.Error())
	}

	req := new(model.CreateStoreRequest)
	if err := c.BodyParser(req); err != nil {
		return helper.Response(c, http.StatusBadRequest, "invalid request", nil, err.Error())
	}

	if err := h.validate.Struct(req); err != nil {
		return helper.Response(c, http.StatusBadRequest, "validation error", nil, err.Error())
	}

	result, err := h.storeUseCase.CreateStore(userID, *req)
	if err != nil {
		return helper.Response(c, http.StatusBadRequest, "create store failed", nil, err.Error())
	}

	return helper.Response(c, http.StatusCreated, "create store success", result, "")
}

func (h *storeHandler) GetMyStore(c *fiber.Ctx) error {
	userID, err := utils.ExtractToken(c)
	if err != nil {
		return helper.Response(c, http.StatusUnauthorized, "unauthorized", nil, err.Error())
	}

	result, err := h.storeUseCase.GetStoreByUserID(userID)
	if err != nil {
		return helper.Response(c, http.StatusNotFound, "get store failed", nil, err.Error())
	}

	return helper.Response(c, http.StatusOK, "get store success", result, "")
}

func (h *storeHandler) UpdateMyStore(c *fiber.Ctx) error {
	userID, err := utils.ExtractToken(c)
	if err != nil {
		return helper.Response(c, http.StatusUnauthorized, "unauthorized", nil, err.Error())
	}

	storeID, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return helper.Response(c, http.StatusBadRequest, "invalid store id", nil, err.Error())
	}

	req := new(model.UpdateStoreRequest)
	if err := c.BodyParser(req); err != nil {
		return helper.Response(c, http.StatusBadRequest, "invalid request", nil, err.Error())
	}

	result, err := h.storeUseCase.UpdateStore(userID, uint(storeID), *req)
	if err != nil {
		return helper.Response(c, http.StatusBadRequest, "update store failed", nil, err.Error())
	}

	return helper.Response(c, http.StatusOK, "update store success", result, "")
}

func (h *storeHandler) GetStoreByID(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return helper.Response(c, http.StatusBadRequest, "invalid store id", nil, err.Error())
	}

	result, err := h.storeUseCase.GetStoreByID(uint(id))
	if err != nil {
		return helper.Response(c, http.StatusNotFound, "get store failed", nil, err.Error())
	}

	return helper.Response(c, http.StatusOK, "get store success", result, "")
}

func (h *storeHandler) GetAllStores(c *fiber.Ctx) error {
	limit := helper.QueryInt(c, "limit", 10)
	offset := helper.QueryInt(c, "offset", 0)
	name := c.Query("name", "")

	result, err := h.storeUseCase.GetAllStores(limit, offset, name)
	if err != nil {
		return helper.Response(c, http.StatusInternalServerError, "get stores failed", nil, err.Error())
	}

	return helper.Response(c, http.StatusOK, "get stores success", result, "")
}

func (h *storeHandler) DeleteMyStore(c *fiber.Ctx) error {
	userID, err := utils.ExtractToken(c)
	if err != nil {
		return helper.Response(c, http.StatusUnauthorized, "unauthorized", nil, err.Error())
	}

	storeID, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return helper.Response(c, http.StatusBadRequest, "invalid store id", nil, err.Error())
	}

	if err := h.storeUseCase.DeleteStore(userID, uint(storeID)); err != nil {
		return helper.Response(c, http.StatusBadRequest, "delete store failed", nil, err.Error())
	}

	return helper.Response(c, http.StatusOK, "delete store success", nil, "")
}
