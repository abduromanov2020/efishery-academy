package usecase

import (
	"ecommerce-project/entity"
	"ecommerce-project/repository"

	"github.com/jinzhu/copier"
)

type IProductDetailUsecase interface {
	CreateProductDetail(req entity.CreateProductDetailRequest) error
	GetListProductDetail() ([]entity.GetProductDetailResponse, error)
	GetProductDetailByID(id int) (entity.GetProductDetailResponse, error)
	UpdateProductDetailByID(id int, req entity.UpdateProductDetailRequest) error
	DeleteProductDetailByID(id int) error
}

type ProductDetailUsecase struct {
	productDetailRepository repository.IProductDetailRepository
}

func NewProductDetailUsecase(productDetailRepository repository.IProductDetailRepository) *ProductDetailUsecase {
	return &ProductDetailUsecase{productDetailRepository: productDetailRepository}
}

func (p ProductDetailUsecase) CreateProductDetail(req entity.CreateProductDetailRequest) error {
	productDetail := entity.Product_Detail{}

	copier.Copy(&productDetail, &req)

	err := p.productDetailRepository.Create(productDetail)

	if err != nil {
		return err
	}

	return nil
}

func (p ProductDetailUsecase) GetListProductDetail() ([]entity.GetProductDetailResponse, error) {
	productDetails, err := p.productDetailRepository.GetAll()

	if err != nil {
		return nil, err
	}

	var productDetailResponse []entity.GetProductDetailResponse
	copier.Copy(&productDetailResponse, &productDetails)

	return productDetailResponse, nil
}

func (p ProductDetailUsecase) GetProductDetailByID(id int) (entity.GetProductDetailResponse, error) {
	productDetail, err := p.productDetailRepository.GetByID(id)

	if err != nil {
		return entity.GetProductDetailResponse{}, err
	}

	var productDetailResponse entity.GetProductDetailResponse
	copier.Copy(&productDetailResponse, &productDetail)

	return productDetailResponse, nil
}

func (p ProductDetailUsecase) UpdateProductDetailByID(id int, req entity.UpdateProductDetailRequest) error {
	productDetail, err := p.productDetailRepository.GetByID(id)

	if err != nil {
		return err
	}

	copier.CopyWithOption(&productDetail, &req, copier.Option{IgnoreEmpty: true})

	err = p.productDetailRepository.Update(productDetail)

	if err != nil {
		return err
	}

	return nil
}

func (p ProductDetailUsecase) DeleteProductDetailByID(id int) error {
	_, err := p.productDetailRepository.GetByID(id)

	if err != nil {
		return err
	}

	err = p.productDetailRepository.Delete(id)

	if err != nil {
		return err
	}

	return nil
}
