package server

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/jotace1/simple-authentication/db"
	"github.com/jotace1/simple-authentication/internal/http"
	"github.com/jotace1/simple-authentication/internal/repository"
	account_usecase "github.com/jotace1/simple-authentication/internal/usecase/account"
)

type Server struct {
	server *fiber.App
}

func NewServer() Server {
	return Server{
		server: fiber.New(),
	}
}

func (s *Server) Run() {

	s.server.Use(requestid.New())
	s.server.Use(recover.New())

	db := db.Init()

	accountRepository := repository.NewAccountRepository(db)
	accountUseCase := account_usecase.NewAccountUseCase(accountRepository)

	port := "8080"
	if value, exists := os.LookupEnv("APP_PORT"); exists {
		port = value
	}

	http.ConfigRoutes(s.server, db, accountUseCase, nil)

	if err := s.server.Listen(fmt.Sprintf(":%s", port)); err != nil {
		panic("Error trying to start the server")
	}
}
