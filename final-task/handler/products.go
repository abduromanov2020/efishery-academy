package handler

import (
	"ecommerce-project/entity"
	"ecommerce-project/entity/response"
	"ecommerce-project/usecase"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ProductsHandler struct {
	productsUsecase usecase.IProductsUsecase
}

func NewProductsHandler(productsUsecase usecase.IProductsUsecase) *ProductsHandler {
	return &ProductsHandler{productsUsecase: productsUsecase}
}

func (p ProductsHandler) CreateProduct(c echo.Context) error {
	productRequest := entity.CreateProductsRequest{}

	c.Bind(&productRequest)

	if productRequest.Name == "" || productRequest.Price == 0 || productRequest.Stock == 0 {
		return c.JSON(http.StatusBadRequest, response.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: "Data Cannot Be Empty",
			Data:    []entity.CreateProductsRequest{},
		})
	}

	err := p.productsUsecase.CreateProduct(productRequest)

	if err != nil {
		return c.JSON(http.StatusBadRequest, response.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: "Invalid Req Body",
			Data:    err.Error(),
		})
	}

	return c.JSON(http.StatusOK, response.BaseResponse{
		Code:    http.StatusOK,
		Message: "Success Create Product",
		Data:    productRequest,
	})
}

func (p ProductsHandler) GetListProduct(c echo.Context) error {
	products, err := p.productsUsecase.GetListProduct()

	if err != nil {
		return c.JSON(http.StatusBadRequest, response.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: "Failed Get List Product",
			Data:    err.Error(),
		})

	}

	if len(products) == 0 {
		return c.JSON(http.StatusBadRequest, response.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: "Empty List Product",
			Data:    []entity.Products{},
		})
	}

	return c.JSON(http.StatusOK, response.BaseResponse{
		Code:    http.StatusOK,
		Message: "Success Get List Product",
		Data:    products,
	})

}

func (p ProductsHandler) GetProductByID(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	product, err := p.productsUsecase.GetProductByID(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, response.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: "Failed Get Product By ID",
			Data:    err.Error(),
		})
	}

	return c.JSON(http.StatusOK, response.BaseResponse{
		Code:    http.StatusOK,
		Message: "Success Get Product By ID",
		Data:    product,
	})
}

func (p ProductsHandler) GetProductByPrice(c echo.Context) error {

	minPrice, _ := strconv.Atoi(c.Param("minPrice"))
	maxPrice, _ := strconv.Atoi(c.Param("maxPrice"))

	products, err := p.productsUsecase.GetProductByPrice(minPrice, maxPrice)

	if err != nil {
		return c.JSON(http.StatusBadRequest, response.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: "Failed Get Product By Price",
			Data:    err.Error(),
		})
	}

	if len(products) == 0 {
		return c.JSON(http.StatusBadRequest, response.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: "Empty List Product",
			Data:    []entity.Products{},
		})
	}

	return c.JSON(http.StatusOK, response.BaseResponse{
		Code:    http.StatusOK,
		Message: "Success Get Product By Price",
		Data:    products,
	})
}

func (p ProductsHandler) GetProductByCategory(c echo.Context) error {
	category := c.Param("category")

	products, err := p.productsUsecase.GetProductByCategory(category)

	if err != nil {
		return c.JSON(http.StatusBadRequest, response.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: "Failed Get Product By Category",
			Data:    err.Error(),
		})
	}

	if len(products) == 0 {
		return c.JSON(http.StatusBadRequest, response.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: "Empty List Product",
			Data:    []entity.Products{},
		})
	}

	return c.JSON(http.StatusOK, response.BaseResponse{
		Code:    http.StatusOK,
		Message: "Success Get Product By Category",
		Data:    products,
	})

}

func (p ProductsHandler) UpdateProductByID(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	productRequest := entity.UpdateProductsRequest{}

	c.Bind(&productRequest)

	err := p.productsUsecase.UpdateProductByID(id, productRequest)

	if err != nil {
		return c.JSON(http.StatusBadRequest, response.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: "Failed Update Product",
			Data:    err.Error(),
		})
	}

	return c.JSON(http.StatusOK, response.BaseResponse{
		Code:    http.StatusOK,
		Message: "Success Update Product",
		Data:    productRequest,
	})
}

func (p ProductsHandler) DeleteProductByID(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	err := p.productsUsecase.DeleteProductByID(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, response.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: "Failed Delete Product",
			Data:    err.Error(),
		})
	}

	return c.JSON(http.StatusOK, response.BaseResponse{
		Code:    http.StatusOK,
		Message: "Success Delete Product",
		Data:    []entity.Products{},
	})
}
