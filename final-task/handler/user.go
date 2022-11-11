package handler

import (
	"ecommerce-project/entity"
	"ecommerce-project/entity/response"
	"ecommerce-project/usecase"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	userUsecase *usecase.UserUsecase
}

func NewUserHandler(userCase *usecase.UserUsecase) *UserHandler {
	return &UserHandler{userUsecase: userCase}
}

func (h UserHandler) CreateUser(c echo.Context) error {
	userRequest := entity.CreateUserRequest{}

	c.Bind(&userRequest)

	if userRequest.Username == "" || userRequest.Email == "" || userRequest.Password == "" {
		return c.JSON(http.StatusBadRequest, response.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: "Data Cannot Be Empty",
			Data:    []entity.CreateUserRequest{},
		})
	}

	err := h.userUsecase.CreateUser(userRequest)

	if err != nil {
		return c.JSON(http.StatusBadRequest, response.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: "Invalid Req Body",
			Data:    err.Error(),
		})
	}

	return c.JSON(http.StatusOK, response.BaseResponse{
		Code:    http.StatusOK,
		Message: "Success Create User",
		Data:    userRequest,
	})
}

func (h UserHandler) GetListUser(c echo.Context) error {
	users, err := h.userUsecase.GetListUser()

	if err != nil {
		return c.JSON(http.StatusBadRequest, response.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: "Failed Get List User",
			Data:    err.Error(),
		})
	}

	if len(users) == 0 {
		return c.JSON(http.StatusNotFound, response.BaseResponse{
			Code:    http.StatusNotFound,
			Message: "Empty List User",
			Data:    []entity.User{},
		})
	}

	return c.JSON(http.StatusOK, response.BaseResponse{
		Code:    http.StatusOK,
		Message: "Success Get List User",
		Data:    users,
	})
}

func (h UserHandler) GetUserByID(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	user, err := h.userUsecase.GetUserByID(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, response.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: "Failed Get User",
			Data:    err.Error(),
		})
	}

	return c.JSON(http.StatusOK, response.BaseResponse{
		Code:    http.StatusOK,
		Message: "Success Get User By ID",
		Data:    user,
	})
}

func (h UserHandler) UpdateUserByID(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	userRequest := entity.UpdateUserRequest{}

	c.Bind(&userRequest)

	if userRequest == (entity.UpdateUserRequest{}) {
		return c.JSON(http.StatusBadRequest, response.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: "At Least One Field Must Be Filled",
			Data:    []entity.UpdateUserRequest{},
		})
	}

	err := h.userUsecase.UpdateUserByID(id, userRequest)

	if err != nil {
		return c.JSON(http.StatusBadRequest, response.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: "Failed Update User",
			Data:    err.Error(),
		})
	}

	return c.JSON(http.StatusOK, response.BaseResponse{
		Code:    http.StatusOK,
		Message: "Success Update User With ID : " + strconv.Itoa(id),
		Data:    userRequest,
	})
}

func (h UserHandler) DeleteUserByID(c echo.Context) error {

	id, _ := strconv.Atoi(c.Param("id"))

	err := h.userUsecase.DeleteUserByID(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, response.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: "Failed Delete User",
			Data:    err.Error(),
		})
	}

	return c.JSON(http.StatusOK, response.BaseResponse{
		Code:    http.StatusOK,
		Message: "Success Delete User",
		Data:    []string{"User With ID : " + strconv.Itoa(id)},
	})
}
