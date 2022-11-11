package handler

import (
	"ecommerce-project/entity"
	"ecommerce-project/entity/response"
	"ecommerce-project/usecase"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type CategoryHandler struct {
	categoryUsecase usecase.ICategoryUsecase
}

func NewCategoryHandler(categoryUsecase usecase.ICategoryUsecase) *CategoryHandler {
	return &CategoryHandler{categoryUsecase: categoryUsecase}
}

func (category CategoryHandler) CreateCategory(c echo.Context) error {
	categoryRequest := entity.CreateCategoryRequest{}

	c.Bind(&categoryRequest)

	if categoryRequest.CategoryName == "" {
		return c.JSON(http.StatusBadRequest, response.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: "Data Cannot Be Empty",
			Data:    nil,
		})
	}

	err := category.categoryUsecase.CreateCategory(categoryRequest)

	if err != nil {
		return c.JSON(http.StatusBadRequest, response.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: "Invalid Req Body",
			Data:    err.Error(),
		})
	}

	return c.JSON(http.StatusOK, response.BaseResponse{
		Code:    http.StatusOK,
		Message: "Success Create Category",
		Data:    categoryRequest,
	})
}

func (category CategoryHandler) GetListCategory(c echo.Context) error {
	categories, err := category.categoryUsecase.GetListCategory()

	if err != nil {
		return c.JSON(http.StatusBadRequest, response.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: "Failed Get List Category",
			Data:    err.Error(),
		})

	}

	if len(categories) == 0 {
		return c.JSON(http.StatusBadRequest, response.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: "Data Not Found",
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, response.BaseResponse{
		Code:    http.StatusOK,
		Message: "Success Get List Category",
		Data:    categories,
	})
}

func (category CategoryHandler) GetCategoryById(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	categoryByID, err := category.categoryUsecase.GetCategoryByID(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, response.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: "Failed Get Category By ID",
			Data:    err.Error(),
		})
	}

	return c.JSON(http.StatusOK, response.BaseResponse{
		Code:    http.StatusOK,
		Message: "Success Get Category By Id",
		Data:    categoryByID,
	})
}

func (category CategoryHandler) UpdateCategoryByID(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	categoryRequest := entity.UpdateCategoryRequest{}

	c.Bind(&categoryRequest)

	if categoryRequest.CategoryName == "" {
		return c.JSON(http.StatusBadRequest, response.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: "Data Cannot Be Empty",
			Data:    nil,
		})
	}

	err := category.categoryUsecase.UpdateCategoryByID(id, categoryRequest)

	if err != nil {
		return c.JSON(http.StatusBadRequest, response.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: "Failed Update Category",
			Data:    err.Error(),
		})
	}

	return c.JSON(http.StatusOK, response.BaseResponse{
		Code:    http.StatusOK,
		Message: "Success Update Category",
		Data:    categoryRequest,
	})
}

func (category CategoryHandler) DeleteCategoryByID(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	err := category.categoryUsecase.DeleteCategoryByID(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, response.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: "Failed Delete Category",
			Data:    err.Error(),
		})
	}

	return c.JSON(http.StatusOK, response.BaseResponse{
		Code:    http.StatusOK,
		Message: "Success Delete Category",
		Data:    nil,
	})
}
