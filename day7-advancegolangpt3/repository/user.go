package repository

import (
	"golangday7/entity"

	"gorm.io/gorm"
)

type IUserRepository interface {
	Create(user entity.User) error
	GetAll() ([]entity.User, error)
	GetById(id int) (entity.User, error)
	Update(user entity.User) (entity.User, error)
	Delete(id int) error
}

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (u UserRepository) Create(user entity.User) error {
	if err := u.db.Create(&user).Error; err != nil {
		return err
	}

	return nil
}

func (u UserRepository) GetAll() ([]entity.User, error) {
	var users []entity.User

	if err := u.db.Find(&users).Error; err != nil {
		return []entity.User{}, err
	}

	return users, nil
}

func (u UserRepository) GetById(id int) (entity.User, error) {
	var user entity.User

	if err := u.db.First(&user, id).Error; err != nil {
		return entity.User{}, err
	}

	return user, nil
}

func (u UserRepository) Update(user entity.User) (entity.User, error) {

	if err := u.db.Save(&user).Error; err != nil {
		return entity.User{}, err
	}

	return user, nil
}

func (u UserRepository) Delete(id int) error {
	var user entity.User
	if err := u.db.Delete(&user, id).Error; err != nil {
		return err
	}

	return nil
}
