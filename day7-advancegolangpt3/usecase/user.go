package usecase

import (
	"golangday7/entity"
	"golangday7/entity/response"
	"golangday7/repository"

	"github.com/jinzhu/copier"
)

type IUserUsecase interface {
	CreateUser(user response.CreateUserRequest) error
	GetListUser() ([]response.GetUserResponse, error)
	GetUserById(id int) (response.GetUserResponse, error)
}

type UserUsecase struct {
	userRepository repository.IUserRepository
}

func NewUserUseCase(userRepository repository.IUserRepository) *UserUsecase {
	return &UserUsecase{userRepository: userRepository}
}

func (u UserUsecase) CreateUser(req response.CreateUserRequest) error {
	users := entity.User{}
	copier.Copy(&users, &req)

	err := u.userRepository.Create(users)
	if err != nil {
		return err
	}

	return nil
}

func (u UserUsecase) GetListUser() ([]response.GetUserResponse, error) {
	users, err := u.userRepository.GetAll()
	if err != nil {
		return nil, err
	}

	userResponse := []response.GetUserResponse{}
	copier.Copy(&userResponse, &users)
	return userResponse, nil

}

func (u UserUsecase) GetUserById(id int) (response.GetUserResponse, error) {
	users, err := u.userRepository.GetById(id)

	if users == (entity.User{}) {
		return response.GetUserResponse{}, err
	}

	if err != nil {
		return response.GetUserResponse{}, err
	}

	userResponse := response.GetUserResponse{}
	copier.Copy(&userResponse, &users)

	return userResponse, nil

}

func (u UserUsecase) UpdateUserById(userRequest response.UpdateUserRequest, id int) (response.GetUserResponse, error) {
	user, err := u.userRepository.GetById(id)

	if err != nil {
		return response.GetUserResponse{}, err
	}

	copier.CopyWithOption(&user, &userRequest, copier.Option{IgnoreEmpty: true})

	user, _ = u.userRepository.Update(user)

	userResponse := response.GetUserResponse{}
	copier.Copy(&userResponse, &user)

	return userResponse, nil
}

func (u UserUsecase) DeleteUserById(id int) error {
	_, err := u.userRepository.GetById(id)

	if err != nil {
		return err
	}

	err = u.userRepository.Delete(id)

	return err
}
