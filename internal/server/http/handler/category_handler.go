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

type categoryHandler struct {
	categoryUseCase usecase.CategoryUseCase
	validate        *validator.Validate
}

func NewCategoryHandler(categoryUseCase usecase.CategoryUseCase, validate *validator.Validate) *categoryHandler {
	return &categoryHandler{
		categoryUseCase: categoryUseCase,
		validate:        validate,
	}
}

func (h *categoryHandler) CreateCategory(c *fiber.Ctx) error {
	// Check if user is admin
	if _, err := utils.ExtractToken(c); err != nil {
		return helper.Response(c, http.StatusUnauthorized, "unauthorized", nil, err.Error())
	}

	isAdmin, err := utils.CheckIsAdmin(c)
	if err != nil || !isAdmin {
		return helper.Response(c, http.StatusForbidden, "forbidden: only admin can create category", nil, "")
	}

	req := new(model.CreateCategoryRequest)
	if err := c.BodyParser(req); err != nil {
		return helper.Response(c, http.StatusBadRequest, "invalid request", nil, err.Error())
	}

	if err := h.validate.Struct(req); err != nil {
		return helper.Response(c, http.StatusBadRequest, "validation error", nil, err.Error())
	}

	result, err := h.categoryUseCase.CreateCategory(*req)
	if err != nil {
		return helper.Response(c, http.StatusBadRequest, "create category failed", nil, err.Error())
	}

	return helper.Response(c, http.StatusCreated, "create category success", result, "")
}

func (h *categoryHandler) GetCategoryByID(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return helper.Response(c, http.StatusBadRequest, "invalid category id", nil, err.Error())
	}

	result, err := h.categoryUseCase.GetCategoryByID(uint(id))
	if err != nil {
		return helper.Response(c, http.StatusNotFound, "get category failed", nil, err.Error())
	}

	return helper.Response(c, http.StatusOK, "get category success", result, "")
}

func (h *categoryHandler) GetAllCategories(c *fiber.Ctx) error {
	limit := helper.QueryInt(c, "limit", 10)
	offset := helper.QueryInt(c, "offset", 0)
	name := c.Query("name", "")

	result, err := h.categoryUseCase.GetAllCategories(limit, offset, name)
	if err != nil {
		return helper.Response(c, http.StatusInternalServerError, "get categories failed", nil, err.Error())
	}

	return helper.Response(c, http.StatusOK, "get categories success", result, "")
}

func (h *categoryHandler) UpdateCategory(c *fiber.Ctx) error {
	// Check if user is admin
	isAdmin, err := utils.CheckIsAdmin(c)
	if err != nil || !isAdmin {
		return helper.Response(c, http.StatusForbidden, "forbidden: only admin can update category", nil, "")
	}

	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return helper.Response(c, http.StatusBadRequest, "invalid category id", nil, err.Error())
	}

	req := new(model.UpdateCategoryRequest)
	if err := c.BodyParser(req); err != nil {
		return helper.Response(c, http.StatusBadRequest, "invalid request", nil, err.Error())
	}

	if err := h.validate.Struct(req); err != nil {
		return helper.Response(c, http.StatusBadRequest, "validation error", nil, err.Error())
	}

	result, err := h.categoryUseCase.UpdateCategory(uint(id), *req)
	if err != nil {
		return helper.Response(c, http.StatusBadRequest, "update category failed", nil, err.Error())
	}

	return helper.Response(c, http.StatusOK, "update category success", result, "")
}

func (h *categoryHandler) DeleteCategory(c *fiber.Ctx) error {
	// Check if user is admin
	isAdmin, err := utils.CheckIsAdmin(c)
	if err != nil || !isAdmin {
		return helper.Response(c, http.StatusForbidden, "forbidden: only admin can delete category", nil, "")
	}

	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return helper.Response(c, http.StatusBadRequest, "invalid category id", nil, err.Error())
	}

	if err := h.categoryUseCase.DeleteCategory(uint(id)); err != nil {
		return helper.Response(c, http.StatusBadRequest, "delete category failed", nil, err.Error())
	}

	return helper.Response(c, http.StatusOK, "delete category success", nil, "")
}
