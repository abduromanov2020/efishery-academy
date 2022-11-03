package main

import (
	"golangday7/config"
	"golangday7/handler"
	"golangday7/repository"
	"golangday7/routes"
	"golangday7/usecase"

	"github.com/gofiber/fiber/v2"
)

func main() {
	config.Database()
	config.AutoMigrate()

	app := fiber.New()

	userRepository := repository.NewUserRepository(config.DB)
	userUseCase := usecase.NewUserUseCase(userRepository)
	userHandler := handler.NewUserHandler(userUseCase)

	routes.Routes(app, userHandler)

	app.Listen(":3000")
}
