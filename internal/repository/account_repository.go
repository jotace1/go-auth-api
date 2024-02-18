package repository

import (
	"errors"

	"github.com/jotace1/simple-authentication/internal/entity"
	"gorm.io/gorm"
)

type AccountRepository interface {
	Create(input entity.Account) (*entity.Account, error)
	GetByEmail(email string, throwNotFoundError bool) (*entity.Account, error)
}

type repository struct {
	db *gorm.DB
}

func NewAccountRepository(db *gorm.DB) AccountRepository {
	return &repository{
		db: db,
	}
}

func (r *repository) Create(input entity.Account) (*entity.Account, error) {
	result := r.db.Create(&input)

	if result.Error != nil {
		return nil, result.Error
	}

	return &input, nil
}

func (r *repository) GetByEmail(email string, throwNotFoundError bool) (*entity.Account, error) {
	var account *entity.Account

	result := r.db.Where("email = ?", email).First(&account)

	if result.Error != nil && throwNotFoundError && errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, errors.New("account not found")
	}

	if result.Error != nil && !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, result.Error
	}

	return account, nil
}
