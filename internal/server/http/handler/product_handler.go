package handler

import (
	"fmt"
	"net/http"
	"path/filepath"
	"strconv"
	"tugas_akhir_example/internal/helper"
	"tugas_akhir_example/internal/pkg/entity"
	"tugas_akhir_example/internal/pkg/model"
	"tugas_akhir_example/internal/pkg/usecase"
	"tugas_akhir_example/internal/utils"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type productHandler struct {
	productUseCase usecase.ProductUseCase
	validate       *validator.Validate
}

func NewProductHandler(productUseCase usecase.ProductUseCase, validate *validator.Validate) *productHandler {
	return &productHandler{
		productUseCase: productUseCase,
		validate:       validate,
	}
}

func (h *productHandler) CreateProduct(c *fiber.Ctx) error {
	userID, err := utils.ExtractToken(c)
	if err != nil {
		return helper.Response(c, http.StatusUnauthorized, "unauthorized", nil, err.Error())
	}

	req := new(model.CreateProductRequest)
	_ = c.BodyParser(req)

	if req.CategoryID == 0 {
		categoryIDStr := c.FormValue("category_id")
		if categoryIDStr == "" {
			return helper.Response(c, http.StatusBadRequest, "invalid request", nil, "missing required fields")
		}

		categoryID, err := strconv.ParseUint(categoryIDStr, 10, 32)
		if err != nil {
			return helper.Response(c, http.StatusBadRequest, "invalid category_id", nil, err.Error())
		}
		req.CategoryID = uint(categoryID)
	}

	if req.Name == "" {
		req.Name = c.FormValue("name")
	}
	if req.Description == "" {
		req.Description = c.FormValue("description")
	}
	if req.Price == 0 {
		priceStr := c.FormValue("price")
		if priceStr != "" {
			if p, err := strconv.ParseFloat(priceStr, 64); err == nil {
				req.Price = p
			}
		}
	}
	if req.Quantity == 0 {
		quantityStr := c.FormValue("quantity")
		if quantityStr != "" {
			if q, err := strconv.Atoi(quantityStr); err == nil {
				req.Quantity = q
			}
		}
	}

	if err := h.validate.Struct(req); err != nil {
		return helper.Response(c, http.StatusBadRequest, "validation error", nil, err.Error())
	}

	// Handle file upload
	var imagePath string
	file, err := c.FormFile("image")
	if err == nil && file != nil {
		// Save file
		filename := fmt.Sprintf("%d_%s", userID, filepath.Base(file.Filename))
		destPath := filepath.Join("uploads", "products", filename)
		if err := c.SaveFile(file, destPath); err != nil {
			return helper.Response(c, http.StatusBadRequest, "save file failed", nil, err.Error())
		}
		imagePath = destPath
	}

	result, err := h.productUseCase.CreateProduct(userID, *req, imagePath)
	if err != nil {
		return helper.Response(c, http.StatusBadRequest, "create product failed", nil, err.Error())
	}

	return helper.Response(c, http.StatusCreated, "create product success", result, "")
}

func (h *productHandler) GetMyProducts(c *fiber.Ctx) error {
	userID, err := utils.ExtractToken(c)
	if err != nil {
		return helper.Response(c, http.StatusUnauthorized, "unauthorized", nil, err.Error())
	}

	limit := helper.QueryInt(c, "limit", 10)
	offset := helper.QueryInt(c, "offset", 0)

	filter := entity.FilterProduct{
		Name:   c.Query("search", ""),
		SortBy: c.Query("sort", "newest"),
	}

	if categoryIDStr := c.Query("category_id"); categoryIDStr != "" {
		if id, err := strconv.ParseUint(categoryIDStr, 10, 32); err == nil {
			filter.CategoryID = uint(id)
		}
	}

	result, err := h.productUseCase.GetProductsByStore(userID, limit, offset, filter)
	if err != nil {
		return helper.Response(c, http.StatusInternalServerError, "get products failed", nil, err.Error())
	}

	return helper.Response(c, http.StatusOK, "get products success", result, "")
}

func (h *productHandler) GetAllProducts(c *fiber.Ctx) error {
	limit := helper.QueryInt(c, "limit", 10)
	offset := helper.QueryInt(c, "offset", 0)
	name := c.Query("search", "")
	sortBy := c.Query("sort", "newest")

	var categoryID, storeID uint
	if categoryIDStr := c.Query("category_id"); categoryIDStr != "" {
		if id, err := strconv.ParseUint(categoryIDStr, 10, 32); err == nil {
			categoryID = uint(id)
		}
	}
	if storeIDStr := c.Query("store_id"); storeIDStr != "" {
		if id, err := strconv.ParseUint(storeIDStr, 10, 32); err == nil {
			storeID = uint(id)
		}
	}

	result, err := h.productUseCase.GetAllProducts(limit, offset, name, categoryID, storeID, sortBy)
	if err != nil {
		return helper.Response(c, http.StatusInternalServerError, "get products failed", nil, err.Error())
	}

	return helper.Response(c, http.StatusOK, "get products success", result, "")
}

func (h *productHandler) GetProductByID(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return helper.Response(c, http.StatusBadRequest, "invalid product id", nil, err.Error())
	}

	result, err := h.productUseCase.GetProductByID(uint(id))
	if err != nil {
		return helper.Response(c, http.StatusNotFound, "get product failed", nil, err.Error())
	}

	return helper.Response(c, http.StatusOK, "get product success", result, "")
}

func (h *productHandler) UpdateProduct(c *fiber.Ctx) error {
	userID, err := utils.ExtractToken(c)
	if err != nil {
		return helper.Response(c, http.StatusUnauthorized, "unauthorized", nil, err.Error())
	}

	productID, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return helper.Response(c, http.StatusBadRequest, "invalid product id", nil, err.Error())
	}

	req := new(model.UpdateProductRequest)
	_ = c.BodyParser(req)

	if categoryIDStr := c.FormValue("category_id"); categoryIDStr != "" {
		if id, err := strconv.ParseUint(categoryIDStr, 10, 32); err == nil {
			req.CategoryID = uint(id)
		}
	}

	if req.Name == "" {
		req.Name = c.FormValue("name")
	}
	if req.Description == "" {
		req.Description = c.FormValue("description")
	}

	if priceStr := c.FormValue("price"); priceStr != "" {
		if p, err := strconv.ParseFloat(priceStr, 64); err == nil {
			req.Price = p
		}
	}

	if quantityStr := c.FormValue("quantity"); quantityStr != "" {
		if q, err := strconv.Atoi(quantityStr); err == nil {
			req.Quantity = q
		}
	}

	// Handle file upload
	var imagePath string
	file, err := c.FormFile("image")
	if err == nil && file != nil {
		filename := fmt.Sprintf("%d_%s", userID, filepath.Base(file.Filename))
		destPath := filepath.Join("uploads", "products", filename)
		if err := c.SaveFile(file, destPath); err != nil {
			return helper.Response(c, http.StatusBadRequest, "save file failed", nil, err.Error())
		}
		imagePath = destPath
	}

	result, err := h.productUseCase.UpdateProduct(userID, uint(productID), *req, imagePath)
	if err != nil {
		return helper.Response(c, http.StatusBadRequest, "update product failed", nil, err.Error())
	}

	return helper.Response(c, http.StatusOK, "update product success", result, "")
}

func (h *productHandler) DeleteProduct(c *fiber.Ctx) error {
	userID, err := utils.ExtractToken(c)
	if err != nil {
		return helper.Response(c, http.StatusUnauthorized, "unauthorized", nil, err.Error())
	}

	productID, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return helper.Response(c, http.StatusBadRequest, "invalid product id", nil, err.Error())
	}

	if err := h.productUseCase.DeleteProduct(userID, uint(productID)); err != nil {
		return helper.Response(c, http.StatusBadRequest, "delete product failed", nil, err.Error())
	}

	return helper.Response(c, http.StatusOK, "delete product success", nil, "")
}
