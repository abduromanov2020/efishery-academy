package usecase

import (
	"ecommerce-project/entity"
	"ecommerce-project/repository"

	"github.com/jinzhu/copier"
)

type IPaymentUsecase interface {
	CreatePayment(req entity.CreatePaymentRequest) error
	GetListPayment() ([]entity.GetPaymentResponse, error)
	GetPaymentByID(id int) (entity.GetPaymentResponse, error)
	DeletePaymentByID(id int) error
}

type PaymentUsecase struct {
	paymentRepository repository.IPaymentRepository
}

func NewPaymentUsecase(paymentRepository repository.IPaymentRepository) *PaymentUsecase {
	return &PaymentUsecase{paymentRepository: paymentRepository}
}

func (p PaymentUsecase) CreatePayment(req entity.CreatePaymentRequest) error {
	payment := entity.Payment{}

	copier.Copy(&payment, &req)
	err := p.paymentRepository.Create(payment)
	if err != nil {
		return err
	}

	return nil
}

func (p PaymentUsecase) GetListPayment() ([]entity.GetPaymentResponse, error) {
	payments, err := p.paymentRepository.GetAll()
	if err != nil {
		return nil, err
	}

	var paymentResponse []entity.GetPaymentResponse
	copier.Copy(&paymentResponse, &payments)

	return paymentResponse, nil
}

func (p PaymentUsecase) GetPaymentByID(id int) (entity.GetPaymentResponse, error) {
	payment, err := p.paymentRepository.GetByID(id)
	if err != nil {
		return entity.GetPaymentResponse{}, err
	}

	var paymentResponse entity.GetPaymentResponse
	copier.Copy(&paymentResponse, &payment)

	return paymentResponse, nil

}

func (p PaymentUsecase) DeletePaymentByID(id int) error {
	_, err := p.paymentRepository.GetByID(id)

	if err != nil {
		return err
	}

	err = p.paymentRepository.Delete(id)

	if err != nil {
		return err
	}

	return nil

}
