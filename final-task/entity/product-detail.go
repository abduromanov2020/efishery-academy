package entity

type Product_Detail struct {
	ProductID          uint     `json:"product_id" gorm:"primaryKey;autoIncrement"`
	Products           Products `gorm:"foreignKey:ProductID"`
	ProductDescription string   `json:"description" `
}

type CreateProductDetailRequest struct {
	ProductID          uint     `json:"product_id"`
	Products           Products `gorm:"foreignKey:ProductID"`
	ProductDescription string   `json:"description" `
}

type UpdateProductDetailRequest struct {
	ProductDescription string `json:"description" `
}

type GetProductDetailResponse struct {
	ProductID          uint     `json:"product_id"`
	Products           Products `gorm:"foreignKey:ProductID"`
	ProductDescription string   `json:"description" `
}
