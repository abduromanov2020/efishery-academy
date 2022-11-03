package routes

import (
	"golangday7/handler"

	"github.com/gofiber/fiber/v2"
)

func Routes(app fiber.Router, userHandler *handler.UserHandler) {
	r := app.Group("api/v1")

	r.Post("/user", userHandler.CreateUser)
	r.Get("/users", userHandler.GetListUser)
	r.Get("/users/:id", userHandler.GetUserById)
	r.Put("/users/:id", userHandler.UpdateUserById)
	r.Delete("/users/:id", userHandler.DeleteUserById)
}
