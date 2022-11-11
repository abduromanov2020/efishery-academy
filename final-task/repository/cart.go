package repository

import (
	"ecommerce-project/entity"

	"gorm.io/gorm"
)

type ICartRepository interface {
	GetAll() ([]entity.Cart, error)
	GetByID(id int) (entity.Cart, error)
}

type CartRepository struct {
	db *gorm.DB
}

func NewCartRepository(db *gorm.DB) *CartRepository {
	return &CartRepository{db: db}
}

func (c CartRepository) GetAll() ([]entity.Cart, error) {
	var carts []entity.Cart

	if err := c.db.Debug().Preload("Cart_Product").Find(&carts).Error; err != nil {
		return carts, err
	}

	return carts, nil
}

func (c CartRepository) GetByID(id int) (entity.Cart, error) {
	var cart entity.Cart

	if err := c.db.First(&cart, id).Error; err != nil {
		return cart, err
	}

	return cart, nil
}
