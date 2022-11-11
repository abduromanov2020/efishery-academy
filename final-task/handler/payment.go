package handler

import (
	"ecommerce-project/entity"
	"ecommerce-project/entity/response"
	"ecommerce-project/usecase"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type PaymentHandler struct {
	paymentUsecase usecase.IPaymentUsecase
}

func NewPaymentHandler(paymentUsecase usecase.IPaymentUsecase) *PaymentHandler {
	return &PaymentHandler{paymentUsecase: paymentUsecase}
}

func (p PaymentHandler) CreatePayment(c echo.Context) error {
	paymentRequest := entity.CreatePaymentRequest{}

	file, _ := c.FormFile("file")
	cId := c.FormValue("cart_id")

	cartId, _ := strconv.Atoi(cId)

	paymentRequest.File = file.Filename
	paymentRequest.CartID = uint(cartId)

	err := p.paymentUsecase.CreatePayment(paymentRequest)

	if err != nil {
		return c.JSON(http.StatusBadRequest, response.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: "Invalid Req Body",
			Data:    err.Error(),
		})
	}

	return c.JSON(http.StatusOK, response.BaseResponse{
		Code:    http.StatusOK,
		Message: "Success Create Payment",
		Data:    paymentRequest,
	})
}

func (p PaymentHandler) GetListPayment(c echo.Context) error {
	payments, err := p.paymentUsecase.GetListPayment()

	if err != nil {
		return c.JSON(http.StatusBadRequest, response.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: "Invalid Req Body",
			Data:    err.Error(),
		})

	}

	if len(payments) == 0 {
		return c.JSON(http.StatusOK, response.BaseResponse{
			Code:    http.StatusOK,
			Message: "Empty Payment",
			Data:    []entity.Payment{},
		})
	}

	return c.JSON(http.StatusOK, response.BaseResponse{
		Code:    http.StatusOK,
		Message: "Success Get List Payment",
		Data:    payments,
	})
}

func (p PaymentHandler) GetPaymentById(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	payment, err := p.paymentUsecase.GetPaymentByID(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, response.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: "Failed Get Payment By Id",
			Data:    err.Error(),
		})
	}

	return c.JSON(http.StatusOK, response.BaseResponse{
		Code:    http.StatusOK,
		Message: "Success Get Payment By Id",
		Data:    payment,
	})

}

func (p PaymentHandler) DeletePaymentByID(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	err := p.paymentUsecase.DeletePaymentByID(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, response.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: "Failed Delete Payment",
			Data:    err.Error(),
		})
	}

	return c.JSON(http.StatusOK, response.BaseResponse{
		Code:    http.StatusOK,
		Message: "Success Delete Payment",
		Data:    []entity.Payment{},
	})
}
