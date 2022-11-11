package handler

import (
	"ecommerce-project/entity/response"
	"ecommerce-project/usecase"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type CartHandler struct {
	cartUsecase usecase.ICartUsecase
}

func NewCartHandler(cartUsecase usecase.ICartUsecase) *CartHandler {
	return &CartHandler{cartUsecase: cartUsecase}
}

func (cart CartHandler) GetListCart(c echo.Context) error {
	carts, err := cart.cartUsecase.GetListCart()

	if err != nil {
		return c.JSON(http.StatusBadRequest, response.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: "Invalid Req Body",
			Data:    err.Error(),
		})

	}

	if len(carts) == 0 {
		return c.JSON(http.StatusOK, response.BaseResponse{
			Code:    http.StatusOK,
			Message: "Empty Cart",
			Data:    []interface{}{},
		})
	}

	return c.JSON(http.StatusOK, response.BaseResponse{
		Code:    http.StatusOK,
		Message: "Success Get List Cart",
		Data:    carts,
	})
}

func (cart CartHandler) GetCartByID(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	cartByID, err := cart.cartUsecase.GetCartByID(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, response.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: "Failed Get Cart By ID",
			Data:    err.Error(),
		})

	}

	return c.JSON(http.StatusOK, response.BaseResponse{
		Code:    http.StatusOK,
		Message: "Success Get Cart",
		Data:    cartByID,
	})
}
