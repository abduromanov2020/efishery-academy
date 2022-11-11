package routes

import (
	"ecommerce-project/handler"

	"github.com/labstack/echo/v4"
)

func CategoryRoutes(e *echo.Echo, categoryHandler *handler.CategoryHandler) {

	g := e.Group("/api/v1")
	g.GET("/category", categoryHandler.GetListCategory)
	g.GET("/category/:id", categoryHandler.GetCategoryById)
	g.POST("/category", categoryHandler.CreateCategory)
	g.PUT("/category/:id", categoryHandler.UpdateCategoryByID)
	g.DELETE("/category/:id", categoryHandler.DeleteCategoryByID)
}
