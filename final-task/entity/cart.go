package entity

type Cart struct {
	ID           uint           `json:"id" gorm:"column:id;primaryKey;autoIncrement"`
	UserID       uint           `json:"user_id"`
	Cart_Product []Cart_Product `gorm:"foreignKey:CartID"`
}

type GetCartResponse struct {
	ID           uint           `json:"id"`
	UserID       uint           `json:"user_id"`
	Cart_Product []Cart_Product `gorm:"foreignKey:CartID"`
}
