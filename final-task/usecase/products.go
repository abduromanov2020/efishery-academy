package usecase

import (
	"ecommerce-project/entity"
	"ecommerce-project/repository"

	"github.com/jinzhu/copier"
)

type IProductsUsecase interface {
	CreateProduct(req entity.CreateProductsRequest) error
	GetListProduct() ([]entity.GetProductsResponse, error)
	GetProductByID(id int) (entity.GetProductsResponse, error)
	GetProductByPrice(minPrice, maxPrice int) ([]entity.GetProductsResponse, error)
	GetProductByCategory(category string) ([]entity.GetProductsResponse, error)
	UpdateProductByID(id int, req entity.UpdateProductsRequest) error
	DeleteProductByID(id int) error
}

type ProductsUsecase struct {
	productRepository repository.IProductsRepository
}

func NewProductsUsecase(productRepository repository.IProductsRepository) *ProductsUsecase {
	return &ProductsUsecase{productRepository: productRepository}
}

func (p ProductsUsecase) CreateProduct(req entity.CreateProductsRequest) error {
	product := entity.Products{}

	copier.Copy(&product, &req)
	err := p.productRepository.Create(product)
	if err != nil {
		return err
	}

	return nil
}

func (p ProductsUsecase) GetListProduct() ([]entity.GetProductsResponse, error) {
	products, err := p.productRepository.GetAll()
	if err != nil {
		return nil, err
	}

	var productResponse []entity.GetProductsResponse
	copier.Copy(&productResponse, &products)

	return productResponse, nil
}

func (p ProductsUsecase) GetProductByID(id int) (entity.GetProductsResponse, error) {
	product, err := p.productRepository.GetByID(id)

	if err != nil {
		return entity.GetProductsResponse{}, err
	}

	var productResponse entity.GetProductsResponse
	copier.Copy(&productResponse, &product)

	return productResponse, nil

}

func (p ProductsUsecase) GetProductByPrice(minPrice, maxPrice int) ([]entity.GetProductsResponse, error) {
	product, err := p.productRepository.GetByPrice(minPrice, maxPrice)

	if err != nil {
		return []entity.GetProductsResponse{}, err
	}

	var productResponse []entity.GetProductsResponse
	copier.CopyWithOption(&productResponse, &product, copier.Option{IgnoreEmpty: true})

	return productResponse, nil
}

func (p ProductsUsecase) GetProductByCategory(category string) ([]entity.GetProductsResponse, error) {
	product, err := p.productRepository.GetByCategory(category)

	if err != nil {
		return []entity.GetProductsResponse{}, err
	}

	var productResponse []entity.GetProductsResponse
	copier.Copy(&productResponse, &product)

	return productResponse, nil
}

func (p ProductsUsecase) UpdateProductByID(id int, req entity.UpdateProductsRequest) error {
	product, err := p.productRepository.GetByID(id)
	if err != nil {
		return err
	}

	copier.CopyWithOption(&product, &req, copier.Option{IgnoreEmpty: true})
	err = p.productRepository.Update(product)
	if err != nil {
		return err
	}

	return nil
}

func (p ProductsUsecase) DeleteProductByID(id int) error {
	_, err := p.productRepository.GetByID(id)

	if err != nil {
		return err
	}

	err = p.productRepository.Delete(id)

	if err != nil {
		return err
	}

	return nil
}
