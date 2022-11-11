package repository

import (
	"ecommerce-project/entity"

	"gorm.io/gorm"
)

type ICategoryRepository interface {
	Create(category entity.Category) error
	GetAll() ([]entity.Category, error)
	GetByID(id int) (entity.Category, error)
	Update(category entity.Category) error
	Delete(id int) error
}

type CategoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) *CategoryRepository {
	return &CategoryRepository{db: db}
}

func (c CategoryRepository) Create(category entity.Category) error {
	err := c.db.Create(&category).Error
	if err != nil {
		return err
	}

	return nil
}

func (c CategoryRepository) GetAll() ([]entity.Category, error) {
	var categories []entity.Category

	if err := c.db.Find(&categories).Error; err != nil {
		return categories, err
	}

	return categories, nil
}

func (c CategoryRepository) GetByID(id int) (entity.Category, error) {
	var category entity.Category

	if err := c.db.First(&category, id).Error; err != nil {
		return category, err
	}

	return category, nil
}

func (c CategoryRepository) Update(category entity.Category) error {
	if err := c.db.Save(&category).Error; err != nil {
		return err
	}

	return nil
}

func (c CategoryRepository) Delete(id int) error {

	if err := c.db.Where("category_id = ?", id).Delete(&entity.Products{}).Error; err != nil {
		return err
	}

	if err := c.db.Delete(&entity.Category{}, id).Error; err != nil {
		return err
	}

	return nil
}
