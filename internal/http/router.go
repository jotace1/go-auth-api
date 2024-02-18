package http

import (
	nethttp "net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/jotace1/simple-authentication/internal/http/handler"
	"github.com/jotace1/simple-authentication/internal/http/middleware"
	accountUsecase "github.com/jotace1/simple-authentication/internal/usecase/account"
	"gorm.io/gorm"
)

func ConfigRoutes(
	router *fiber.App,
	db *gorm.DB,
	accountUsecase accountUsecase.AccountUseCase,
	httpClient *nethttp.Client,
) *fiber.App {

	accountHandler := handler.NewAccountHandler(accountUsecase)

	router.Post("/account", accountHandler.CreateAccount)
	router.Post("/account/login", accountHandler.Login)

	authenticatedGroup := router.Group("/auth", middleware.AuthMiddleware)
	authenticatedGroup.Get("test", func(c *fiber.Ctx) error {
		return c.JSON("You are authenticated")
	})

	return router
}
