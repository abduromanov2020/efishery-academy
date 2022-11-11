package repository

import (
	"ecommerce-project/entity"

	"gorm.io/gorm"
)

type IProductDetailRepository interface {
	Create(productDetail entity.Product_Detail) error
	GetAll() ([]entity.Product_Detail, error)
	GetByID(id int) (entity.Product_Detail, error)
	Update(productDetail entity.Product_Detail) error
	Delete(id int) error
}

type ProductDetailRepository struct {
	db *gorm.DB
}

func NewProductDetailRepository(db *gorm.DB) *ProductDetailRepository {
	return &ProductDetailRepository{db: db}
}

func (p ProductDetailRepository) Create(productDetail entity.Product_Detail) error {

	err := p.db.Create(&productDetail).Error
	if err != nil {
		return err
	}

	return nil
}

func (p ProductDetailRepository) GetAll() ([]entity.Product_Detail, error) {
	var productDetails []entity.Product_Detail

	if err := p.db.Joins("Products").Find(&productDetails).Error; err != nil {
		return productDetails, err
	}

	return productDetails, nil
}

func (p ProductDetailRepository) GetByID(id int) (entity.Product_Detail, error) {
	var productDetail entity.Product_Detail

	if err := p.db.Joins("Products").First(&productDetail, id).Error; err != nil {
		return productDetail, err
	}

	return productDetail, nil
}

func (p ProductDetailRepository) Update(productDetail entity.Product_Detail) error {
	if err := p.db.Save(&productDetail).Error; err != nil {
		return err
	}

	return nil
}

func (p ProductDetailRepository) Delete(id int) error {
	if err := p.db.Delete(&entity.Product_Detail{}, id).Error; err != nil {
		return err
	}

	return nil
}
