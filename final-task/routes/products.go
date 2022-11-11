package routes

import (
	"ecommerce-project/handler"

	"github.com/labstack/echo/v4"
)

func ProductsRoutes(e *echo.Echo, productsHandler *handler.ProductsHandler) {

	g := e.Group("/api/v1")
	g.GET("/product", productsHandler.GetListProduct)
	g.GET("/product/:id", productsHandler.GetProductByID)
	g.GET("/product/price/:minPrice/:maxPrice", productsHandler.GetProductByPrice)
	g.GET("/product/category/:category", productsHandler.GetProductByCategory)
	g.POST("/product", productsHandler.CreateProduct)
	g.PUT("/product/:id", productsHandler.UpdateProductByID)
	g.DELETE("/product/:id", productsHandler.DeleteProductByID)
}
