package repository

import (
	"ecommerce-project/entity"

	"gorm.io/gorm"
)

type IUserRepositroy interface {
	Create(user entity.User) error
	GetAll() ([]entity.User, error)
	GetByID(id int) (entity.User, error)
	Update(user entity.User) error
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

	uid := user.ID

	if err := u.db.Create(&entity.Cart{UserID: uid}).Error; err != nil {
		return err
	}

	return nil
}

func (u UserRepository) GetAll() ([]entity.User, error) {
	var users []entity.User

	if err := u.db.Find(&users).Error; err != nil {
		return users, err
	}

	return users, nil
}

func (u UserRepository) GetByID(id int) (entity.User, error) {
	var user entity.User

	if err := u.db.First(&user, id).Error; err != nil {
		return user, err
	}

	return user, nil
}

func (u UserRepository) Update(user entity.User) error {
	if err := u.db.Save(&user).Error; err != nil {
		return err
	}

	return nil
}

func (u UserRepository) Delete(id int) error {
	var cart entity.Cart
	var payment entity.Payment

	if err := u.db.Where("cart_id = ?", id).Delete(&payment).Error; err != nil {
		return err
	}

	if err := u.db.Where("id = ?", id).Delete(&cart).Error; err != nil {
		return err
	}

	if err := u.db.Delete(&entity.User{}, id).Error; err != nil {
		return err
	}

	return nil
}
