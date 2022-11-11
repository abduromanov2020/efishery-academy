package routes

import (
	"ecommerce-project/handler"

	"github.com/labstack/echo/v4"
)

func PaymentRoutes(e *echo.Echo, paymentHandler *handler.PaymentHandler) {

	g := e.Group("/api/v1")
	g.GET("/payment", paymentHandler.GetListPayment)
	g.GET("/payment/:id", paymentHandler.GetPaymentById)
	g.POST("/payment", paymentHandler.CreatePayment)
	g.DELETE("/payment/:id", paymentHandler.DeletePaymentByID)

}
