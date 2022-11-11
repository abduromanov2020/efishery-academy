package usecase

import (
	"ecommerce-project/entity"
	"ecommerce-project/repository"

	"github.com/jinzhu/copier"
)

type ICartProductUsecase interface {
	CreateCartProduct(req entity.CreateCartProductRequest) error
	GetListCartProduct() ([]entity.GetCartProductResponse, error)
	GetCartProductByID(id int) (entity.GetCartProductResponse, error)
	UpdateCartProductByID(id int, req entity.UpdateCartProductRequest) error
	DeleteCartProductByID(id int) error
}

type CartProductUsecase struct {
	cartProductRepository repository.ICartProductRepository
}

func NewCartProductUsecase(cartProductRepository repository.ICartProductRepository) *CartProductUsecase {
	return &CartProductUsecase{cartProductRepository: cartProductRepository}
}

func (c CartProductUsecase) CreateCartProduct(req entity.CreateCartProductRequest) error {
	var cartProduct entity.Cart_Product
	copier.Copy(&cartProduct, &req)

	err := c.cartProductRepository.Create(cartProduct)
	if err != nil {
		return err
	}

	return nil
}

func (c CartProductUsecase) GetListCartProduct() ([]entity.GetCartProductResponse, error) {
	cartProduct, err := c.cartProductRepository.GetAll()
	if err != nil {
		return []entity.GetCartProductResponse{}, err
	}

	var cartProductResponse []entity.GetCartProductResponse
	copier.Copy(&cartProductResponse, &cartProduct)

	return cartProductResponse, nil
}

func (c CartProductUsecase) GetCartProductByID(id int) (entity.GetCartProductResponse, error) {
	cartProduct, err := c.cartProductRepository.GetByID(id)
	if err != nil {
		return entity.GetCartProductResponse{}, err
	}

	var cartProductResponse entity.GetCartProductResponse
	copier.Copy(&cartProductResponse, &cartProduct)

	return cartProductResponse, nil
}

func (c CartProductUsecase) UpdateCartProductByID(id int, req entity.UpdateCartProductRequest) error {
	cartProduct, err := c.cartProductRepository.GetByID(id)
	if err != nil {
		return err
	}

	copier.CopyWithOption(&cartProduct, &req, copier.Option{IgnoreEmpty: true})

	err = c.cartProductRepository.Update(cartProduct)
	if err != nil {
		return err
	}

	return nil
}

func (c CartProductUsecase) DeleteCartProductByID(id int) error {
	err := c.cartProductRepository.Delete(id)
	if err != nil {
		return err
	}

	return nil
}
