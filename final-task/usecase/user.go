package usecase

import (
	"ecommerce-project/entity"
	"ecommerce-project/repository"

	"github.com/jinzhu/copier"
)

type IUserUsecase interface {
	CreateUser(user entity.CreateUserRequest)
	GetListUser() ([]entity.GetUserResponse, error)
	GetUserByID(id int) (entity.GetUserResponse, error)
	UpdateUserByID(id int, user entity.UpdateUserRequest) error
	DeleteUserByID(id int) error
}

type UserUsecase struct {
	userRepository repository.IUserRepositroy
}

func NewUserUsecase(userRepository repository.IUserRepositroy) *UserUsecase {
	return &UserUsecase{userRepository: userRepository}
}

func (u UserUsecase) CreateUser(req entity.CreateUserRequest) error {
	users := entity.User{}
	copier.Copy(&users, &req)

	err := u.userRepository.Create(users)
	if err != nil {
		return err
	}

	return nil
}

func (u UserUsecase) GetListUser() ([]entity.GetUserResponse, error) {
	users, err := u.userRepository.GetAll()
	if err != nil {
		return nil, err
	}

	var userResponse []entity.GetUserResponse
	copier.Copy(&userResponse, &users)

	return userResponse, nil
}

func (u UserUsecase) GetUserByID(id int) (entity.GetUserResponse, error) {
	user, err := u.userRepository.GetByID(id)
	if err != nil {
		return entity.GetUserResponse{}, err
	}

	var userResponse entity.GetUserResponse
	copier.Copy(&userResponse, &user)

	return userResponse, nil
}

func (u UserUsecase) UpdateUserByID(id int, req entity.UpdateUserRequest) error {
	user, err := u.userRepository.GetByID(id)

	if err != nil {
		return err
	}

	copier.CopyWithOption(&user, &req, copier.Option{IgnoreEmpty: true, DeepCopy: true})

	err = u.userRepository.Update(user)

	if err != nil {
		return err
	}

	return nil
}
func (u UserUsecase) DeleteUserByID(id int) error {
	_, err := u.userRepository.GetByID(id)

	if err != nil {
		return err
	}

	err = u.userRepository.Delete(id)

	if err != nil {
		return err
	}

	return nil
}
