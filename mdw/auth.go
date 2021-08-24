package mdw

import (
	"crypto/subtle"

	"github.com/labstack/echo/v4"
)

func BasicAuth(username, password string, e echo.Context) (bool, error) {
	if subtle.ConstantTimeCompare([]byte(username), []byte("hoangdo")) == 1 &&
		subtle.ConstantTimeCompare([]byte(password), []byte("123")) == 1 {
		e.Set("username", username)
		return true, nil
	}
	return false, nil
}
