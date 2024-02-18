package account_usecase

import (
	"errors"

	"github.com/jotace1/simple-authentication/internal/entity"
)

func (u *accountUseCase) Create(username string, password string, email string) (*entity.Account, error) {
	existingAccount, err := u.accountRepository.GetByEmail(email, false)

	if err != nil {
		return nil, err
	}

	if existingAccount.AccountId != "" {
		return nil, errors.New("account already exists")
	}

	accountEntity := entity.Account{
		Username: username,
		Password: password,
		Email:    email,
	}

	instantiatedAccount, err := accountEntity.Instantiate()

	if err != nil {
		return nil, err
	}

	return u.accountRepository.Create(instantiatedAccount)
}
