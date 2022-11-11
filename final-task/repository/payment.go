package repository

import (
	"ecommerce-project/entity"

	"gorm.io/gorm"
)

type IPaymentRepository interface {
	Create(payment entity.Payment) error
	GetAll() ([]entity.Payment, error)
	GetByID(id int) (entity.Payment, error)
	Update(payment entity.Payment) error
	Delete(id int) error
}

type PaymentRepository struct {
	db *gorm.DB
}

func NewPaymentRepository(db *gorm.DB) *PaymentRepository {
	return &PaymentRepository{db: db}
}

func (p PaymentRepository) Create(payment entity.Payment) error {
	err := p.db.Create(&payment).Error
	if err != nil {
		return err
	}

	return nil
}

func (p PaymentRepository) GetAll() ([]entity.Payment, error) {
	var payments []entity.Payment

	if err := p.db.Find(&payments).Error; err != nil {
		return payments, err
	}

	return payments, nil
}

func (p PaymentRepository) GetByID(id int) (entity.Payment, error) {
	var payment entity.Payment

	if err := p.db.First(&payment, id).Error; err != nil {
		return payment, err
	}

	return payment, nil
}

func (p PaymentRepository) Update(payment entity.Payment) error {

	if err := p.db.Save(&payment).Error; err != nil {
		return err
	}

	return nil
}

func (p PaymentRepository) Delete(id int) error {

	if err := p.db.Delete(&entity.Payment{}, id).Error; err != nil {
		return err
	}

	return nil
}
