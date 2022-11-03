package handler

import (
	"golangday7/entity/response"
	"golangday7/usecase"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	userUsecase *usecase.UserUsecase
}

func NewUserHandler(userCase *usecase.UserUsecase) *UserHandler {
	return &UserHandler{userUsecase: userCase}
}

func (h UserHandler) CreateUser(ctx *fiber.Ctx) error {
	userRequest := response.CreateUserRequest{}
	if err := ctx.BodyParser(&userRequest); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.BaseResponse{
			Code:    fiber.StatusBadRequest,
			Message: "Invalid req Body",
			Data:    err.Error(),
		})
	}

	if err := h.userUsecase.CreateUser(userRequest); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.BaseResponse{
			Code:    fiber.StatusBadRequest,
			Message: "Failed to Create Server",
			Data:    err.Error(),
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(response.BaseResponse{
		Code:    fiber.StatusCreated,
		Message: "Success Create User",
	})

}

func (h UserHandler) GetListUser(ctx *fiber.Ctx) error {
	users, err := h.userUsecase.GetListUser()
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.BaseResponse{
			Code:    fiber.StatusBadRequest,
			Message: "Failed to Get List User",
			Data:    nil,
		})
	}

	if len(users) == 0 {
		return ctx.Status(fiber.StatusNoContent).JSON(response.BaseResponse{
			Code:    fiber.StatusNoContent,
			Message: "User Not Found",
			Data:    nil,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(response.BaseResponse{
		Code:    fiber.StatusOK,
		Message: "Success Get List User",
		Data:    users,
	})
}

func (h UserHandler) GetUserById(ctx *fiber.Ctx) error {
	userId, _ := strconv.Atoi(ctx.Params("id"))

	user, err := h.userUsecase.GetUserById(userId)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.BaseResponse{
			Code:    fiber.StatusBadRequest,
			Message: "Failed to Get User By Id",
			Data:    err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(response.BaseResponse{
		Code:    fiber.StatusOK,
		Message: "Success Get User By Id",
		Data:    user,
	})
}

func (h UserHandler) UpdateUserById(ctx *fiber.Ctx) error {
	userId, _ := strconv.Atoi(ctx.Params("id"))

	userRequest := response.UpdateUserRequest{}

	if err := ctx.BodyParser(&userRequest); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.BaseResponse{
			Code:    fiber.StatusBadRequest,
			Message: "Invalid Request Body",
			Data:    err.Error(),
		})
	}
	user, err := h.userUsecase.UpdateUserById(userRequest, userId)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.BaseResponse{
			Code:    fiber.StatusBadRequest,
			Message: "Failed to Update User",
			Data:    err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(response.BaseResponse{
		Code:    fiber.StatusOK,
		Message: "Success Update User",
		Data:    user,
	})
}

func (h UserHandler) DeleteUserById(ctx *fiber.Ctx) error {
	userId, _ := strconv.Atoi(ctx.Params("id"))

	err := h.userUsecase.DeleteUserById(userId)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.BaseResponse{
			Code:    fiber.StatusBadRequest,
			Message: "Failed to Delete User",
			Data:    err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(response.BaseResponse{
		Code:    fiber.StatusOK,
		Message: "Success Delete User",
	})
}
