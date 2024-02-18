package handler_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/jotace1/simple-authentication/internal/entity"
	"github.com/jotace1/simple-authentication/internal/http/handler"
	"github.com/jotace1/simple-authentication/mocks"
	"github.com/stretchr/testify/assert"
)

func Test_AccountHandler_Create(t *testing.T) {
	mockedAccountUsecase := &mocks.AccountUseCase{}
	accountHandler := handler.NewAccountHandler(mockedAccountUsecase)
	router := fiber.New()
	router.Post("/account", accountHandler.CreateAccount)

	t.Run("should create account succesfully", func(t *testing.T) {
		input := &handler.CreateAccountInput{
			Username: "jhondoe",
			Password: "123456",
			Email:    "jhondoe@email.com",
		}

		createMethodOutput := &entity.Account{
			AccountId: "123",
			Username:  "jhondoe",
			Email:     "jhondoe@email.com",
			Password:  "somehashedpassword",
		}

		output := handler.CreateAccountOutput{
			AccountId: "123",
			Username:  "jhondoe",
			Email:     "jhondoe@email.com",
		}

		byteInput, _ := json.Marshal(input)
		mockedAccountUsecase.On("Create", input.Username, input.Password, input.Email).Return(createMethodOutput, nil).Once()

		request, err := http.NewRequest(http.MethodPost, "/account", strings.NewReader(string(byteInput)))
		if err != nil {
			t.Errorf("expected no error, got %s", err)
		}

		resp, err := router.Test(request)
		if err != nil {
			t.Errorf("expected no error, got %s", err)
		}

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			t.Errorf("expected no error, got %s", err)
		}

		respBody, err := json.Marshal(output)
		assert.NoError(t, err)
		assert.Equal(t, bytes.TrimSpace(respBody), bytes.TrimSpace(body))
		assert.Equal(t, http.StatusCreated, resp.StatusCode)
		mockedAccountUsecase.AssertExpectations(t)
	})

	t.Run("should not create an account when there is some field missing", func(t *testing.T) {
		input := &handler.CreateAccountInput{
			Username: "jhondoe",
			Password: "123456",
		}

		byteInput, _ := json.Marshal(input)

		request, err := http.NewRequest(http.MethodPost, "/account", strings.NewReader(string(byteInput)))
		if err != nil {
			t.Errorf("expected no error, got %s", err)
		}

		resp, err := router.Test(request)
		if err != nil {
			t.Errorf("expected no error, got %s", err)
		}

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			t.Errorf("expected no error, got %s", err)
		}

		expectedError := `{"message":"Validation error","errors":{"Email":"Field Email is required"}}`

		assert.NoError(t, err)
		assert.Equal(t, bytes.TrimSpace([]byte(expectedError)), bytes.TrimSpace(body))
		assert.Equal(t, http.StatusUnprocessableEntity, resp.StatusCode)
		mockedAccountUsecase.AssertExpectations(t)
	})

	t.Run("should not create an account when there is an error on create method", func(t *testing.T) {
		input := &handler.CreateAccountInput{
			Username: "jhondoe",
			Password: "123456",
			Email:    "jhondoe@email.com",
		}

		byteInput, _ := json.Marshal(input)
		mockedAccountUsecase.On("Create", input.Username, input.Password, input.Email).Return(nil, errors.New("error")).Once()

		request, err := http.NewRequest(http.MethodPost, "/account", strings.NewReader(string(byteInput)))
		if err != nil {
			t.Errorf("expected no error, got %s", err)
		}

		resp, err := router.Test(request)
		if err != nil {
			t.Errorf("expected no error, got %s", err)
		}

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			t.Errorf("expected no error, got %s", err)
		}

		expectedError := `{"message":"Error","errors":"error"}`

		assert.NoError(t, err)
		assert.Equal(t, bytes.TrimSpace([]byte(expectedError)), bytes.TrimSpace(body))
		assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
		mockedAccountUsecase.AssertExpectations(t)
	})
}

func Test_AccountHandler_Login(t *testing.T) {
	mockedAccountUsecase := &mocks.AccountUseCase{}
	accountHandler := handler.NewAccountHandler(mockedAccountUsecase)
	router := fiber.New()
	router.Post("/account/login", accountHandler.Login)

	t.Run("should login an user succesfully", func(t *testing.T) {
		input := &handler.LoginInput{
			Email:    "jhondoe@email.com",
			Password: "123456",
		}

		token := "validtoken"

		output := handler.LoginOutput{
			Token: token,
		}

		byteInput, _ := json.Marshal(input)
		mockedAccountUsecase.On("Login", input.Email, input.Password).Return(token, nil).Once()

		request, err := http.NewRequest(http.MethodPost, "/account/login", strings.NewReader(string(byteInput)))
		if err != nil {
			t.Errorf("expected no error, got %s", err)
		}

		resp, err := router.Test(request)
		if err != nil {
			t.Errorf("expected no error, got %s", err)
		}

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			t.Errorf("expected no error, got %s", err)
		}

		respBody, err := json.Marshal(output)
		assert.NoError(t, err)
		assert.Equal(t, bytes.TrimSpace(respBody), bytes.TrimSpace(body))
		assert.Equal(t, http.StatusOK, resp.StatusCode)
		mockedAccountUsecase.AssertExpectations(t)
	})

	t.Run("should not login an user if there is an error in login method", func(t *testing.T) {
		input := &handler.LoginInput{
			Email:    "jhondoe@email.com",
			Password: "123456",
		}

		byteInput, _ := json.Marshal(input)
		mockedAccountUsecase.On("Login", input.Email, input.Password).Return("", errors.New("error")).Once()

		request, err := http.NewRequest(http.MethodPost, "/account/login", strings.NewReader(string(byteInput)))
		if err != nil {
			t.Errorf("expected no error, got %s", err)
		}

		resp, err := router.Test(request)
		if err != nil {
			t.Errorf("expected no error, got %s", err)
		}

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			t.Errorf("expected no error, got %s", err)
		}

		expectedError := `{"message":"Error","errors":"error"}`

		assert.Equal(t, bytes.TrimSpace([]byte(expectedError)), bytes.TrimSpace(body))
		assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
		mockedAccountUsecase.AssertExpectations(t)
	})

	t.Run("should not login an user if there is some field missing", func(t *testing.T) {
		input := &handler.LoginInput{
			Email: "jhondoe@email.com",
		}

		byteInput, _ := json.Marshal(input)

		request, err := http.NewRequest(http.MethodPost, "/account/login", strings.NewReader(string(byteInput)))
		if err != nil {
			t.Errorf("expected no error, got %s", err)
		}

		resp, err := router.Test(request)
		if err != nil {
			t.Errorf("expected no error, got %s", err)
		}

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			t.Errorf("expected no error, got %s", err)
		}

		expectedError := `{"message":"Validation error","errors":{"Password":"Field Password is required"}}`

		assert.Equal(t, bytes.TrimSpace([]byte(expectedError)), bytes.TrimSpace(body))
		assert.Equal(t, http.StatusUnprocessableEntity, resp.StatusCode)
		mockedAccountUsecase.AssertExpectations(t)
	})
}
