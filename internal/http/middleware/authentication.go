package middleware

import (
	"errors"
	"net/http"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/jotace1/simple-authentication/pkg/shared"
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

func AuthMiddleware(c *fiber.Ctx) error {
	var res shared.Response
	tknStr := c.GetReqHeaders()["Authorization"]

	formattedTokenWithoutBearer := tknStr[7:]

	if formattedTokenWithoutBearer == "" {
		err := errors.New("you must provide a token in the Authorization header")
		res.BuildError(err, http.StatusUnauthorized)
		c.Status(res.StatusCode).JSON(res.Data)
		return nil
	}

	claims := &Claims{}

	tkn, err := jwt.ParseWithClaims(formattedTokenWithoutBearer, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SIGNATURE_SECRET")), nil
	})

	if err != nil && err == jwt.ErrSignatureInvalid {
		res.BuildError(err, http.StatusUnauthorized)
		c.Status(res.StatusCode).JSON(res.Data)
		return nil
	}

	if err != nil {
		res.BuildError(err, http.StatusInternalServerError)
		c.Status(res.StatusCode).JSON(res.Data)
		return nil
	}

	if !tkn.Valid {
		err := errors.New("the token you provided is invalid")
		res.BuildError(err, http.StatusUnauthorized)
		c.Status(http.StatusUnauthorized).JSON(res.Data)
		return nil
	}

	c.Next()
	return nil
}
