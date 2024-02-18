package account_usecase_test

import (
	"errors"
	"testing"

	"github.com/jotace1/simple-authentication/internal/entity"
	account_usecase "github.com/jotace1/simple-authentication/internal/usecase/account"
	"github.com/jotace1/simple-authentication/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateAccountUseCase(t *testing.T) {
	mockedAccountRepository := &mocks.AccountRepository{}
	usecase := account_usecase.NewAccountUseCase(mockedAccountRepository)

	t.Run("should create an account successfully", func(t *testing.T) {
		testEmail := "jhondoe@email.com"
		testUsername := "jhondoe"
		testPassword := "123456"

		createMethodInput, err := entity.Account{
			Username: testUsername,
			Password: testPassword,
			Email:    testEmail,
		}.Instantiate()

		if err != nil {
			t.Errorf("expected no error, got %s", err)
		}

		createMethodOutput := &entity.Account{
			AccountId: createMethodInput.AccountId,
			Username:  createMethodInput.Username,
			Password:  createMethodInput.Password,
			Email:     createMethodInput.Email,
		}

		mockedAccountRepository.On("GetByEmail", testEmail, false).Return(&entity.Account{}, nil).Once()
		mockedAccountRepository.On("Create", mock.Anything).Return(createMethodOutput, nil).Once()

		createdAccount, err := usecase.Create(testUsername, testPassword, testEmail)
		if err != nil {
			t.Errorf("expected no error, got %s", err)
		}

		assert.Equal(t, createMethodOutput, createdAccount)
		assert.Equal(t, nil, err)
		mockedAccountRepository.AssertExpectations(t)
	})

	t.Run("should not create an account when there is already an account with the same email", func(t *testing.T) {
		testEmail := "jhondoe@email.com"
		testUsername := "jhondoe"
		testPassword := "123456"

		existingAccount := entity.Account{
			AccountId: "123",
			Username:  testUsername,
			Password:  testPassword,
			Email:     testEmail,
		}

		mockedAccountRepository.On("GetByEmail", testEmail, false).Return(&existingAccount, nil).Once()

		_, err := usecase.Create(testUsername, testPassword, testEmail)

		expectedError := errors.New("account already exists")

		assert.Equal(t, expectedError, err)
		mockedAccountRepository.AssertExpectations(t)
	})

	t.Run("should not create an account when get by email method throws an error", func(t *testing.T) {
		testEmail := "jhondoe@email.com"
		testUsername := "jhondoe"
		testPassword := "123456"

		existingAccount := entity.Account{
			AccountId: "123",
			Username:  testUsername,
			Password:  testPassword,
			Email:     testEmail,
		}

		mockedAccountRepository.On("GetByEmail", testEmail, false).Return(&existingAccount, errors.New("error")).Once()

		_, err := usecase.Create(testUsername, testPassword, testEmail)

		expectedError := errors.New("error")

		assert.Equal(t, expectedError, err)
		mockedAccountRepository.AssertExpectations(t)
	})
}
