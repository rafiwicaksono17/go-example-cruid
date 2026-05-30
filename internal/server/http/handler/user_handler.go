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

type userHandler struct {
	userUseCase usecase.UserUseCase
	validate    *validator.Validate
}

func NewUserHandler(userUseCase usecase.UserUseCase, validate *validator.Validate) *userHandler {
	return &userHandler{
		userUseCase: userUseCase,
		validate:    validate,
	}
}

func (h *userHandler) GetMe(c *fiber.Ctx) error {
	userID, err := utils.ExtractToken(c)
	if err != nil {
		return helper.Response(c, http.StatusUnauthorized, "unauthorized", nil, err.Error())
	}

	result, err := h.userUseCase.GetUserByID(userID)
	if err != nil {
		return helper.Response(c, http.StatusNotFound, "get user failed", nil, err.Error())
	}

	return helper.Response(c, http.StatusOK, "get user success", result, "")
}

func (h *userHandler) UpdateMe(c *fiber.Ctx) error {
	userID, err := utils.ExtractToken(c)
	if err != nil {
		return helper.Response(c, http.StatusUnauthorized, "unauthorized", nil, err.Error())
	}

	req := new(model.UpdateUserRequest)
	if err := c.BodyParser(req); err != nil {
		return helper.Response(c, http.StatusBadRequest, "invalid request", nil, err.Error())
	}

	result, err := h.userUseCase.UpdateUser(userID, *req)
	if err != nil {
		return helper.Response(c, http.StatusBadRequest, "update user failed", nil, err.Error())
	}

	return helper.Response(c, http.StatusOK, "update user success", result, "")
}

func (h *userHandler) GetUserByID(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return helper.Response(c, http.StatusBadRequest, "invalid user id", nil, err.Error())
	}

	result, err := h.userUseCase.GetUserByID(uint(id))
	if err != nil {
		return helper.Response(c, http.StatusNotFound, "get user failed", nil, err.Error())
	}

	return helper.Response(c, http.StatusOK, "get user success", result, "")
}

func (h *userHandler) GetAllUsers(c *fiber.Ctx) error {
	limit := helper.QueryInt(c, "limit", 10)
	offset := helper.QueryInt(c, "offset", 0)
	name := c.Query("name", "")

	result, err := h.userUseCase.GetAllUsers(limit, offset, name)
	if err != nil {
		return helper.Response(c, http.StatusInternalServerError, "get users failed", nil, err.Error())
	}

	return helper.Response(c, http.StatusOK, "get users success", result, "")
}

func (h *userHandler) DeleteMe(c *fiber.Ctx) error {
	userID, err := utils.ExtractToken(c)
	if err != nil {
		return helper.Response(c, http.StatusUnauthorized, "unauthorized", nil, err.Error())
	}

	if err := h.userUseCase.DeleteUser(userID); err != nil {
		return helper.Response(c, http.StatusBadRequest, "delete user failed", nil, err.Error())
	}

	return helper.Response(c, http.StatusOK, "delete user success", nil, "")
}
