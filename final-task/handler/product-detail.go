package handler

import (
	"ecommerce-project/entity"
	"ecommerce-project/entity/response"
	"ecommerce-project/usecase"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ProductDetailHandler struct {
	productDetailUsecase usecase.IProductDetailUsecase
}

func NewProductDetailHandler(productDetailUsecase usecase.IProductDetailUsecase) *ProductDetailHandler {
	return &ProductDetailHandler{productDetailUsecase: productDetailUsecase}
}

func (p ProductDetailHandler) CreateProductDetail(c echo.Context) error {
	productDetailRequest := entity.CreateProductDetailRequest{}

	c.Bind(&productDetailRequest)

	if productDetailRequest.ProductID == 0 || productDetailRequest.ProductDescription == "" {
		return c.JSON(http.StatusBadRequest, response.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: "Data Cannot Be Empty",
			Data:    []entity.CreateProductDetailRequest{},
		})
	}

	err := p.productDetailUsecase.CreateProductDetail(productDetailRequest)

	if err != nil {
		return c.JSON(http.StatusBadRequest, response.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: "Invalid Req Body",
			Data:    err.Error(),
		})
	}

	return c.JSON(http.StatusOK, response.BaseResponse{
		Code:    http.StatusOK,
		Message: "Success Create Product Detail",
		Data:    productDetailRequest,
	})
}

func (p ProductDetailHandler) GetListProductDetail(c echo.Context) error {
	product, err := p.productDetailUsecase.GetListProductDetail()

	if err != nil {
		return c.JSON(http.StatusBadRequest, response.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: "Failed Get List Product Detail",
			Data:    err.Error(),
		})

	}

	if len(product) == 0 {
		return c.JSON(http.StatusOK, response.BaseResponse{
			Code:    http.StatusOK,
			Message: "Empty List Product Detail",
			Data:    []entity.Product_Detail{},
		})
	}

	return c.JSON(http.StatusOK, response.BaseResponse{
		Code:    http.StatusOK,
		Message: "Success Get List Product Detail",
		Data:    product,
	})

}

func (p ProductDetailHandler) GetProductDetailByID(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	product, err := p.productDetailUsecase.GetProductDetailByID(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, response.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: "Failed Get Product Detail By ID",
			Data:    err.Error(),
		})

	}

	return c.JSON(http.StatusOK, response.BaseResponse{
		Code:    http.StatusOK,
		Message: "Success Get Product Detail By ID",
		Data:    product,
	})

}

func (p ProductDetailHandler) UpdateProductDetailByID(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	productDetailRequest := entity.UpdateProductDetailRequest{}

	c.Bind(&productDetailRequest)

	err := p.productDetailUsecase.UpdateProductDetailByID(id, productDetailRequest)

	if err != nil {
		return c.JSON(http.StatusBadRequest, response.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: "Failed Update Product Detail By ID",
			Data:    err.Error(),
		})
	}

	return c.JSON(http.StatusOK, response.BaseResponse{
		Code:    http.StatusOK,
		Message: "Success Update Product Detail",
		Data:    productDetailRequest,
	})
}

func (p ProductDetailHandler) DeleteProductDetailByID(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	err := p.productDetailUsecase.DeleteProductDetailByID(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, response.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: "Failed Delete Product Detail By ID",
			Data:    err.Error(),
		})
	}

	return c.JSON(http.StatusOK, response.BaseResponse{
		Code:    http.StatusOK,
		Message: "Success Delete Product Detail",
		Data:    []entity.Product_Detail{},
	})
}
