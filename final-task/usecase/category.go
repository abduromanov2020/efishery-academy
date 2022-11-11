package usecase

import (
	"ecommerce-project/entity"
	"ecommerce-project/repository"

	"github.com/jinzhu/copier"
)

type ICategoryUsecase interface {
	CreateCategory(category entity.CreateCategoryRequest) error
	GetListCategory() ([]entity.GetCategoryResponse, error)
	GetCategoryByID(id int) (entity.GetCategoryResponse, error)
	UpdateCategoryByID(id int, category entity.UpdateCategoryRequest) error
	DeleteCategoryByID(id int) error
}

type CategoryUsecase struct {
	categoryRepository repository.ICategoryRepository
}

func NewCategoryUsecase(categoryRepository repository.ICategoryRepository) *CategoryUsecase {
	return &CategoryUsecase{categoryRepository: categoryRepository}
}

func (c CategoryUsecase) CreateCategory(req entity.CreateCategoryRequest) error {
	categoryEntity := entity.Category{}

	copier.Copy(&categoryEntity, &req)
	err := c.categoryRepository.Create(categoryEntity)
	if err != nil {
		return err
	}

	return nil
}

func (c CategoryUsecase) GetListCategory() ([]entity.GetCategoryResponse, error) {
	categories, err := c.categoryRepository.GetAll()
	if err != nil {
		return nil, err
	}

	var categoryResponse []entity.GetCategoryResponse
	copier.Copy(&categoryResponse, &categories)

	return categoryResponse, nil
}

func (c CategoryUsecase) GetCategoryByID(id int) (entity.GetCategoryResponse, error) {
	category, err := c.categoryRepository.GetByID(id)
	if err != nil {
		return entity.GetCategoryResponse{}, err
	}

	var categoryResponse entity.GetCategoryResponse
	copier.Copy(&categoryResponse, &category)

	return categoryResponse, nil
}

func (c CategoryUsecase) UpdateCategoryByID(id int, req entity.UpdateCategoryRequest) error {
	category, err := c.categoryRepository.GetByID(id)
	if err != nil {
		return err
	}

	copier.CopyWithOption(&category, &req, copier.Option{IgnoreEmpty: true})

	err = c.categoryRepository.Update(category)

	if err != nil {
		return err
	}

	return nil
}

func (c CategoryUsecase) DeleteCategoryByID(id int) error {
	_, err := c.categoryRepository.GetByID(id)

	if err != nil {
		return err
	}

	err = c.categoryRepository.Delete(id)

	if err != nil {
		return err
	}

	return nil
}
