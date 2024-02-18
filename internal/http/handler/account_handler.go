package handler

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	account_usecase "github.com/jotace1/simple-authentication/internal/usecase/account"
	"github.com/jotace1/simple-authentication/pkg/shared"
)

type accountHandler struct {
	Usecase account_usecase.AccountUseCase
}

func NewAccountHandler(
	usecase account_usecase.AccountUseCase,
) accountHandler {
	return accountHandler{
		Usecase: usecase,
	}
}

type CreateAccountInput struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
	Email    string `json:"email" validate:"required"`
}

type CreateAccountOutput struct {
	AccountId string `json:"account_id"`
	Username  string `json:"username"`
	Email     string `json:"email"`
}

func (h accountHandler) CreateAccount(c *fiber.Ctx) error {
	var input CreateAccountInput
	var res shared.Response

	err := json.Unmarshal(c.Body(), &input)
	if err != nil {
		res.BuildError(err, http.StatusBadRequest)
		return c.Status(res.StatusCode).JSON(res.Data)
	}

	validate := validator.New()

	err = validate.Struct(input)
	if err != nil {
		res.BuildError(err, http.StatusBadRequest)
		return c.Status(res.StatusCode).JSON(res.Data)
	}

	result, err := h.Usecase.Create(input.Username, input.Password, input.Email)

	if err != nil {
		res.BuildError(err, http.StatusBadRequest)
		return c.Status(res.StatusCode).JSON(res.Data)
	}

	output := CreateAccountOutput{
		AccountId: result.AccountId,
		Username:  result.Username,
		Email:     result.Email,
	}

	res.BuildSuccess(output, http.StatusCreated)

	return c.Status(res.StatusCode).JSON(res.Data)
}

type LoginInput struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type LoginOutput struct {
	Token string `json:"token"`
}

func (h accountHandler) Login(c *fiber.Ctx) error {
	var input LoginInput
	var res shared.Response

	err := json.Unmarshal(c.Body(), &input)
	if err != nil {
		res.BuildError(err, http.StatusBadRequest)
		return c.Status(res.StatusCode).JSON(res.Data)
	}

	validate := validator.New()

	err = validate.Struct(input)
	if err != nil {
		res.BuildError(err, http.StatusBadRequest)
		return c.Status(res.StatusCode).JSON(res.Data)
	}

	result, err := h.Usecase.Login(input.Email, input.Password)

	if err != nil {
		res.BuildError(err, http.StatusBadRequest)
		return c.Status(res.StatusCode).JSON(res.Data)
	}

	output := LoginOutput{
		Token: result,
	}

	res.BuildSuccess(output, http.StatusOK)

	return c.Status(res.StatusCode).JSON(res.Data)
}
