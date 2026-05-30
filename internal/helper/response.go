package helper

import (
	"github.com/gofiber/fiber/v2"
)

type ResponseData struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Error   string      `json:"error,omitempty"`
}

type PaginatedResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Error   string      `json:"error,omitempty"`
}

func Response(c *fiber.Ctx, code int, message string, data interface{}, errorMsg string) error {
	return c.Status(code).JSON(ResponseData{
		Code:    code,
		Message: message,
		Data:    data,
		Error:   errorMsg,
	})
}
