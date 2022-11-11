package repository

import (
	"ecommerce-project/entity"

	"gorm.io/gorm"
)

type IProductsRepository interface {
	Create(product entity.Products) error
	GetAll() ([]entity.Products, error)
	GetByID(id int) (entity.Products, error)
	GetByPrice(minPrice, maxPrice int) ([]entity.Products, error)
	GetByCategory(categoryName string) ([]entity.Products, error)
	Update(product entity.Products) error
	Delete(id int) error
}

type ProductsRepository struct {
	db *gorm.DB
}

func NewProductsRepository(db *gorm.DB) *ProductsRepository {
	return &ProductsRepository{db: db}
}

func (p ProductsRepository) Create(product entity.Products) error {
	err := p.db.Create(&product).Error
	if err != nil {
		return err
	}

	return nil
}

func (p ProductsRepository) GetAll() ([]entity.Products, error) {
	var products []entity.Products

	if err := p.db.Joins("Category").Find(&products).Error; err != nil {
		return products, err
	}

	return products, nil
}

func (p ProductsRepository) GetByID(id int) (entity.Products, error) {
	var product entity.Products

	if err := p.db.Joins("Category").First(&product, id).Error; err != nil {
		return product, err
	}

	return product, nil
}

func (p ProductsRepository) GetByPrice(minPrice, maxPrice int) ([]entity.Products, error) {
	var products []entity.Products

	if err := p.db.Debug().Where("price BETWEEN ? AND ?", minPrice, maxPrice).Find(&products).Error; err != nil {
		return products, err
	}

	return products, nil
}

func (p ProductsRepository) GetByCategory(categoryName string) ([]entity.Products, error) {
	var products []entity.Products

	if err := p.db.Debug().Joins("Category").Where("category_name = ?", categoryName).Find(&products).Error; err != nil {
		return products, err
	}

	return products, nil
}

func (p ProductsRepository) Update(product entity.Products) error {

	if err := p.db.Save(&product).Error; err != nil {
		return err
	}

	return nil
}

func (p ProductsRepository) Delete(id int) error {
	var product entity.Products
	var productDetail entity.Product_Detail

	if err := p.db.Where("product_id = ?", id).Delete(&productDetail).Error; err != nil {
		return err
	}

	if err := p.db.Delete(&product, id).Error; err != nil {
		return err
	}

	return nil
}
