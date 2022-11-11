package routes

import (
	"ecommerce-project/handler"

	"github.com/labstack/echo/v4"
)

func CartRoutes(c *echo.Echo, cartHandler *handler.CartHandler) {
	g := c.Group("/api/v1")
	g.GET("/cart", cartHandler.GetListCart)
	g.GET("/cart/:id", cartHandler.GetCartByID)
}
