package account_usecase

import (
	"github.com/jotace1/simple-authentication/internal/entity"
	"github.com/jotace1/simple-authentication/internal/repository"
)

type AccountUseCase interface {
	Create(username string, password string, email string) (*entity.Account, error)
	Login(password string, email string) (string, error)
}

type accountUseCase struct {
	accountRepository repository.AccountRepository
}

func NewAccountUseCase(accountRepository repository.AccountRepository) AccountUseCase {
	return &accountUseCase{
		accountRepository: accountRepository,
	}
}
