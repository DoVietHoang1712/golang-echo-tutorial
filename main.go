package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/DoVietHoang1712/golang-echo-tutorial/handler"
	"github.com/DoVietHoang1712/golang-echo-tutorial/mdw"
	"github.com/go-playground/validator"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type User struct {
	Name  string `json:"name" form:"name" query:"name" validate:"required"`
	Email string `json:"email" form:"email" query:"email" validate:"required,email"`
}

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

func main() {
	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "user_agent=${user_agent}, method=${method}, uri=${uri}, status=${status}\n",
	}))

	isLogedIn := middleware.JWT([]byte("123"))
	// e.Use(middleware.BasicAuth(mdw.BasicAuth))
	e.GET("/", hello, isLogedIn)
	e.GET("/product/:id", handler.GetProduct)
	e.POST("/login", handler.Login, middleware.BasicAuth((mdw.BasicAuth)))
	e.POST("product", func(context echo.Context) (err error) {
		u := new(User)
		if err = context.Bind(u); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		if err = context.Validate(u); err != nil {
			return err
		}
		return context.JSON(http.StatusOK, u)
	})
	e.POST("/test", func(context echo.Context) error {
		u := new(User)
		if err := context.Bind(u); err != nil {
			return err
		}
		return context.JSON(http.StatusOK, u)
	})
	log.Fatal(e.Start(":1232"))
}

func hello(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	username := claims["username"]
	fmt.Println(username)
	return c.JSON(http.StatusOK, map[string]string{"data": "123"})
}
