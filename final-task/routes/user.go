package routes

import (
	"ecommerce-project/handler"

	"github.com/labstack/echo/v4"
)

func UserRoutes(e *echo.Echo, userHandler *handler.UserHandler) {

	g := e.Group("/api/v1")
	g.GET("/user", userHandler.GetListUser)
	g.GET("/user/:id", userHandler.GetUserByID)
	g.POST("/user", userHandler.CreateUser)
	g.PUT("/user/:id", userHandler.UpdateUserByID)
	g.DELETE("/user/:id", userHandler.DeleteUserByID)
}
