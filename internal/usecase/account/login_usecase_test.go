package account_usecase_test

import (
	"errors"
	"testing"

	"github.com/jotace1/simple-authentication/internal/entity"
	account_usecase "github.com/jotace1/simple-authentication/internal/usecase/account"
	"github.com/jotace1/simple-authentication/mocks"
	"github.com/stretchr/testify/assert"
)

func TestLoginUseCase(t *testing.T) {
	mockedAccountRepository := &mocks.AccountRepository{}
	usecase := account_usecase.NewAccountUseCase(mockedAccountRepository)

	t.Run("should login when there is an existing account successfully", func(t *testing.T) {
		userPassword := "123456"

		existingAccount, err := entity.Account{
			AccountId: "123",
			Username:  "jhondoe",
			Password:  userPassword,
			Email:     "jhondoe@email.com",
		}.Instantiate()

		if err != nil {
			t.Errorf("expected no error, got %s", err)
		}

		mockedAccountRepository.On("GetByEmail", existingAccount.Email, true).Return(&existingAccount, nil).Once()

		token, err := usecase.Login(existingAccount.Email, userPassword)
		if err != nil {
			t.Errorf("expected no error, got %s", err)
		}

		assert.NotNil(t, token)
		assert.Equal(t, nil, err)
		mockedAccountRepository.AssertExpectations(t)
	})

	t.Run("should not login when there password does not match", func(t *testing.T) {

		existingAccount, err := entity.Account{
			AccountId: "123",
			Username:  "jhondoe",
			Password:  "123456",
			Email:     "jhondoe@email.com",
		}.Instantiate()

		if err != nil {
			t.Errorf("expected no error, got %s", err)
		}

		mockedAccountRepository.On("GetByEmail", existingAccount.Email, true).Return(&existingAccount, nil).Once()

		_, err = usecase.Login(existingAccount.Email, "1234")

		expectedError := errors.New("invalid password")

		assert.Equal(t, expectedError, err)
		mockedAccountRepository.AssertExpectations(t)
	})

	t.Run("should not login when there is not an existing account", func(t *testing.T) {

		mockedAccountRepository.On("GetByEmail", "jhondoe@email.com", true).Return(nil, errors.New("account not found")).Once()

		_, err := usecase.Login("jhondoe@email.com", "123456")

		expectedError := errors.New("account not found")

		assert.Equal(t, expectedError, err)
		mockedAccountRepository.AssertExpectations(t)
	})

}
