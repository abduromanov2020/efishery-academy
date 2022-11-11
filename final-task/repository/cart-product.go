package repository

import (
	"ecommerce-project/entity"

	"gorm.io/gorm"
)

type ICartProductRepository interface {
	Create(cart entity.Cart_Product) error
	GetAll() ([]entity.Cart_Product, error)
	GetByID(id int) (entity.Cart_Product, error)
	Update(cart entity.Cart_Product) error
	Delete(id int) error
}

type CartProductRepository struct {
	db *gorm.DB
}

func NewCartProductRepository(db *gorm.DB) *CartProductRepository {
	return &CartProductRepository{db: db}
}

func (c CartProductRepository) Create(cartProduct entity.Cart_Product) error {

	err := c.db.Create(&cartProduct).Error
	if err != nil {
		return err
	}

	return nil
}

func (c CartProductRepository) GetAll() ([]entity.Cart_Product, error) {
	var cartProduct []entity.Cart_Product

	if err := c.db.Joins("Products").Find(&cartProduct).Error; err != nil {
		return cartProduct, err
	}

	return cartProduct, nil
}

func (c CartProductRepository) GetByID(id int) (entity.Cart_Product, error) {
	var cartProduct entity.Cart_Product

	if err := c.db.First(&cartProduct, id).Error; err != nil {
		return cartProduct, err
	}

	return cartProduct, nil
}

func (c CartProductRepository) Update(cartProduct entity.Cart_Product) error {
	if err := c.db.Save(&cartProduct).Error; err != nil {
		return err
	}

	return nil
}

func (c CartProductRepository) Delete(id int) error {
	if err := c.db.Delete(&entity.Cart_Product{}, id).Error; err != nil {
		return err
	}

	return nil
}
