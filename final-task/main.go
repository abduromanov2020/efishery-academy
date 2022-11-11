package main

import (
	"ecommerce-project/config"
	"ecommerce-project/handler"
	"ecommerce-project/repository"
	"ecommerce-project/routes"
	"ecommerce-project/usecase"

	"github.com/labstack/echo/v4"
)

func main() {
	config.Database()
	config.AutoMigrate()

	e := echo.New()

	userRepository := repository.NewUserRepository(config.DB)
	userUseCase := usecase.NewUserUsecase(userRepository)
	userHandler := handler.NewUserHandler(userUseCase)

	routes.UserRoutes(e, userHandler)

	productsRepository := repository.NewProductsRepository(config.DB)
	productsUseCase := usecase.NewProductsUsecase(productsRepository)
	productsHandler := handler.NewProductsHandler(productsUseCase)

	routes.ProductsRoutes(e, productsHandler)

	productDetailRepository := repository.NewProductDetailRepository(config.DB)
	productDetailUseCase := usecase.NewProductDetailUsecase(productDetailRepository)
	productDetailHandler := handler.NewProductDetailHandler(productDetailUseCase)

	routes.ProductDetailRoutes(e, productDetailHandler)

	categoryRepository := repository.NewCategoryRepository(config.DB)
	categoryUseCase := usecase.NewCategoryUsecase(categoryRepository)
	categoryHandler := handler.NewCategoryHandler(categoryUseCase)

	routes.CategoryRoutes(e, categoryHandler)

	paymentRepository := repository.NewPaymentRepository(config.DB)
	paymentUseCase := usecase.NewPaymentUsecase(paymentRepository)
	paymentHandler := handler.NewPaymentHandler(paymentUseCase)

	routes.PaymentRoutes(e, paymentHandler)

	cartRepository := repository.NewCartRepository(config.DB)
	cartUseCase := usecase.NewCartUsecase(cartRepository)
	cartHandler := handler.NewCartHandler(cartUseCase)

	routes.CartRoutes(e, cartHandler)

	cartProductRepository := repository.NewCartProductRepository(config.DB)
	cartProductUseCase := usecase.NewCartProductUsecase(cartProductRepository)
	cartProductHandler := handler.NewCartProductHandler(cartProductUseCase)

	routes.CartProductRoutes(e, cartProductHandler)

	e.Logger.Fatal(e.Start(":1323"))
}
