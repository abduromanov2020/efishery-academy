package entity

type Payment struct {
	CartID uint   `json:"cart_id" gorm:"primaryKey"`
	Cart   Cart   `gorm:"foreignKey:CartID"`
	File   string `json:"file"`
}

type CreatePaymentRequest struct {
	CartID uint   `json:"cart_id"`
	File   string `json:"file" `
}

type GetPaymentResponse struct {
	CartID uint   `json:"cart_id"`
	File   string `json:"file" `
}
