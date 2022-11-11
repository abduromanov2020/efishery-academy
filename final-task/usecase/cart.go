package usecase

import (
	"ecommerce-project/entity"
	"ecommerce-project/repository"

	"github.com/jinzhu/copier"
)

type ICartUsecase interface {
	GetListCart() ([]entity.Cart, error)
	GetCartByID(id int) (entity.GetCartResponse, error)
}

type CartUsecase struct {
	cartRepository repository.ICartRepository
}

func NewCartUsecase(cartRepository repository.ICartRepository) *CartUsecase {
	return &CartUsecase{cartRepository: cartRepository}
}

func (c CartUsecase) GetListCart() ([]entity.Cart, error) {
	carts, err := c.cartRepository.GetAll()
	if err != nil {
		return carts, err
	}

	cartResponse := []entity.GetCartResponse{}
	copier.Copy(&cartResponse, &carts)

	return carts, nil
}

func (c CartUsecase) GetCartByID(id int) (entity.GetCartResponse, error) {
	cart, err := c.cartRepository.GetByID(id)

	if err != nil {
		return entity.GetCartResponse{}, err
	}

	var cartResponse entity.GetCartResponse
	copier.Copy(&cartResponse, &cart)

	return cartResponse, nil

}
