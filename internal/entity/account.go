package entity

import (
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type Account struct {
	AccountId string    `json:"account_id" gorm:"primary_key"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at" gorm:"autoUpdateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoCreateTime"`
}

func (ac Account) TableName() string {
	return "accounts"
}

func (ac Account) Instantiate() (Account, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(ac.Password), bcrypt.DefaultCost)

	if err != nil {
		return Account{}, err
	}

	return Account{
		AccountId: uuid.New().String(),
		Username:  ac.Username,
		Password:  string(hashedPassword),
		Email:     ac.Email,
	}, nil
}

func (ac Account) CheckPassword(password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(ac.Password), []byte(password))

	if err != nil {
		return err
	}

	return nil
}
