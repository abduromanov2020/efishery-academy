package entity

type Payment struct {
	ID     int    `gorm:"primaryKey" json:"id"`
	CartID uint   `json:"cart_id"`
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
