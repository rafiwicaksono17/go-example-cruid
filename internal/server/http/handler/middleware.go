package handler

import (
	"net/http"
	"tugas_akhir_example/internal/helper"
	"tugas_akhir_example/internal/utils"

	"github.com/gofiber/fiber/v2"
)

func MiddlewareAuth(ctx *fiber.Ctx) error {
	authHeader := ctx.Get("Authorization")
	if authHeader == "" {
		return helper.Response(ctx, http.StatusUnauthorized, "unauthorized", nil, "missing authorization header")
	}

	// Extract Bearer token
	tokenString := authHeader
	if len(authHeader) > 7 && authHeader[:7] == "Bearer " {
		tokenString = authHeader[7:]
	}

	claims, err := utils.ValidateToken(tokenString)
	if err != nil {
		return helper.Response(ctx, http.StatusUnauthorized, "unauthorized", nil, "invalid or expired token")
	}

	ctx.Locals("user_id", claims.UserID)
	ctx.Locals("email", claims.Email)
	ctx.Locals("is_admin", claims.IsAdmin)

	return ctx.Next()
}
