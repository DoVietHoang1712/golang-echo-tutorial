package handler

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetProduct(e echo.Context) error {
	var products = []map[int]string{{1: "laptop"}, {2: "tv"}, {3: "mobile"}}
	var product map[int]string
	for _, p := range products {
		for k := range p {
			pID, err := strconv.Atoi(e.Param("id"))
			if err != nil {
				return err
			}
			if pID == k {
				product = p
			}
		}
	}
	if product == nil {
		return e.JSON(http.StatusNotFound, "product not found")
	}
	return e.JSON(http.StatusOK, product)
}
