package http

import (
	route "tugas_akhir_example/internal/server/http/handler"

	"tugas_akhir_example/internal/infrastructure/container"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func HTTPRouteInit(r *fiber.App, containerConf *container.Container) {
	r.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("API running")
	})

	validate := validator.New()

	api := r.Group("/api/v1")

	// Auth routes (no middleware)
	authHandler := route.NewAuthHandler(containerConf.AuthUsc, validate)
	api.Post("/auth/register", authHandler.Register)
	api.Post("/auth/login", authHandler.Login)

	// Protected routes (with JWT middleware)
	protected := api.Group("", route.MiddlewareAuth)

	// User routes
	userHandler := route.NewUserHandler(containerConf.UserUsc, validate)
	protected.Get("/users/me", userHandler.GetMe)
	protected.Patch("/users/me", userHandler.UpdateMe)
	protected.Delete("/users/me", userHandler.DeleteMe)
	protected.Get("/users", userHandler.GetAllUsers)
	protected.Get("/users/:id", userHandler.GetUserByID)

	// Store routes
	storeHandler := route.NewStoreHandler(containerConf.StoreUsc, validate)
	protected.Post("/stores", storeHandler.CreateStore)
	protected.Get("/stores/me", storeHandler.GetMyStore)
	protected.Patch("/stores/:id", storeHandler.UpdateMyStore)
	protected.Delete("/stores/:id", storeHandler.DeleteMyStore)
	protected.Get("/stores", storeHandler.GetAllStores)
	protected.Get("/stores/:id", storeHandler.GetStoreByID)

	// Address routes
	addressHandler := route.NewAddressHandler(containerConf.AddressUsc, validate)
	protected.Post("/addresses", addressHandler.CreateAddress)
	protected.Get("/addresses", addressHandler.GetMyAddresses)
	protected.Get("/addresses/:id", addressHandler.GetMyAddressByID)
	protected.Patch("/addresses/:id", addressHandler.UpdateMyAddress)
	protected.Delete("/addresses/:id", addressHandler.DeleteMyAddress)

	// Category routes
	categoryHandler := route.NewCategoryHandler(containerConf.CategoryUsc, validate)
	protected.Post("/categories", categoryHandler.CreateCategory)
	protected.Get("/categories", categoryHandler.GetAllCategories)
	protected.Get("/categories/:id", categoryHandler.GetCategoryByID)
	protected.Patch("/categories/:id", categoryHandler.UpdateCategory)
	protected.Delete("/categories/:id", categoryHandler.DeleteCategory)

	// Product routes
	productHandler := route.NewProductHandler(containerConf.ProductUsc, validate)
	protected.Post("/products", productHandler.CreateProduct)
	protected.Get("/products/me", productHandler.GetMyProducts)
	protected.Get("/products", productHandler.GetAllProducts)
	protected.Get("/products/:id", productHandler.GetProductByID)
	protected.Patch("/products/:id", productHandler.UpdateProduct)
	protected.Delete("/products/:id", productHandler.DeleteProduct)

	// Transaction routes
	transactionHandler := route.NewTransactionHandler(containerConf.TransactionUsc, validate)
	protected.Post("/transactions", transactionHandler.CreateTransaction)
	protected.Get("/transactions/me", transactionHandler.GetMyTransactions)
	protected.Get("/transactions/me/:id", transactionHandler.GetMyTransactionByID)
	protected.Patch("/transactions/me/:id", transactionHandler.UpdateTransactionStatus)
	protected.Get("/transactions", transactionHandler.GetAllTransactions)
}
