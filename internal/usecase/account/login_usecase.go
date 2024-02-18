package account_usecase

import (
	"errors"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type AccountInfo struct {
	AccountId string
	Username  string
	Email     string
}

type Claims struct {
	AccountInfo AccountInfo
	jwt.StandardClaims
}

func (u *accountUseCase) Login(email string, password string) (string, error) {
	existingAccount, err := u.accountRepository.GetByEmail(email, true)

	if err != nil {
		return "", err
	}

	err = existingAccount.CheckPassword(password)

	if err != nil {
		err := errors.New("invalid password")
		return "", err
	}

	expirationTime := time.Now().Add(10 * time.Minute)

	claims := &Claims{
		AccountInfo: AccountInfo{
			AccountId: existingAccount.AccountId,
			Username:  existingAccount.Username,
			Email:     existingAccount.Email,
		},
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	jwtSecret := []byte(os.Getenv("SIGNATURE_SECRET"))

	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
