package routes

import (
	"ecommerce-project/handler"

	"github.com/labstack/echo/v4"
)

func ProductDetailRoutes(e *echo.Echo, productDetailHandler *handler.ProductDetailHandler) {

	g := e.Group("/api/v1")
	g.GET("/productDetail", productDetailHandler.GetListProductDetail)
	g.GET("/productDetail/:id", productDetailHandler.GetProductDetailByID)
	g.POST("/productDetail", productDetailHandler.CreateProductDetail)
	g.PUT("/productDetail/:id", productDetailHandler.UpdateProductDetailByID)
	g.DELETE("/productDetail/:id", productDetailHandler.DeleteProductDetailByID)
}
