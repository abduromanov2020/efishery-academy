package handler

import (
	"ecommerce-project/entity"
	"ecommerce-project/entity/response"
	"ecommerce-project/usecase"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type CartProductHandler struct {
	cartProductUsecase usecase.ICartProductUsecase
}

func NewCartProductHandler(cartProductUsecase usecase.ICartProductUsecase) *CartProductHandler {
	return &CartProductHandler{cartProductUsecase: cartProductUsecase}
}

func (cart CartProductHandler) CreateCartProduct(c echo.Context) error {
	var req entity.CreateCartProductRequest

	c.Bind(&req)

	err := cart.cartProductUsecase.CreateCartProduct(req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: "Invalid Req Body",
			Data:    err.Error(),
		})
	}

	return c.JSON(http.StatusOK, response.BaseResponse{
		Code:    http.StatusOK,
		Message: "Success Create Cart",
		Data:    req,
	})
}

func (cart CartProductHandler) GetListCartProduct(c echo.Context) error {
	cartProduct, err := cart.cartProductUsecase.GetListCartProduct()

	if err != nil {
		return c.JSON(http.StatusBadRequest, response.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: "Invalid Req Body",
			Data:    err.Error(),
		})

	}

	if len(cartProduct) == 0 {
		return c.JSON(http.StatusOK, response.BaseResponse{
			Code:    http.StatusOK,
			Message: "Empty Cart",
			Data:    []entity.Cart_Product{},
		})
	}

	return c.JSON(http.StatusOK, response.BaseResponse{
		Code:    http.StatusOK,
		Message: "Success Get List Cart",
		Data:    cartProduct,
	})
}

func (cart CartProductHandler) GetCartProductByID(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	cartProductById, err := cart.cartProductUsecase.GetCartProductByID(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, response.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: "Failed Get Cart Product By Id",
			Data:    err.Error(),
		})
	}

	return c.JSON(http.StatusOK, response.BaseResponse{
		Code:    http.StatusOK,
		Message: "Success Get Cart By Id",
		Data:    cartProductById,
	})
}

func (cartProduct CartProductHandler) UpdateCartProductByID(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	cartRequest := entity.UpdateCartProductRequest{}

	c.Bind(&cartRequest)

	err := cartProduct.cartProductUsecase.UpdateCartProductByID(id, cartRequest)

	if err != nil {
		return c.JSON(http.StatusBadRequest, response.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: "Failed Update Cart",
			Data:    err.Error(),
		})
	}

	return c.JSON(http.StatusOK, response.BaseResponse{
		Code:    http.StatusOK,
		Message: "Success Update Cart",
		Data:    cartRequest,
	})
}

func (cart CartProductHandler) DeleteCartProductByID(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	err := cart.cartProductUsecase.DeleteCartProductByID(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, response.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: "Failed Delete Cart",
			Data:    err.Error(),
		})
	}

	return c.JSON(http.StatusOK, response.BaseResponse{
		Code:    http.StatusOK,
		Message: "Success Delete Cart",
		Data:    []entity.Cart_Product{},
	})
}
