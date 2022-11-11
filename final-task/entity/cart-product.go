package entity

type Cart_Product struct {
	ID        uint     `json:"id" gorm:"column:id;primaryKey;autoIncrement"`
	ProductID uint     `json:"product_id" gorm:"unique"`
	Products  Products `gorm:"foreignKey:ProductID"`
	Quantity  uint     `json:"quantity"`
	CartID    uint     `json:"cart_id"`
}

type CreateCartProductRequest struct {
	ProductID uint `json:"product_id" gorm:"unique"`
	Quantity  uint `json:"quantity" gorm:"not null"`
	CartID    uint `json:"cart_id"`
}

type UpdateCartProductRequest struct {
	ProductID uint `json:"product_id" gorm:"unique"`
	Quantity  uint `json:"quantity" gorm:"not null"`
	CartID    uint `json:"cart_id"`
}

type GetCartProductResponse struct {
	ID       uint `json:"id"`
	Products []Products
	Quantity uint `json:"quantity" gorm:"not null"`
	CartID   uint `json:"cart_id"`
}
