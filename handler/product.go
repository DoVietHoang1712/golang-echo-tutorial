package handler

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type Product struct {
	ID primitive.ObjectID `json:"_id" bson:"_id"`
	Name string `json:"product_name" bson:"product_name"`
	Price int `json:"price" bson:"price"`
	Currency string `json:"currency" bson:"currency"`
	Quantity int `json:"quantity" bson:"quantity"`
	Discount int `json:"discount,omitempty" bson:"discount,omitempty"`
	Vendor string `json:"vendor" bson:"vendor"`
	Accessories []string `json:"accessories,omitempty" bson:"accessories,omitempty"`
	SkuID string `json:"sku_id" bson:"sku_id"`
}

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
