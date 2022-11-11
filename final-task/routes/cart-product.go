package routes

import (
	"ecommerce-project/handler"

	"github.com/labstack/echo/v4"
)

func CartProductRoutes(c *echo.Echo, cartProductHandler *handler.CartProductHandler) {
	g := c.Group("/api/v1")
	g.GET("/cartProduct", cartProductHandler.GetListCartProduct)
	g.GET("/cartProduct/:id", cartProductHandler.GetCartProductByID)
	g.POST("/cartProduct", cartProductHandler.CreateCartProduct)
	g.PUT("/cartProduct/:id", cartProductHandler.UpdateCartProductByID)
	g.DELETE("/cartProduct/:id", cartProductHandler.DeleteCartProductByID)
}
