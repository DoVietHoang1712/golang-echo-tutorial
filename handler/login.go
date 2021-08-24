package handler

import (
	"log"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func Login(c echo.Context) error {
	username := c.Get("username").(string)
	claim := jwt.MapClaims{}
	claim["username"] = username
	claim["exp"] = time.Now().Add(5 * time.Minute).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	t, err := token.SignedString([]byte("123"))
	if err != nil {
		log.Fatalln(err)
		return err
	}
	return c.JSON(http.StatusOK, map[string]string{
		"access_token": t,
	})
}
