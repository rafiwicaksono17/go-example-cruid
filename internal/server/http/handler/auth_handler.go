package handler

import (
	"net/http"
	"tugas_akhir_example/internal/helper"
	"tugas_akhir_example/internal/pkg/model"
	"tugas_akhir_example/internal/pkg/usecase"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type authHandler struct {
	authUseCase usecase.AuthUseCase
	validate    *validator.Validate
}

func NewAuthHandler(authUseCase usecase.AuthUseCase, validate *validator.Validate) *authHandler {
	return &authHandler{
		authUseCase: authUseCase,
		validate:    validate,
	}
}

func (h *authHandler) Register(c *fiber.Ctx) error {
	req := new(model.RegisterRequest)
	if err := c.BodyParser(req); err != nil {
		return helper.Response(c, http.StatusBadRequest, "invalid request", nil, err.Error())
	}

	if err := h.validate.Struct(req); err != nil {
		return helper.Response(c, http.StatusBadRequest, "validation error", nil, err.Error())
	}

	result, err := h.authUseCase.Register(*req)
	if err != nil {
		return helper.Response(c, http.StatusBadRequest, "register failed", nil, err.Error())
	}

	return helper.Response(c, http.StatusCreated, "register success", result, "")
}

func (h *authHandler) Login(c *fiber.Ctx) error {
	req := new(model.LoginRequest)
	if err := c.BodyParser(req); err != nil {
		return helper.Response(c, http.StatusBadRequest, "invalid request", nil, err.Error())
	}

	if err := h.validate.Struct(req); err != nil {
		return helper.Response(c, http.StatusBadRequest, "validation error", nil, err.Error())
	}

	result, err := h.authUseCase.Login(*req)
	if err != nil {
		return helper.Response(c, http.StatusUnauthorized, "login failed", nil, err.Error())
	}

	return helper.Response(c, http.StatusOK, "login success", result, "")
}
