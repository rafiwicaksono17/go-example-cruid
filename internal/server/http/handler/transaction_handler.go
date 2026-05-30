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

type transactionHandler struct {
	transactionUseCase usecase.TransactionUseCase
	validate           *validator.Validate
}

func NewTransactionHandler(transactionUseCase usecase.TransactionUseCase, validate *validator.Validate) *transactionHandler {
	return &transactionHandler{
		transactionUseCase: transactionUseCase,
		validate:           validate,
	}
}

func (h *transactionHandler) CreateTransaction(c *fiber.Ctx) error {
	userID, err := utils.ExtractToken(c)
	if err != nil {
		return helper.Response(c, http.StatusUnauthorized, "unauthorized", nil, err.Error())
	}

	req := new(model.CreateTransactionRequest)
	if err := c.BodyParser(req); err != nil {
		return helper.Response(c, http.StatusBadRequest, "invalid request", nil, err.Error())
	}

	if err := h.validate.Struct(req); err != nil {
		return helper.Response(c, http.StatusBadRequest, "validation error", nil, err.Error())
	}

	result, err := h.transactionUseCase.CreateTransaction(userID, *req)
	if err != nil {
		return helper.Response(c, http.StatusBadRequest, "create transaction failed", nil, err.Error())
	}

	return helper.Response(c, http.StatusCreated, "create transaction success", result, "")
}

func (h *transactionHandler) GetMyTransactions(c *fiber.Ctx) error {
	userID, err := utils.ExtractToken(c)
	if err != nil {
		return helper.Response(c, http.StatusUnauthorized, "unauthorized", nil, err.Error())
	}

	limit := helper.QueryInt(c, "limit", 10)
	offset := helper.QueryInt(c, "offset", 0)
	status := c.Query("status", "")

	result, err := h.transactionUseCase.GetTransactionsByUser(userID, limit, offset, status)
	if err != nil {
		return helper.Response(c, http.StatusInternalServerError, "get transactions failed", nil, err.Error())
	}

	return helper.Response(c, http.StatusOK, "get transactions success", result, "")
}

func (h *transactionHandler) GetMyTransactionByID(c *fiber.Ctx) error {
	userID, err := utils.ExtractToken(c)
	if err != nil {
		return helper.Response(c, http.StatusUnauthorized, "unauthorized", nil, err.Error())
	}

	transactionID, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return helper.Response(c, http.StatusBadRequest, "invalid transaction id", nil, err.Error())
	}

	result, err := h.transactionUseCase.GetTransactionByID(userID, uint(transactionID))
	if err != nil {
		return helper.Response(c, http.StatusNotFound, "get transaction failed", nil, err.Error())
	}

	return helper.Response(c, http.StatusOK, "get transaction success", result, "")
}

func (h *transactionHandler) UpdateTransactionStatus(c *fiber.Ctx) error {
	userID, err := utils.ExtractToken(c)
	if err != nil {
		return helper.Response(c, http.StatusUnauthorized, "unauthorized", nil, err.Error())
	}

	transactionID, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return helper.Response(c, http.StatusBadRequest, "invalid transaction id", nil, err.Error())
	}

	req := new(model.UpdateTransactionStatusRequest)
	if err := c.BodyParser(req); err != nil {
		return helper.Response(c, http.StatusBadRequest, "invalid request", nil, err.Error())
	}

	if err := h.validate.Struct(req); err != nil {
		return helper.Response(c, http.StatusBadRequest, "validation error", nil, err.Error())
	}

	result, err := h.transactionUseCase.UpdateTransactionStatus(userID, uint(transactionID), *req)
	if err != nil {
		return helper.Response(c, http.StatusBadRequest, "update transaction status failed", nil, err.Error())
	}

	return helper.Response(c, http.StatusOK, "update transaction status success", result, "")
}

func (h *transactionHandler) GetAllTransactions(c *fiber.Ctx) error {
	// Check if user is admin
	isAdmin, err := utils.CheckIsAdmin(c)
	if err != nil || !isAdmin {
		return helper.Response(c, http.StatusForbidden, "forbidden: only admin can view all transactions", nil, "")
	}

	limit := helper.QueryInt(c, "limit", 10)
	offset := helper.QueryInt(c, "offset", 0)
	status := c.Query("status", "")

	result, err := h.transactionUseCase.GetAllTransactions(limit, offset, status)
	if err != nil {
		return helper.Response(c, http.StatusInternalServerError, "get transactions failed", nil, err.Error())
	}

	return helper.Response(c, http.StatusOK, "get transactions success", result, "")
}
